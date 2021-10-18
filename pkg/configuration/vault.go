package configuration

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// loadVaultSecrets attempts to load secrets from Vault.
func loadVaultSecrets(v *viper.Viper) error {
	// Validate whether Vault is configured.
	vaultAddr := v.GetString(varVaultAddr)
	if len(vaultAddr) == 0 {
		// Do nothing.
		return nil
	}

	// Attempt to connect to Vault.
	log.Println("[Vault] Attempting to load secrets...")

	// Build the configuration.
	config := vault.DefaultConfig()
	config.Address = vaultAddr

	// Build a client.
	client, err := vault.NewClient(config)
	if err != nil {
		return fmt.Errorf("[Vault] could not initialise Vault: %w", err)
	}

	// Perform authentication.
	client.SetToken(v.GetString(varVaultToken))

	// Load the secret.
	secret, err := client.Logical().Read(v.GetString(varVaultSecretsPath))
	if err != nil {
		return fmt.Errorf("[Vault] could not read secrets: %w", err)
	} else if secret == nil {
		log.Println("[Vault] Secret engine was not found.")
		return nil
	} else if secret.Data == nil {
		log.Println("[Vault] No secrets were found at the given path.")
		return nil
	}

	// Extract the data.
	dataRaw := secret.Data["data"]
	data, ok := dataRaw.(map[string]interface{})
	if !ok {
		return fmt.Errorf("[Vault] could not load data: %T %#v", dataRaw, dataRaw)
	}

	// Override the configured variables.
	loaded := 0
	for _, keyRaw := range v.AllKeys() {
		// Check whether this variable was overridden.
		key := strings.ToUpper(keyRaw)
		value, ok := data[key]
		if ok {
			v.Set(key, value)
			loaded += 1
		}
	}

	log.Printf("[Vault] Loaded %d secrets.", loaded)

	return nil
}
