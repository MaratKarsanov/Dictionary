package words

import (
	"database/sql"
	"errors"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetWordById ищем слово по id
func (r *Repo) RGetWordById(id int) (*Word, error) {
	var word Word
	err := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id = $1`, id).
		Scan(&word.Id, &word.Title, &word.Translation)
	if err != nil {
		return nil, err
	}

	return &word, nil
}

// CreateNewWords добавляет новые переводы в базу даных
func (r *Repo) CreateNewWords(word, translate string) error {
	_, err := r.db.Exec(`INSERT INTO ru_en (title, translation) VALUES ($1, $2)`, word, translate)
	if err != nil {
		return err
	}

	return nil
}

// UpdateWord обновляет слово
func (r *Repo) UpdateWord(id int, word, translate string) error {
	res, err := r.db.Exec(`UPDATE ru_en SET title = $1, translation = $2 WHERE id = $3`, word, translate, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("WordNotFound")
	}

	return nil
}

// DeleteWordById удаляет слово по id
func (r *Repo) DeleteWordById(id int) error {
	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) SearchWords(query string) ([]Word, error) {
	rows, err := r.db.Query(`
        SELECT id, title, translation
        FROM ru_en
        ORDER BY similarity(title, $1) DESC
        LIMIT 100;
    `, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Word
	for rows.Next() {
		var word Word
		err := rows.Scan(&word.Id, &word.Title, &word.Translation)
		if err != nil {
			return nil, err
		}
		result = append(result, word)
	}
	return result, rows.Err()
}
