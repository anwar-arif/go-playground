package atm

type WithdrawalTransaction struct {
	BaseTransaction
}

func NewWithdrawalTransaction(transactionId string, account *Account, amount float64) *WithdrawalTransaction {
	return &WithdrawalTransaction{
		BaseTransaction{
			TransactionId: transactionId,
			Account:       account,
			Amount:        amount,
		},
	}
}

func (wt *WithdrawalTransaction) Execute() error {
	if err := wt.Account.Debit(wt.Amount); err != nil {
		return err
	}
	return nil
}
