package logging

import "github.com/sirupsen/logrus"

func AddressComponent() logrus.Fields {
	return logrus.Fields{"component": "address"}
}

func DatabaseComponent() logrus.Fields {
	return logrus.Fields{"component": "db"}
}

func GoogleComponent() logrus.Fields {
	return logrus.Fields{"component": "google"}
}

func VaultComponent() logrus.Fields {
	return logrus.Fields{"component": "vault"}
}