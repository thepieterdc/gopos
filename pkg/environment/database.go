package environment

import "os"

// MongoUri URI to the MongoDB instance, used to cache Google Place responses.
var MongoUri = os.Getenv("MONGO_URI")
