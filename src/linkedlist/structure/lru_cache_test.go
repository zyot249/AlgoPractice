package structure

import "testing"

func Test1(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 0)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)
	obj.Put(4, 4)
	obj.Get(1)
	obj.Get(3)
	obj.Get(4)
}
