package deque

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DequeTestSuite struct {
	suite.Suite
}

func (suite *DequeTestSuite) SetUpTest() {
}

func TestDequeTestSuite(t *testing.T) {
	suite.Run(t, new(DequeTestSuite))
}

func (suite *DequeTestSuite) TestNewDeque() {
	deque := New[int]()
	suite.NotNil(deque)
}

func (suite *DequeTestSuite) TestNewDequeWithItems() {
	deque := New[int](1, 2, 3)
	suite.NotNil(deque)
	suite.Equal(3, deque.Size())
}

func (suite *DequeTestSuite) TestPushFront() {
	deque := New[int]()
	deque.PushFront(1)
	deque.PushFront(2)
	deque.PushFront(3)
	suite.Equal(3, deque.Size())
}

func (suite *DequeTestSuite) TestPushBack() {
	deque := New[int]()
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)
	suite.Equal(3, deque.Size())
}

func (suite *DequeTestSuite) TestPopFront() {
	deque := New[int](1, 2, 3)
	item, ok := deque.PopFront()
	suite.Equal(1, item)
	suite.True(ok)
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestPopFrontEmpty() {
	deque := New[int]()
	item, ok := deque.PopFront()
	suite.Equal(0, item)
	suite.False(ok)
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestPopBack() {
	deque := New[int](1, 2, 3)
	item, ok := deque.PopBack()
	suite.Equal(3, item)
	suite.True(ok)
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestPopBackEmpty() {
	deque := New[int]()
	item, ok := deque.PopBack()
	suite.Equal(0, item)
	suite.False(ok)
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestFront() {
	deque := New[int](1, 2, 3)
	item, ok := deque.Front()
	suite.Equal(1, item)
	suite.True(ok)
	suite.Equal(3, deque.Size())
}

func (suite *DequeTestSuite) TestFrontEmpty() {
	deque := New[int]()
	item, ok := deque.Front()
	suite.Equal(0, item)
	suite.False(ok)
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestBack() {
	deque := New[int](1, 2, 3)
	item, ok := deque.Back()
	suite.Equal(3, item)
	suite.True(ok)
	suite.Equal(3, deque.Size())
}

func (suite *DequeTestSuite) TestBackEmpty() {
	deque := New[int]()
	item, ok := deque.Back()
	suite.Equal(0, item)
	suite.False(ok)
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestEmpty() {
	deque := New[int]()
	suite.True(deque.Empty())
}

func (suite *DequeTestSuite) TestEmptyNotEmpty() {
	deque := New[int](1)
	suite.False(deque.Empty())
}

func (suite *DequeTestSuite) TestSize() {
	deque := New[int]()
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestSizeNotEmpty() {
	deque := New[int](1)
	suite.Equal(1, deque.Size())
}

func (suite *DequeTestSuite) TestClear() {
	deque := New[int](1, 2, 3)
	deque.Clear()
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestClearEmpty() {
	deque := New[int]()
	deque.Clear()
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestAt() {
	deque := New[int](1, 2, 3)
	item, ok := deque.At(1)
	suite.Equal(2, item)
	suite.True(ok)
}

func (suite *DequeTestSuite) TestAtEmpty() {
	deque := New[int]()
	item, ok := deque.At(1)
	suite.Equal(0, item)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestAtOutOfRange() {
	deque := New[int](1, 2, 3)
	item, ok := deque.At(3)
	suite.Equal(0, item)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestSet() {
	deque := New[int](1, 2, 3)
	ok := deque.Set(1, 4)
	suite.True(ok)
	suite.Equal(4, deque.data[1])
}

func (suite *DequeTestSuite) TestSetEmpty() {
	deque := New[int]()
	ok := deque.Set(1, 4)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestSetOutOfRange() {
	deque := New[int](1, 2, 3)
	ok := deque.Set(3, 4)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestInsert() {
	deque := New[int](1, 2, 3)
	ok := deque.Insert(1, 4)
	suite.True(ok)
	suite.Equal(4, deque.data[1])
}

func (suite *DequeTestSuite) TestInsertEmpty() {
	deque := New[int]()
	ok := deque.Insert(1, 4)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestInsertOutOfRange() {
	deque := New[int](1, 2, 3)
	ok := deque.Insert(4, 4)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestInsertRange() {
	deque := New[int](1, 2, 3)
	ok := deque.InsertRange(1, []int{4, 5})
	suite.True(ok)
	suite.Equal(4, deque.data[1])
	suite.Equal(5, deque.data[2])
}

func (suite *DequeTestSuite) TestInsertRangeEmpty() {
	deque := New[int]()
	ok := deque.InsertRange(1, []int{4, 5})
	suite.False(ok)
}

func (suite *DequeTestSuite) TestInsertRangeOutOfRange() {
	deque := New[int](1, 2, 3)
	ok := deque.InsertRange(4, []int{4, 5})
	suite.False(ok)
}

func (suite *DequeTestSuite) TestErase() {
	deque := New[int](1, 2, 3)
	ok := deque.Erase(1)
	suite.True(ok)
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestEraseEmpty() {
	deque := New[int]()
	ok := deque.Erase(1)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestEraseOutOfRange() {
	deque := New[int](1, 2, 3)
	ok := deque.Erase(4)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestEraseRange() {
	deque := New[int](1, 2, 3, 4)
	ok := deque.EraseRange(1, 3)
	suite.True(ok)
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestEraseRangeEmpty() {
	deque := New[int]()
	ok := deque.EraseRange(1, 2)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestEraseRangeOutOfRange() {
	deque := New[int](1, 2, 3)
	ok := deque.EraseRange(4, 5)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestAppendRange() {
	deque := New[int](1, 2, 3)
	deque.AppendRange([]int{4, 5}...)
	suite.Equal(5, deque.Size())
}

func (suite *DequeTestSuite) TestAppendRangeEmpty() {
	deque := New[int]()
	deque.AppendRange([]int{4, 5}...)
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestAppendRangeWithItems() {
	deque := New[int](1, 2, 3)
	deque.AppendRange([]int{4, 5}...)
	suite.Equal(5, deque.Size())
}

func (suite *DequeTestSuite) TestPrependRange() {
	deque := New[int](1, 2, 3)
	deque.PrependRange([]int{4, 5}...)
	suite.Equal(5, deque.Size())
}

func (suite *DequeTestSuite) TestPrependRangeEmpty() {
	deque := New[int]()
	deque.PrependRange([]int{4, 5}...)
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestPrependRangeWithItems() {
	deque := New[int](1, 2, 3)
	deque.PrependRange([]int{4, 5}...)
	suite.Equal(5, deque.Size())
}

func (suite *DequeTestSuite) TestEraseIf() {
	deque := New[int](1, 2, 3)
	deque.EraseIf(func(item int) bool {
		return item == 2
	})
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestEraseIfEmpty() {
	deque := New[int]()
	deque.EraseIf(func(item int) bool {
		return item == 2
	})
	suite.Equal(0, deque.Size())
}

func (suite *DequeTestSuite) TestEraseIfWithItems() {
	deque := New[int](1, 2, 3)
	deque.EraseIf(func(item int) bool {
		return item == 2
	})
	suite.Equal(2, deque.Size())
}

func (suite *DequeTestSuite) TestContains() {
	deque := New[int](1, 2, 3)
	ok := Contains(deque, 2)
	suite.True(ok)
}

func (suite *DequeTestSuite) TestContainsEmpty() {
	deque := New[int]()
	ok := Contains(deque, 2)
	suite.False(ok)
}

func (suite *DequeTestSuite) TestContainsWithItems() {
	deque := New[int](1, 2, 3)
	ok := Contains(deque, 2)
	suite.True(ok)
}

func (suite *DequeTestSuite) TestIterator() {
	deque := New[int](1, 2, 3)
	iterator := deque.Iterator()
	suite.NotNil(iterator)
}

// Benchmark pour PushBack sur votre Deque
func BenchmarkDequePushBack(b *testing.B) {
	deque := New[int]()
	for i := 0; i < b.N; i++ {
		deque.PushBack(i)
	}

	for i := 0; i < b.N; i++ {
		if deque.data[i] != i {
			b.Errorf("deque.data[%d] != %d", deque.data[i], i)
		}
	}
}

// Benchmark pour PushBack sur une slice
func BenchmarkSliceAppend(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}

	for i := 0; i < b.N; i++ {
		if slice[i] != i {
			b.Errorf("slice[%d] != %d", slice[i], i)
		}
	}
}
