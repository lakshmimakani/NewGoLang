package main

import (
	"fmt"
	"sync"
	"time"
)

// Data represents a key-value pair with a TTL
type Data struct {
	Value     string
	ExpiresAt time.Time
}

// KeyValueStore represents the key-value store
type KeyValueStore struct {
	data map[string]Data
	mu   sync.Mutex
}

// NewKeyValueStore initializes a new key-value store
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]Data),
	}
}

// Set adds or updates a key-value pair with a TTL
func (kvs *KeyValueStore) Set(key, value string, ttl time.Duration) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	kvs.data[key] = Data{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// Get retrieves the value associated with a key
func (kvs *KeyValueStore) Get(key string) (string, error) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	data, ok := kvs.data[key]
	if !ok {
		return "", fmt.Errorf("key '%s' not found", key)
	}

	if time.Now().After(data.ExpiresAt) {
		delete(kvs.data, key)
		return "", fmt.Errorf("key '%s' has expired", key)
	}

	return data.Value, nil
}

// Delete removes a key from the key-value store
func (kvs *KeyValueStore) Delete(key string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	delete(kvs.data, key)
}

func main() {
	kvs := NewKeyValueStore()

	// Set key-value pairs with TTL
	kvs.Set("name", "John Doe", time.Second*5)
	kvs.Set("email", "john@example.com", time.Second*10)

	// Retrieve values
	val1, err1 := kvs.Get("name")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Name:", val1)
	}

	val2, err2 := kvs.Get("email")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Email:", val2)
	}

	// Wait for some time to allow values to expire
	time.Sleep(time.Second * 6)

	// Try retrieving expired value
	_, err3 := kvs.Get("name")
	if err3 != nil {
		fmt.Println(err3)
	}

	// Try deleting a key
	kvs.Delete("email")

	// Attempt to retrieve the deleted key
	_, err4 := kvs.Get("email")
	if err4 != nil {
		fmt.Println(err4)
	}
}
