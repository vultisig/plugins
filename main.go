package main

import (
	"fmt"
	"log"

	"github.com/vultisig/plugins/plugin"
	"github.com/vultisig/plugins/verifier"
)

func main() {
	// Create a new verifier
	v := verifier.NewVerifier()

	// Create and register a payroll plugin
	payrollplugin := plugin.NewPayrollplugin(
		"PayrollPlugin",
		"0xEmployeeAddress123",
		1000,
	)

	err := v.Registerplugin("PayrollPlugin", payrollplugin)
	if err != nil {
		log.Fatalf("Failed to register plugin: %v", err)
	}

	// Simulate plugin execution
	tx, err := payrollplugin.InitiateTransaction()
	if err != nil {
		log.Fatalf("Failed to initiate transaction: %v", err)
	}

	// Validate the transaction
	err = v.ValidateTransaction(tx)
	if err != nil {
		log.Fatalf("Transaction validation failed: %v", err)
	}

	fmt.Printf("Transaction validated successfully:\n")

}
