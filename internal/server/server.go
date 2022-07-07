package server

import (
	"strconv"

	"github.com/gin-gonic/gin"

	contract "gitlab.com/energi/blockchain-challenge/internal/contracts/biddingwar"
)

func Serve(biddingwarContract *contract.BiddingWarHelper) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we're here",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/getcommissions", func(c *gin.Context) {

			amount, err := biddingwarContract.GetCommissions()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"commissions": amount,
			})

		})

		v1.GET("/getcurrentroundnumber", func(c *gin.Context) {
			roundNumber, err := biddingwarContract.GetCurrentRoundNumber()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"roundNumber": roundNumber,
			})

		})

		v1.GET("/getbidat/:round", func(c *gin.Context) {
			round := c.Param("round")

			rountInt, err := strconv.ParseInt(round, 10, 64)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			bidround, err := biddingwarContract.GetBidAt(rountInt)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, bidround)
		})

		v1.GET("/getlastbid", func(c *gin.Context) {
			bidround, err := biddingwarContract.GetLastBid()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, bidround)
		})

		v1.GET("/restartgame", func(c *gin.Context) {

			tx, err := biddingwarContract.RestartGame()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"txHash": tx.Hash().Hex(),
			})

		})

		v1.GET("/withdraw", func(c *gin.Context) {

			tx, err := biddingwarContract.Withdraw()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"txHash": tx.Hash().Hex(),
			})

		})

		v1.GET("/paywiner/:round", func(c *gin.Context) {

			round := c.Param("round")

			rountInt, err := strconv.ParseInt(round, 10, 64)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			tx, err := biddingwarContract.PayWinner(rountInt)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"txHash": tx.Hash().Hex(),
			})

		})
	}

	r.Run()
}
