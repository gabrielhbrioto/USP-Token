// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../script/DeployUSP.s.sol";

contract DeployUSPTest is Test {
    function test_Run_DeploysAndConfiguresContracts() public {
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 deployerPrivateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.addr(deployerPrivateKey);

        vm.deal(deployer, 20 ether);
        vm.setEnv("PRIVATE_KEY", mnemonic);

        DeployUSP script = new DeployUSP();
        script.run();
    }
}