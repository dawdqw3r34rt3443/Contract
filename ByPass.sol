// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract EthReceiver {
    
    event Received(address sender, uint amount);
    
    constructor(address payable receiver) payable {

        receiver.transfer(msg.value);
        
        emit Received(msg.sender, msg.value);
        
        selfdestruct(payable(address(this)));
    }
}