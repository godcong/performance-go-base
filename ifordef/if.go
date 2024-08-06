package ifordef

type ConditionIf struct {
	caller func() bool
}

func (c ConditionIf) Call() bool {
	if c.caller == nil {
		return false
	}
	return c.caller()
}

func NewIf(caller func() bool) *ConditionIf {
	return &ConditionIf{caller: caller}
}
