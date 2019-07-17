package kata

type ContainerOfWaterCalculator interface {
	maxArea(height []int) int
}

type ContainerOfWater struct {}

func (c ContainerOfWater) maxArea(height []int) (area int) {
	maxVal := 0

	for x, y := range height {
		x++
		for k, l := range height {
			k++
			if y == l && x == k {
				continue
			}

			width := k - x
			height := 0

			if (y < l) {
				height = y
			} else {
				height = l
			}

			volume := width * height

			if (volume > maxVal) {
				maxVal = volume
			}
		}
	}
	return maxVal
}
