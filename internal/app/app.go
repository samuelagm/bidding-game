package app

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env"
	"github.com/ethereum/go-ethereum/crypto"
	contract "gitlab.com/energi/blockchain-challenge/internal/contracts/biddingwar"
	"gitlab.com/energi/blockchain-challenge/internal/logger"
	"gitlab.com/energi/blockchain-challenge/internal/manager"
	"gitlab.com/energi/blockchain-challenge/internal/server"
)

type envconfig struct {
	Host              string        `env:"HOST" envDefault:"https://nodeapi.test.energi.network/v1/jsonrpc"`
	WSS_Host          string        `env:"WSS_HOST" envDefault:"wss://nodeapi.test.energi.network/ws"`
	PrivateKey        string        `env:"PRIVATEKEY" envDefault:"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"`
	ContractAddress   string        `env:"CONTRACT" envDefault:"0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6"`
	GameCheckInterval time.Duration `env:"GAME_CHECK_INTERVAL" envDefault:"5m"`
}

var cfg envconfig
var privateKey *ecdsa.PrivateKey

func init() {
	cfg = envconfig{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	var err error
	privateKey, err = crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	ctx, cancelCTX := context.WithCancel(context.Background())

	biddingWarHelper := contract.NewBiddingWarHelper(
		cfg.Host,
		privateKey,
		cfg.ContractAddress,
	)

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
