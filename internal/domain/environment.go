package domain

// Production Serves end-users/clients
const Production = "prod"

// Stage Mirror of production environment
const Stage = "stage"

// Development server acting as a sandbox where unit testing may be performed by the developer
const Development = "dev"

// Integration CI build target, or for developer testing of side effects
const Integration = "integration"

// Test The environment where interface testing is performed
const Test = "test"

// SanitizeEnvironment cleans the current kernel's environment and set dev environment if
// value is not valid
func SanitizeEnvironment(env string) string {
	if env != Production && env != Development && env != Integration && env != Stage &&
		env != Test {
		return Development
	}

	return env
}
