---
tags:
  - golang-util
---
```go
import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min, max float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := min + r.Intn(max-min+1)
	return n
}
```