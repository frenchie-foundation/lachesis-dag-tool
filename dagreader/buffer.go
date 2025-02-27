package main

import (
	"fmt"
	"sync"

	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/lachesis-base/gossip/dagordering"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/frenchie-foundation/lachesis-dag-tool/dagreader/internal"
)

type EventsBuffer struct {
	db internal.Db

	events struct {
		info      map[hash.Event]*internal.EventInfo
		processed map[idx.Epoch]map[hash.Event]dag.Event
	}

	ordering *dagordering.EventsBuffer

	output chan *internal.EventInfo
	busy   sync.WaitGroup
	sync.RWMutex

	logger.Instance
}

func NewEventsBuffer(db internal.Db, done <-chan struct{}) *EventsBuffer {
	const count = 3000

	s := &EventsBuffer{
		db:       db,
		output:   make(chan *internal.EventInfo, 10),
		Instance: logger.New("buffer"),
	}

	s.events.processed = make(map[idx.Epoch]map[hash.Event]dag.Event, 3)
	s.events.info = make(map[hash.Event]*internal.EventInfo, count)

	go db.Load(s.output)

	s.ordering = dagordering.New(dag.Metric{
		Num:  count,
		Size: cachescale.Identity.U64(10 * opt.MiB),
	}, dagordering.Callback{
		Process: func(e dag.Event) error {
			id := e.ID()
			epoch := id.Epoch()
			info := s.events.info[id]
			if info == nil {
				panic("event info not found")
			}
			if _, exists := s.events.processed[epoch]; !exists {
				s.events.processed[epoch] = make(map[hash.Event]dag.Event, count)
				delete(s.events.processed, epoch-2)
			}

			s.Log.Debug("completed event", "id", id)
			select {
			case s.output <- info:
				s.events.processed[epoch][id] = e
				delete(s.events.info, id)
			case <-done:
				return fmt.Errorf("Interrupted")
			}

			return nil
		},

		Exists: func(e hash.Event) bool {
			if ee, ok := s.events.processed[e.Epoch()]; ok {
				if _, exists := ee[e]; exists {
					return true
				}
			}

			if len(s.events.processed) < 2 {
				return s.db.HasEvent(e)
			}

			return false
		},

		Get: func(e hash.Event) dag.Event {
			if ee, ok := s.events.processed[e.Epoch()]; ok {
				if event, exists := ee[e]; exists {
					return event
				}
			}

			if len(s.events.processed) < 2 {
				info := s.db.GetEvent(e)
				if info != nil {
					return info.Event
				}
			}

			return nil
		},

		Check: func(e dag.Event, parents dag.Events) error {
			// trust to all
			return nil
		},
	})

	return s
}

func (s *EventsBuffer) Push(e *internal.EventInfo) {
	s.Lock()
	defer s.Unlock()

	s.events.info[e.Event.ID()] = e
	s.ordering.PushEvent(e.Event, "")
}

func (s *EventsBuffer) Close() {
	s.Lock()
	defer s.Unlock()

	close(s.output)
	s.ordering.Clear()
}
