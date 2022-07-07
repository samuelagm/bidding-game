import { expect } from "chai";
import { profile } from "console";
import { BigNumber } from "ethers";
import { ethers, waffle, network } from "hardhat";
import { BiddingWar } from "../../typechain";


const getContract = async (): Promise<BiddingWar> => {
    const bw = await ethers.getContractFactory("BiddingWar");
    const dbw = await bw.deploy();
    return dbw;
}

describe("Contract Basic Unit Tests", function () {

    it("should be biddable", async function () {
        const contract = await getContract()
        await contract.deployed()
        await expect(contract.bid({ value: ethers.utils.parseEther("10") }))
            .to.emit(contract, "PotentialWinner")
    });

    it("should be able to restart game", async function () {
        const contract = await getContract()
        await contract.deployed()
        await network.provider.send('evm_increaseTime', [7200]);
        await network.provider.send('evm_mine', []);
        await expect(contract.restart())
            .to.emit(contract, "NewBidRound")
    });

    it("should be able to pay winner", async function () {
        const provider = waffle.provider;
        const signers = await ethers.getSigners();
        const contract = await getContract()
        await contract.deployed()
        await (await contract.bid({ value: ethers.utils.parseEther("10") })).wait()
        await network.provider.send('evm_increaseTime', [7200]);
        await network.provider.send('evm_mine', []);
        await expect(contract.payWinner(0))
            .to.emit(contract, "CashOut")
    });

    it("contract owner should be able to withdraw profit", async function () {
        const provider = waffle.provider;
        const signers = await ethers.getSigners();
        const contract = await getContract()
        await contract.deployed()
        await (await contract.bid({ value: ethers.utils.parseEther("10") })).wait()
        const prevBalance = await provider.getBalance(
            signers[0].address
        )
        const txr2 = await (await contract.withdraw()).wait()
        const txFee2 = txr2.effectiveGasPrice.mul(txr2.cumulativeGasUsed)
        const profit = BigNumber.from(5).div(100).mul(ethers.utils.parseEther("10"))
        const currBalance = await provider.getBalance(
            signers[0].address
        )
        expect(currBalance).to.be.equal(
            prevBalance.add(profit).sub(txFee2)
        )
    });

});


describe("Edge Case Tests", function () {

    it("bid must be greater than zero", async function () {
        const contract = await getContract()
        await contract.deployed()
        await expect(contract.bid({ value: ethers.utils.parseEther("0") }))
        .to.be.revertedWith("invalid bid")
    });

    it("prevent a lower bid", async function () {
        const contract = await getContract()
        await contract.deployed()
        await contract.bid({ value: ethers.utils.parseEther("10") })
        await expect(contract.bid({ value: ethers.utils.parseEther("9") }))
        .to.be.revertedWith("bid lower than current bid")
    });

    it("prevent bid when time elapses", async function () {
        const contract = await getContract()
        await contract.deployed()
        await network.provider.send('evm_increaseTime', [7200]);
        await network.provider.send('evm_mine', []);
        await expect(contract.bid({ value: ethers.utils.parseEther("10") }))
        .to.be.revertedWith("bid round has ended")
    });

    it("prevent restart in the middle of a bidding war", async function () {
        const contract = await getContract()
        await contract.deployed()
        await expect(contract.restart())
        .to.be.revertedWith("a bidding war is ongoing")
    });

    it("prevent payment in the middle of a bidding war", async function () {
        const contract = await getContract()
        await contract.deployed()
        await expect(contract.payWinner(0))
        .to.be.revertedWith("a bidding war is ongoing")
    });
});
