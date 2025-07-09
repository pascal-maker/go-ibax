/*----------------------------------------------------------------
- Copyright (c) IBAX. All rights reserved.
- See LICENSE in the project root for license information.
---------------------------------------------------------------*/

package storage

import "sync"

type Table struct {
	name string
}

var (
	balancesMu sync.Mutex
	balances   = make(map[string]int)
)

func Init() {
	// Storage initialization
}

func GetTable(name string) *Table {
	return &Table{name: name}
}

func (t *Table) Get(keys []string) (map[string]interface{}, error) {
	if t.name == "balances" && len(keys) > 0 {
		balancesMu.Lock()
		defer balancesMu.Unlock()
		bal := balances[keys[0]]
		return map[string]interface{}{"balance": bal}, nil
	}
	return map[string]interface{}{"balance": 0}, nil
}

func (t *Table) Update(keys []string, values map[string]interface{}) error {
	if t.name == "balances" && len(keys) > 0 {
		balancesMu.Lock()
		defer balancesMu.Unlock()
		if v, ok := values["balance"]; ok {
			if bal, ok := v.(int); ok {
				balances[keys[0]] = bal
			} else if bal64, ok := v.(int64); ok {
				balances[keys[0]] = int(bal64)
			}
		}
	}
	return nil
}
