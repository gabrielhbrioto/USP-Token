// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@account-abstraction/contracts/interfaces/IPaymaster.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";
import "./IIdentityRegistry.sol";

contract USPPaymaster is IPaymaster, Ownable {
    IIdentityRegistry public identityRegistry;
    IEntryPoint public immutable entryPoint;
    
    mapping(address => uint256) public gasSpentByStudent;
    
    uint256 public maxGasPerStudent = 0.05 ether;

    mapping(address => bool) public whitelistedTargets;
    event GasCharged(address indexed student, uint256 actualGasCost, PostOpMode mode);
    event GasLimitUpdated(uint256 newLimit);

    constructor(IEntryPoint _entryPoint, address _identityRegistry, address _admin) Ownable(_admin) {
        entryPoint = _entryPoint;
        identityRegistry = IIdentityRegistry(_identityRegistry);
    }

    // Função Setter controlada pela Governança (que será o Owner)
    function setMaxGasPerStudent(uint256 _newLimit) external onlyOwner {
        maxGasPerStudent = _newLimit;
        emit GasLimitUpdated(_newLimit);
    }

    // setTargetWhitelist já existe e é onlyOwner, então a Governança já pode usá-la
    function setTargetWhitelist(address target, bool status) external onlyOwner {
        whitelistedTargets[target] = status;
    }

    function validatePaymasterUserOp(
        PackedUserOperation calldata userOp,
        bytes32,
        uint256 maxCost
    ) external view override returns (bytes memory context, uint256 validationData) {
        if (!identityRegistry.isStudentActive(userOp.sender)) return ("", 1);

        address target = _getTargetFromCallData(userOp.callData);
        if (!whitelistedTargets[target]) return ("", 1);

        // Verifica contra a variável maxGasPerStudent
        if (gasSpentByStudent[userOp.sender] > maxGasPerStudent) return ("", 1); 

        if (entryPoint.balanceOf(address(this)) < maxCost) return ("", 1);

        return (abi.encode(userOp.sender), 0);
    }

    // ... (postOp, makeDeposit, withdrawTo, _getTargetFromCallData mantidos iguais) ...
    function postOp(PostOpMode mode, bytes calldata context, uint256 actualGasCost, uint256) external override {
        address student = abi.decode(context, (address));
        gasSpentByStudent[student] += actualGasCost;
        emit GasCharged(student, actualGasCost, mode);
    }

    function makeDeposit() external payable {
        entryPoint.depositTo{value: msg.value}(address(this));
    }

    function withdrawTo(address payable withdrawAddress, uint256 amount) external onlyOwner {
        entryPoint.withdrawTo(withdrawAddress, amount);
    }

    function _getTargetFromCallData(bytes calldata data) internal pure returns (address target) {
        if (data.length < 20) return address(0);
        return address(bytes20(data[0:20])); 
    }
}