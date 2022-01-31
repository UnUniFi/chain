const fs = require("fs");

const data = fs.readFileSync("genesis-pretty.json");
const json = data.toString("utf-8");
const obj = JSON.parse(json);

// remove all accounts info
obj.app_state.auth.accounts.splice(0);
obj.app_state.bank.balances.splice(0);
obj.app_state.genutil.gen_txs.splice(0);
obj.app_state.pricefeed.params.markets.forEach(market => {
  market.oracles.splice(0)
})

fs.writeFileSync("genesis-pretty.json", JSON.stringify(obj, null, "  "));
