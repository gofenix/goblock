package main

type BlockchainIterator struct {
	currentHash []byte
	db          DB
}

func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.lastHash, bc.db}
	return bci
}

func (i *BlockchainIterator) Next() *Block {
	block := i.db.GetBlock(i.currentHash)
	i.currentHash = block.PrevBlockHash
	return block
}
