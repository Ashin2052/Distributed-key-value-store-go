package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// Node represents a single node in the distributed key-value store
type Node struct {
	data map[string]string
}

// DistributedStore represents a distributed key-value store consisting of multiple nodes
type DistributedStore struct {
	nodes []*Node
	mutex sync.RWMutex
}

// NewNode creates a new Node instance
func NewNode() *Node {
	return &Node{
		data: make(map[string]string),
	}
}

// getValue retrieves the value for a given key from the node
func (n *Node) getValue(key string) (string, bool) {
	value, ok := n.data[key]
	return value, ok
}

// setValue sets a key-value pair in the node
func (n *Node) setValue(key, value string) {
	n.data[key] = value
}

// deleteValue deletes a key-value pair from the node
func (n *Node) deleteValue(key string) {
	delete(n.data, key)
}

// newDistributedStore creates a new DistributedStore instance
func newDistributedStore() *DistributedStore {
	return &DistributedStore{
		nodes: []*Node{},
	}
}

// addNode adds a new node to the distributed store
func (d *DistributedStore) addNode() {
	node := NewNode()
	d.nodes = append(d.nodes, node)
}

// getNodeForKey returns the node responsible for a given key
func (d *DistributedStore) getNodeForKey(key string) *Node {
	nodeIndex := len(key) % len(d.nodes)
	return d.nodes[nodeIndex]
}

// put stores a key-value pair in the distributed key-value store
func (d *DistributedStore) put(key, value string) {
	node := d.getNodeForKey(key)
	node.setValue(key, value)
}

// get retrieves the value for a given key from the distributed key-value store
func (d *DistributedStore) get(key string) (string, bool) {
	node := d.getNodeForKey(key)
	return node.getValue(key)
}

// delete removes a key-value pair from the distributed key-value store
func (d *DistributedStore) delete(key string) {
	node := d.getNodeForKey(key)
	node.deleteValue(key)
}

func (d *DistributedStore) replicate() {
	if len(d.nodes) < 2 {
		return
	}

	lastNode := d.nodes[len(d.nodes)-1]
	for key, value := range lastNode.data {
		for _, node := range d.nodes[:len(d.nodes)-1] {
			node.data[key] = value
		}
	}
}

func main() {
	store := newDistributedStore()
	store.addNode() // Add a node to the distributed store
	store.addNode()

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
		if len(parts) < 2 && parts[0] != "replicate" {
			fmt.Println("Invalid command. Usage: <put|get|delete> <key> [<value>]")
			continue
		}

		switch parts[0] {
		case "put":
			if len(parts) != 3 {
				fmt.Println("Invalid command. Usage: put <key> <value>")
				continue
			}
			store.put(parts[1], parts[2])
			fmt.Println("Key-value pair stored.")
		case "get":
			if len(parts) != 2 {
				fmt.Println("Invalid command. Usage: get <key>")
				continue
			}
			value, ok := store.get(parts[1])
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
			store.delete(parts[1])
			fmt.Println("Key-value pair deleted.")
		case "replicate":
			store.replicate()
			fmt.Println("Node Replicated.")
		default:
			fmt.Println("Invalid command. Usage: <put|get|delete> <key> [<value>]")
		}
	}
}
