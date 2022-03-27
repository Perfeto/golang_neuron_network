package main

import (
	"fmt"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"math"
	"neurons/internal/figures"
	"neurons/internal/machine_learn"
	"time"
)

const (
	width       = 1280
	height      = 720
	predictStep = 20
)

func main() {

	wnd, cv, err := sdlcanvas.CreateWindow(width, height, "Neuron Network")
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

	neuronNetwork := machine_learn.NewNetwork(2, 1, 3, 3, 0.1)

	jobChan := make(chan Job, 30_000)
	resultChan := make(chan Result, 30_000)
	for i := 0; i < 8; i++ {
		go DoWork(jobChan, resultChan, func(job Job) Result {
			neuronInputX := float32(job.X) / width * float32(predictStep)
			neuronInputY := float32(job.Y) / height * float32(predictStep)

			predict := neuronNetwork.Predict([]float32{neuronInputX, neuronInputY})

			groupID := figures.FirstGroupID
			if math.Round(float64(predict[0])) != 1 {
				groupID = figures.SecondGroupID
			}

			return Result{
				X:       job.X,
				Y:       job.Y,
				Predict: groupID,
			}
		})
	}

	ticker := time.NewTicker(3 * time.Second)

	go func() {
		for range ticker.C {
			fitNeuronOnDots(neuronNetwork, dotsList)
		}
	}()

	wnd.MainLoop(func() {
		toDoJobs := 0

		for x := 0; x < cv.Width()/predictStep; x++ {
			for y := 0; y < cv.Height()/predictStep; y++ {
				toDoJobs++
				jobChan <- Job{
					X: float32(x),
					Y: float32(y),
				}
			}
		}

		for result := range resultChan {
			toDoJobs--

			figures.PrintPoint(
				figures.Point{
					X:       int(result.X * predictStep),
					Y:       int(result.Y * predictStep),
					GroupID: result.Predict,
				},
				cv,
				18,
			)
			if toDoJobs == 0 {
				break
			}
		}

		figures.PrintPoints(dotsList, cv)
	})
}
