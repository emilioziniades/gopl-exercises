package main

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

type withdrawal struct {
	amount int
	status chan bool
}

var withdrawals = make(chan withdrawal)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	w := withdrawal{
		amount: amount,
		status: make(chan bool),
	}
	withdrawals <- w
	return <-w.status

}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdrawals:
			if w.amount > balance {
				w.status <- false
				continue
			}
			balance -= w.amount
			w.status <- true

		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
