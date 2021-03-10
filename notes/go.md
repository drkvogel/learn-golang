Thu Feb 15 04:56:23 2018
gofmt
```
$ go tool tour
```
or
```
$ go get golang.org/x/tour/gotour
$ gotour
```
```go
package main
import "fmt"

func main() {
	fmt.Println()
	// time.Now()
}
```

A return statement without arguments returns the named return values.
This is known as a "naked" return.

`var` e.g. 

```go
var i int
```

### initialisers

#### var

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```go
var i, j int = 1, 2
var c, python, java = true, false, "no!"

var i, j int = 1, 2
k := 3
c, python, java := true, false, "no!"
```

Go's basic types are

```go
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32, represents a Unicode code point
float32 float64
complex64 complex128
```

Variable declaration block

```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

zero values: `0`, `false`, `""`.

```go
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
```

#### const

```go
const Pi = 3.14
const World = "世界"
fmt.Println("Hello", World)
fmt.Println("Happy", Pi, "Day")

const Truth = true
fmt.Println("Go rules?", Truth)
```

#### for

```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}

sum := 1
for ; sum < 1000; { // The init and post statement are optional
	sum += sum
}

sum := 1
for sum < 1000 { // like `while`
	sum += sum
}

for { // infinite loop
}

```

### if

```go
if x < 0 {
	return sqrt(-x) + "i"
}

if v := math.Pow(x, n); v < lim { // short statement, scoped to if block
	return v
}

```


#### sqrt

```go
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	//z := x
	for i := 0; i < 10; i++ {
		fmt.Println(i, ": z: ", z)
		if z*z == x {
			return z
		}
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	z := 22.0
	fmt.Println("my Sqrt: ", Sqrt(z))
	fmt.Println("math.Sqrt: ", math.Sqrt(z))
}
```

#### switch

```go
fmt.Print("Go runs on ")
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.", os)
}
```

### defer

```go
func test() {
	fmt.Println("world2")
}

func main() {
	//defer fmt.Println("world")
	defer test()

	fmt.Println("hello")
}
```

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

#### pointers

```go
i, j := 42, 2701

p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j
```

#### struct

```go
type Vertex struct {
	X int
	Y int
}

v := Vertex{1, 2}
v.X = 4
fmt.Println(v.X)

p := &v 	// pointer to struct
p.X = 1e9	// dot notation OK
fmt.Println(v)

v1 = Vertex{1, 2}  // has type Vertex
v2 = Vertex{X: 1}  // Y:0 is implicit
v3 = Vertex{}      // X:0 and Y:0
p  = &Vertex{1, 2} // has type *Vertex

```


#### arrays

```go
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1])
fmt.Println(a)

primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes)
```

#### slice

```go
primes := [6]int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4]
fmt.Println(s) // [3 5 7]

names := [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
}
fmt.Println(names)

a := names[0:2]
b := names[1:3]
fmt.Println(a, b)

b[0] = "XXX" // changes underlying array
fmt.Println(a, b)
fmt.Println(names)

// slice literals
q := []int{2, 3, 5, 7, 11, 13}
fmt.Println(q)

r := []bool{true, false, true, true, false, true}
fmt.Println(r)

s := []struct {
	i int
	b bool
}{
	{2, true},
	{3, false},
	{5, true},
	{7, true},
	{11, false},
	{13, true},
}
fmt.Println(s)

// default low and high bounds
s := []int{2, 3, 5, 7, 11, 13}

s = s[1:4]
fmt.Println(s) // [3 5 7]

s = s[:2]
fmt.Println(s) // [3 5]

s = s[1:]
fmt.Println(s) // [5]

// slice length and capacity
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// nil slice
// The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.
var s []int
fmt.Println(s, len(s), cap(s))
if s == nil {
	fmt.Println("nil!")
}


```

#### make

```go
func main() {
    a := make([]int, 5)
    printSlice("a", a)

    b := make([]int, 0, 5)
    printSlice("b", b)

    c := b[:2]
    printSlice("c", c)

    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
```

#### append() to slice

```go
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
```

The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
(To learn more about slices, read the [Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) article.)

#### range

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

#### `_`

The underscore character is used for a variabe that can be assigned to but not read, to avoid a compiler error if a return value from a range or a function is not used.

```go
pow := make([]int, 10)

for i := range pow { // only use index returned from range
	pow[i] = 1 << uint(i) // == 2**i
}

for _, value := range pow { // ignore index returned from range
	fmt.Printf("%d\n", value)
}
```

### exercise: slices

```go
package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := range p {
		p[y] = make([]uint8, dx)
	}
	for y := range p {
		for x := range p[y] {
			//p[x][y] = (uint8(x)+uint8(y))/2
			//p[x][y] = uint8(x) * uint8(y)
			p[x][y] = uint8(math.Pow(float64(x), float64(y)))
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
```

#### map

```go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use.

#### map literal

```go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// If the top-level type is just a type name, you can omit it from the elements of the literal.
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

```

#### mutating maps

```go
m := make(map[string]int)

m["Answer"] = 42
fmt.Println("The value:", m["Answer"])

m["Answer"] = 48
fmt.Println("The value:", m["Answer"])

delete(m, "Answer")
fmt.Println("The value:", m["Answer"])

v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```

Test that a key is present with a two-value assignment:

```go
elem, ok = m[key]
```

If key is in m, ok is true. If not, ok is false.
If key is not in the map, then elem is the zero value for the map's element type.
Note: if elem or ok have not yet been declared you could use a short declaration form:

```go
elem, ok := m[key]
```

```js
// what's this?
function createOperations(){
  var operations = [];
  for (var i = 0; i < 5; i++) {
   operations.push(function(input){
    return input * this.i;
   })
  }
  return operations;
}

var ops = createOperations();
for (var i = 0; i < 5; i++) {            
        alert(ops[i](i));
}
```
