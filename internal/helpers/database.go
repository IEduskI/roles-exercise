package helpers

import (
	"encoding/json"
	"io"
	"os"
	"roles-exercise/internal"
)

func GetData() []internal.Grade {
	db, _ := os.Open("evaluation.json")
	defer db.Close()

	gradesB, _ := io.ReadAll(db)

	var grades []internal.Grade

	_ = json.Unmarshal(gradesB, &grades)

	return grades
}
