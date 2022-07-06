#!/bin/bash

#Import Smartcontracts into project as go

mkdir ./internal/contracts/BiddingWar &> /dev/null

abigen --abi ../assets/BiddingWar.abi \
--bin ../assets/BiddingWar.bin \
--pkg contract \
--out ../internal/contracts/BiddingWar/biddingwar.go