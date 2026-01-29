// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IIdentityRegistry {
    function isStudentActive(address studentAddress) external view returns (bool);
}
