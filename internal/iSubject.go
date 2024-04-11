package internal

type Subject interface {
	// GetResults return the information of the evaluations
	GetResults(grades []Grade)
}
