package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env"
	contract "gitlab.com/energi/blockchain-challenge/internal/contracts/biddingwar"
	"gitlab.com/energi/blockchain-challenge/internal/logger"
	"gitlab.com/energi/blockchain-challenge/internal/manager"
	"gitlab.com/energi/blockchain-challenge/internal/server"
)

type envconfig struct {
	Host              string        `env:"HOST" envDefault:"https://nodeapi.test.energi.network/v1/jsonrpc"`
	WSS_Host          string        `env:"WSS_HOST" envDefault:"wss://nodeapi.test.energi.network/ws"`
	ContractAddress   string        `env:"CONTRACT" envDefault:"0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6"`
	GameCheckInterval time.Duration `env:"GAME_CHECK_INTERVAL" envDefault:"5m"`
}

var cfg envconfig

func init() {
	cfg = envconfig{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func Run() {
	ch := make(chan os.Signal, 1)
	ctx, cancelCTX := context.WithCancel(context.Background())
	biddingWarHelper := contract.NewBiddingWarHelper(
		cfg.Host,
		cfg.ContractAddress,
	)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go logger.LogEvents(ctx, cfg.WSS_Host, cfg.ContractAddress)
	go server.Serve(biddingWarHelper)
	go manager.Manage(cfg.GameCheckInterval, biddingWarHelper)

	for {
		signal := <-ch
		fmt.Printf("\nGot signal:%s, exiting...\n", signal)
		cancelCTX()
		os.Exit(0)
	}
}
