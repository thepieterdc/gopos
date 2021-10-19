package configuration

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thepieterdc/gopos/pkg/logging"
)

var (
	varGoogleApiBase    = "GOOGLE_API_BASE"
	varGoogleApiKey     = "GOOGLE_API_KEY"
	varMongoUri         = "MONGO_URI"
	varVaultAddr        = "VAULT_ADDR"
	varVaultSecretsPath = "VAULT_SECRETS_PATH"
	varVaultToken       = "VAULT_TOKEN"
)

// config holds the configuration.
var config *Configuration

func initialise() (*Configuration, error) {
	// Create a new Viper instance.
	v := viper.New()
	v.AutomaticEnv()

	// Configure the supported environment variables.
	v.SetDefault(varGoogleApiBase, "https://maps.googleapis.com")
	v.SetDefault(varGoogleApiKey, "")
	v.SetDefault(varMongoUri, "")
	v.SetDefault(varVaultAddr, "")
	v.SetDefault(varVaultSecretsPath, "")
	v.SetDefault(varVaultToken, "")

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
