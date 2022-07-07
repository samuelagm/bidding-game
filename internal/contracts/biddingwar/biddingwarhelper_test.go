package contract

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/caarlos0/env"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
var ownerAddress common.Address
var biddingwarContract *Contract
var client *ethclient.Client

func initVars() {
	cfg = envconfig{}
	var err error
	if err = env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	privateKey, err = crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	client, err = ethclient.Dial(cfg.Host)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("error casting public key to ECDSA")
	}

	ownerAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	biddingwarContract, err = NewContract(common.HexToAddress(cfg.ContractAddress), client)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
}

func TestMain(m *testing.M) {
	initVars()
	os.Exit(m.Run())
}

func TestBiddingWarHelper_GetCommissions(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			want:    big.NewInt(0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.GetCommissions()
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.GetCommissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("BiddingWarHelper.GetCommissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiddingWarHelper_GetCurrentRoundNumber(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			want:    big.NewInt(0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.GetCurrentRoundNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.GetCurrentRoundNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) < 0 {
				t.Errorf("BiddingWarHelper.GetCurrentRoundNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiddingWarHelper_GetLastBid(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    BiddingWarBidRound
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			want:    BiddingWarBidRound{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.GetLastBid()
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.GetLastBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, nil) {
				t.Errorf("BiddingWarHelper.GetLastBid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiddingWarHelper_GetBidAt(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	type args struct {
		round int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    BiddingWarBidRound
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			args: args{
				round: 0,
			},
			want:    BiddingWarBidRound{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.GetBidAt(tt.args.round)
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.GetBidAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, nil) {
				t.Errorf("BiddingWarHelper.GetBidAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiddingWarHelper_GameIsRunning(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.GameIsRunning()
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.GameIsRunning() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got).Kind() != reflect.Bool {
				t.Errorf("BiddingWarHelper.GameIsRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiddingWarHelper_Withdraw(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    *types.Transaction
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			want:    &types.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.Withdraw()
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, nil) {
				t.Errorf("BiddingWarHelper.Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiddingWarHelper_RestartGame(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    *types.Transaction
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			want:    &types.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.RestartGame()
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.RestartGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, nil) {
				t.Errorf("BiddingWarHelper.RestartGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiddingWarHelper_PayWinner(t *testing.T) {
	type fields struct {
		Host             string
		PrivateKey       *ecdsa.PrivateKey
		ContractInstance *Contract
		Client           *ethclient.Client
		OwnerAddress     common.Address
	}
	type args struct {
		round int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Transaction
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				Host:             cfg.Host,
				PrivateKey:       privateKey,
				ContractInstance: biddingwarContract,
				Client:           client,
				OwnerAddress:     ownerAddress,
			},
			want:    &types.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BiddingWarHelper{
				Host:             tt.fields.Host,
				PrivateKey:       tt.fields.PrivateKey,
				ContractInstance: tt.fields.ContractInstance,
				Client:           tt.fields.Client,
				OwnerAddress:     tt.fields.OwnerAddress,
			}
			got, err := c.PayWinner(tt.args.round)
			if (err != nil) != tt.wantErr {
				t.Errorf("BiddingWarHelper.PayWinner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, nil) {
				t.Errorf("BiddingWarHelper.PayWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
