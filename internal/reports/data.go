package reports

import (
	"database/sql"
	"errors"
)

type ReportsRepository struct {
	db *sql.DB
}

func NewReportsRepository(db *sql.DB) *ReportsRepository {
	return &ReportsRepository{db: db}
}

func (r *ReportsRepository) GetReport(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.Id, &report.Title, &report.Description, &report.CreatedAt, &report.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *ReportsRepository) CreateReport(title, description string) error {
	_, err := r.db.Exec(
		`INSERT INTO reports (title, description, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())`,
		title,
		description)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReportsRepository) UpdateReport(id int, title, description string) error {
	res, err := r.db.Exec(
		`UPDATE reports SET title = $1, description = $2, updated_at = NOW() WHERE id = $3`,
		title,
		description,
		id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("ReportNotFound")
	}

	return nil
}

func (r *ReportsRepository) DeleteReport(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
