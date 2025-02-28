package atm

type DepositTransaction struct {
	BaseTransaction
}

func NewDepositTransaction(transactionId string, account *Account, amount float64) *DepositTransaction {
	return &DepositTransaction{
		BaseTransaction{
			TransactionId: transactionId,
			Account:       account,
			Amount:        amount,
		},
	}
}

func (dt *DepositTransaction) Execute() error {
	return dt.Account.Credit(dt.Amount)
}
