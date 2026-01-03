package storage

import (
	"encoding/json"
	"log"

	"github.com/dgraph-io/badger/v4"
	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
)

const (
	dbPath       = "./data/blockchain"
	blockPrefix  = "block_"
	lastHashKey  = "lh"
)

type BadgerStore struct {
	DB *badger.DB
}

func NewBadgerStore() *BadgerStore {
	opts := badger.DefaultOptions(dbPath).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return &BadgerStore{DB: db}
}

func (bs *BadgerStore) Close() {
	bs.DB.Close()
}

func (bs *BadgerStore) SaveBlock(block *domain.Block) {
	err := bs.DB.Update(func(txn *badger.Txn) error {
		data, err := json.Marshal(block)
		if err != nil {
			return err
		}

		key := append([]byte(blockPrefix), block.Hash...)
		if err := txn.Set(key, data); err != nil {
			return err
		}

		return txn.Set([]byte(lastHashKey), block.Hash)
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (bs *BadgerStore) LoadBlockchain() []*domain.Block {
	var blocks []*domain.Block

	err := bs.DB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()

			if len(key) < len(blockPrefix) || string(key[:len(blockPrefix)]) != blockPrefix {
				continue
			}

			err := item.Value(func(val []byte) error {
				var block domain.Block
				if err := json.Unmarshal(val, &block); err != nil {
					return err
				}
				blocks = append(blocks, &block)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return blocks
}
