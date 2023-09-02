package pokecach

import (
	"testing"
	"time"
)

func TestCreateCache( t *testing.T){
	interval := time.Millisecond * 10
	cach := NewCache(interval)
	if cach.cache == nil {
		t.Error("cache non created")
	}
}

func TestAGCache( t *testing.T){
	interval := time.Millisecond * 10
	cach := NewCache(interval)
	cach.Add("key1", []byte("val1"))
	ere, i := cach.Get("key1")
	if string(ere)!="val1"{
		t.Error("no match")
	}

	if !i{
		t.Error("key !found")
	}
}

func TestReap (t *testing.T) {
	interval := time.Millisecond * 10
	cach := NewCache(interval)
	cach.Add("key", []byte("vval"))
	time.Sleep(interval/2)
	_,er := cach.Get("key")
	if !er {
		t.Error("no reap work")
	}

}