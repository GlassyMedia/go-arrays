package arrays

import (
	. "launchpad.net/gocheck"
)

func (s *S) TestToStringArray(c *C) {
	obj := []interface{}{"A", "b", "c"}
	strs, err := StringArray(obj)
	c.Assert(err, IsNil)
	c.Check(strs, DeepEquals, []string{"A", "b", "c"})
}

func (s *S) TestToStringArrayFail(c *C) {
	obj := []interface{}{"A", 300, "c"} // not all strings
	strs, err := StringArray(obj)
	c.Assert(err, ErrorMatches, "Error converting int to string")
	c.Check(strs, DeepEquals, []string{"A"})
}
