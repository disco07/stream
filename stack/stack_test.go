package stack

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type StackTestSuite struct {
	suite.Suite
}

func (suite *StackTestSuite) SetUpTest() {
}

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}

func (suite *StackTestSuite) TestNewStack() {
	stack := New[int]()
	suite.NotNil(stack)
}

func (suite *StackTestSuite) TestNewStackWithItems() {
	stack := New[int](1, 2, 3)
	suite.NotNil(stack)
	suite.Equal(3, stack.Size())
}

func (suite *StackTestSuite) TestPush() {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	suite.Equal(3, stack.Size())
}

func (suite *StackTestSuite) TestPop() {
	stack := New[int](1, 2, 3)
	item, b := stack.Pop()
	suite.Equal(3, item)
	suite.Equal(2, stack.Size())
	suite.True(b)
}

func (suite *StackTestSuite) TestPopEmpty() {
	stack := New[int]()
	item, b := stack.Pop()
	suite.Equal(0, item)
	suite.Equal(0, stack.Size())
	suite.False(b)
}

func (suite *StackTestSuite) TestIterator() {
	stack := New[int](1, 2, 3)
	it := stack.Iterator()
	log.Println(stack.data)
	for it.Next() {
		fmt.Println(it.Value())
	}
}
