package structure

/*
Problem:
Ref: https://leetcode.com/problems/min-stack/
Design the stack that contains:
- push
- pop
- top
- getMin

Constraints:
- All the supported function have the time complexity of O(1)
- Element value: [-2^31, 2^31 - 1]
- pop, top, getMin can be called even with empty stack
- At most 3 * 10^4 calls will be made to push, pop, top and getMin
*/

/*
First approach:
Firstly, we can think about the data structure used to store all the elements in stack.
- array

Then, we handle each supported function of stack and think about how to serve it with time complexity of O(1):
- Push
- Pop
- Top
We just append new element to last of arr

- GetMin
In order to get the minimum of stack in O(1), we need to store the position of minimum element and update it each time push and pop are called
But pop function will be O(n)
*/

//type MinStack struct {
//	store       []int
//	minPosition int
//}
//
//func Constructor() MinStack {
//	stack := &MinStack{
//		store:       []int{},
//		minPosition: -1,
//	}
//
//	return *stack
//}
//
//func (this *MinStack) Push(val int) {
//	this.store = append(this.store, val)
//
//	if this.minPosition == -1 {
//		this.minPosition = 0
//	} else {
//		if val < this.store[this.minPosition] {
//			this.minPosition = len(this.store) - 1
//		}
//	}
//}
//
//func (this *MinStack) Pop() {
//	if len(this.store) == 0 {
//		return
//	}
//
//	if this.minPosition == len(this.store)-1 {
//		if len(this.store) == 1 {
//			this.minPosition = -1
//		} else {
//			this.minPosition = 0
//			for i := 1; i < len(this.store)-1; i++ {
//				if this.store[i] < this.store[this.minPosition] {
//					this.minPosition = i
//				}
//			}
//		}
//	}
//
//	this.store = this.store[:len(this.store)-1]
//}
//
//func (this *MinStack) Top() int {
//	if len(this.store) == 0 {
//		return 0
//	}
//
//	return this.store[len(this.store)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	if len(this.store) == 0 {
//		return 0
//	}
//
//	return this.store[this.minPosition]
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/*
Second approach:
We only improve to serve getMin function in O(1)
If we use only 1 variant to store the position of minimum element of all stack
The result is getMin function will have time complexity of O(1)
But it make pop function be slower and have O(n) to find new minimum element
Instead of that, we store minimum element of stack for every pushed elements
*/

type StackValue struct {
	val int
	min int
}

type MinStack struct {
	store []StackValue
}

func Constructor() MinStack {
	stack := &MinStack{
		store: []StackValue{},
	}

	return *stack
}

func (this *MinStack) Push(val int) {
	if len(this.store) == 0 {
		this.store = append(this.store, StackValue{val: val, min: val})
	} else {
		curMin := this.GetMin()
		if curMin <= val {
			this.store = append(this.store, StackValue{val: val, min: curMin})
		} else {
			this.store = append(this.store, StackValue{val: val, min: val})
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.store) == 0 {
		return
	}

	this.store = this.store[:len(this.store)-1]
}

func (this *MinStack) Top() int {
	if len(this.store) == 0 {
		return 0
	}

	return this.store[len(this.store)-1].val
}

func (this *MinStack) GetMin() int {
	if len(this.store) == 0 {
		return 0
	}

	return this.store[len(this.store)-1].min
}
