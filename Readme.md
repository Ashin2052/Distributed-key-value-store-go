## Distributed key-value-store
## Installation


```bash
Install Go on your machine (https://golang.org/doc/install).
Clone the repository or create a new Go module.
Navigate to the project directory.
```

## To Run the application
```go run main.go```

## To Test the application
```go test -v```

##Usage
The CLI supports the following commands:

put <key> <value>: Stores the given key-value pair in the distributed key-value store.
get <key>: Retrieves the value associated with the given key from the distributed key-value store.
delete <key>: Deletes the key-value pair associated with the given key from the distributed key-value store.
To start the CLI, run the following command:

shell
Copy code
go run main.go
After starting the CLI, you can enter commands:

```
> put name John
Key-value pair stored.

> get name
Value: John

> delete name
Key-value pair deleted.

> get name
Key not found.
```
## Implementation Details

The implementation consists of two main types: Node and DistributedStore.

#### Node
The Node type represents a single node in the distributed key-value store. It contains a data map for storing key-value pairs. The GetValue, SetValue, and DeleteValue methods provide the functionality to retrieve, set, and delete key-value pairs from the node.

#### DistributedStore
The DistributedStore type represents the distributed key-value store. It is responsible for managing multiple nodes and distributing key-value pairs among them. The AddNode method adds a new node to the distributed store. The GetNodeForKey method determines the node responsible for a given key based on a simple hashing algorithm. The Put, Get, and Delete methods interact with the appropriate node to perform the respective operations.

#### Command-Line Interface
The program's main function sets up the CLI. It creates a DistributedStore instance and adds a node to it. It then reads commands from the standard input, parses them, and performs the corresponding operations using the DistributedStore methods. The CLI supports the put, get, and delete commands as described earlier.