```go
package main

import (
	"fmt"
	"time"

	p "github.com/devMYC/go-promise"
)

func slowDouble(n int) int {
	time.Sleep(500 * time.Millisecond)
	return 2 * n
}

func main() {
	start := time.Now()

	x := slowDouble(1)
	y := slowDouble(2)
	z := slowDouble(3)

	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
	fmt.Println(time.Since(start))

	start = time.Now()

	px := p.NewPromise(func(resolve chan<- int, reject chan<- error) {
		resolve <- slowDouble(4)
	})
	py := p.NewPromise(func(resolve chan<- int, reject chan<- error) {
		resolve <- slowDouble(5)
	})
	pz := p.NewPromise(func(resolve chan<- int, reject chan<- error) {
		resolve <- slowDouble(6)
	})

	x, _ = px.Await()
	y, _ = py.Await()
	z, _ = pz.Await()

	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
	fmt.Println(time.Since(start))
}
```

```
x = 2, y = 4, z = 6
1.502222092s
x = 8, y = 10, z = 12
501.108038ms
```
