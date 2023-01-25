import { ethers } from "hardhat";

const hre = require("hardhat");

async function main() {
  await hre.run('compile');

  const uniV3Router = "0xE592427A0AEce92De3Edee1F18E0157C05861564" //goerli and mainnet
  const zarJoin = "0x0216C83b8C7984AF8c10d95A656a003DfF8D2266" //TODO: this is SimJoin on goerli (not ZarJoin)

  const UniswapV3Callee = await ethers.getContractFactory("UniswapV3Callee");
  const uniswapV3Callee = await UniswapV3Callee.deploy(uniV3Router, zarJoin);

  await uniswapV3Callee.deployed();

  console.log(`UniswapV3Callee contract deployed to ${uniswapV3Callee.address}`);

  // verify price oracle
  console.log('----- Verifying Contract ------')
  await hre.run("verify:verify", {
    address: uniswapV3Callee.address,
    constructorArguments: [uniV3Router, zarJoin],
  }).catch((error) => {
    console.log("error in verifying: ", error)
  });
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
