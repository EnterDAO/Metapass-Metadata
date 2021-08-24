package dlt

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func ConnectToEthereum() *EthereumClient {

	nodeURL := os.Getenv("NODE_URL")

	client, err := NewEthereumClient(nodeURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Successfully connected to ethereum client")

	return client
}
