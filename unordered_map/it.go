package unordered_map

type Iterator[K comparable, V any] struct {
	m       *Map[K, V]
	index   uint64
	count   uint64
	found   bool
	founded uint64
}

func (it *Iterator[K, V]) Next() bool {
	if it.found {
		return false
	}

	for it.index < it.m.capacity {
		if !it.found && it.m.items[it.count].filled {
			it.index = it.count
			it.count++
			it.founded++

			if it.founded == it.m.size {
				it.found = true
			}

			return true
		}

		it.count++
	}

	return false
}

func (it *Iterator[K, V]) Value() (K, V) {
	if it.index < it.m.capacity && it.m.items[it.index].filled {
		return it.m.items[it.index].key, it.m.items[it.index].value
	}

	var k K
	var v V
	return k, v
}

func (it *Iterator[K, V]) SetValue(value V) {
	if it.index < 0 || it.index >= it.m.capacity {
		return
	}

	it.m.items[it.index].value = value
}

//func (it *Iterator[K, V]) Remove() {
//	if it.index < 0 || it.index >= it.m.capacity {
//		return
//	}
//
//	it.m.remove(it.index)
//	it.index--
//}
