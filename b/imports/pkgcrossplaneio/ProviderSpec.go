package pkgcrossplaneio


// ProviderSpec specifies details about a request to install a provider to Crossplane.
type ProviderSpec struct {
	// Package is the name of the package that is being requested.
	Package *string `field:"required" json:"package" yaml:"package"`
	// ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.
	ControllerConfigRef *ProviderSpecControllerConfigRef `field:"optional" json:"controllerConfigRef" yaml:"controllerConfigRef"`
	// IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package.
	//
	// Default is false.
	// Default: false.
	//
	IgnoreCrossplaneConstraints *bool `field:"optional" json:"ignoreCrossplaneConstraints" yaml:"ignoreCrossplaneConstraints"`
	// PackagePullPolicy defines the pull policy for the package.
	//
	// Default is IfNotPresent.
	// Default: IfNotPresent.
	//
	PackagePullPolicy *string `field:"optional" json:"packagePullPolicy" yaml:"packagePullPolicy"`
	// PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries.
	PackagePullSecrets *[]*ProviderSpecPackagePullSecrets `field:"optional" json:"packagePullSecrets" yaml:"packagePullSecrets"`
	// RevisionActivationPolicy specifies how the package controller should update from one revision to the next.
	//
	// Options are Automatic or Manual. Default is Automatic.
	// Default: Automatic.
	//
	RevisionActivationPolicy *string `field:"optional" json:"revisionActivationPolicy" yaml:"revisionActivationPolicy"`
	// RevisionHistoryLimit dictates how the package controller cleans up old inactive package revisions.
	//
	// Defaults to 1. Can be disabled by explicitly setting to 0.
	// Default: 1. Can be disabled by explicitly setting to 0.
	//
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package.
	//
	// Setting this value to true may have unintended consequences. Default is false.
	// Default: false.
	//
	SkipDependencyResolution *bool `field:"optional" json:"skipDependencyResolution" yaml:"skipDependencyResolution"`
}

