#!/bin/sh

cd "$(dirname "$0")" || exit 1


baseDir="../../ioporaclecontracts"

solc --optimize --abi $baseDir/contracts/RegistryContract.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts --include-path $baseDir/node_modules
solc --optimize --abi $baseDir/contracts/OracleContract.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts --include-path $baseDir/node_modules

./abigen --abi $baseDir/build/contracts/abi/RegistryContract.abi --pkg iop --type RegistryContract --out ../pkg/iop/registrycontract.go
./abigen --abi $baseDir/build/contracts/abi/OracleContract.abi --pkg iop --type OracleContract --out ../pkg/iop/oraclecontract.go
