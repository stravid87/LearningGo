const fs = require('fs');
const crypto = require("crypto").webcrypto;
globalThis.crypto = crypto;
require('./golang/wasm/wasm_exec.js');

function loadWebAssembly() {
  const wasmModule = fs.readFileSync('./golang/wasm/main.wasm');
  const go = new Go();
  const importObject = go.importObject;
  WebAssembly.instantiate(wasmModule, importObject).then((results) => {
    const instance = results.instance
    go.run(instance);
  });
}

loadWebAssembly();

// Your formatNumber() function
function formatNumber(q, b) {
  return addTwoNumbers(q, b);
}

const formattedNumber = formatNumber(2, 3);
console.log(`Formatted Number: ${formattedNumber}`);

module.exports = formatNumber