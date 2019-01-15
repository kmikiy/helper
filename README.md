# helper

helper consists of two packages `ptr` & `labaled`. 

`ptr` makes it easy to create pointers of common built-in types. `labeled` is useful for adding additional information when using built-in types.


## Installation

```sh
go get github.com/kmikiy/helper/...
```

## Example

### ptr

```go
import (
	"log"
	"time"

	"github.com/kmikiy/helper/ptr"
)

type Subscription struct {
	Name         string
	SubscribedAt *time.Time
}

func main() {
    // with "github.com/kmikiy/helper/ptr"
	sub := Subscription{
		Name:         "pro",
		SubscribedAt: ptr.Time(time.Now()),
	}

    log.Printf("sub: %#+v\n", sub)
    
    // without "github.com/kmikiy/helper/ptr"
    now := time.Now() 
    sub := Subscription{
		Name:         "pro",
		SubscribedAt: &now,
	}

    log.Printf("sub: %#+v\n", sub)
}

```



### labeled

```go
import (
	"log"
	"time"

	l "github.com/kmikiy/helper/labeled"
)

func fetch(limit int, offset int) {
	log.Printf("limit: %#+v\n", offset)
	log.Printf("offset: %#+v\n", offset)
}

func main() {
    // with "github.com/kmikiy/helper/labeled"
    fetch(l.Int("limit", 1), l.Int("offset", 20))
    
    // without "github.com/kmikiy/helper/labeled"
	fetch(1, 20)
}

```