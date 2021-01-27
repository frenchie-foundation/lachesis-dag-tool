module github.com/Fantom-foundation/lachesis-dag-tool

go 1.14

require (
	github.com/Fantom-foundation/go-lachesis v0.8.0-rc.2
	github.com/Fantom-foundation/lachesis-base v0.0.0-20201210130850-782ad52d6c4a
	github.com/deckarep/golang-set v1.7.1
	github.com/ethereum/go-ethereum v1.9.22
	github.com/hashicorp/golang-lru v0.5.4
	github.com/neo4j/neo4j-go-driver v1.8.3
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/stretchr/testify v1.6.1
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca
	github.com/urfave/cli v1.22.1
)

replace (
	github.com/Fantom-foundation/go-lachesis => ../go-lachesis
	github.com/ethereum/go-ethereum => github.com/Fantom-Foundation/go-ethereum v1.9.22-ftm-0.1
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0
)
