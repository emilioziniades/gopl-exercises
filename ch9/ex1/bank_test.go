// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdrawal(t *testing.T) {
	done := make(chan struct{})
	var status bool
	//Alice deposits
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()
	//Bob withdraws
	go func() {
		status = Withdraw(100)
		fmt.Println(status)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := status, true; got != want {
		t.Errorf("Withdraw status = %t, want %t", got, want)
	}
	if got, want := Balance(), 400; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

}
