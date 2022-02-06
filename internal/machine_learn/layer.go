package machine_learn

type Layer struct {
	ListNeurons []Neuron
}

func NewHiddenLayer(neuronsCount uint, previousLayer *Layer, learningRate float32) *Layer {
	newLayer := &Layer{ListNeurons: make([]Neuron, 0, neuronsCount)}

	for i := uint(0); i < neuronsCount; i++ {
		newLayer.ListNeurons = append(newLayer.ListNeurons, NewHiddenNeuron(previousLayer, learningRate))
	}

	return newLayer
}

func NewReceptorsLayer(neuronsCount uint) *Layer {
	newLayer := &Layer{ListNeurons: make([]Neuron, 0, neuronsCount)}

	for i := uint(0); i < neuronsCount; i++ {
		newLayer.ListNeurons = append(newLayer.ListNeurons, NewReceptor(float32(ReceptorInitValue)))
	}

	return newLayer
}
