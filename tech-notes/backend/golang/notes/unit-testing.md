## Best practices

### Table Driven Testing

### Test Suite

### Use Interfaces And Avoid file I/O

Now we have a function that reads player’s data from a file and prints it:

```go
package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "os"
)

type Player struct {
   Name   string `json:"name"`
   Age    int    `json:"Age"`
}

func main() {
   processData("data.txt")
}

func processData(file string) error {
   f, err := os.Open(file)
   if err != nil {
      return err
   }
   defer f.Close()
   data, err := ioutil.ReadAll(f)
   if err != nil {
      return err
   }
   var players []Player

   err = json.Unmarshal(data, &players)
   if err != nil {
      return err
   }

   for _, p := range players {
      fmt.Println("Player name: ", p.Name)
   }
   return nil
}
```

What’s wrong with it?

Not easily testable, because you need to make sure that the data.txt file exists and contains data before running the tests
What if there is a large amount of data and the server that performs the tests has some lower grade hardware?
You also don’t want to run into race conditions where concurrent tests use/modify the same file
You want your tests to be fast, independent, isolated, consistent and not flaky.

A better approach would be this:

```go
package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "os"
   "io"
)

type Player struct {
   Name   string `json:"name"`
   Age    int    `json:"age"`
}

func main() {
   processData("data.txt")
}

func processData(file string) error {
   f, err := os.Open(file)
   if err != nil {
      return err
   }
   defer f.Close()
   return unmarshalAndPrint(f)
}

func unmarshalAndPrint(f io.Reader) error {
   data, err := ioutil.ReadAll(f)
   if err != nil {
      return err
   }
   var players []Player

   err = json.Unmarshal(data, &players)
   if err != nil {
      return err
   }

   for _, p := range players {
      fmt.Println("Player name: ", p.Name)
   }
   return nil
}
```

Test file:

```go
package main

import (
   "github.com/stretchr/testify/assert"
   "strings"
   "testing"
)
```

```go
func TestUnmarshalAndPrint(t *testing.T) {
   t.Run("testing unmarshalAndPrint()", func(t *testing.T) {
      err := unmarshalAndPrint(strings.NewReader(`[{"name": "John", "age": 900}]`))
      assert.Nil(t, err)
   })
}
```

Now for testing, instead of preparing data and opening a file, we just pass a literal JSON string to strings.NewReader

However, if for some reason you still have to interact with a filesystem, consider using the following resources:
AFERO: A filesystem abstraction system for Go
fstest — testing implementations and users of file systems
TempFile, TempDir — create a temporary file/directory
