const fs = require("fs");

const data = fs.readFileSync("genesis-pretty.json");
const json = data.toString("utf-8");
const obj = JSON.parse(json);
const unprettified = JSON.stringify(obj);
fs.writeFileSync("genesis.json", unprettified);
