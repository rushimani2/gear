//go:build no_runtime_type_checking

package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateProvider_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProvider_ManifestParameters(props *ProviderProps) error {
	return nil
}

func validateProvider_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewProviderParameters(scope constructs.Construct, id *string, props *ProviderProps) error {
	return nil
}

