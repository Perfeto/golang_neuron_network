package main

import (
	"fmt"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"math"
	"neurons/internal/figures"
	"neurons/internal/machine_learn"
)

const (
	width  = 1280
	height = 720
)

func main() {

	wnd, cv, err := sdlcanvas.CreateWindow(width, height, "Hello")
	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()

	var dotsList []figures.Point

	wnd.MouseDown = func(button, x, y int) {
		fmt.Println(button, x, y)

		if button == 1 {
			dotsList = append(dotsList, figures.Point{X: x, Y: y, GroupID: figures.FirstGroupID})
		} else if button == 3 {
			dotsList = append(dotsList, figures.Point{X: x, Y: y, GroupID: figures.SecondGroupID})
		}
	}

	neuronNetwork := machine_learn.NewNetwork(2, 1, 1, 0.4)

	wnd.MainLoop(func() {
		if len(dotsList) > 0 {
			fitNeuronOnDots(neuronNetwork, dotsList)
		}

		predictStep := 15
		for x := 0; x < cv.Width()/predictStep; x++ {
			for y := 0; y < cv.Height()/predictStep; y++ {
				neuronInputX := float32(x) / width * float32(predictStep)
				neuronInputY := float32(y) / height * float32(predictStep)

				predict := neuronNetwork.Predict([]float32{neuronInputX, neuronInputY})

				groupID := figures.FirstGroupID
				if math.Round(float64(predict[0])) != 1 {
					groupID = figures.SecondGroupID
				}

				figures.PrintPoint(
					figures.Point{
						X:       x * predictStep,
						Y:       y * predictStep,
						GroupID: groupID,
					},
					cv,
					18,
				)
			}
		}

		figures.PrintPoints(dotsList, cv)
	})
}
