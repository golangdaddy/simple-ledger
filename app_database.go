package main

import (
	"encoding/json"
	//
	"github.com/dgraph-io/badger"
)

// Update creates or updates the specified key
func (app *App) Update(k string, v []byte) error {

	// Start a writable transaction.
	txn := app.badgerDB.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	err := txn.Set([]byte(k), v)
	if err != nil {
		return err
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(nil); err != nil {
		return err
	}

	return nil
}

// Get returns a specified key from the database
func (app *App) Get(k string, dst interface{}) ([]byte, error) {

	// Start a writable transaction.
	txn := app.badgerDB.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	item, err := txn.Get([]byte(k))
	if err != nil {
		return nil, err
	}

	var value []byte
	if err := item.Value(
		func(v []byte) {
			value = append([]byte{}, v...)
		},
	); err != nil {
		return nil, err
	}

	if dst != nil {
		if err := json.Unmarshal(value, dst); err != nil {
			return nil, err
		}
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(nil); err != nil {
		return nil, err
	}

	return value, nil
}

// PrefixIterate will return results based of keys containing the supplied prefix
func (app *App) PrefixIterate(prefix string, process func (k string, v []byte) error) error {

	p := []byte(prefix)

	return app.badgerDB.View(
		func(txn *badger.Txn) error {
			it := txn.NewIterator(badger.DefaultIteratorOptions)
			defer it.Close()
			for it.Seek(p); it.ValidForPrefix(p); it.Next() {

				item := it.Item()
				k := item.Key()
				if err := item.Value(
					func(v []byte) {
						process(string(k), v)
					},
				); err != nil {
					return err
				}
			}
			return nil
		},
	)
}
