package apiextensionscrossplaneio


// FromFieldPath specifies how to patch from a field path.
//
// The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.
type CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath string

const (
	// Optional.
	CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath_OPTIONAL CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath = "OPTIONAL"
	// Required.
	CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath_REQUIRED CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath = "REQUIRED"
)

