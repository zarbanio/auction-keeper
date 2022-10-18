import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("UniswapV2CalleeIRDT", function () {
  async function deployOneYearLockFixture() {
    const [owner, otherAccount] = await ethers.getSigners();

    const UniswapV2CalleeIRDT = await ethers.getContractFactory("UniswapV2CalleeIRDT");
    const uniswapV2CalleeIrdt = await UniswapV2CalleeIRDT.deploy();

    return { uniswapV2CalleeIrdt, owner, otherAccount };
  }

});
