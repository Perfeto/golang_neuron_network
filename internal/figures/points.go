package figures

import "github.com/tfriedel6/canvas"

type GroupID uint

const (
	FirstGroupID  GroupID = 1
	SecondGroupID         = 2

	circleBorderSize = 15
	circleRadius     = 30
)

type Point struct {
	X, Y    int
	GroupID GroupID
}

func PrintPoints(dotsList []Point, cv *canvas.Canvas) {
	for _, point := range dotsList {
		PrintPoint(point, cv, circleRadius)
	}
}

func PrintPoint(point Point, cv *canvas.Canvas, circleRadius float64) {
	cv.SetFillStyle("#FFFFFF")
	cv.BeginPath()
	cv.Ellipse(float64(point.X*2), float64(point.Y*2), circleRadius+(circleRadius/10), circleRadius+(circleRadius/10), 0, 0, 6, false)
	cv.ClosePath()
	cv.Fill()

	cv.SetFillStyle("#52ed09")
	if point.GroupID == FirstGroupID {
		cv.SetFillStyle("#0346ff")
	}

	cv.BeginPath()
	cv.Ellipse(float64(point.X*2), float64(point.Y*2), circleRadius, circleRadius, 0, 0, 6, false)
	cv.ClosePath()
	cv.Fill()
}
