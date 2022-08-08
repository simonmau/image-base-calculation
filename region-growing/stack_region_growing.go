package regiongrowing

import (
	"errors"

	"github.com/simonmau/spacial-base-calculation/point"
)

type StackRegionGrowing struct {
	width  int32
	height int32
	size   int32

	startPts []point.T
}

func GenStackRegionGrowing(width, height int32, startPoints *[]point.T) StackRegionGrowing {
	return StackRegionGrowing{
		width:    width,
		height:   height,
		size:     width * height,
		startPts: *startPoints,
	}
}

func (s *StackRegionGrowing) FindAllWithStartPoints(image *[]bool) (*[][]point.T, error) {
	img := *image

	if len(img) != int(s.size) {
		return nil, errors.New("wrong size for stack-region-growing")
	}

	searchedPoints := make([]bool, s.size)
	foundRegions := make([][]point.T, 0)

	for _, startPt := range s.startPts {
		currentRegion := make([]point.T, 0)

		s.growRegionWithStartPoint(image, &searchedPoints, &startPt, &currentRegion)

		if len(currentRegion) > 0 {
			foundRegions = append(foundRegions, currentRegion)
		}
	}

	return &foundRegions, nil
}

//expects already checked data&searched arrays as an empty foundpoints array to push the results in
func (s *StackRegionGrowing) growRegionWithStartPoint(image, searchedPoints *[]bool, startPoint *point.T, foundPoints *[]point.T) {
	img := *image

	if !startPoint.InRange(&s.width, &s.height) || !img[startPoint.ToIndex(&s.width, 1)] {
		return
	}

	stack := Stack{}
	stack.Push(startPoint)

	sp := *searchedPoints

	for len(stack) > 0 {
		pt, _ := stack.Pop()

		neighbours := []point.T{
			{pt[0] + 1, pt[1] + 1},
			{pt[0] + 1, pt[1]},
			{pt[0] + 1, pt[1] - 1},
			{pt[0], pt[1] + 1},
			// {pt[0], pt[1]}, //do not include center point
			{pt[0], pt[1] - 1},
			{pt[0] - 1, pt[1] + 1},
			{pt[0] - 1, pt[1]},
			{pt[0] - 1, pt[1] - 1},
		}

		for _, neighbourPoint := range neighbours {
			if !neighbourPoint.InRange(&s.width, &s.height) {
				continue
			}

			index := neighbourPoint.ToIndex(&s.width, 1)

			if !img[index] || sp[index] {
				continue
			}

			sp[index] = true
			*foundPoints = append(*foundPoints, neighbourPoint)

			stack.Push(&neighbourPoint)
		}
	}
}
