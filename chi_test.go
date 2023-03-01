package chi

import (
	"log"
	"os"
	"testing"
)

func simplePanicFunc(i ...interface{}) {
	panic(i)
}

func TestNewPool(t *testing.T) {
	pool := NewPool(5, nil, simplePanicFunc)
	if pool == nil {
		t.Errorf("pool is nil")
	}

	pool = NewPool(0, nil, simplePanicFunc)
	if pool != nil {
		t.Errorf("pool is not nil")
	}

	pool = NewPool(-5, nil, simplePanicFunc)
	if pool != nil {
		t.Errorf("pool is not nil")
	}
	t.Log("Test Done")
}

func TestFuncPanic(t *testing.T) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "===Error===", log.Ldate|log.Ltime|log.Lshortfile)
	pool := NewPool(1, logger, simplePanicFunc)
	for i := 0; i < 2; i++ {
		pool.Process(1, 2, 3)
	}
	pool.Wait()
	t.Log("Test Done")
}

func TestPool(t *testing.T) {
	pool := NewPool(3, nil, func(i ...interface{}) {
		x := i[0].(int)
		y := i[1].(int)
		t.Logf("x:%d, y:%d, sum:%d", x, y, x+y)
	})
	for i := 0; i < 50; i++ {
		pool.Process(i, i-30)
	}
	pool.Wait()
	t.Log("Test Done")
}
