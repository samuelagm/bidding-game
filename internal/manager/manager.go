package manager

import (
	"fmt"
	"time"

	contract "gitlab.com/energi/blockchain-challenge/internal/contracts/biddingwar"
)

//Helps to restart the bidding war game after it ends.
func Manage(interval time.Duration, biddingWarHelper *contract.BiddingWarHelper) {

	for range time.Tick(interval) {
		running, _ := biddingWarHelper.GameIsRunning()
		if !running {
			_, err := biddingWarHelper.RestartGame()
			if err != nil {
				fmt.Println("Error restarting game:", err)
			}
		}
	}
}
