const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const OracleContract = artifacts.require("OracleContract");

module.exports = async function () {
  let oracleContract = await OracleContract.deployed();
  let address = ["0x24E6d8c3507cDe498F3E8eB1bE042D8A0E827650"];

  await oracleContract.urge(address);

};
