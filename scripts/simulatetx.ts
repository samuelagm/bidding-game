// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { ethers, network } from "hardhat";

async function main() {
    const bw = await ethers.getContractFactory("BiddingWar");
    const bidwarInstance = bw.attach("0x712516e61C8B383dF4A63CFe83d7701Bce54B03e")

    //console.log(bidwarInstance)
    //await bidwarInstance.restart()
    const lastBid = await bidwarInstance.getLastBid()
    console.log(lastBid)
    const tx = await (await bidwarInstance.bid({ value: lastBid.winningBid.add(ethers.utils.parseEther("10"))  })).wait()
    console.log(tx)

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
