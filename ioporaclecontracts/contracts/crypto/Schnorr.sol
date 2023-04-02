// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./BN256G1.sol";

/**
 * @title Schnorr Signature Verify on BN256 Elliptic Curve
 * @author W1tnezz
 */
library  Schnorr {
  uint256 public constant GX = 1;
  uint256 public constant GY = 2;

  struct ECPoint {
    uint256 x;
    uint256 y;
  }

  function toString(address account) public pure returns (string memory) {
    return toString(abi.encodePacked(account));
  }

  function toString(uint256 value) public pure returns (string memory) {
    return toString(abi.encodePacked(value));
  }

  function toString(bytes32 value) public pure returns (string memory) {
    return toString(abi.encodePacked(value));
  }

  function toString(bytes memory data) public pure returns (string memory) {
    bytes memory alphabet = "0123456789abcdef";
    bytes memory str = new bytes(2 + data.length * 2);
    str[0] = "0";
    str[1] = "x";
    for (uint i = 0; i < data.length; i++) {
      str[2 + i * 2] = alphabet[uint(uint8(data[i] >> 4))];
      str[3 + i * 2] = alphabet[uint(uint8(data[i] & 0x0f))];
    }
    return string(str);
  }

  function verify(uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) internal returns (bool) {
    uint256[3] memory input1 =
    [
    GX,
    GY,
    signature
    ];
    // S1 = sig * G;
    ECPoint memory S1;
    (S1.x, S1.y) = BN256G1.mulPoint(input1);

    uint256[3] memory input2 =
    [
    pubKeyX,
    pubKeyY,
    _hash
    ];

    // S2 = R + _hash * PubKey;
    ECPoint memory S2;
    (S2.x, S2.y) = BN256G1.mulPoint(input2);

    uint256[4] memory input3 =
    [
    RX,
    RY,
    S2.x,
    S2.y
    ];
    (S2.x, S2.y) = BN256G1.addPoint(input3);
    // S1 == S2;
    return S1.x == S2.x && S1.y == S2.y;
  }
}