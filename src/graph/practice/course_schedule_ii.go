package practice

/*
Problem:
Ref: https://leetcode.com/problems/course-schedule-ii/description/

Given:
- Total of courses (numCourses)
- Courses are labeled from 0 -> numCourses - 1
- An array (prerequisites) where prerequisites[i] = (ai, bi) indicates that you must take course bi first if you want to take course ai

Return:
- The ordering of courses you should take to finish all courses, if impossible to finish all courses => empty

Constraints:
- numCourses: [1, 2000]
- length of prerequisites: [0, numCourses * (numCourses - 1)]
- ai, bi: [0, numCourses)
- ai != bi
- all prerequisites are distinct
*/

/*
First approach:
From prerequisites, I will find the courses that can be taken immediately and add them to a queue
From them, I will find next courses I can take after.
Redo it until no more course can be taken
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	nextCourses := make([][]int, numCourses)
	courseRequirement := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		courseRequirement[prerequisite[0]]++
		nextCourses[prerequisite[1]] = append(nextCourses[prerequisite[1]], prerequisite[0])
	}

	queue := make([]int, 0)
	for course, requirement := range courseRequirement {
		if requirement == 0 {
			queue = append(queue, course)
		}
	}

	var order []int
	for len(queue) > 0 {
		count := len(queue)
		for _, course := range queue {
			order = append(order, course)
			if nextCourses[course] == nil {
				continue
			}

			for _, next := range nextCourses[course] {
				courseRequirement[next]--
				if courseRequirement[next] == 0 {
					queue = append(queue, next)
				}
			}
		}

		queue = queue[count:]
	}

	if len(order) == numCourses {
		return order
	} else {
		return []int{}
	}
}
