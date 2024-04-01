package practice

import (
	"sort"
)

/*
Problem:
Ref: https://leetcode.com/problems/car-fleet/
Given:
- all cars are going on a one-land road
- target destination
- a position array of cars
- a speed array of cars
- when cars meet each other before reaching target, they will become a fleet and adjust to the same min speed and continue go to target

Return:
- num of fleets

Constraints:
- num of cars: [1, 10^5]
- target: [0, 10^6]
- position[i]: [0, target)
- all position is unique
- speed[i]: (0, 10^6]
*/

/*
First approach:
Firstly, we sort list of cars by position in decent order.
We use a stack to store the time to finish of these cars
For each car in sorted list, we calculate the time needed for it to reach target.
If its needed time is smaller than the needed time of the car at the top of stack, don't push this time to stack and vice versa
*/

type Car struct {
	position int
	speed    int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i, p := range position {
		cars[i] = Car{position: p, speed: speed[i]}
	}

	sort.Slice(cars, func(a int, b int) bool {
		return cars[a].position > cars[b].position
	})

	var stack []float32
	for _, car := range cars {
		timeToFinish := float32(target-car.position) / float32(car.speed)
		if len(stack) == 0 || timeToFinish > stack[len(stack)-1] {
			stack = append(stack, timeToFinish)
		}
	}

	return len(stack)
}
