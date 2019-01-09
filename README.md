A simple thread-safe internal k-v store with the ability to expire a key.

[![GoDoc](https://godoc.org/github.com/dotSlashLu/cache?status.svg)](https://godoc.org/github.com/dotSlashLu/cache)

## EXAMPLES
```go
c := New()
// set cache with a ttl
c.Set("k", 1, 10)
		                     
// set cache with no ttl
c.Set("k1", "v1")
```

## SYNOPSIS
```
package cache
    import "github.com/dotSlashLu/cache"

    A simple internal k-v store with the ability to expire a key

TYPES

type Cache struct {
    // contains filtered or unexported fields
}

func New() (c *Cache)
    Create a new DB

func (c *Cache) Close()
    Unset DB After this call, all methods will be invalid

func (c *Cache) Del(k string)
    Delete k from db

func (c *Cache) Exists(k string) bool
    Test if k exists in db

func (c *Cache) Flush()
    Flush all keys 
    
        * design considerations:

        best effort is made to reduce the gc cycle by manually close all the keys
        that are going to be deleted or replaced but doing a loop or making a
        storage for stopping all the timers are not worth the gain so they are
        spared when flushing

func (c *Cache) Get(k string) interface{}
    Get value from k return nil if the key is not set

func (c *Cache) Keys() (keys []string)
    Return all keys *unsorted*

func (c *Cache) Set(k string, v interface{}, ttl ...int64)
    Set k with v with an optional ttl in seconds
```
