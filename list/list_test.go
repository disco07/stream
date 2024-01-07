package list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ListTestSuite struct {
	suite.Suite

	l *List[int]
}

func (suite *ListTestSuite) SetupTest() {
	suite.l = New[int]()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

func (suite *ListTestSuite) TestList() {
	suite.Equal(0, suite.l.Size())
	suite.True(suite.l.Empty())

	suite.l.PushBack(0)
	suite.l.PushBack(1)
	suite.l.PushBack(2)
	suite.l.PushBack(3)

	suite.Equal(4, suite.l.Size())
	suite.False(suite.l.Empty())

}
