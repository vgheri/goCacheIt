# goCacheIt
Basic in memory key-value store accessible by a HTTP RESTful interface, written in Go.

##Modules
1. The module responsible for managing the key-value data structure. Main responsibilities include:
  1. Creation of the data structure at start-up
  2. Add a key-value couple
  3. Convert the JSON input to be stored as BSON
  4. Retrieve a value by key
  5. Convert the stored BSON to JSON
2. The module managing the interface with external clients through a  RESTful HTTP API which provides the following methods
  1. POST /api/store/   ⇒ MIME type application/json
  2. body { “key”: stringValue, “value”: { … } }
  3. GET /api/store/?key=my-key ⇒ accepting data type application/json . Result is just the value part {...}

##Implementation details

###Key-value store data structure
goCacheIt uses a Splay-tree to store keys and values in memory.
The advantage of using a Splay-tree is that frequently “managed” nodes are much quicker to retrieve than less ones, making it perfect for a cache system similar to how CPU cache work.
This should also allow for a faster LRU cache eviction policy, as least recently used nodes would obviously be at the bottom of the tree and we could therefore skip many checks in between the root and the leafs.

###Keys
Keys are of type string and can have a max length of 255 chars.

###Values
The value object is binary encoded as bson to keep spatial overhead to a minimum and consume as few memory resources as possible.

###Data structure details

```go
type Tree struct {
  root *Node  
}
```

```go
type Node struct {
  *parent, *left, *right Node,
  key string,
  value byte[]
}
```
