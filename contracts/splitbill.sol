// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {ETHToUSDConverter} from "./converter.sol";

error NoPayers();
error NotOwner();
error IncorrectAmount();
error WithdrawFailed();
error BillAlreadyPayed();

contract SplitBill {
    using ETHToUSDConverter for uint256;

    event Payed(
        address from,
        uint256 amount
    );

    event Completed(
        uint256 amount
    );

    address public immutable i_owner;
    uint256 public immutable i_totalAmountInUSD;
    uint256 public immutable i_amountPerPayerInUSD;

    uint256 public s_amountPayedInUSD;

    mapping (address payer => bool hasPayed) public payerStatus;

    constructor(address[] memory payers, uint256 amountInUSD){
        if(payers.length == 0) { revert NoPayers(); }        

        i_owner = msg.sender;
        i_totalAmountInUSD = amountInUSD;
        i_amountPerPayerInUSD = amountInUSD/payers.length;

        for (uint i=0; i < payers.length; i++) 
        {
            payerStatus[payers[i]] = false;
        }
    }

    function pay() public payable {
        if (msg.value.convertETHToUSD() < i_amountPerPayerInUSD * 1e18) { revert IncorrectAmount();}
        payerStatus[msg.sender] = true;    
        s_amountPayedInUSD += i_amountPerPayerInUSD;


        if (s_amountPayedInUSD == i_totalAmountInUSD) {
           (bool callSuccess, ) = payable(i_owner).call{value: address(this).balance}("");
           if (callSuccess == false) {revert WithdrawFailed(); }
           emit Completed(i_totalAmountInUSD);
        }

        emit Payed(msg.sender, i_amountPerPayerInUSD);
    }
   
    function withdraw() public payable onlyOwner {
         if(address(this).balance == 0) { revert BillAlreadyPayed(); }     

        (bool callSuccess, ) = payable(msg.sender).call{value: address(this).balance}("");

        if (callSuccess == false) {revert WithdrawFailed(); }
    }

     modifier onlyOwner {
        if (msg.sender != i_owner) revert NotOwner();
        _;
    }
    
    fallback() external payable {
        pay();
    }

    receive() external payable {
        pay();
    }
}