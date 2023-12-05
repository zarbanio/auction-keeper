
# Auction kepper

A brief description of what this project does


## Installation

clone project

```bash
  git clone git@github.com:zarbanio/auction-keeper.git
  cd auction-keeper
```
download realse of this lib depends on what system you use and extract it in the root project directory.
  - [ X - library](https://github.com/zarbanio/auction-keeper/x/releases)
for example
```bash
  tar -xvf eh-gen-v0.3.5-darwin-amd64.tar.gz
  cp -r ./eh-gen ./auction-keeper
```

then 
```bash
make
go Install
go mod tidy
```
after make file need to config project.
###### here is example file for goerli ethereum testnet network just create file and name it to config.yaml and copy and paste configs below.

```yaml

Indexer:
  BlockInterval: '10s'
  PoolSize: 2
  BlockPtr: 8485879
Redis:
  URL: 'auction-keeper-redis:6379'
postgres:
  Host: 'localhost:5432'
  User: 'postgres'
  Password: 'postgres'
  DB: 'postgres'
  MigrationsPath: './store/migrations'
Network:
  ChainId: 5
  NativeAsset:
    Name: 'Ethereum'
    Symbol: 'ETH'
    Decimals: 18
    MockAddress: '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
  Node:
    Api: 'https://rpc.ankr.com/eth_goerli'
Wallet:
  Private: ''
  Address: ''
Times:
  VaultTicker: 60
  FlopperTicker: 60
  LiquidatorTicker: 60
ZarJoin: '0xe752ba573bb3ee4840ad424f9c386f7696084e6a'
Vat: '0xe33f43032db170bac2903115e3afb6db981ac822'
Vow: '0x24020386d0162ee2d940111de81a89d0868a0b78'
Dog: '0xc61157c26cdcb5f88cd8d81d61c9fd97023c0cd0'
Flopper: '0x370699470ad9f3b80d42c81e981ee9b2b5ed1801'
UniswapV3QuoterAddress: '0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6'
Collaterals:
  -
    Name: 'ETHA'
    Erc20addr: '0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6'
    Decimals: 18
    Clipper: '0xE3FDbF840476c8A0A1cd327Ad5DBcd77Dd94fEb5'
    GemJoinAdapter: '0x77C0154cF96039519e5BEffaDcfe6eD3c6354b6d'
    UniswapV3Callee: '0x040bC072eE59551aFc67DDF3a265c42e762728fe'
    UniswapV3Path:
      - # WETH -> DAI
        Fee: 10000
        TokenA: '0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6'
        TokenB: '0x7df1d674021792630e15eca46467516660e021af'
      - # DAI -> Zar
        Fee: 10000
        TokenA: '0x7df1d674021792630e15eca46467516660e021af'
        TokenB: '0x91f53a3c281cd82531ea540ebfbb9a64e13ebc8f'
  -
    Name: 'ETHB'
    Erc20addr: '0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6'
    Decimals: 18
    Clipper: '0x5bfD87c9212738834B800FaCb58e073cb3E475c0'
    GemJoinAdapter: '0x361eE781A33560BD8649EDf77ec65f49eB9C4542'
    UniswapV3Callee: '0x040bC072eE59551aFc67DDF3a265c42e762728fe'
    UniswapV3Path:
      - # WETH -> DAI
        Fee: 10000
        TokenA: '0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6'
        TokenB: '0x7df1d674021792630e15eca46467516660e021af'
      - # DAI -> ZAR
        Fee: 10000
        TokenA: '0x7df1d674021792630e15eca46467516660e021af'
        TokenB: '0x91f53a3c281cd82531ea540ebfbb9a64e13ebc8f'
Processor:
  MinProfitPercentage: 1.025
  MinLotZarValue: 0
  MaxLotZarValue: 1000000

```
now we need to replace wallet.Private and wallet.Address with our wallet information.

```yaml
Wallet:
  Private: 'your private key'
  Address: 'your public key'
```
##### then 
```bash 
go build
```

## Commands

### account
Manage account - check balance, join pool, etc.

account balance         # Check account ETH balance
account join <amount>   # Join pool with specified amount 
account hope <address>  # Check account's hope balance

### auctions

Manage auctions - list auctions, redo, take, etc.

auctions clippers       # List auction clippers
auctions list           # List all auctions 
auctions redo <ilk> <id> # Redo specific auction
auctions take <ilk> <id>  # Take funds from specific auction

### vaults

View and list Maker vaults

vaults ls               # List default vault set
vaults ls --all         # List all vaults
vaults ls <id>          # List specific vault by ID

### run

Run the keeper across different modes

run                     # Runs in default bark, redo and take modes
run --modes bark        # Run only in bark mode
run --modes redo        # Run only in redo mode
run --ilks ethb,daib    # Run on specific ilks