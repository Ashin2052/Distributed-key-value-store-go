package main

import "testing"

func TestDistributedStore(t *testing.T) {
	store := NewDistributedStore()
	store.AddNode()

	// Test Put and Get
	store.Put("name", "John")
	value, ok := store.Get("name")
	if !ok || value != "John" {
		t.Errorf("Failed to store or retrieve key-value pair: expected value 'John', got '%s'", value)
	}

	// Test Delete
	store.Delete("name")
	_, ok = store.Get("name")
	if ok {
		t.Errorf("Failed to delete key-value pair: key 'name' still exists")
	}

	// Test multiple nodes
	store.AddNode()
	store.Put("age", "30")
	value, ok = store.Get("age")
	if !ok || value != "30" {
		t.Errorf("Failed to store or retrieve key-value pair in a multi-node setup: expected value '30', got '%s'", value)
	}
}

func TestNode(t *testing.T) {
	node := NewNode()

	// Test SetValue and GetValue
	node.SetValue("name", "John")
	value, ok := node.GetValue("name")
	if !ok || value != "John" {
		t.Errorf("Failed to store or retrieve key-value pair in node: expected value 'John', got '%s'", value)
	}

	// Test DeleteValue
	node.DeleteValue("name")
	_, ok = node.GetValue("name")
	if ok {
		t.Errorf("Failed to delete key-value pair in node: key 'name' still exists")
	}
}

func TestGetNodeForKey(t *testing.T) {
	store := NewDistributedStore()
	store.AddNode()

	node := store.GetNodeForKey("name")
	if node == nil {
		t.Errorf("Failed to get node for key 'name'")
	}
}
