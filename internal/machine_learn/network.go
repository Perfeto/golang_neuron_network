package machine_learn

type Network struct {
	Layers []*Layer
}

func NewNetwork(inputSize, outputSize, hiddenLayersCount, layersSize uint, learningRate float32) *Network {
	newNetwork := &Network{Layers: make([]*Layer, 0, hiddenLayersCount+2)}

	newNetwork.Layers = append(newNetwork.Layers, NewReceptorsLayer(inputSize))

	for i := uint(0); i < hiddenLayersCount; i++ {
		newNetwork.Layers = append(newNetwork.Layers, NewHiddenLayer(layersSize, newNetwork.Layers[len(newNetwork.Layers)-1], learningRate))
	}

	newNetwork.Layers = append(newNetwork.Layers, NewHiddenLayer(outputSize, newNetwork.Layers[len(newNetwork.Layers)-1], learningRate))

	return newNetwork
}

func (n *Network) GetPrediction() []float32 {
	outputNeurons := n.Layers[len(n.Layers)-1].ListNeurons

	predict := make([]float32, 0, len(outputNeurons))
	for _, neuron := range outputNeurons {
		predict = append(predict, neuron.GetValue())
	}

	return predict
}

func (n *Network) SetInputValues(inputs []float32) {
	for neuronIndex, neuron := range n.Layers[0].ListNeurons {
		neuron.(*Receptor).SetValue(inputs[neuronIndex])
	}
}

type TrainExercise struct {
	IncomeValues   []float32
	OutgoingValues []float32
}

func (n *Network) TrainOnce(trainSet []TrainExercise) {
	for _, exercise := range trainSet {
		n.SetInputValues(exercise.IncomeValues)

		for i, predict := range n.GetPrediction() {
			n.Layers[len(n.Layers)-1].ListNeurons[i].(*HiddenNeuron).SetError(predict - exercise.OutgoingValues[i])
		}
	}
}

func (n *Network) Fit(trainSet []TrainExercise, epochs uint) {
	for i := uint(0); i < epochs; i++ {
		n.TrainOnce(trainSet)
	}
}

func (n *Network) Predict(inputs []float32) []float32 {
	n.SetInputValues(inputs)

	return n.GetPrediction()
}
