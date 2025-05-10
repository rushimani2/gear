package apiextensionscrossplaneio


// Policy configures the specifics of patching behaviour.
type CompositionRevisionSpecPatchSetsPatchesPolicy struct {
	// FromFieldPath specifies how to patch from a field path.
	//
	// The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.
	FromFieldPath CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath `field:"optional" json:"fromFieldPath" yaml:"fromFieldPath"`
}

