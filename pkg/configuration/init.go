package configuration

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thepieterdc/gopos/pkg/logging"
)

const (
	envGoogleApiBase    = "GOOGLE_API_BASE"
	envGoogleApiKey     = "GOOGLE_API_KEY"
	envongoUri          = "MONGO_URI"
	envVaultAddr        = "VAULT_ADDR"
	envVaultSecretsPath = "VAULT_SECRETS_PATH"
	envVaultToken       = "VAULT_TOKEN"
)

// config holds the configuration.
var config *Configuration

func initialise() (*Configuration, error) {
	// Create a new Viper instance.
	v := viper.New()
	v.AutomaticEnv()

	// Configure the supported environment variables.
	v.SetDefault(envGoogleApiBase, "https://maps.googleapis.com")
	v.SetDefault(envGoogleApiKey, "")
	v.SetDefault(envongoUri, "")
	v.SetDefault(envVaultAddr, "")
	v.SetDefault(envVaultSecretsPath, "")
	v.SetDefault(envVaultToken, "")

	// Load environment variables from Vault.
	err := loadVaultSecrets(v)
	if err != nil {
		return nil, err
	}

	// Wrap the configuration.
	c := new(Configuration)
	c.config = v
	return c, err
}

// Configure initialises the configuration and returns it.
func Configure() *Configuration {
	// Initialise the configuration if needed.
	if config == nil {
		var err error
		config, err = initialise()
		if err != nil {
			log.WithFields(logging.BootStage()).WithFields(logging.VaultComponent()).Fatal(err)
		}
	}

	return config
}
