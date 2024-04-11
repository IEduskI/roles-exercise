package main

import (
	"fmt"
	"roles-exercise/internal"
	"roles-exercise/internal/helpers"
	"roles-exercise/internal/implementations"
	"strings"
)

func main() {
	instructor := internal.NewUser("instructor")
	professor := internal.NewUser("professor")
	rector := internal.NewUser("rector")

	instructorImpl := implementations.NewInstructorImpl(instructor)
	professorImpl := implementations.NewProfessorImpl(professor)
	rectorImpl := implementations.NewRectorImpl(rector)

	// Insert type of user
	user := "rector"

	grades := helpers.GetData()

	switch {
	case strings.EqualFold(user, instructor.GetUsername()):
		instructorImpl.GetResults(grades)
	case strings.EqualFold(user, professor.GetUsername()):
		professorImpl.GetResults(grades)
	case strings.EqualFold(user, rector.GetUsername()):
		rectorImpl.GetResults(grades)
	default:
		fmt.Println("invalid username")
	}

}
