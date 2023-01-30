# Uniswap V3 Calle

### contracts .env template:
```
# Add wallet private key
Private=

# Add Alchemy or Infura provider keys, alchemy takes preference at the config level
ALCHEMY_KEY=
INFURA_KEY=

# Optional Etherscan key, for automatize the verification of the contracts at Etherscan
ETHERSCAN_KEY=

```

### Deploy Uniswap v3 Callee
* testnet  (goerli)
```
npx hardhat run scripts/deploy-uniswap-v3-callee.ts --network goerli
``` 

* mainnet
```
npx hardhat run scripts/deploy-uniswap-v3-callee.ts --network main
``` 