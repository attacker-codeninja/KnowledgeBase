// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

contract MoodDiary{
    string mood;

    function setMood(string memory _mood) public{
        mood = _mood;
    }

    function getMood() public view returns(string memory){
        return mood;
    } 
}
 // 0xe39A51DB6aeB8B45e72f51f5E65c7df55981aa5a
