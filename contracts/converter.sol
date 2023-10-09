// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

library ETHToUSDConverter {
    function getETHPriceInUSD() internal view returns (uint256) {
        // Sepolia ETH / USD Address
        // https://docs.chain.link/data-feeds/price-feeds/addresses
        AggregatorV3Interface priceFeed = AggregatorV3Interface(
            0x694AA1769357215DE4FAC081bf1f309aDC325306
        );

        (, int256 answer, , , ) = priceFeed.latestRoundData();

        if (answer == 0) {
            // Handle the error: return a default value or throw an error
            return 0; // or throw an error
        }

        // ETH/USD rate in 18 digit
        return uint256(answer * 1e10);
    }

    function convertETHToUSD(
        uint256 ethAmount
    ) internal view returns (uint256) {
        uint256 ethPrice = getETHPriceInUSD();
        uint256 ethAmountInUsd = (ethPrice * ethAmount) / 1e18;
        // the actual ETH/USD conversion rate, after adjusting the extra 0s.
        return ethAmountInUsd;
    }

    // function convertUSDToETH(uint256 amount) internal view returns (uint256) {
    //     uint256 ethPrice = getETHPriceInUSD(); // ETH/USD
    //     //USD /ETH
    //     // 1722 * 10**18
    //     uint256 usdToEthRate = (1e18 * 1e18) / ethPrice;
    //     uint256 ethAmount = amount * usdToEthRate;
    //     // the actual ETH/USD conversion rate, after adjusting the extra 0s.
    //     return ethAmount;
    // }
}
