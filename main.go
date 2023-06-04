package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Node represents a single node in the distributed key-value store
type Node struct {
	data map[string]string
}

// NewNode creates a new Node instance
func NewNode() *Node {
	return &Node{
		data: make(map[string]string),
	}
}

// GetValue retrieves the value for a given key from the node
func (n *Node) GetValue(key string) (string, bool) {
	value, ok := n.data[key]
	return value, ok
}

// SetValue sets a key-value pair in the node
func (n *Node) SetValue(key, value string) {
	n.data[key] = value
}

// DeleteValue deletes a key-value pair from the node
func (n *Node) DeleteValue(key string) {
	delete(n.data, key)
}

// DistributedStore represents a distributed key-value store consisting of multiple nodes
type DistributedStore struct {
	nodes []*Node
}

// NewDistributedStore creates a new DistributedStore instance
func NewDistributedStore() *DistributedStore {
	return &DistributedStore{
		nodes: []*Node{},
	}
}

// AddNode adds a new node to the distributed store
func (d *DistributedStore) AddNode() {
	node := NewNode()
	d.nodes = append(d.nodes, node)
}

// GetNodeForKey returns the node responsible for a given key
func (d *DistributedStore) GetNodeForKey(key string) *Node {
	nodeIndex := len(key) % len(d.nodes)
	return d.nodes[nodeIndex]
}

// Put stores a key-value pair in the distributed key-value store
func (d *DistributedStore) Put(key, value string) {
	node := d.GetNodeForKey(key)
	node.SetValue(key, value)
}

// Get retrieves the value for a given key from the distributed key-value store
func (d *DistributedStore) Get(key string) (string, bool) {
	node := d.GetNodeForKey(key)
	return node.GetValue(key)
}

// Delete removes a key-value pair from the distributed key-value store
func (d *DistributedStore) Delete(key string) {
	node := d.GetNodeForKey(key)
	node.DeleteValue(key)
}

func main() {
	store := NewDistributedStore()
	store.AddNode() // Add a node to the distributed store

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		command = strings.TrimSpace(command)
		parts := strings.SplitN(command, " ", 3)
		if len(parts) < 2 {
			fmt.Println("Invalid command. Usage: <put|get|delete> <key> [<value>]")
			continue
		}

		switch parts[0] {
		case "put":
			if len(parts) != 3 {
				fmt.Println("Invalid command. Usage: put <key> <value>")
				continue
			}
			store.Put(parts[1], parts[2])
			fmt.Println("Key-value pair stored.")
		case "get":
			if len(parts) != 2 {
				fmt.Println("Invalid command. Usage: get <key>")
				continue
			}
			value, ok := store.Get(parts[1])
			if ok {
				fmt.Printf("Value: %s\n", value)
			} else {
				fmt.Println("Key not found.")
			}
		case "delete":
			if len(parts) != 2 {
				fmt.Println("Invalid command. Usage: delete <key>")
				continue
			}
			store.Delete(parts[1])
			fmt.Println("Key-value pair deleted.")
		default:
			fmt.Println("Invalid command. Usage: <put|get|delete> <key> [<value>]")
		}
	}
}
