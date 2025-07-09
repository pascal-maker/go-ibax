package contract

import (
	"fmt"

	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/storage"
)

type Context struct {
	Sender string
	Args   []interface{}
}

var globalContext = &Context{Sender: "test_sender"}

func GetContext() *Context {
	return globalContext
}

func SetSender(sender string) {
	globalContext.Sender = sender
}

func Init() {
	// Initialization code
	globalContext = &Context{Sender: "test_sender"}
}

// Execute executes a contract function
func Execute(contract string, function string, args ...interface{}) (interface{}, error) {
	if contract == "TokenContract" {
		switch function {
		case "mint":
			if len(args) != 2 {
				return nil, fmt.Errorf("mint requires 2 arguments: to, amount")
			}
			to, ok := args[0].(string)
			if !ok {
				return nil, fmt.Errorf("invalid to address")
			}
			amount, ok := args[1].(int)
			if !ok {
				return nil, fmt.Errorf("invalid amount")
			}

			// Get current balance
			table := storage.GetTable("balances")
			current, err := table.Get([]string{to})
			if err != nil {
				return nil, fmt.Errorf("error getting balance: %v", err)
			}

			// Update balance
			newBalance := current["balance"].(int) + amount
			if err := table.Update([]string{to}, map[string]interface{}{"balance": newBalance}); err != nil {
				return nil, fmt.Errorf("error updating balance: %v", err)
			}
			return true, nil

		case "transfer":
			if len(args) != 2 {
				return nil, fmt.Errorf("transfer requires 2 arguments: to, amount")
			}
			to, ok := args[0].(string)
			if !ok {
				return nil, fmt.Errorf("invalid to address")
			}
			amount, ok := args[1].(int)
			if !ok {
				return nil, fmt.Errorf("invalid amount")
			}

			// Get sender's balance
			table := storage.GetTable("balances")
			sender := GetContext().Sender
			senderBalance, err := table.Get([]string{sender})
			if err != nil {
				return nil, fmt.Errorf("error getting sender balance: %v", err)
			}

			// Check if sender has enough balance
			if senderBalance["balance"].(int) < amount {
				return nil, consts.ErrInsufficientBalance
			}

			// Update sender's balance
			newSenderBalance := senderBalance["balance"].(int) - amount
			if err := table.Update([]string{sender}, map[string]interface{}{"balance": newSenderBalance}); err != nil {
				return nil, fmt.Errorf("error updating sender balance: %v", err)
			}

			// Get recipient's balance
			recipientBalance, err := table.Get([]string{to})
			if err != nil {
				return nil, fmt.Errorf("error getting recipient balance: %v", err)
			}

			// Update recipient's balance
			newRecipientBalance := recipientBalance["balance"].(int) + amount
			if err := table.Update([]string{to}, map[string]interface{}{"balance": newRecipientBalance}); err != nil {
				return nil, fmt.Errorf("error updating recipient balance: %v", err)
			}
			return true, nil

		case "burn":
			if len(args) != 1 {
				return nil, fmt.Errorf("burn requires 1 argument: amount")
			}
			amount, ok := args[0].(int)
			if !ok {
				return nil, fmt.Errorf("invalid amount")
			}

			// Get current balance
			table := storage.GetTable("balances")
			sender := GetContext().Sender
			current, err := table.Get([]string{sender})
			if err != nil {
				return nil, fmt.Errorf("error getting balance: %v", err)
			}

			// Check if sender has enough balance
			if current["balance"].(int) < amount {
				return nil, consts.ErrInsufficientBalance
			}

			// Update balance
			newBalance := current["balance"].(int) - amount
			if err := table.Update([]string{sender}, map[string]interface{}{"balance": newBalance}); err != nil {
				return nil, fmt.Errorf("error updating balance: %v", err)
			}
			return true, nil

		case "getBalance":
			if len(args) != 1 {
				return nil, fmt.Errorf("getBalance requires 1 argument: address")
			}
			address, ok := args[0].(string)
			if !ok {
				return nil, fmt.Errorf("invalid address")
			}

			// Get balance
			table := storage.GetTable("balances")
			balance, err := table.Get([]string{address})
			if err != nil {
				return nil, fmt.Errorf("error getting balance: %v", err)
			}
			return balance["balance"].(int), nil

		default:
			return nil, fmt.Errorf("unknown function: %s", function)
		}
	}
	return nil, fmt.Errorf("unknown contract: %s", contract)
}
