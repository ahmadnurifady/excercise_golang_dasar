package soal

import (
	"fmt"
	"runtime"
	"sync"
)

type Account struct {
	Id      string
	Name    string
	Balance int
	sync.Mutex
}

type AccountMethod interface {
	AddBalance(amount int)
	GetBalance(amount int)
}

func (a *Account) AddBalance(amount int) {
	a.Lock()
	a.Balance += amount
	a.Unlock()
}

func (a *Account) GetBalance(amount int) {
	a.Balance -= amount
}

func StartRaceCondition() {
	// mtx := sync.Mutex{}

	runtime.GOMAXPROCS(4)

	var count int
	wg := sync.WaitGroup{}

	account1 := Account{
		Id:      "001",
		Name:    "person 1",
		Balance: 10,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 10; j++ {
				// mtx.Lock()
				account1.AddBalance(1)
				fmt.Println(account1.Balance)
				// count++
				// mtx.Unlock()
			}

			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("var count", count)
}
