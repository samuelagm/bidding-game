package logger

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	eventtypes "gitlab.com/energi/blockchain-challenge/internal/logger/types"
)

func Watch(ctx context.Context, config eventtypes.Config) (ch <-chan error, err error) {

	client, err := ethclient.Dial(config.Host)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	// content, err := ioutil.ReadFile(config.ContractFile)
	// if err != nil {
	// 	return nil, err
	// }
	// contractAbi, err := abi.JSON(strings.NewReader(string(content)))
	// if err != nil {
	// 	return nil, err
	// }
	query := ethereum.FilterQuery{
		FromBlock: nil,
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		fmt.Println("sub error", err)
		return nil, err
	}

	errChan := make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("My watch has ended.")
				client.Close()
				return
			case err := <-sub.Err():
				client.Close()
				errChan <- err
			case vLog := <-logs:
				fmt.Println(vLog)
				if err != nil {
					log.Println("contract event unpacking:", err)
				}

			}
		}

	}()

	return errChan, nil

}
