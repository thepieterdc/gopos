package configuration

// MongoUri gets the uri to connect to the database.
func (c Configuration) MongoUri() string {
	return c.config.GetString(envongoUri)
}
