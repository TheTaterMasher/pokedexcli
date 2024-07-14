package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s doesn't match %s", actual, cas.inputVal)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 100

	cache := NewCache(interval)

	key1 := "key"
	cache.Add(key1, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key1)
	if ok {
		t.Errorf("%s should ave been reaped", key1)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 100

	cache := NewCache(interval)

	key1 := "key"
	cache.Add(key1, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(key1)
	if !ok {
		t.Errorf("%s shouldn't ave been reaped", key1)
	}
}
