import {HardhatUserConfig} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {NETWORKS_RPC_URL, NETWORKS_DEFAULT_GAS,} from './helper-hardhat-config';

require('dotenv').config()

const Private = process.env.Private || '';
const ETHERSCAN_KEY = process.env.ETHERSCAN_KEY || '';
const DEFAULT_BLOCK_GAS_LIMIT = 8000000;

const config: HardhatUserConfig = {
    solidity: {
        version: '0.6.12',
        settings: {
            optimizer: {enabled: true, runs: 200},
        },
    },
    typechain: {
        outDir: 'types',
        target: 'ethers-v5',
    },
    etherscan: {
        apiKey: ETHERSCAN_KEY,
    },
    networks: {
        main: {
            url: NETWORKS_RPC_URL['main'],
            // gasPrice: NETWORKS_DEFAULT_GAS['main'],
            chainId: 1,
            accounts: [Private]
        },
        goerli: {
            url: NETWORKS_RPC_URL['goerli'],
            // blockGasLimit: DEFAULT_BLOCK_GAS_LIMIT,
            // gas: DEFAULT_BLOCK_GAS_LIMIT,
            // gasPrice: NETWORKS_DEFAULT_GAS['goerli'],
            chainId: 5,
            accounts: [Private]
        },
    },
};

export default config;
