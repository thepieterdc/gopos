package configuration

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thepieterdc/gopos/internal/pkg/logging"
	"strings"
)

// loadVaultSecrets attempts to load secrets from Vault.
func loadVaultSecrets(v *viper.Viper) error {
	// Initialise the logging fields.
	logger := log.WithFields(logging.BootStage()).WithFields(logging.VaultComponent())

	// Validate whether Vault is configured.
	vaultAddr := v.GetString(envVaultAddr)
	if len(vaultAddr) == 0 {
		// Do nothing.
		logger.Info("Not configured. Skipping.")
		return nil
	}

	// Attempt to connect to Vault.
	logger.Info("Attempting to load secrets.")

	// Build the configuration.
	config := vault.DefaultConfig()
	config.Address = vaultAddr

	// Build a client.
	client, err := vault.NewClient(config)
	if err != nil {
		return fmt.Errorf("could not initialise: %w", err)
	}

	// Perform authentication.
	client.SetToken(v.GetString(envVaultToken))

	// Load the secret.
	secret, err := client.Logical().Read(v.GetString(envVaultSecretsPath))
	if err != nil {
		return fmt.Errorf("could not read secrets: %w", err)
	} else if secret == nil {
		return fmt.Errorf("secret engine was not found")
	} else if secret.Data == nil {
		logger.Warn("No secrets were found at the given path.")
		return nil
	}

	// Extract the data.
	dataRaw := secret.Data["data"]
	data, ok := dataRaw.(map[string]interface{})
	if !ok {
		return fmt.Errorf("could not load data: %T %#v", dataRaw, dataRaw)
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

	logger.Infof("Loaded %d secrets.", loaded)

	return nil
}
