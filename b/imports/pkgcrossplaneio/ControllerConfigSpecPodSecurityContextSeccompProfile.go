package pkgcrossplaneio


// The seccomp options to use by the containers in this pod.
//
// Note that this field cannot be set when spec.os.name is windows.
type ControllerConfigSpecPodSecurityContextSeccompProfile struct {
	// type indicates which kind of seccomp profile will be applied.
	//
	// Valid options are:
	// Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.
	Type *string `field:"required" json:"type" yaml:"type"`
	// localhostProfile indicates a profile defined in a file on the node should be used.
	//
	// The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is "Localhost".
	LocalhostProfile *string `field:"optional" json:"localhostProfile" yaml:"localhostProfile"`
}

