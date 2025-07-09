package smart

import (
    "github.com/IBAX-io/go-ibax/packages/consts"
    "github.com/IBAX-io/go-ibax/packages/contract"
    "github.com/IBAX-io/go-ibax/packages/storage"
)

// TokenContractTransfer transfers tokens from one account to another
func TokenContractTransfer(args ...interface{}) (interface{}, error) {
    to := args[0].(string)
    amount := args[1].(int64)

    // Get sender's address from context
    ctx := contract.GetContext()
    sender := ctx.Sender

    // Get sender's balance
    senderBalance, err := storage.GetTable("balances").Get([]string{sender})
    if err != nil {
        return false, err
    }

    // Check if sender has enough balance
    if senderBalance["balance"].(int64) < amount {
        return false, consts.ErrInsufficientBalance
    }

    // Update sender's balance
    newSenderBalance := senderBalance["balance"].(int64) - amount
    if err := storage.GetTable("balances").Update([]string{sender}, map[string]interface{}{"balance": newSenderBalance}); err != nil {
        return false, err
    }

    // Get recipient's balance
    recipientBalance, err := storage.GetTable("balances").Get([]string{to})
    if err != nil {
        return false, err
    }

    // Update recipient's balance
    newRecipientBalance := recipientBalance["balance"].(int64) + amount
    if err := storage.GetTable("balances").Update([]string{to}, map[string]interface{}{"balance": newRecipientBalance}); err != nil {
        return false, err
    }

    return true, nil
}

// TokenContractMint mints new tokens to an account
func TokenContractMint(args ...interface{}) (interface{}, error) {
    to := args[0].(string)
    amount := args[1].(int64)

    // Get recipient's balance
    recipientBalance, err := storage.GetTable("balances").Get([]string{to})
    if err != nil {
        return false, err
    }

    // Update recipient's balance
    newBalance := recipientBalance["balance"].(int64) + amount
    if err := storage.GetTable("balances").Update([]string{to}, map[string]interface{}{"balance": newBalance}); err != nil {
        return false, err
    }

    return true, nil
}

// TokenContractBurn burns tokens from the sender's account
func TokenContractBurn(args ...interface{}) (interface{}, error) {
    amount := args[0].(int64)

    // Get sender's address from context
    ctx := contract.GetContext()
    sender := ctx.Sender

    // Get sender's balance
    senderBalance, err := storage.GetTable("balances").Get([]string{sender})
    if err != nil {
        return false, err
    }

    // Check if sender has enough balance
    if senderBalance["balance"].(int64) < amount {
        return false, consts.ErrInsufficientBalance
    }

    // Update sender's balance
    newBalance := senderBalance["balance"].(int64) - amount
    if err := storage.GetTable("balances").Update([]string{sender}, map[string]interface{}{"balance": newBalance}); err != nil {
        return false, err
    }

    return true, nil
}

// TokenContractGetBalance gets the balance of an account
func TokenContractGetBalance(args ...interface{}) (interface{}, error) {
    address := args[0].(string)

    // Get account balance
    balance, err := storage.GetTable("balances").Get([]string{address})
    if err != nil {
        return 0, err
    }

    return balance["balance"].(int64), nil
}
