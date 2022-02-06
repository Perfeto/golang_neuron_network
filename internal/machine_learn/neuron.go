package machine_learn

import (
	"math/rand"
)

const (
	LearningRage      = 0.2
	ReceptorInitValue = 0
)

type Tether struct {
	OutgoingNeuron Neuron
	Weight         float32
}

type Neuron interface {
	GetValue() float32
	SetError(errValue float32)
}

type Receptor struct {
	Value float32
}

func (r *Receptor) SetError(_ float32) {
	return
}

func NewReceptor(value float32) *Receptor {
	return &Receptor{Value: value}
}

func (r *Receptor) GetValue() float32 {
	return r.Value
}

func (r *Receptor) SetValue(val float32) {
	r.Value = val
}

type HiddenNeuron struct {
	LearningRate    float32
	IncomingTithers []*Tether
	Value           float32
	ActivationFunc  func(x float32) float32
}

func NewHiddenNeuron(previousLayer *Layer, learningRate float32) *HiddenNeuron {
	neuron := &HiddenNeuron{
		LearningRate:    learningRate,
		IncomingTithers: make([]*Tether, 0, len(previousLayer.ListNeurons)),
		ActivationFunc:  sigmoid,
	}

	for previousNeuronIndex := range previousLayer.ListNeurons {
		neuron.IncomingTithers = append(neuron.IncomingTithers, &Tether{
			OutgoingNeuron: previousLayer.ListNeurons[previousNeuronIndex],
			Weight:         rand.Float32() - 0.5,
		})
	}

	return neuron
}

func (h *HiddenNeuron) GetValue() float32 {
	return h.ActivationFunc(h.GetInputSum())
}

func (h *HiddenNeuron) GetInputSum() float32 {
	valueAccum := float32(0)

	for _, tither := range h.IncomingTithers {
		valueAccum += tither.Weight * tither.OutgoingNeuron.GetValue()
	}

	return valueAccum
}

func (h *HiddenNeuron) SetError(errValue float32) {
	delta := errValue * sigmoidDerivative(h.GetInputSum())

	for inputIndex := range h.IncomingTithers {
		h.IncomingTithers[inputIndex].Weight -= h.IncomingTithers[inputIndex].OutgoingNeuron.GetValue() * delta * h.LearningRate
		(h.IncomingTithers[inputIndex].OutgoingNeuron).SetError(h.IncomingTithers[inputIndex].Weight * delta)
	}
}
