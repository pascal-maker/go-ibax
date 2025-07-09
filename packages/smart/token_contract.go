package smart

// TokenContract demonstrates basic token transfer functionality
func TokenContract() string {
    return `{
        "name": "TokenContract",
        "version": "1.0",
        "description": "Basic token transfer contract",
        "functions": [
            {
                "name": "transfer",
                "args": [
                    {"name": "to", "type": "string"},
                    {"name": "amount", "type": "integer"}
                ],
                "returns": [
                    {"name": "success", "type": "boolean"}
                ]
            },
            {
                "name": "mint",
                "args": [
                    {"name": "to", "type": "string"},
                    {"name": "amount", "type": "integer"}
                ],
                "returns": [
                    {"name": "success", "type": "boolean"}
                ]
            },
            {
                "name": "burn",
                "args": [
                    {"name": "amount", "type": "integer"}
                ],
                "returns": [
                    {"name": "success", "type": "boolean"}
                ]
            },
            {
                "name": "getBalance",
                "args": [
                    {"name": "address", "type": "string"}
                ],
                "returns": [
                    {"name": "balance", "type": "integer"}
                ]
            }
        ],
        "tables": [
            {
                "name": "balances",
                "fields": [
                    {"name": "address", "type": "string", "key": true},
                    {"name": "balance", "type": "integer"}
                ]
            }
        ]
    }`
}
