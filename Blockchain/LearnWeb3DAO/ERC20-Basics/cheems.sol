// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

// import "./openzeppelin-contracts/contracts/token/ERC20/ERC20.sol"
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol";

contract Cheems is ERC20 {
    constructor(string memory _name, string memory _symbol) ERC20(_name, _symbol) {
        _mint(msg.sender, 10 * 10 ** 18);
    }
}

// 0x503FcB8e0CEE41e51E6B96AAbe89f751B3FD354c