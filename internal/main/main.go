package main

import (
	"fmt"
	"math/rand"
	"neurons/internal/machine_learn"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	neuronsNetwork := machine_learn.NewNetwork(2, 1, 2, 0.2)

	neuronsNetwork.Fit([]machine_learn.TrainExercise{
		{
			IncomeValues:   []float32{0, 0},
			OutgoingValues: []float32{0},
		},
		{
			IncomeValues:   []float32{0, 1},
			OutgoingValues: []float32{1},
		},
		{
			IncomeValues:   []float32{1, 0},
			OutgoingValues: []float32{1},
		},
		{
			IncomeValues:   []float32{1, 1},
			OutgoingValues: []float32{0},
		},
	}, 1_000_000)

	fmt.Printf("%d XOR %d = %f \n", 0, 0, neuronsNetwork.Predict([]float32{0, 0})[0])
	fmt.Printf("%d XOR %d = %f \n", 0, 1, neuronsNetwork.Predict([]float32{0, 1})[0])
	fmt.Printf("%d XOR %d = %f \n", 1, 0, neuronsNetwork.Predict([]float32{1, 0})[0])
	fmt.Printf("%d XOR %d = %f \n", 1, 1, neuronsNetwork.Predict([]float32{1, 1})[0])
}
