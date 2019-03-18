pragma solidity ^0.5;

contract example1 {

  address contractOwner;

  function example() public {
    contractOwner = msg.sender;
  }
}