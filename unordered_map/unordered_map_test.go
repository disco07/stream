package unordered_map

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UnorderedMapTestSuite struct {
	suite.Suite
}

func (test *UnorderedMapTestSuite) SetupTest() {
}

func TestUnorderedMapTestSuite(t *testing.T) {
	suite.Run(t, new(UnorderedMapTestSuite))
}

func (test *UnorderedMapTestSuite) TestNewMap() {
	m := New[int, int](HashInt)
	test.NotNil(m)
}

func (test *UnorderedMapTestSuite) TestSet() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	test.Equal(uint64(1), m.Size())
}

func (test *UnorderedMapTestSuite) TestGet() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	v, ok := m.Get(1)
	test.Equal(2, v)
	test.True(ok)
}

func (test *UnorderedMapTestSuite) TestGetMissing() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	v, ok := m.Get(2)
	test.Equal(0, v)
	test.False(ok)
}

func (test *UnorderedMapTestSuite) TestSetOverwrite() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	m.Set(1, 3)
	v, ok := m.Get(1)
	test.Equal(3, v)
	test.True(ok)
}

func (test *UnorderedMapTestSuite) TestSetMany() {
	m := New[int, int](HashInt)
	for i := 0; i < 1000; i++ {
		m.Set(i, i)
	}

	test.Equal(uint64(1000), m.Size())
}

func (test *UnorderedMapTestSuite) TestDelete() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	m.Delete(1)
	test.Equal(uint64(0), m.Size())
}

func (test *UnorderedMapTestSuite) TestDeleteMissing() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	m.Delete(2)
	test.Equal(uint64(1), m.Size())
}

func (test *UnorderedMapTestSuite) TestDeleteMany() {
	m := New[int, int](HashInt)
	for i := 0; i < 100; i++ {
		m.Set(i, i)
	}

	for i := 0; i < 100; i++ {
		m.Delete(i)
	}

	test.Equal(uint64(0), m.Size())
}

func (test *UnorderedMapTestSuite) TestEmpty() {
	m := New[int, int](HashInt)
	test.True(m.Empty())
}

func (test *UnorderedMapTestSuite) TestEmptyNotEmpty() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	test.False(m.Empty())
}

func (test *UnorderedMapTestSuite) TestSize() {
	m := New[int, int](HashInt)
	test.Equal(uint64(0), m.Size())
}

func (test *UnorderedMapTestSuite) TestSizeNotEmpty() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	test.Equal(uint64(1), m.Size())
}

func (test *UnorderedMapTestSuite) TestClear() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	m.Clear()
	test.Equal(uint64(0), m.Size())
}

func (test *UnorderedMapTestSuite) TestClearEmpty() {
	m := New[int, int](HashInt)
	m.Clear()
	test.Equal(uint64(0), m.Size())
}

func (test *UnorderedMapTestSuite) TestEraseIf() {
	m := New[int, int](HashInt)
	m.Set(1, 2)
	m.EraseIf(func(k int, v int) bool {
		return k == 1
	})
	test.Equal(uint64(0), m.Size())
}
