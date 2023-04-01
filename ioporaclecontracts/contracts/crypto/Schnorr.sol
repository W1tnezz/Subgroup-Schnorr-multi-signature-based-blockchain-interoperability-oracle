// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./EllipticCurve.sol";

/**
 * @title Schnorr Signature Verify on Secp256k1 Elliptic Curve
 * @author W1tnezz
 */
library  Schnorr {
  uint256 public constant GX = 0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798;
  uint256 public constant GY = 0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8;
  uint256 public constant AA = 0;
  uint256 public constant BB = 7;
  uint256 public constant PP = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F;

  struct ECPoint {
    uint256 x;
    uint256 y;
  }

  function verify(uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) internal pure returns (bool) {
    require(EllipticCurve.isOnCurve(pubKeyX, pubKeyY, AA, BB, PP), "Invalid public key!");
    require(EllipticCurve.isOnCurve(RX, RY, AA, BB, PP), "Invalid R!");

    // S1 = sig * G;
    ECPoint memory S1;
    (S1.x, S1.y) = EllipticCurve.ecMul(signature, GX, GY, AA, PP);

    // S2 = R + _hash * PubKey;
    ECPoint memory S2;
    (S2.x, S2.y) = EllipticCurve.ecMul(_hash, pubKeyX, pubKeyY, AA, PP);
    (S2.x, S2.y) = EllipticCurve.ecAdd(RX, RY, S2.x, S2.y, AA, PP);
    // S1 == S2;

    return S1.x == S2.x;
  }
}