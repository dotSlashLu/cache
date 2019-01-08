// A simple internal k-v store with the ability to expire a key
package cache

import (
	"sync"
	"time"
)

type DB struct {
	lock *sync.Mutex
	db   map[string]interface{}
}

// Create a new DB
func New() (db *DB) {
	db = new(DB)
	db.lock = new(sync.Mutex)
	db.db = make(map[string]interface{})
	return db
}

// Delete k from db
func (db *DB) Del(k string) {
	delete(db.db, k)
}

// Set k with v with an optional ttl in seconds
func (db *DB) Set(k string, v interface{}, ttl ...int64) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.db[k] = v
	if len(ttl) == 1 {
		c := time.Tick(time.Duration(ttl[0]) * time.Second)
		go func() {
			<-c
			db.Del(k)
		}()
	}
}

// Get value from k
// return nil if the key is not set
func (db *DB) Get(k string) interface{} {
	v := db.db[k]
	return v
}

// Return all keys *unsorted*
func (db *DB) Keys() (keys []string) {
	keys = make([]string, 0, len(db.db))
	for k := range db.db {
		keys = append(keys, k)
	}
	return keys
}

// Test if k exists in db
func (db *DB) Exists(k string) bool {
	_, exists := db.db[k]
	return exists
}

// Flush all keys
func (db *DB) Flush() {
	db.db = make(map[string]interface{})
}

// Unset DB
// After this call, all methods will be invalid
func (db *DB) Close() {
	db.db = nil
}
