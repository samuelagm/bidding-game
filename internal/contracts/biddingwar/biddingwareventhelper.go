package contract

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BiddingWarEventHelper struct {
	Host             string
	ContractInstance *Contract
	Client           *ethclient.Client
}

func NewBiddingWarEventHelper(wsshost string, address string) *BiddingWarHelper {
	client, err := ethclient.Dial(wsshost)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	c, err := NewContract(common.HexToAddress(address), client)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return &BiddingWarHelper{
		Host:             wsshost,
		ContractInstance: c,
		Client:           client,
	}

}
