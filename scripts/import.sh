#!/bin/bash

#Import Smartcontracts into project as golang packages

abigen --abi ../assets/BiddingWar.abi \
--bin ../assets/BiddingWar.bin \
--pkg contract \
--out ../internal/contracts/biddingwar/biddingwar.go