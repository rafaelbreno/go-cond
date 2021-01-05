package cond

type Condition struct {
	CurrentVal bool
}

type Wrapper struct {
	Value interface{}
}

func NewCondition() Condition {
	return Condition{}
}

func (c *Condition) And(conds ...bool) *Condition {
	l := len(conds)

	c.CurrentVal = conds[0] && conds[1]

	if l == 2 {
		c.CurrentVal = conds[0] && conds[1]
	}

	for i := 2; i < l; i++ {
		c.CurrentVal = c.CurrentVal && conds[i]
	}

	return c
}

// Go from https://play.golang.org/p/WOnE31ldpb
// Lines: 70~81
