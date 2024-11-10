
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
... config ...

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
```
account balance               # Check wallet balance and account balance in system
account join <amount>         # Deposit ZAR to the system for executing 'take' actions
account hope <address>        # Check account's hope
account exit <amount>         # Withdraw ZAR from system
account exit <ilk> <amount>   # Withdraw tokens from ilk
```


### auctions

Manage auctions - list auctions, redo, take, etc.
```
auctions clippers                     # List auction clippers
auctions list                         # List all auctions 
auctions redo <ilk> <id>              # Redo specific auction
auctions take <ilk> <id> [--uniswap]  # Take funds from specific auction
```

### vaults

View and list Maker vaults
```
vaults ls               # List default vault set
vaults ls --all         # List all vaults
vaults ls <id>          # List specific vault by ID
```


### run

Run the keeper across different modes
```
run [--uniswap]                     # Runs in default bark, redo and take modes
run --modes bark                    # Run only in bark mode
run --modes redo                    # Run only in redo mode
run --modes take [--uniswap]        # Run only in take mode
run --ilks ethb,daib [--uniswap]    # Run on specific ilks
```
If you set the `uniswap` flag, then for each `take` action, the released collateral from the under-collateralized vault would be swapped to the ZAR in UniswapV3 DEX and the vault debt is paid by the swapped tokens and the leftover will be transferred to your wallet, so, in this mode you only need to pay transaction gas fee and don't need to have any ERC20 tokens in your wallet, otherwise, if you don't set the `uniswap` flag, then the vaults will be liquidated by your personal ZAR funds in system.

## Initializing
At the beginning, execute this command to get the list of all clippers:
```bash
go run main.go auctions clippers --config=/PATH/TO/YOUR/CONFIG
```
After it, you need to give access to the clippers that you want to run `take` action on vaults that are on those ilks. You need to run it once and for all. So, we suggest to give access to each clipper address by running the following command:
```bash
go run main.go account hope [clipper-address-1] [clipper-address-2] ... --config=/PATH/TO/YOUR/CONFIG
```
If you want to run bot in the `take` mode and don't set `uniswap` flag, then you have to deposit some ZAR to the system to use it for repaying debt. To do so, run this command:
```bash
go run main.go account join <amount> --config=/PATH/TO/YOUR/CONFIG
```
If you want to exit your ZAR or other ERC20s from the system, you need to first give access to the `join` contracts by running hope command.
Now, you've done all initializations you need. It's time to run bot in the mode that you want.