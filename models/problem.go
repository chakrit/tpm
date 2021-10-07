package models

type Problem struct {
	ID          int    `db:"id"`
	Summary     string `db:"summary"`
	Description string `db:"description"`
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
