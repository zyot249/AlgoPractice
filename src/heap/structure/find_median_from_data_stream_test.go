package structure

import "testing"

func Test1(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.FindMedian()
	obj.AddNum(3)
	obj.FindMedian()

}
