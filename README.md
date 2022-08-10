# eth-tools

## contract-fetch

```
Usage: contract-fetch <addr> <outDir>

addr	- contract address (required)
outDir	- output directory (optional, default: ./lib/)

Flags:
  ETH_TOOLS_APIKEY (required)
  ETH_TOOLS_URL (optional, default: https://api.etherscan.io)
```

### install

`go install github.com/seblw/eth-tools/cmd/contract-fetch`


### example

```
$ contract-fetch 0xB92336759618F55bd0F8313bd843604592E27bd8
lib/Replica/node_modules/@summa-tx/memview-sol/contracts/TypedMemView.sol
lib/Replica/node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol
lib/Replica/node_modules/@summa-tx/memview-sol/contracts/SafeMath.sol
lib/Replica/node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol
lib/Replica/node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol
lib/Replica/packages/contracts-core/contracts/Replica.sol
lib/Replica/packages/contracts-core/contracts/NomadBase.sol
lib/Replica/packages/contracts-core/contracts/libs/Message.sol
lib/Replica/packages/contracts-core/contracts/libs/Merkle.sol
lib/Replica/node_modules/@openzeppelin/contracts-upgradeable/proxy/Initializable.sol
lib/Replica/packages/contracts-core/contracts/interfaces/IMessageRecipient.sol
lib/Replica/node_modules/@openzeppelin/contracts/cryptography/ECDSA.sol
lib/Replica/packages/contracts-core/contracts/Version0.sol
lib/Replica/packages/contracts-core/contracts/libs/TypeCasts.sol
```
