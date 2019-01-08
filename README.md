A simple internal k-v store with the ability to expire a key

## SYNOPSIS
```
type DB struct {
    // contains filtered or unexported fields
}

func New() (db *DB)
    Create a new DB

func (db *DB) Close()
    Unset DB After this call, all methods will be invalid

func (db *DB) Del(k string)
    Delete k from db

func (db *DB) Exists(k string) bool
    Test if k exists in db

func (db *DB) Flush()
    Flush all keys

func (db *DB) Get(k string) interface{}
    Get value from k return nil if the key is not set

func (db *DB) Keys() (keys []string)
    Return all keys *unsorted*

func (db *DB) Set(k string, v interface{}, ttl ...int64)
    Set k with v with an optional ttl in seconds
```
