# Go-IBAX Token Contract Demo

A simple token contract implementation with mint, transfer, burn, and getBalance functionality. This demo showcases a basic token system with in-memory storage.

## Features

- **Mint**: Create new tokens and assign them to an account
- **Transfer**: Move tokens between accounts
- **Burn**: Remove tokens from an account
- **GetBalance**: Query token balance for any account
- **In-Memory Storage**: Simple, thread-safe storage for demo purposes

## Project Structure

```
go-ibax/
├── packages/
│   ├── contract/
│   │   └── context.go          # Contract execution and context management
│   ├── storage/
│   │   └── storage.go          # In-memory storage implementation
│   └── smart/
│       ├── token_contract.go   # Token contract definition
│       └── token_contract_impl.go # Token contract implementation
├── test/
│   └── test_token.go           # Demo script
└── README.md                   # This file
```

## How It Works

### Contract Functions

1. **Mint(to, amount)**
   - Adds tokens to the specified account
   - No authorization required (for demo purposes)

2. **Transfer(to, amount)**
   - Moves tokens from sender to recipient
   - Requires sufficient balance in sender's account

3. **Burn(amount)**
   - Removes tokens from sender's account
   - Requires sufficient balance

4. **GetBalance(address)**
   - Returns the current balance for the specified address

### Context Management

The contract uses a global context to track the current sender:
- `SetSender(address)`: Set the current sender
- `GetContext()`: Get the current context

### Storage

Uses a simple in-memory map with thread-safe access:
- Balances are stored as `map[string]int`
- Thread-safe operations with mutex protection

## Running the Demo

### Prerequisites

- Go 1.16 or later
- Git

### Quick Start

1. **Clone the repository:**
   ```bash
   git clone <your-repo-url>
   cd go-ibax
   ```

2. **Run the demo:**
   ```bash
   go run test/test_token.go
   ```

### Expected Output

```
Minted 1000 tokens to test_account_1: true
Account test_account_1 balance: 1000
Transferred 500 tokens from test_account_1 to test_account_2: true
Account test_account_1 balance: 500
Account test_account_2 balance: 500
Burned 250 tokens from test_account_2: true
Final balances:
Account test_account_1: 500
Account test_account_2: 250
```

## Demo Flow

The test script demonstrates:

1. **Initialization**: Sets up contract and storage
2. **Minting**: Creates 1000 tokens for account1
3. **Balance Check**: Verifies the minted amount
4. **Transfer**: Moves 500 tokens from account1 to account2
5. **Balance Check**: Shows balances after transfer
6. **Burn**: Removes 250 tokens from account2
7. **Final Check**: Displays final balances

## Code Examples

### Setting the Sender Context

```go
contract.SetSender("my_account")
```

### Minting Tokens

```go
result, err := contract.Execute("TokenContract", "mint", "recipient", 1000)
```

### Transferring Tokens

```go
contract.SetSender("sender_account")
result, err := contract.Execute("TokenContract", "transfer", "recipient", 500)
```

### Checking Balance

```go
balance, err := contract.Execute("TokenContract", "getBalance", "account_address")
```

### Burning Tokens

```go
contract.SetSender("account_to_burn_from")
result, err := contract.Execute("TokenContract", "burn", 250)
```

## Architecture

### Contract Layer (`packages/contract/`)
- Handles contract execution
- Manages sender context
- Provides the main interface for contract operations

### Storage Layer (`packages/storage/`)
- Provides in-memory storage
- Thread-safe operations
- Simple key-value store for balances

### Smart Contract Layer (`packages/smart/`)
- Contains contract definitions
- Implements contract logic
- Provides contract metadata

## Future Enhancements

- [ ] Persistent storage (database/file-based)
- [ ] Access control for minting/burning
- [ ] Event logging for operations
- [ ] ERC-20 style approve/allowance system
- [ ] Multi-token support
- [ ] API endpoints for web integration
- [ ] Proper Go testing framework
- [ ] CLI tool for interaction

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is part of the IBAX ecosystem. See LICENSE file for details.

## Acknowledgments

Built as a demonstration of basic token contract functionality for the IBAX blockchain platform.



