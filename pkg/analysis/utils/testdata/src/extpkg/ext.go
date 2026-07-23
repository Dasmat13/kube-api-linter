package extpkg

type NoIsZero struct {
	Field string `json:"field,omitempty"`
}

type ValueReceiver struct{
	Field string `json:"field,omitempty"`
}
func (v ValueReceiver) IsZero() bool { return true }

type PointerReceiver struct{
	Field string `json:"field,omitempty"`
}
func (p *PointerReceiver) IsZero() bool { return true }
