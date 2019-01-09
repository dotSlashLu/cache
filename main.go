// A simple internal k-v store with the ability to expire a key
package cache

import (
	"sync"
	"time"
)

type dbEntry struct {
	value interface{}
	timer *time.Timer
}

func (e *dbEntry) getValue() interface{} {
	return e.value
}

// stop timer for faster gc
//	https://golang.org/pkg/time/#After
func (e *dbEntry) stopTimer() {
	if e.timer == nil {
		return
	}
	e.timer.Stop()
}

type Cache struct {
	// an RWMutex will be good since we don't want to block concurrent read
	l  *sync.RWMutex
	db map[string]*dbEntry
}

// Create a new DB
func New() (c *Cache) {
	c = new(Cache)
	c.l = new(sync.RWMutex)
	c.db = make(map[string]*dbEntry)
	return c
}

// Set k with v with an optional ttl in seconds
func (c *Cache) Set(k string, v interface{}, ttl ...int64) {
	c.lock()
	defer c.unlock()
	e := c.getEntry(k)
	if e == nil {
		e = new(dbEntry)
	} else {
		e.stopTimer()
	}
	e.value = v

	if len(ttl) == 1 {
		d := time.Duration(ttl[0]) * time.Second
		e.timer = time.AfterFunc(d, func() {
			c.Del(k)
		})
	}
	c.db[k] = e
}

func (c *Cache) getEntry(k string) *dbEntry {
	return c.db[k]
}

func (c *Cache) lock() {
	c.l.Lock()
}

func (c *Cache) unlock() {
	c.l.Unlock()
}

// Get value from k
// return nil if the key is not set
func (c *Cache) Get(k string) interface{} {
	c.lock()
	defer c.unlock()
	e := c.getEntry(k)
	if e == nil {
		return nil
	}
	return e.getValue()
}

// Delete k from db
func (c *Cache) Del(k string) {
	c.lock()
	defer c.unlock()
	e := c.getEntry(k)
	if e == nil {
		return
	}
	e.stopTimer()
	delete(c.db, k)
}

// Return all keys *unsorted*
func (c *Cache) Keys() (keys []string) {
	db := c.db
	keys = make([]string, 0, len(db))
	for k := range db {
		keys = append(keys, k)
	}
	return keys
}

// Test if k exists in db
func (c *Cache) Exists(k string) bool {
	c.lock()
	defer c.unlock()
	_, exists := c.db[k]
	return exists
}

// Flush all keys
//
// * design considerations:
//	best effort is made to reduce the gc cycle by manually close all the keys
//	that are going to be deleted or replaced but doing a loop or making a
// 	storage for stopping all the timers are not worth the gain so they are
//	spared when flushing
func (c *Cache) Flush() {
	c.db = make(map[string]*dbEntry)
}

// Unset DB
// After this call, all methods will be invalid
func (c *Cache) Close() {
	c.db = nil
}
