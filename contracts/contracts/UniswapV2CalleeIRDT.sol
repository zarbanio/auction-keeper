// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.6.12;

interface VatLike {
    function hope(address) external;
}

interface GemJoinLike {
    function dec() external view returns (uint256);

    function gem() external view returns (TokenLike);

    function exit(address, uint256) external;
}

interface IrdtJoinLike {
    function irdt() external view returns (TokenLike);

    function vat() external view returns (VatLike);

    function join(address, uint256) external;
}

interface TokenLike {
    function approve(address, uint256) external;

    function transfer(address, uint256) external;

    function balanceOf(address) external view returns (uint256);
}

interface CharterManagerLike {
    function exit(address crop, address usr, uint256 val) external;
}

interface UniswapV2Router02Like {
    function swapExactTokensForTokens(uint256, uint256, address[] calldata, address, uint256) external returns (uint[] memory);
}

// Simple Callee Example to interact with MatchingMarket
// This Callee contract exists as a standalone contract
contract UniswapV2Callee {
    UniswapV2Router02Like   public uniRouter02;
    IrdtJoinLike             public irdtJoin;
    TokenLike               public irdt;

    uint256                 public constant RAY = 10 ** 27;

    function add(uint x, uint y) internal pure returns (uint z) {
        require((z = x + y) >= x, "ds-math-add-overflow");
    }

    function sub(uint x, uint y) internal pure returns (uint z) {
        require((z = x - y) <= x, "ds-math-sub-underflow");
    }

    function divup(uint256 x, uint256 y) internal pure returns (uint256 z) {
        z = add(x, sub(y, 1)) / y;
    }

    function setUp(address uniRouter02_, address irdtJoin_) internal {
        uniRouter02 = UniswapV2Router02Like(uniRouter02_);
        irdtJoin = IrdtJoinLike(irdtJoin_);
        irdt = irdtJoin.irdt();

        irdt.approve(irdtJoin_, uint256(- 1));
    }

    function _fromWad(address gemJoin, uint256 wad) internal view returns (uint256 amt) {
        amt = wad / 10 ** (sub(18, GemJoinLike(gemJoin).dec()));
    }
}

// Uniswapv2Router02 route directs swaps from one pool to another
contract UniswapV2CalleeIRDT is UniswapV2Callee {
    constructor(address uniRouter02_, address irdtJoin_) public {
        setUp(uniRouter02_, irdtJoin_);
    }

    function clipperCall(
        address sender, // Clipper Caller and Dai deliveryaddress
        uint256 daiAmt, // Dai amount to payback[rad]
        uint256 gemAmt, // Gem amount received [wad]
        bytes calldata data     // Extra data needed (gemJoin)
    ) external {
        (
        address to,            // address to send remaining IRDT to
        address gemJoin,       // gemJoin adapter address
        uint256 minProfit,     // minimum profit in IRDT to make [wad]
        address[] memory path, // Uniswap pool path
        address charterManager // pass address(0) if no manager
        ) = abi.decode(data, (address, address, uint256, address[], address));

        // Convert gem amount to token precision
        gemAmt = _fromWad(gemJoin, gemAmt);

        // Exit collateral to token version
        if (charterManager != address(0)) {
            CharterManagerLike(charterManager).exit(gemJoin, address(this), gemAmt);
        } else {
            GemJoinLike(gemJoin).exit(address(this), gemAmt);
        }

        // Approve uniRouter02 to take gem
        TokenLike gem = GemJoinLike(gemJoin).gem();
        gem.approve(address(uniRouter02), gemAmt);

        // Calculate amount of IRDT to Join (as erc20 WAD value)
        uint256 irdtToJoin = divup(daiAmt, RAY);

        // Do operation and get irdt amount bought (checking the profit is achieved)
        uniRouter02.swapExactTokensForTokens(
            gemAmt,
            add(irdtToJoin, minProfit),
            path,
            address(this),
            block.timestamp
        );

        // Although Uniswap will accept all gems, this check is a sanity check, just in case
        // Transfer any lingering gem to specified address
        if (gem.balanceOf(address(this)) > 0) {
            gem.transfer(to, gem.balanceOf(address(this)));
        }

        // Convert DAI bought to internal vat value of the msg.sender of Clipper.take
        irdtJoin.join(sender, irdtToJoin);

        // Transfer remaining DAI to specified address
        irdt.transfer(to, irdt.balanceOf(address(this)));
    }
}