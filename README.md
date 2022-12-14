# eth-tools

eth-tools is a set of smart contract utilities that can be easly composed with each other.

### install all tools in one go
`go install github.com/seblw/eth-tools/cmd/...@latest`

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

`go install github.com/seblw/eth-tools/cmd/contract-fetch@latest`


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

## contract-remappings

accepts list of files as standard input (stdin), prints foundry compatible dependency remappings for @-prefixed entries to standard output (stdout).

### install

`go install github.com/seblw/eth-tools/cmd/contract-remappings@latest`

### example

```
$ contract-fetch 0xB92336759618F55bd0F8313bd843604592E27bd8 | contract-remappings
@summa-tx=lib/Replica/node_modules/@summa-tx
@openzeppelin=lib/Replica/node_modules/@openzeppelin
```

## contract-abi

prints contract's ABI with additional metadata in three first rows.

### install

`go install github.com/seblw/eth-tools/cmd/contract-abi@latest`


### example

```
$ contract-abi 0xB92336759618F55bd0F8313bd843604592E27bd8 | head -n 12
// ContractName: Replica
// CompilerVersion: v0.7.6+commit.7338295f
// LicenseType: 
[
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "_localDomain",
        "type": "uint32"
      }
    ],
```

## contract-interface

accepts contract ABI in JSON format as standard input (stdin), prints solidity interface representation to standard output (stdout).

### install

`go install github.com/seblw/eth-tools/cmd/contract-interface@latest`

### example

```
$ contract-abi 0xB92336759618F55bd0F8313bd843604592E27bd8 | contract-interface
constructor(uint32 _localDomain) public
event NewUpdater(address oldUpdater, address newUpdater)
function process(bytes _message) external nonpayable returns (bool _success)
function prove(bytes32 _leaf, bytes32[32] _proof, uint256 _index) external nonpayable returns (bool)
```