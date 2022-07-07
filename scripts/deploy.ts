import { ethers } from "hardhat";

async function main() {
  const bw = await ethers.getContractFactory("BiddingWar");
  const dbw = await bw.deploy();
  await dbw.deployed();
  console.log("contract address:", dbw.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
