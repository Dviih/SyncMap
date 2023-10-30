# SyncMap

### Sometimes it can happen to have a memory race error within Go while using maps, so Synchronous is a solution!

## Install: `go get -u github.com/Dviih/SyncMap`

## Usage
- `Make` - makes a new map.
- `Get`  - gets a key in the map.
- `Add`  - pushes a key to map, if the map is nil it will create.
- `Push` - pushes a key to map.
- `Del`  - deletes a key from a map.

## Example

```go
package main

import (
	"fmt"
	"github.com/Dviih/SyncMap"
)

func main() {
	data := syncmap.SyncMap[string, string]{}

	data.Add("key", "value")
	fmt.Println(data.Get("key"))
	
	data.Del("key")
}
```

---

#### Made by @Dviih
