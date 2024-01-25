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
	suite.l = New[int](0, 1)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

//func (suite *ListTestSuite) TestList() {
//	suite.l.PushBack(2)
//	suite.l.PushBack(3)
//
//	suite.Equal(4, suite.l.Size())
//	suite.False(suite.l.Empty())
//}
//
//func (suite *ListTestSuite) TestListPushFront() {
//	suite.l.PushFront(2)
//	suite.l.PushFront(3)
//
//	suite.Equal(4, suite.l.Size())
//	suite.False(suite.l.Empty())
//}
//
//func (suite *ListTestSuite) TestListPushBack() {
//	suite.l.PushBack(2)
//	suite.l.PushBack(3)
//
//	suite.Equal(4, suite.l.Size())
//	suite.False(suite.l.Empty())
//}
//
//func (suite *ListTestSuite) TestListPopFront() {
//	suite.l.PushBack(2)
//	suite.l.PushBack(3)
//
//	val, _ := suite.l.PopFront()
//	suite.Equal(0, val)
//	val, _ = suite.l.PopFront()
//	suite.Equal(1, val)
//	val, _ = suite.l.PopFront()
//	suite.Equal(2, val)
//	val, _ = suite.l.PopFront()
//	suite.Equal(3, val)
//
//	suite.Equal(0, suite.l.Size())
//	suite.True(suite.l.Empty())
//}
//
//func (suite *ListTestSuite) TestListPopBack() {
//	suite.l.PushBack(2)
//	suite.l.PushBack(3)
//
//	val, _ := suite.l.PopBack()
//	suite.Equal(3, val)
//	val, _ = suite.l.PopBack()
//	suite.Equal(2, val)
//	val, _ = suite.l.PopBack()
//	suite.Equal(1, val)
//	val, _ = suite.l.PopBack()
//	suite.Equal(0, val)
//
//	suite.Equal(0, suite.l.Size())
//	suite.True(suite.l.Empty())
//}
//
//func (suite *ListTestSuite) TestListClear() {
//	suite.l.PushBack(2)
//	suite.l.PushBack(3)
//
//	suite.l.Clear()
//	suite.Equal(0, suite.l.Size())
//	suite.True(suite.l.Empty())
//}
//
//func (suite *ListTestSuite) TestListFront() {
//	suite.l.PushBack(2)
//	suite.l.PushBack(3)
//
//	val, _ := suite.l.Front()
//	suite.Equal(0, val)
//}
//
//func (suite *ListTestSuite) TestListBack() {
//	suite.l.PushBack(2)
//	suite.l.PushBack(3)
//
//	val, _ := suite.l.Back()
//	suite.Equal(3, val)
//}
//
//func (suite *ListTestSuite) TestListInsert() {}
