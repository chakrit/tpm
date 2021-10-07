package models

import "time"

type Problem struct {
	ID          int       `db:"id"`
	Summary     string    `db:"summary"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func ListAllProblems() ([]*Problem, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	var result []*Problem
	err = db.Select(&result, `SELECT * FROM problems ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateProblem(p *Problem) (*Problem, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	err = db.Get(p, `
		INSERT INTO problems (summary, description)
		VALUES ($1, $2)
		RETURNING *
	`, p.Summary, p.Description)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetProblem(problemId int) (*Problem, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	prob := &Problem{}
	err = db.Get(prob, `SELECT * FROM problems WHERE id = $1`, problemId)
	if err != nil {
		return nil, err
	}

	return prob, nil
}
