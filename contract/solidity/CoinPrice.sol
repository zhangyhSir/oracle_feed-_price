pragma solidity ^0.8.18;

interface IAggregatorV3 {
    function decimals() external view returns (uint8);

    function description() external view returns (string memory);

    function version() external view returns (uint256);

    // getRoundData and latestRoundData should both raise "No data present"
    // if they do not have data to report, instead of returning unset values
    // which could be misinterpreted as actual reported values.
    function getRoundData(uint80 _roundId)
        external
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        );

    function latestRoundData()
        external
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        );
}


contract btcPrice is IAggregatorV3 {

    address public contractOwner;
    int256 public btcCoinPrice;
    mapping(uint80=>int256) idToBTCPrice;
    constructor(){
        contractOwner=msg.sender;
    }

    function setBtcCoinPrice(uint80 id,int256 price) public {
        require(msg.sender==contractOwner,'you have no acquire to change');
        idToBTCPrice[id]=price;
        btcCoinPrice = price;
    }

    function decimals() public view returns (uint8) {
        return 0;
    }

    function description() public view returns (string memory) {
        return "";
    }

    function version() public view returns (uint256) {
        return 0;
    }

    // getRoundData and latestRoundData should both raise "No data present"
    // if they do not have data to report, instead of returning unset values
    // which could be misinterpreted as actual reported values.
    function getRoundData(uint80 id)
        public
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        )
    {
        return (0,idToBTCPrice[id] , 0, 0, 0);
    }

    function latestRoundData()
        public
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        )
    {
        return (0, btcCoinPrice, 0, 0, 0);
    }
}

contract ethPrice is IAggregatorV3 {
    address public contractOwner;
    mapping(uint80=>int256) idToETHPrice; 
    int256  public ethCoinPrice;
      constructor(){
        contractOwner=msg.sender;
    }

    function setEthCoinPrice(uint80 id,int256 price) public {
        require(msg.sender==contractOwner,'you have no acquire to change');
        idToETHPrice[id]=price;
        ethCoinPrice = price;
    }

    function decimals() public view returns (uint8) {
        return 0;
    }

    function description() public view returns (string memory) {
        return "";
    }

    function version() public view returns (uint256) {
        return 0;
    }

    // getRoundData and latestRoundData should both raise "No data present"
    // if they do not have data to report, instead of returning unset values
    // which could be misinterpreted as actual reported values.
    function getRoundData(uint80 id)
        public
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        )
    {
        return (0, idToETHPrice[id], 0, 0, 0);
    }

    function latestRoundData()
        public
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        )
    {
        return (0, ethCoinPrice, 0, 0, 0);
    }
}
