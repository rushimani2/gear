package pkgcrossplaneio


// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type ProviderRevisionSpecPackagePullSecrets struct {
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
	Name *string `field:"optional" json:"name" yaml:"name"`
}

