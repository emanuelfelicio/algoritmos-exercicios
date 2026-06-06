package algoritmosexercicios

func maxArea(heights []int) int {

	left, right := 0, len(heights)-1
	area := 0
	for left < right {
		leftV, rightV := heights[left], heights[right]
		distance := right - left
		currentArea := 0
		if leftV <= rightV {
			currentArea = leftV * distance
			left++
		} else {
			currentArea = rightV * distance
			right--
		}

		if currentArea > area {
			area = currentArea
		}
	}
	return area
}
