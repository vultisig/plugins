package verifier

import (
	"fmt"

	"github.com/vultisig/plugins/plugin"
)

// userPlugins maps user IDs to their authorized plugin names
type userPlugins map[string]map[string]bool

type Verifier struct {
	registeredPlugins map[string]plugin.Plugin
}

func NewVerifier() *Verifier {
	return &Verifier{
		registeredPlugins: make(map[string]plugin.Plugin),
	}
}

func (v *Verifier) Registerplugin(name string, p plugin.Plugin) error {
	//perform resharing with the plugin, once this is done we add it to the authorized list
	if _, exists := v.registeredPlugins[name]; exists {
		return fmt.Errorf("plugin %s already registered", name)
	}
	v.registeredPlugins[name] = p
	return nil
}

func (v *Verifier) ValidateTransaction(tx *plugin.Transaction) error {
	if _, exists := v.registeredPlugins[tx.PluginId]; !exists {
		return fmt.Errorf("plugin %s not allowed", tx.PluginId)
	}

	//perform transaction

	return nil
}
