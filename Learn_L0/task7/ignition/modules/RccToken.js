const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { ethers, upgrades} = require("hardhat");

module.exports = buildModule("RccToken", (m) => {
  const rccToken = m.contract("RccToken");

  return { rccToken };
});
