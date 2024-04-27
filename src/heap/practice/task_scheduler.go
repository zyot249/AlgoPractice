package practice

import "container/heap"

/*
Problem:
Ref: https://leetcode.com/problems/task-scheduler/description/

Given:
- An array of CPU tasks represented by letters from A - Z
- A cooling time n
- Identical tasks must be separated by at least n intervals due to cooling time
- Each cycle or interval allows the completion of one task
- Tasks can be completed in any order

Return:
- The minimum number of intervals required to complete all tasks

Constraints:
- tasks length: [1, 10^4]
- n: [0, 100]
- tasks[i]: uppercase English letter
*/

/*
First approach:
I will use a heap as a priority queue to find the next task need to be completed
Task with smaller waiting intervals will be completed first.
If some tasks have the same waiting intervals, task with greater waiting tasks will be completed first.
*/

type TaskScheduler struct {
	Name         byte
	WaitingTasks int
	NextInterval int
}

type TaskPriorityQueue []*TaskScheduler

func (h *TaskPriorityQueue) Len() int { return len(*h) }
func (h *TaskPriorityQueue) Less(i, j int) bool {
	return (*h)[i].WaitingTasks > (*h)[j].WaitingTasks
}

func (h *TaskPriorityQueue) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *TaskPriorityQueue) Push(x interface{}) { *h = append(*h, x.(*TaskScheduler)) }
func (h *TaskPriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	taskCountByName := make([]int, 26)
	for _, task := range tasks {
		taskCountByName[task-'A']++
	}

	var taskQueue []*TaskScheduler
	for i := 0; i < len(taskCountByName); i++ {
		taskName := byte('A' + i)
		if taskCountByName[i] == 0 {
			continue
		}

		taskQueue = append(taskQueue, &TaskScheduler{
			Name:         taskName,
			WaitingTasks: taskCountByName[i],
		})
	}

	queue := TaskPriorityQueue(taskQueue)
	heap.Init(&queue)

	count := 0
	var coolingTasks []*TaskScheduler
	for queue.Len() > 0 || len(coolingTasks) > 0 {
		count++

		if len(coolingTasks) > 0 && coolingTasks[0].NextInterval <= count {
			t := coolingTasks[0]
			heap.Push(&queue, t)
			coolingTasks = coolingTasks[1:]
		}

		if queue.Len() > 0 {
			task := heap.Pop(&queue).(*TaskScheduler)

			if task.WaitingTasks > 1 {
				task.WaitingTasks--
				task.NextInterval = count + n + 1
				coolingTasks = append(coolingTasks, task)
			}
		}
	}

	return count
}

/*
Solution:
Use the formula:
Create a frequency array to store the count of each task
Find the maximum frequency among all tasks
Count the number of tasks with the maximum frequency
Calculate the minimum number of intervals required
    intervals := (maxFreq - 1) * (n + 1) + maxCount
Return the maximum of intervals and the total number of tasks
*/
