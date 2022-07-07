# Solution
***
This project contains a smart contract that implements the game described in #1, it also contains the golang service that does the following:
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

### Assumptions
1. Restarting the game and paying winners are mutually exclusive.
2. The game is started automatically on contract deployment.
3. Logging events mean writing smart contract events to stdout/console. 

### Prerequisites

1. Install npm, typescript and go dependencies  
   
   `npm install`  
   `go install`

### ENV
`HOST=https://nodeapi.test.energi.network/v1/jsonrpc`  
`WSS_HOST=wss://nodeapi.test.energi.network/ws`  
`PRIVATEKEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`  
`CONTRACT=0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6`  
`GAME_CHECK_INTERVAL=5m`  

### Tests

   - `npx hardhat test`
   - `PRIVATE_KEY=... go test ./...`

### Execution

`go run main.go`


