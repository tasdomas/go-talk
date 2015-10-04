package package_test

import (
	"testing"

	gc "gopkg.in/check.v1"
)

// START OMIT
// Boilerplate.
func Test(t *testing.T) {
	gc.TestingT(t)
}

// The test suite.
type TSuite struct{}

var _ = gc.Suite(&TSuite{})

func (t *TSuite) SetUpSuite(c *gc.C) {}

func (t *TSuite) TearDownSuite(c *gc.C) {}

func (t *TSuite) SetUpTest(c *gc.C) {}

func (t *TSuite) TearDownTest(c *gc.C) {}

func (t *TSuite) TestSomething(c *gc.C) {}

func (t *TSuite) BenchmarkSomething(c *gc.C) {}

// END OMIT
