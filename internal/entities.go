package internal

type Grade struct {
	Name             string  `json:"name"`
	FirstEvaluation  float64 `json:"firstEvaluation"`
	SecondEvaluation float64 `json:"secondEvaluation"`
	ThirdEvaluation  float64 `json:"thirdEvaluation"`
}
