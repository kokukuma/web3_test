pragma solidity ^0.8.11;

contract Inbox {
	string public message;
	address payable owner;

    constructor(string memory initialMessage) {
        message = initialMessage;
        owner = payable(msg.sender);
    }

	function setMessage(string memory newMessage) public {
		message = newMessage;
	}
}
