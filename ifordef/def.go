package ifordef

type ConditionDef struct {
	caller func() bool
}

func (c ConditionDef) Call() bool {
	return c.caller()
}

func NewDef(caller func() bool) *ConditionDef {
	if caller == nil {
		caller = func() bool { return false }
	}
	return &ConditionDef{
		caller: caller,
	}
}
