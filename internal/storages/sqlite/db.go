package sqllite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Check tasks per day
func (l *LiteDB) CheckTaskPerDay(ctx context.Context, userID sql.NullString, t *storages.Task) (bool, error) {
	now := time.Now()
	stmt := `SELECT COUNT(*) FROM tasks WHERE created_date = ? and user_id = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, now.Format("2006-01-02"), userID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	maxTodo := l.GetMaxToDo(ctx, userID)
	return count <= maxTodo, nil
}

// AddTask adds a new task to DB
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	fmt.Println(userID, pwd)
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row, err := l.DB.Query(stmt, userID, pwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer row.Close()
	u := &storages.User{}
	err = row.Scan(&u.ID)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (l *LiteDB) GetMaxToDo(ctx context.Context, userID sql.NullString) int {
	stmt := `SELECT max_todo FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	var maxTodo int
	err := row.Scan(&maxTodo)
	if err != nil {
		return 0
	}
	return maxTodo
}
