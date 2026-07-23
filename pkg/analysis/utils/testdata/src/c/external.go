package c

import (
	"extpkg"
	v1 "k8s.io/api/core/v1"
)

type ExternalTestStructs struct {
	// Should flag: external struct without IsZero
	TargetNoIsZero extpkg.NoIsZero `json:"targetNoIsZero,omitzero"` // want "zero value is valid" "validation is not complete"

	// Should not flag: external struct with IsZero on value receiver
	TargetValueRecv extpkg.ValueReceiver `json:"targetValueRecv,omitzero"` // want "zero value is valid" "validation is complete"

	// Should not flag: external struct with IsZero on pointer receiver
	TargetPtrRecv extpkg.PointerReceiver `json:"targetPtrRecv,omitzero"` // want "zero value is valid" "validation is complete"

	// Should not flag: hardcoded exception
	TargetLocalObjRef v1.LocalObjectReference `json:"targetLocalObjRef,omitzero"` // want "zero value is valid" "validation is complete"

	// Type aliases
	TargetAliasToNoIsZero AliasToNoIsZero `json:"targetAliasToNoIsZero,omitzero"` // want "zero value is valid" "validation is not complete"
}

type AliasToNoIsZero = extpkg.NoIsZero
