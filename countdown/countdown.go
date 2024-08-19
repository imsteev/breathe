package countdown

func New(from int) *CountDown {
	return &CountDown{
		from: from,
		current: from,
	}
}

type CountDown struct {
	from      int
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
	c.current = c.from
}
