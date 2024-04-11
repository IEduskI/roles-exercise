package implementations

import (
	"fmt"
	"roles-exercise/internal"
)

type InstructorImpl struct {
	user internal.User
}

func NewInstructorImpl(user internal.User) *InstructorImpl {
	return &InstructorImpl{
		user: user,
	}
}

func (i *InstructorImpl) GetResults(grades []internal.Grade) {
	fmt.Printf("Name   ||  firstEvaluation  ||  secondEvaluation  ||  thirdEvaluation\n")
	for _, grade := range grades {
		fmt.Printf("%v    ||   %v  ||  %v  ||  %v\n", grade.Name, grade.FirstEvaluation, grade.SecondEvaluation, grade.ThirdEvaluation)
	}
}
