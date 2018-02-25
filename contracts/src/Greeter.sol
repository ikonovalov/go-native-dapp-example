pragma solidity ^0.4.20;

contract Greeter {
    string public name;
    uint256 public count;

    event _GreetEv(string name, uint256 count);

    function greet(string _name) public {
        name = _name;
        count += 1;
        _GreetEv(_name, count);
    }
}