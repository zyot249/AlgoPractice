package practice

/*
Problem:
Ref: https://leetcode.com/problems/course-schedule/description/

Given:
- A total number of courses (numCourses) labeled from 0 -> numCourses - 1
- An array (prerequisites) where prerequisites[i] = (ai, bi) indicates that you must take course bi first if you want to take course ai

Return:
- True if you can finish all courses and false otherwise

Constraints:
- numCourses: [1, 200]
- prerequisites length: [1, 5000]
- all prerequisites are unique
*/

/*
First approach:
Firstly, I will enqueue all the courses that I can take immediately.
After that, I get dequeue all of them and enqueue all the next courses I can take.
Redo until no next course can be taken
If total taken courses == total courses => true
false otherwise

[
	[0, 1]
	[2, 1]
	[2, 3]
]
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	prerequisitesMap := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		prerequisitesMap[prerequisite[0]] = append(prerequisitesMap[prerequisite[0]], prerequisite[1])
	}

	totalTakenCourses := 0
	takenCourses := make([]bool, numCourses)
	for course := 0; course < numCourses; course++ {
		_, exists := prerequisitesMap[course]
		if !exists {
			totalTakenCourses += 1
			takenCourses[course] = true
		}
	}

	haveNext := true
	for haveNext {
		haveNext = false

		for course, prerequisite := range prerequisitesMap {
			canTake := true
			for _, requireCourse := range prerequisite {
				if !takenCourses[requireCourse] {
					canTake = false
					break
				}
			}

			if canTake {
				takenCourses[course] = true
				delete(prerequisitesMap, course)
				totalTakenCourses += 1
				haveNext = true
			}
		}
	}

	if totalTakenCourses == numCourses {
		return true
	}

	return false
}
