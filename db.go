package main

import (
	"os"
)

const dbFile = "chain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type DB interface {
	LastHash() []byte
	AddGenesis(address string, genesis *Block) []byte
	AddBlock(newBlock *Block) []byte
	GetBlock(currentHash []byte) *Block
	Close()
}

var d *DB

func NewDB() DB {
	if d != nil {
		return *d
	} else {
		return NewBoltDB()
	}
}

func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}
