# Crypton Blockchain

Crypton is a simple blockchain implementation in Go, designed to provide a basic understanding of blockchain concepts. It includes functionalities to add blocks, validate proof-of-work, and traverse the blockchain.

## Features

- Add blocks to the blockchain.
- Validate proof-of-work for blocks.
- Serialize and deserialize blocks.
- Iterate through the blockchain.
- Persistent storage using BadgerDB.

## Prerequisites

- [Go](https://golang.org/) (version 1.18 or later).
- [BadgerDB](https://github.com/dgraph-io/badger) for persistent storage.

## Installation

1.  git clone cd crypton
2.  go mod tidy
3.  go run main.go

## Usage

`go run cmd/main.go get-balance --address ""`

`go run cmd/main.go create --address ""`

`go run cmd/main.go print`

`go run cmd/main.go send -from "FROM" -to "TO" -amount AMOUNT`
