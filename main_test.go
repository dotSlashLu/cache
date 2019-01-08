package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	k := "test"
	v := "asdf"
	db := New()
	fmt.Println("set k", k, " = ", v, "with a ttl of 1")
	db.Set(k, v, 1)

	desc := "got value should equal to set value"
	fmt.Println(desc)
	gotV := db.Get(k)
	assert.Equal(t, gotV, v, desc)

	desc = "got value should equal to nil after expiration"
	time.Sleep(2 * time.Second)
	gotV = db.Get(k)
	assert.Equal(t, nil, gotV, desc)
}

func TestGet(t *testing.T) {
	db := New()
	v := db.Get("asdf")
	desc := "get non-exist key should return nil"
	fmt.Println(desc)
	assert.Equal(t, nil, v, desc)
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
