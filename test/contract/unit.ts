import { expect } from "chai";
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

    it("should be able to pay winner", async function () {
        const provider = waffle.provider;
        const signers = await ethers.getSigners();
        const contract = await getContract()
        await contract.deployed()
        // const prevBalance = await provider.getBalance(
        //     signers[0].address
        //   )
        const tx = await contract.bid({ value: ethers.utils.parseEther("10") })
        const rcp = await tx.wait()

        await network.provider.send('evm_increaseTime', [7200]);
        await network.provider.send('evm_mine', []);


        // const rx = await provider.send("evm_increaseTime", [3600])
        // await provider.send("evm_mine", [])
        // console.log(rx)

        // const rcp2 = await (await contract.withdraw()).wait()
        // console.log(rcp2)

        // const currBalance = await provider.getBalance(
        //     signers[0].address
        //   )

        // console.log("prev balance:", ethers.utils.formatUnits(prevBalance))
        // console.log("cur balance:", ethers.utils.formatUnits(currBalance))

        await expect(contract.payWinner(0))
            .to.emit(contract, "CashOut")
    });

    it("contract owner should be able to withdraw funds", async function () {
        const provider = waffle.provider;
        const signers = await ethers.getSigners();
        const contract = await getContract()
        await contract.deployed()

        const prevBalance = await provider.getBalance(
            signers[0].address
        )

        const txr1 = await (await contract.bid({ value: ethers.utils.parseEther("10") })).wait()
        const txFee1 = txr1.effectiveGasPrice.mul(txr1.cumulativeGasUsed)
        const txr2 = await (await contract.withdraw()).wait()
        const txFee2 = txr2.effectiveGasPrice.mul(txr2.cumulativeGasUsed)
        const currBalance = await provider.getBalance(
            signers[0].address
        )

        expect(currBalance).to.be.equal(
            prevBalance.sub(txFee1).sub(txFee2)
        )


        // const currBalance = await provider.getBalance(
        //     signers[0].address
        //   )

        // console.log("prev balance:", ethers.utils.formatUnits(prevBalance))
        // console.log("cur balance:", ethers.utils.formatUnits(currBalance))


    });



});
