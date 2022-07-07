
import { ethers } from "hardhat";

async function main() {
    const bw = await ethers.getContractFactory("BiddingWar");
    const bidwarInstance = bw.attach("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")
    const lastBid = await bidwarInstance.getLastBid()
    const tx = await bidwarInstance.bid(
        { value: lastBid.winningBid.add(ethers.utils.parseEther("1")) }
    )
    const txReceipt = await tx.wait()
    console.log(txReceipt)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
