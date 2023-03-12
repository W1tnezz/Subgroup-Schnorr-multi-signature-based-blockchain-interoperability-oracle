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
  let tx = "0x1256901778428453331bd44b1619e05350b853610968f54da7338e4c331acbe2";
  let fee = await oracleContract.TOTAL_FEE();

  // await web3.eth.subscribe(
  //   "logs",
  //   {
  //     address: OracleContract.address,
  //     topics: [topic],
  //   },
  //   async function (error, result) {
  //     if (!error) {
  //       let receipt = await web3.eth.getTransactionReceipt(
  //         result.transactionHash
  //       );
  //       records.push({ id: counter, gas: receipt.cumulativeGasUsed });

  //       if (counter === 10) {
  //         await csvWriter.writeRecords(records).then(() => {
  //           console.log("...Done");
  //         });
  //         return;
  //       }

  //       counter++;
  //     }
  //   }
  // );

  await oracleContract.validateTransaction(tx, {
    value: fee,
  });

  let count = await oracleContract.countEnrollNodes();
  console.log(count)
};
