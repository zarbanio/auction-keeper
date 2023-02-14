import {HardhatUserConfig} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {NETWORKS_RPC_URL} from './helper-hardhat-config';

require('dotenv').config()

const Private = process.env.Private || '';
const ETHERSCAN_KEY = process.env.ETHERSCAN_KEY || '';

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
            chainId: 1,
            accounts: [Private]
        },
        goerli: {
            url: NETWORKS_RPC_URL['goerli'],
            chainId: 5,
            accounts: [Private]
        },
    },
};

export default config;
