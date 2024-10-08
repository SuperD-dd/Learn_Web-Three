# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a Hardhat Ignition module that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat ignition deploy ./ignition/modules/Lock.js
```


# 合约部署到测试网的code；

```shell
安装下面两个
npm i @openzeppelin/contracts  
npm i dotenv
编译好部署
npx hardhat ignition deploy ./ignition/modules/RccToken.js --network sepolia
```

