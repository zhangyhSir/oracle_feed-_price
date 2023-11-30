// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract getBtcPrice {
    AggregatorV3Interface internal dataFeedBtc;

    constructor() {
        dataFeedBtc = AggregatorV3Interface(
            0xA39434A63A52E749F02807ae27335515BA4b07F7
        );
    }

    function getLatestData() public view returns (uint80,int256) {
        (
            uint80 roundID,
            int256 answer,
            /*uint startedAt*/,
            /*uint timeStamp*/,
            /*uint80 answeredInRound*/
        ) = dataFeedBtc.latestRoundData();
        return (roundID,answer);
    }
}
contract getEthPrice{
     AggregatorV3Interface internal dataFeedEth;
      constructor() {
        dataFeedEth = AggregatorV3Interface(
            0xD4a33860578De61DBAbDc8BFdb98FD742fA7028e
        );
    }

    function getLatestData() public view returns (uint80,int256) {
        (
            uint80 roundID,
            int256 answer,
            /*uint startedAt*/,
            /*uint timeStamp*/,
            /*uint80 answeredInRound*/
        ) = dataFeedEth.latestRoundData();
        return (roundID,answer);
    }
}
