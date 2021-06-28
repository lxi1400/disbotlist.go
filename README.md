# disbotlist.go
## Wrapper for https://disbots.xyz

# Example
```golang

package main
import (
	"fmt"
	"github.com/lxi1400/disbotlist.go"
)


func main() {
	DisBotListSession :=  disbotlist.DisBotListClient{Token: "TokenHere"}
	result := DisBotListSession.HasVoted("UserID")
	fmt.Println(result)
}
```
