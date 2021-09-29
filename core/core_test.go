package core

import "testing"

func TestPut(t *testing.T) {
	const key = "chave1"
	const value = "valor1"

	myStore := NewStorage()

	defer myStore.Delete(key)
	// testing put
	err := myStore.Put(key, value)
	if err != nil {
		t.Error(err)
	}
	// checking if the key is in the map
	val, contains := myStore.m[key]
	if !contains {
		t.Error("key not found")
	}
	// checking if the value is the same
	if val != value {
		t.Error("value not found")
	}

}

func TestGet(t *testing.T) {
	const key = "chave1"
	const value = "valor1"

	myStore := NewStorage()

	defer myStore.Delete(key)
	err := myStore.Put(key, value)
	if err != nil {
		t.Error(err)
	}

	val, err := myStore.Get(key)
	if err != nil {
		t.Error(err)
	}

	if val != value {
		t.Error("value not found")
	}

}

func TestDelete(t *testing.T) {
	const key = "chave1"
	const value = "valor1"

	myStore := NewStorage()

	err := myStore.Put(key, value)
	if err != nil {
		t.Error(err)
	}

	err = myStore.Delete(key)
	if err != nil {
		t.Error(err)
	}

	val, err := myStore.Get(key)
	if err == nil {
		t.Error(err)
	}

	if val != "" {
		t.Error("value not found")
	}

}
