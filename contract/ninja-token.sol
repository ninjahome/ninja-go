//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.4.24;

import "./ERC20.sol";
import "./owned.sol";

contract NinjaToken is ERC20, owned{

    string  public constant  name = "Ninja Chat Coin";
    string  public constant  symbol = "NCC";
    uint8   public constant  decimals = 18;
    uint256 public constant INITIAL_SUPPLY = 4.2e8 * (10 ** uint256(decimals));

    constructor() {
        _mint(msg.sender, INITIAL_SUPPLY);
    }
}



