import {HardhatUserConfig} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {NETWORKS_RPC_URL} from './helper-hardhat-config';

require('dotenv').config()

const PRIVATE_KEY = process.env.PRIVATE_KEY || '';
const ETHERSCAN_KEY = process.env.ETHERSCAN_KEY || '';

const config: HardhatUserConfig = {
    solidity: {
        version: '0.6.12',
        settings: {
            optimizer: {enabled: true, runs: 200},
            evmVersion: 'istanbul',
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
            chainId: 42161,
            accounts: [PRIVATE_KEY]
        },
        goerli: {
            url: NETWORKS_RPC_URL['goerli'],
            chainId: 5,
            accounts: [PRIVATE_KEY]
        },
    },
};

export default config;
