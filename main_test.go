package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	assert := assert.New(t)
	k := "test"
	v := "asdf"
	db := New()
	fmt.Println("set k", k, " = ", v, "with a ttl of 1")
	db.Set(k, v, 1)

	k1 := "nottl"
	v1 := "nottlv"
	fmt.Println("set k", k1, " = ", v1, "with no ttl")
	db.Set(k1, v1)

	desc := "got value should equal to set value"
	fmt.Println(desc)
	gotV := db.Get(k)
	assert.Equal(gotV, v, desc)

	desc = "got value should equal to nil after expiration"
	fmt.Println(desc)
	time.Sleep(2 * time.Second)
	gotV = db.Get(k)
	assert.Equal(nil, gotV, desc)

	desc = "got value of key with no ttl should equal to set value"
	fmt.Println(desc)
	gotV = db.Get(k1)
	assert.Equal(v1, gotV, desc)

	desc = "reset key should reset it's value and timer"
	fmt.Println(desc)
	k2 := "k2"
	v2 := "v2"
	desc = "	got value should be the same with set"
	fmt.Println(desc)
	db.Set(k2, v2, 2)
	ch := time.After(1 * time.Second)
	<-ch
	assert.Equal(v2, db.Get(k2), desc)
	desc = "	got value should be the same with " +
		"the one newly set"
	fmt.Println(desc)
	v2New := "v2new"
	db.Set(k2, v2New, 3)
	ch = time.After(2 * time.Second)
	<-ch
	assert.Equal(v2New, db.Get(k2), desc)
	desc = "	got value should be nil after expiration"
	ch = time.After(1*time.Second + 10*time.Microsecond)
	// ch = time.After(1 * time.Second)
	<-ch
	assert.Equal(nil, db.Get(k2), desc)
}

func TestGet(t *testing.T) {
	db := New()
	v := db.Get("asdf")

	desc := "get non-exist key should return nil"
	fmt.Println(desc)
	assert.Equal(t, nil, v, desc)
}

func TestDel(t *testing.T) {
	db := New()
	desc := "del non-existing key should not panic"
	db.Del("asdf")
	fmt.Println(desc)
}

func TestKeys(t *testing.T) {
	db := New()
	keys := []string{"a", "b", "c", "d"}
	for v, k := range keys {
		db.Set(k, v)
	}
	dbKeys := db.Keys()

	desc := "Keys() should equal to set keys"
	fmt.Println(desc)
	sort.Strings(keys)
	sort.Strings(dbKeys)
	assert.Equal(t, keys, dbKeys, desc)
}

func TestExists(t *testing.T) {
	db := New()
	db.Set("exists", "exists")
	desc := "Exists() be true for existing keys"
	fmt.Println(desc)
	assert.Equal(t, true, db.Exists("exists"), desc)

	desc = "Exists() should be false for non-existing keys"
	fmt.Println(desc)
	assert.Equal(t, false, db.Exists("non-existing"), desc)
}

func TestFlush(t *testing.T) {
	db := New()
	keys := []string{"e", "f", "g", "h"}
	for v, k := range keys {
		db.Set(k, v)
	}
	desc := "Keys() should be empty after flush"
	fmt.Println(desc)
	db.Flush()
	assert.Equal(t, []string{}, db.Keys(), desc)
}

func ExampleCache_Set() {
	k := "a"
	v := "b"
	ttl := 10
	c := New()
	// set cache with a ttl
	c.Set(k, v, ttl)

	// set cache without a ttl
	c.Set("k1", "v1")
}
