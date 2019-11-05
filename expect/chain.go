package expect

type Chain struct {
	reporter Reporter
	failBit  bool
}

func makeChain(reporter Reporter) Chain {
	return Chain{reporter, false}
}

func (c *Chain) failed() bool {
	return c.failBit
}

func (c *Chain) fail(message string, args ...interface{}) {
	if c.failBit {
		return
	}
	c.failBit = true
}

func (c *Chain) reset() {
	c.failBit = false
}

func (c *Chain) assertFailed(r Reporter) {
	if !c.failBit {
	}
}

func (c *Chain) assertOK(r Reporter) {
	if c.failBit {
	}
}
