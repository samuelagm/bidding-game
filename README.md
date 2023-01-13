# Bidding Game
***
This project contains a smart contract that implements a bidding game, it also contains the golang service that does the following:
1. Serves endpoint for the following routes:
   - `/getcommissions` returns accumulated commissions.
   - `/getcurrentroundnumber` returns the latest bid round number.
   - `/getbidat/:round` returns the winning bid at the specified round.
   - `/getlastbid` returns the very latest bid.
   - `/restartgame` restarts the game if time has elapsed.
   - `/withdraw` allows the application withdraw commissions to it's wallet.
   - `/paywiner/:round` pays the specified bid round winner.
2. Watches for smart contract events and logs them in json format
3. Tracks game state and restarts it when the game time elapses.

### Prerequisites

1. Install npm, typescript and go dependencies  
   
   `npm install`  
   `go install`

### ENV
`HOST=https://nodeapi.test.energi.network/v1/jsonrpc`  
`WSS_HOST=wss://nodeapi.test.energi.network/ws`  
`PRIVATE_KEY=...`  
`CONTRACT=...`  
`GAME_CHECK_INTERVAL=5m`  

### Tests

   - `npx hardhat test`
   - `PRIVATE_KEY=... go test ./...`

### Execution

`go run main.go`


