package contract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {

}

type BiddingWarHelper struct {
	Host             string
	PrivateKey       *ecdsa.PrivateKey
	ContractInstance *Contract
	Client           *ethclient.Client
	OwnerAddress     common.Address
}

func NewBiddingWarHelper(host string, privateKey *ecdsa.PrivateKey, address string) *BiddingWarHelper {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	publicKey := privateKey.Public()
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
		PrivateKey:       privateKey,
		ContractInstance: c,
		Client:           client,
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

	estimatedGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &c.OwnerAddress,
		Data: []byte{0},
	})
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

	auth, err := bind.NewKeyedTransactorWithChainID(c.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(estimatedGas * 2)
	auth.GasPrice = gasPrice

	return auth, nil
}
