//go:build no_runtime_type_checking

package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateControllerConfigSpecResourcesRequests_FromNumberParameters(value *float64) error {
	return nil
}

func validateControllerConfigSpecResourcesRequests_FromStringParameters(value *string) error {
	return nil
}

