package address

import (
	log "github.com/sirupsen/logrus"
	"github.com/thepieterdc/gopos/pkg/logging"
)

// Initialise the logging fields.
var logger = log.WithFields(logging.RunningStage()).WithFields(logging.AddressComponent())
