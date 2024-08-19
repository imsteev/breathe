package countdown

func New(to int) *CountDown {
	return &CountDown{
		to: to,
		current: to,
	}
}

type CountDown struct {
	to      int
	current int
}

func (c *CountDown) Next() bool {
	if c.current > 0  {
		c.current--
		return true
	}
	return false
}

func (c *CountDown) Reset() {
	c.current = c.to 
}
