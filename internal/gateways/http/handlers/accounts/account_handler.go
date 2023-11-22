package accounts

type AccountHandler struct {
	createAccountUC     CreateAccountUC
	getAccountBalanceUC GetAccountBalanceUC
	getAccountstUC GetAccountstUC
}

func NewAccountHandler(createAccountUC CreateAccountUC) *AccountHandler {
	return &AccountHandler{
		createAccountUC: createAccountUC,
	}
}
