package plugin

import (
	"time"
)

// Transaction represents a money transfer
type Transaction struct {
	/*From      string
	To        string
	...
	...
	Amount    uint64
	Timestamp time.Time*/
	PluginId string
}

// plugin interface defines what a plugin can do
type Plugin interface {
	InitiateTransaction() (*Transaction, error)
	GetId() string
}

// Payrollplugin implements the plugin interface
type Payrollplugin struct {
	PluginId        string
	employeeAddress string
	monthlyAmount   uint64
	lastPaid        time.Time
}

func NewPayrollplugin(PluginId string, employee string, amount uint64) *Payrollplugin {
	return &Payrollplugin{
		PluginId:        "PayrollPlugin",
		employeeAddress: employee,
		monthlyAmount:   amount,
		lastPaid:        time.Now(),
	}
}

func (p *Payrollplugin) GetId() string {
	return p.PluginId
}

func (p *Payrollplugin) InitiateTransaction() (*Transaction, error) {
	now := time.Now()

	// dummy condition
	/*if now.Sub(p.lastPaid) < 30*24*time.Hour {
		return nil, fmt.Errorf("time error")
	}*/

	tx := &Transaction{
		/*From:      "COMPANY_WALLET",
		To:        p.employeeAddress,
		...
		...
		Amount:    p.monthlyAmount,
		Timestamp: now,*/
		PluginId: p.GetId(),
	}

	p.lastPaid = now
	return tx, nil
}
