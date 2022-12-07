package main

import "time"

type Cache interface {
	// Get cached value by key.
	Get(key string) interface{}
	// Set cached value with key and expire time.
	Put(key string, val interface{}, timeout time.Duration) error
	// Delete cached value by key.
	Delete(key string) error
	// Increase cached int value by key, as a counter.
	Incr(key string) error
	// Decrease cached int value by key, as a counter.
	Decr(key string) error
	// Check if cached value exists or not.
	IsExist(key string) bool
	// Clear all cache.
	ClearAll() error
	// Start gc routine based on config string settings.
	StartAndGC(config string) error
}
