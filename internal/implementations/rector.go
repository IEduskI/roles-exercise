package implementations

import (
	"fmt"
	"roles-exercise/internal"
	"roles-exercise/internal/helpers"
)

type RectorImpl struct {
	user internal.User
}

func NewRectorImpl(user internal.User) *RectorImpl {
	return &RectorImpl{
		user: user,
	}
}

func (r *RectorImpl) GetResults(grades []internal.Grade) {
	fmt.Printf("Name || firstEvaluation || secondEvaluation || thirdEvaluation || Subject Status\n")
	for _, grade := range grades {
		fmt.Printf("%v    ||   %v  ||  %v  ||  %v  || %v\n", grade.Name, grade.FirstEvaluation, grade.SecondEvaluation, grade.ThirdEvaluation, helpers.GetSubjectStatus(grade))
	}

	fEval, sEval, tEval := helpers.GetAverages(grades)
	fmt.Println("\nGeneral Averages:")
	fmt.Printf("First Evaluation: %.2f\n", fEval)
	fmt.Printf("Second Evaluation: %.2f\n", sEval)
	fmt.Printf("Third Evaluation: %.2f\n", tEval)

	fEvalHS, sEvalHS, tEvalHS := helpers.GetHigScores(grades)
	fmt.Println("\nHigh Scores:")
	fmt.Printf("First Evaluation: %v\n", fEvalHS)
	fmt.Printf("Second Evaluation: %v\n", sEvalHS)
	fmt.Printf("Third Evaluation: %v\n", tEvalHS)
}
