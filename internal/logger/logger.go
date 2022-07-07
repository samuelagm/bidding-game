package logger

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"
	contract "gitlab.com/energi/blockchain-challenge/internal/contracts/biddingwar"
)

func Watch(ctx context.Context,
	biddingwarContract *contract.BiddingWarHelper) (ch <-chan error, err error) {
	client := biddingwarContract.Client
	fmt.Println("watching events..")
	watchOpt := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}

	cashoutChannel := make(chan *contract.ContractCashOut)
	newRoundChannel := make(chan *contract.ContractNewBidRound)
	newBidChannel := make(chan *contract.ContractNewBid)

	cashoutSub, err := biddingwarContract.ContractInstance.WatchCashOut(
		watchOpt, cashoutChannel)
	if err != nil {
		return nil, err
	}
	newroundSub, err := biddingwarContract.ContractInstance.WatchNewBidRound(
		watchOpt, newRoundChannel)
	if err != nil {
		return nil, err
	}
	newbidSub, err := biddingwarContract.ContractInstance.WatchNewBid(
		watchOpt, newBidChannel)
	if err != nil {
		return nil, err
	}

	errChan := make(chan error)
	go func() {
		for {
			select {
			case cashout := <-cashoutChannel:
				log.Info().Str("event", "cashout").
					Str("winner", cashout.Winner.Hex()).
					Str("round", cashout.Round.String()).
					Str("amount", cashout.Amount.String()).Send()

			case newround := <-newRoundChannel:
				log.Info().Str("event", "new_round").
					Str("round", newround.Round.String()).Send()

			case newbid := <-newBidChannel:
				log.Info().Str("event", "new_bid").
					Str("bidder", newbid.Bidder.Hex()).
					Str("round", newbid.Round.String()).
					Str("amount", newbid.BidAmount.String()).
					Str("timeleft", newbid.TimeLeft.String()).Send()

			case err := <-cashoutSub.Err():
				client.Close()
				errChan <- err

			case err := <-newroundSub.Err():
				client.Close()
				errChan <- err

			case err := <-newbidSub.Err():
				client.Close()
				errChan <- err

			case <-ctx.Done():
				client.Close()
				return
			}
		}
	}()

	return errChan, nil

}
