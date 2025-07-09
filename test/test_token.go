package test

import (
    "fmt"
    "github.com/IBAX-io/go-ibax/packages/contract"
    "github.com/IBAX-io/go-ibax/packages/storage"
    "github.com/IBAX-io/go-ibax/packages/smart"
)

func main() {
    // Initialize contract and storage
    contract.Init()
    storage.Init()
    
    // Create test accounts
    account1 := "test_account_1"
    account2 := "test_account_2"
    
    // Get the token contract
    tokenContract := smart.TokenContract()
    
    // Set sender context
    ctx := contract.GetContext()
    ctx.Sender = account1
    
    // Mint initial tokens to account1
    mintResult, err := contract.Execute(tokenContract, "mint", account1, 1000)
    if err != nil {
        fmt.Printf("Error minting tokens: %v\n", err)
        return
    }
    fmt.Printf("Minted 1000 tokens to %s: %v\n", account1, mintResult)
    
    // Check initial balance
    balance1, err := contract.Execute(tokenContract, "getBalance", account1)
    if err != nil {
        fmt.Printf("Error getting balance: %v\n", err)
        return
    }
    fmt.Printf("Account %s balance: %d\n", account1, balance1)
    
    // Transfer tokens from account1 to account2
    ctx.Sender = account1
    transferResult, err := contract.Execute(tokenContract, "transfer", account2, 500)
    if err != nil {
        fmt.Printf("Error transferring tokens: %v\n", err)
        return
    }
    fmt.Printf("Transferred 500 tokens from %s to %s: %v\n", account1, account2, transferResult)
    
    // Check balances after transfer
    balance1, err = contract.Execute(tokenContract, "getBalance", account1)
    if err != nil {
        fmt.Printf("Error getting balance: %v\n", err)
        return
    }
    balance2, err := contract.Execute(tokenContract, "getBalance", account2)
    if err != nil {
        fmt.Printf("Error getting balance: %v\n", err)
        return
    }
    
    fmt.Printf("Account %s balance: %d\n", account1, balance1)
    fmt.Printf("Account %s balance: %d\n", account2, balance2)
    
    // Burn tokens from account2
    ctx.Sender = account2
    burnResult, err := contract.Execute(tokenContract, "burn", 250)
    if err != nil {
        fmt.Printf("Error burning tokens: %v\n", err)
        return
    }
    fmt.Printf("Burned 250 tokens from %s: %v\n", account2, burnResult)
    
    // Check final balances
    balance1, err = contract.Execute(tokenContract, "getBalance", account1)
    if err != nil {
        fmt.Printf("Error getting balance: %v\n", err)
        return
    }
    balance2, err = contract.Execute(tokenContract, "getBalance", account2)
    if err != nil {
        fmt.Printf("Error getting balance: %v\n", err)
        return
    }
    
    fmt.Printf("Final balances:\n")
    fmt.Printf("Account %s: %d\n", account1, balance1)
    fmt.Printf("Account %s: %d\n", account2, balance2)
}
