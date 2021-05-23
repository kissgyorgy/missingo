package missingo

type ternary struct {
	conditionIsTrue bool
	trueValue       interface{}
}

func If(condition bool) *ternary {
	return &ternary{conditionIsTrue: condition}
}

func (t *ternary) Then(trueValue interface{}) *ternary {
	if t.conditionIsTrue {
		t.trueValue = trueValue
	}
	return t
}

func (t *ternary) Else(falseValue interface{}) interface{} {
	if t.conditionIsTrue {
		return t.trueValue
	} else {
		return falseValue
	}
}
