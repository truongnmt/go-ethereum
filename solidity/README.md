# Solidity
Solidity is a Turing complete programming language for writing smart contracts. Solidity gets compiled to bytecode which is what the Ethereum virtual machine executes.

Some snippet to compile solidity file to opcodes and binary:
```
solc -o . --opcodes example.sol
solc -o . --bin example.sol
```