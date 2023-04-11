const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const OracleContract = artifacts.require("OracleContract");

module.exports = async function () {
  let dir = "./data";
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir);
  }
  const csvWriter = createCsvWriter({
    path: "./data/bls-cost.csv",
    header: [
      { id: "id", title: "id" },
      { id: "gas", title: "gas" },
    ],
  });
  let records = [];

  let topic;
  let eventName = "SubmitValidationResultLog";
  for (const [key, value] of Object.entries(OracleContract.events)) {
    if (value.name === eventName) {
      topic = key;
    }
  }

  let counter = 0;
  let oracleContract = await OracleContract.deployed();
  let tx = "0xf4bb9a6843194a5f9d10c70202b445733ad8feb5b5f9bce2c699e4cd0abb4bd2";
  let fee = await oracleContract.TOTAL_FEE();

  await oracleContract.validateTransaction(tx, {
    value: fee,
  });

  let count = await oracleContract.countEnrollNodes();
  console.log(count);
};
