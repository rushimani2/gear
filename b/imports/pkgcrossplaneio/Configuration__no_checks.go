//go:build no_runtime_type_checking

package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateConfiguration_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfiguration_ManifestParameters(props *ConfigurationProps) error {
	return nil
}

func validateConfiguration_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewConfigurationParameters(scope constructs.Construct, id *string, props *ConfigurationProps) error {
	return nil
}

