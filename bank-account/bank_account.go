package account

import "sync"

type Account struct{
	safe sync.Mutex
	balance int64
	status bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance:initialDeposit, status: true}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.safe.Lock()
	defer a.safe.Unlock()
	if !a.status {
		return 0, false
	}
	payout = a.balance
	a.balance = 0
	a.status = false
	return payout, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.safe.Lock()
	defer a.safe.Unlock()
	if !a.status {
		return a.balance, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.safe.Lock()
	defer a.safe.Unlock()
	if !a.status {
		return a.balance, false
	}
	if a.balance+amount < 0 {
		return a.balance, false
	}

	a.balance += amount
	return   a.balance, true
}
