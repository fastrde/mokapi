package sortedmap

type item[K comparable, V any] struct {
	key   K
	value V
}

type Iterator[K comparable, V any] struct {
	items   []item[K, V]
	current int
}

func (i *Iterator[K, V]) Next() bool {
	if i.current >= len(i.items) {
		return false
	}
	i.current++
	return true
}

func (i *Iterator[K, V]) Item() (K, V) {
	if i.current == 0 {
		panic("current is nil")
	}
	item := i.items[i.current-1]
	return item.key, item.value
}

func (i *Iterator[K, V]) Key() K {
	k, _ := i.Item()
	return k
}

func (i *Iterator[K, V]) Value() V {
	_, v := i.Item()
	return v
}
