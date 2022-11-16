## Problem

c.JSON gin.H{()} outputs empty objects

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	id uint
	todo string
	isCompleted bool
}

func main() {
	todos := []todo{
		{Id: 0, Todo: "Call Billy", IsCompleted: false},
		{Id: 1, Todo: "Push commit to github", IsCompleted: false},
		{Id: 2, Todo: "Clean code", IsCompleted: false},
	}

	r := gin.Default()

	r.GET("/todos", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "Get todos success",
			"data": todos,
		})
	})

	r.Run()
}
```

## Solution

Lowercase fields are considered private, and will not be serialized by the standard json serializer.

Change the fields of your todo type, so that they start with an uppercase letter:

```go
type todo struct {
	Id uint `json:"id"`
	Todo string `json:"todo"`
	IsCompleted bool `json:"is_completed"`
}
```

## Reference

https://stackoverflow.com/questions/59509505/c-json-gin-h-outputs-empty-objects
