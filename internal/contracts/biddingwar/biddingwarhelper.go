package contract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BiddingWarHelper struct {
	Host             string
	ContractInstance *Contract
	Client           *ethclient.Client
	ContractAddress  common.Address
	OwnerAddress     common.Address
}

func NewBiddingWarHelper(host string, address string) *BiddingWarHelper {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal("unable to find PRIVATE_KEY environmental variable")
	}

	publicKey := privateKey.Public()
	privateKey = nil
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	c, err := NewContract(common.HexToAddress(address), client)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return &BiddingWarHelper{
		Host:             host,
		ContractInstance: c,
		Client:           client,
		ContractAddress:  common.HexToAddress(address),
		OwnerAddress:     fromAddress,
	}
}

func (c *BiddingWarHelper) GetCommissions() (*big.Int, error) {
	opts := bind.CallOpts{
		Pending: false,
		From:    c.OwnerAddress,
		Context: context.Background(),
	}
	return c.ContractInstance.GetCommissions(&opts)
}

func (c *BiddingWarHelper) GetCurrentRoundNumber() (*big.Int, error) {
	opts := bind.CallOpts{
		Pending: false,
		From:    c.OwnerAddress,
		Context: context.Background(),
	}
	return c.ContractInstance.GetCurrentRoundNumber(&opts)
}

func (c *BiddingWarHelper) GetLastBid() (BiddingWarBidRound, error) {
	opts := bind.CallOpts{
		Pending: false,
		From:    c.OwnerAddress,
		Context: context.Background(),
	}
	return c.ContractInstance.GetLastBid(&opts)
}

func (c *BiddingWarHelper) GetBidAt(round int64) (BiddingWarBidRound, error) {
	opts := bind.CallOpts{
		Pending: false,
		From:    c.OwnerAddress,
		Context: context.Background(),
	}
	return c.ContractInstance.GetBidAtRound(&opts, big.NewInt(round))
}

func (c *BiddingWarHelper) GameIsRunning() (bool, error) {
	opts := bind.CallOpts{
		Pending: false,
		From:    c.OwnerAddress,
		Context: context.Background(),
	}
	return c.ContractInstance.GameIsRunning(&opts)
}

func (c *BiddingWarHelper) Withdraw() (*types.Transaction, error) {
	auth, err := getAuth(c)
	if err != nil {
		return nil, err
	}
	return c.ContractInstance.Withdraw(auth)
}

func (c *BiddingWarHelper) RestartGame() (*types.Transaction, error) {
	auth, err := getAuth(c)
	if err != nil {
		return nil, err
	}
	return c.ContractInstance.Restart(auth)
}

func (c *BiddingWarHelper) PayWinner(round int64) (*types.Transaction, error) {
	auth, err := getAuth(c)
	if err != nil {
		return nil, err
	}
	return c.ContractInstance.PayWinner(auth, big.NewInt(round))
}

func getAuth(c *BiddingWarHelper) (*bind.TransactOpts, error) {
	client := c.Client
	nonce, err := client.PendingNonceAt(context.Background(), c.OwnerAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	privateKey = nil

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(200000)
	auth.GasPrice = gasPrice

	return auth, nil
}
