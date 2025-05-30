package secretscrossplaneio


// Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.
type StoreConfigSpecVaultAuthTokenEnv struct {
	// Name is the name of an environment variable.
	Name *string `field:"required" json:"name" yaml:"name"`
}

