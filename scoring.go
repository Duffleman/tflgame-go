package tflgame

type Calculations struct {
	Start  float64            `json:"start"`
	Base   float64            `json:"base"`
	End    float64            `json:"end"`
	Final  int                `json:"final"`
	Events []CalculationEvent `json:"events"`
}

type CalculationEvent struct {
	Item   *CalculationItem `json:"item"`
	Effect *string          `json:"effect"`
	Score  float64          `json:"score"`
	Note   string           `json:"note"`
}

type CalculationItem struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (c *Calculations) Add(s CalculationEvent) {
	c.Events = append(c.Events, s)
}
