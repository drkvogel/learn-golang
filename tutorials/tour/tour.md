
[A Tour of Go ](https://tour.golang.org/basics/1)
[2009-11-10 23:00:00](https://www.google.com/search?q=2009-11-10+23%3A00%3A00&ie=UTF-8)

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
  rand.Seed(123)
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

factored (parenthesized) import statement:
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

### Exported names

>In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package. pizza and pi do not start with a capital letter, so they are not exported. When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Go's Declaration Syntax - The Go Blog (https://blog.golang.org/declaration-syntax)
>Newcomers to Go wonder why the declaration syntax is different from the tradition established in the C family.
