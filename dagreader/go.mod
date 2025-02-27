module github.com/frenchie-foundation/lachesis-dag-tool/dagreader

go 1.14

require (
	github.com/Fantom-foundation/go-opera v0.0.0-20210820105149-07604c23d28c
	github.com/Fantom-foundation/lachesis-base v0.0.0-20210721130657-54ad3c8a18c1
	github.com/ethereum/go-ethereum v1.9.22
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/neo4j/neo4j-go-driver v1.8.3
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20210305035536-64b5b1c73954
	github.com/urfave/cli v1.22.1
	gopkg.in/urfave/cli.v1 v1.22.1 // indirect
)

replace (
	github.com/Fantom-foundation/lachesis-base => github.com/frenchie-foundation/lachesis-base v0.0.0-20210420092627-c16f01e35562
	github.com/Fantom-foundation/go-opera => github.com/frenchie-foundation/go-opera v0.0.0-20210621102035-55aaa977f8f5
	github.com/ethereum/go-ethereum => github.com/frenchie-foundation/go-ethereum v1.9.7-0.20210531094457-b859cd9c4511
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0
)
