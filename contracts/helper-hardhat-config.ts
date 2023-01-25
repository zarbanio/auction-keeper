require('dotenv').config();

const INFURA_KEY = process.env.INFURA_KEY || '';
const ALCHEMY_KEY = process.env.ALCHEMY_KEY || '';

const GWEI = 1000 * 1000 * 1000;


export const NETWORKS_RPC_URL = {
  ['main']: ALCHEMY_KEY
    ? `https://eth-mainnet.g.alchemy.com/v2/${ALCHEMY_KEY}`
    : `https://mainnet.infura.io/v3/${INFURA_KEY}`,
  ['goerli']: ALCHEMY_KEY
    ? `https://eth-goerli.g.alchemy.com/v2/${ALCHEMY_KEY}`
    : `https://goerli.infura.io/v3/${INFURA_KEY}`,
};

export const NETWORKS_DEFAULT_GAS = {
  ['main']: 65 * GWEI,
  ['goerli']: 24 * GWEI,
};
