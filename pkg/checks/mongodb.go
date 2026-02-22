package checks

import (
	"github.com/scorify/schema"

	mongodb "github.com/scorify/mongodb"
)

func init() {
	schema, err := schema.Describe(mongodb.Schema{})
	if err != nil {
		panic(err)
	}

	Checks["mongodb"] = Check{
		Func:     mongodb.Run,
		Validate: mongodb.Validate,
		Schema:   schema,
	}
}
