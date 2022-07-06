package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env"
	"github.com/jpillora/backoff"
	"gitlab.com/energi/blockchain-challenge/internal/logger"
	blockchaintypes "gitlab.com/energi/blockchain-challenge/internal/logger/types"
	"gitlab.com/energi/blockchain-challenge/internal/server"
)

type envconfig struct {
	Host       string `env:"HOST" envDefault:"wss://nodeapi.test.energi.network/ws"`
	PrivateKey string `env:"PRIVATEKEY"`
}

var cfg envconfig

func init() {

	cfg = envconfig{
		PrivateKey: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

}

func Run() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	ctx, cancelCTX := context.WithCancel(context.Background())
	watchCtx, watchCancel := context.WithCancel(ctx)

	b := &backoff.Backoff{
		Jitter: true,
	}

	const contractAddress = "0x4826533B4897376654Bb4d4AD88B7faFD0C98528"
	var watchErr <-chan error = nil
	var err error = nil

	var fn = func() {
		watchErr, err = logger.Watch(watchCtx, blockchaintypes.Config{
			Host:            cfg.Host,
			ContractFile:    "../assets/BiddingWar.abi",
			ContractAddress: contractAddress,
		})
	}

	fn()

	if watchErr == nil && err != nil {
		fmt.Println("unable to start blockchain watch")
		log.Fatal(watchErr)
	}

	go server.Serve(contractAddress, cfg.PrivateKey)

	for {
		select {
		case watchErr := <-watchErr:
			reconnectIn := b.Duration()
			fmt.Println(watchErr, "reconnecting in ", reconnectIn)
			time.Sleep(reconnectIn)
			fn()
		case signal := <-ch:
			fmt.Printf("\nGot signal:%s, exiting...\n", signal)
			watchCancel()
			cancelCTX()
			time.Sleep(100 * time.Millisecond)
			os.Exit(0)
		}
	}

}
