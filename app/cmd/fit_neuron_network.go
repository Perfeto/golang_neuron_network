package main

import (
	"neurons/internal/figures"
	"neurons/internal/machine_learn"
)

func fitNeuronOnDots(neuronNetwork *machine_learn.Network, dots []figures.Point) {
	trainSet := make([]machine_learn.TrainExercise, 0, len(dots))

	for _, point := range dots {
		neuronInputX := float32(point.X) / width
		neuronInputY := float32(point.Y) / height

		resultTeam := 0
		if point.GroupID == figures.SecondGroupID {
			resultTeam = 1
		}

		trainSet = append(trainSet, machine_learn.TrainExercise{
			IncomeValues:   []float32{neuronInputX, neuronInputY},
			OutgoingValues: []float32{float32(resultTeam)},
		})
	}

	neuronNetwork.Fit(trainSet, 10_000)
}
