package helpers

import (
	"roles-exercise/internal"
)

const (
	PASSED = "Passed"
	FAILED = "Failed"
)

func GetHigScores(grades []internal.Grade) (fEval, sEval, tEval string) {
	var highGradeF, highGradeS, highGradeT float64
	for _, grade := range grades {
		if grade.FirstEvaluation > highGradeF {
			highGradeF = grade.FirstEvaluation
			fEval = grade.Name
		}

		if grade.SecondEvaluation > highGradeS {
			highGradeS = grade.SecondEvaluation
			sEval = grade.Name
		}

		if grade.ThirdEvaluation > highGradeT {
			highGradeT = grade.ThirdEvaluation
			tEval = grade.Name
		}
	}
	return
}

func GetSubjectStatus(grade internal.Grade) string {
	avg := (grade.FirstEvaluation * .35) + (grade.SecondEvaluation * .325) + (grade.ThirdEvaluation * .325)

	if avg < 6.0 {
		return FAILED
	}

	return PASSED
}

func GetAverages(grades []internal.Grade) (fEval, sEval, tEval float64) {
	var totalGrades float64
	for _, grade := range grades {
		fEval = fEval + grade.FirstEvaluation
		sEval = sEval + grade.SecondEvaluation
		tEval = tEval + grade.ThirdEvaluation
		totalGrades++
	}

	fEval = fEval / totalGrades
	sEval = sEval / totalGrades
	tEval = tEval / totalGrades

	return
}
