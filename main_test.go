package main

import "testing"

func TestDistributedStore(t *testing.T) {
	store := newDistributedStore()
	store.addNode()

	// Test put and get
	store.put("name", "John")
	value, ok := store.get("name")
	if !ok || value != "John" {
		t.Errorf("Failed to store or retrieve key-value pair: expected value 'John', got '%s'", value)
	}

	// Test delete
	store.delete("name")
	_, ok = store.get("name")
	if ok {
		t.Errorf("Failed to delete key-value pair: key 'name' still exists")
	}

	// Test multiple nodes
	store.addNode()
	store.put("age", "30")
	value, ok = store.get("age")
	if !ok || value != "30" {
		t.Errorf("Failed to store or retrieve key-value pair in a multi-node setup: expected value '30', got '%s'", value)
	}
}

func TestReplicate(t *testing.T) {
	store := newDistributedStore()
	store.addNode()
	store.addNode()
	store.addNode()

	store.put("key1", "value1")
	store.put("key2", "value2")

	store.replicate()

	for _, node := range store.nodes[:len(store.nodes)-1] {
		value, ok := node.getValue("key1")
		if !ok || value != "value1" {
			t.Errorf("Failed to replicate key1 to node %v", node)
		}

		value, ok = node.getValue("key2")
		if !ok || value != "value2" {
			t.Errorf("Failed to replicate key2 to node %v", node)
		}
	}
}
func TestNode(t *testing.T) {
	node := NewNode()

	// Test setValue and getValue
	node.setValue("name", "John")
	value, ok := node.getValue("name")
	if !ok || value != "John" {
		t.Errorf("Failed to store or retrieve key-value pair in node: expected value 'John', got '%s'", value)
	}

	// Test deleteValue
	node.deleteValue("name")
	_, ok = node.getValue("name")
	if ok {
		t.Errorf("Failed to delete key-value pair in node: key 'name' still exists")
	}
}

func TestGetNodeForKey(t *testing.T) {
	store := newDistributedStore()
	store.addNode()

	node := store.getNodeForKey("name")
	if node == nil {
		t.Errorf("Failed to get node for key 'name'")
	}
}
