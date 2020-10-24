package tflgame

type Calculations struct {
	Start  float64            `json:"start"`
	Base   float64            `json:"base"`
	End    float64            `json:"end"`
	Final  int                `json:"final"`
	Events []CalculationEvent `json:"events"`
}

type CalculationEvent struct {
	PromptID *string `json:"prompt_id"`
	Effect   *string `json:"effect"`
	Score    float64 `json:"score"`
	Note     string  `json:"note"`
}

func (c *Calculations) Add(s CalculationEvent) {
	c.Events = append(c.Events, s)
}
