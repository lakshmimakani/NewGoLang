package accountopen

type Account struct {
	Username string
	Password string
	Balance  float64
}

var (
	accounts map[string]Account

	totalAccounts int
)

func init() {
	accounts = make(map[string]Account)
}
func OpenAccount(username, password string, initialBalance float64) Account {

	newAccount := Account{
		Username: username,
		Password: password,
		Balance:  initialBalance,
	}

	accounts[username] = newAccount

	totalAccounts++

	return newAccount

}
func (acc1 *Account) AddAmount(amount float64) {
	acc1.Balance += amount
}
func (acc1 *Account) DeductAmount(amount float64) {
	acc1.Balance -= amount
}

func Transfer(sender, receiver *Account, amount float64) error {

	// Deduct amount from sender
	sender.Balance -= amount

	// Add amount to receiver
	receiver.Balance += amount

	return nil
}

func GetTotalAccounts() int {

	return totalAccounts
}

func CloseAccount(username string) string {

	delete(accounts, username)
	totalAccounts--

	return "Account is closed"
}
