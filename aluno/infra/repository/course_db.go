package repository

import (
	"aluno/entity"
	"database/sql"
)

type CourseMySQLRepository struct {
	Db *sql.DB
}

func (c CourseMySQLRepository) Insert(course entity.Course) error {
	stmt, err := c.Db.Prepare(`INSERT into courses(id,name,description,status) values (?,?,?,?,)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	return nil
}
