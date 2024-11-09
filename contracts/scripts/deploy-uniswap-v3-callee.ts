import { ethers } from "hardhat";
import {verifyEtherscanContract} from "./verify";
import {GetConfig} from "../config/get-config";

const hre = require("hardhat");

async function main() {
  let deploymentConfig = GetConfig("uniswapV3Callee", hre.network.name)

  const uniV3Router = deploymentConfig.UniV3Router
  const zarJoin = deploymentConfig.ZarJoin

  const uniswapV3CalleeFactory = await ethers.getContractFactory("UniswapV3Callee");
  const uniswapV3Callee = await uniswapV3CalleeFactory.deploy(uniV3Router, zarJoin);

  await uniswapV3Callee.deployed();

  console.log(`UniswapV3Callee contract deployed to ${uniswapV3Callee.address}`);

  // try to verify
  await verifyEtherscanContract(uniswapV3Callee.address, [uniV3Router, zarJoin])

}


main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
