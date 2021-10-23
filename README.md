# in-memory cache

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ayd-ar/cache"
)

func main() {
	cache := cache.New()
	cache.Set("userID", 42, time.Second * 5)

	value, err := cache.Get("userID")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

	time.Sleep(time.Second * 6)

	value, err = cache.Get("userID")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
```