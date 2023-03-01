# Chi: A super simple goroutine pool
## Introduction
Chi is a super simple goroutine pool.  
It can be used to control the maximum number of coroutines running at the same time.
## Install
```sh
go get github.com/wyyadd/chi
```
## Usage
Let's see how to use this library.
```go
package main

import (
	"github.com/wyyadd/chi"
	"log"
	"os"
)

func main() {
	// Create a logger, which is mainly used to log err when panic happens in the pool.
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "===Info===", log.Ldate|log.Ltime|log.Lshortfile)

	// Create a simple function
	simpleFunc := func(i ...interface{}) {
		x := i[0].(int)
		y := i[1].(int)
		logger.Printf("x:%d, y:%d, sum:%d", x, y, x+y)
	}

	// Create a NewPool with max 3 workers running at the same time
	pool := chi.NewPool(3, logger, simpleFunc)

	// start running 
	for i := 0; i < 10; i++ {
		pool.Process(i+30, i-30)
	}

	// wait all workers finish their jobs
	pool.Wait()
}
```
## Note
As I say, Chi is a super simple goroutine pool. For more complex goroutine pool, you can refer to:  
https://github.com/panjf2000/ants  
https://github.com/Jeffail/tunny
