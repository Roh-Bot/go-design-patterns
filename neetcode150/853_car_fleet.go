package neetcode150

import "sort"

type Car struct {
	Position int
	Speed    int
}

func (c Car) TimeToReachDestination(destination int) float64 {
	return float64(destination-c.Position) / float64(c.Speed)
}

// Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
// Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
// Output: 3
func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))

	for i := 0; i < len(position); i++ {
		cars[i] = Car{
			Position: position[i],
			Speed:    speed[i],
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position < cars[j].Position
	})

	stack := make(Stack[Car], 0, len(cars))
	for i := len(cars) - 1; i >= 0; i-- {
		if !stack.IsEmpty() && stack.Peek().TimeToReachDestination(target) > cars[i].TimeToReachDestination(target) {
			continue
		}
		stack.Push(cars[i])
	}
	return len(stack)
}
