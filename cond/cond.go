package cond

type Condition struct {
	Vars []bool
}

func NewCondition() Condition {
	return Condition{}
}

//func (c *Condition) CondOR ()
