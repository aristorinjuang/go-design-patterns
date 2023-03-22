package main

import "fmt"

type PaymentFacade struct {
	acc      *Account
	wallet   *Wallet
	security *Security
}

func (p *PaymentFacade) MakePayment(amount float64) error {
	p.acc.checkBalance(amount)
	p.wallet.debit(amount)
	p.security.verify()
	fmt.Printf("Payment of $%.2f has been made\n", amount)
	return nil
}

type Account struct {
	balance float64
}

func (a *Account) checkBalance(amount float64) error {
	if a.balance < amount {
		return fmt.Errorf("not enough balance")
	}
	return nil
}

type Wallet struct {
	balance float64
}

func (w *Wallet) debit(amount float64) error {
	if w.balance < amount {
		return fmt.Errorf("not enough balance")
	}
	w.balance -= amount
	return nil
}

type Security struct{}

func (s *Security) verify() error {
	fmt.Println("Security verification complete")
	return nil
}

func main() {
	acc := &Account{1000.0}
	wallet := &Wallet{500.0}
	security := &Security{}
	payment := &PaymentFacade{acc, wallet, security}

	err := payment.MakePayment(750.0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
