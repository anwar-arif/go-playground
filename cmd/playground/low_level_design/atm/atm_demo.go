package atm

import "fmt"

func Run() {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atmMachine := NewATM(cashDispenser, bankingService)

	// create account
	firstAccount, _ := bankingService.CreateAccount("first123", 1000)
	secondAccount, _ := bankingService.CreateAccount("second123", 500)

	card := NewCard(firstAccount.AccountNumber, "1234")
	if err := atmMachine.AuthenticateUser(card); err != nil {
		fmt.Printf("authentication failed %s\n", err.Error())
		return
	}

	// check balance
	if balance, err := atmMachine.CheckBalance(card.CardNumber); err != nil {
		fmt.Printf("check balance failed %s\n", err.Error())
	} else {
		fmt.Printf("balance is %f\n", balance)
	}

	// withdraw cash
	if err := atmMachine.WithdrawCash(firstAccount.AccountNumber, 700); err != nil {
		fmt.Printf("withdraw cash failed %s\n", err.Error())
	} else {
		fmt.Printf("withdraw cash succeed from account: %s\n", firstAccount.AccountNumber)
	}

	// deposit cash
	if err := atmMachine.WithdrawCash(secondAccount.AccountNumber, 200); err != nil {
		fmt.Printf("withdraw cash failed %s\n", err.Error())
	} else {
		fmt.Printf("withdraw cash succeed from account: %s\n", secondAccount.AccountNumber)
	}

	// check updated balance
	if balance, err := atmMachine.CheckBalance(secondAccount.AccountNumber); err != nil {
		fmt.Printf("check balance failed %s\n", err.Error())
	} else {
		fmt.Printf("balance is %f for account: %s\n", balance, secondAccount.AccountNumber)
	}
}
