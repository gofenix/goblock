package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	db *bolt.DB
}

func NewBoltDB() *BoltDB {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	return &BoltDB{
		db: db,
	}
}

func (d *BoltDB) LastHash() []byte {
	var tip []byte
	d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))
		return nil
	})
	return tip
}

func (d *BoltDB) AddBlock(newBlock *Block) []byte {
	var tip []byte
	d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		tip = newBlock.Hash
		return nil
	})
	return tip
}

func (d *BoltDB) AddGenesis(address string, genesis *Block) []byte {
	var tip []byte

	d.db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesis.Hash

		return nil

	})

	return tip
}

func (d *BoltDB) GetBlock(currentHash []byte) *Block {
	var block *Block

	d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})

	return block
}

func (d *BoltDB) Close() {
	d.db.Close()
}
