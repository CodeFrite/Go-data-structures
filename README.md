# Motivation

After completing the Consensys Blockchain Developer Bootcamp, I wanted to dive deeper and try to understand how the EVM actually manages the different data that it stores in its state, more precisely:

- How smart contracts code are stored into the state
- How the transactions composing a block are stored
- How accounts (POAs & smart contracts) are stored

In order to do so, I have decided to dive into the GO-Ethereum implementation. Since Go is a new language to me and since I was expecting the GO-Ethereum implementation to be top edge, I decided to begin by first recoding all the basic data structures that we traditionnaly use.

Each data structure code comes along with test cases. This will allow me to:
- Test my initial code version 
- Allow me to spot any regression in the future if I decide to change/improve it

# Covered data structures

## Heaps

In this implementation, I have used an array to represent the heap. In order to support both, min & max heaps, I have 

## Binary Trees
