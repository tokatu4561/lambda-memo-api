package infrastructure

import (
	"database/sql"
)

type TaskRepositoryGateway struct {
	databaseHandler *sql.DB
}

type DatabaseHandler struct {
	Conn *sql.DB
}

// func NewTaskRepository(db *sql.DB) domain.TaskRepositoryInterface {
// 	return &TaskRepositoryGateway{
// 		databaseHandler: db,
// 	}
// }

// func (t *TaskRepositoryGateway) AddTask(task *domain.Task) (*domain.Task, error) {
// 	newTask, err := Insert(t.databaseHandler, task)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newTask, nil
// }

// func (t *TaskRepositoryGateway) GetTasks() ([]*domain.Task, error) {
// 	tasks, err := GetAll(t.databaseHandler)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return tasks, nil
// }

// func Insert(db *sql.DB, task *domain.Task) (*domain.Task, error) {
// 	stmt := `insert into tasks (user_id, title, created_at, updated_at)
// 		values ($1, $2, $3, $4) returning id`

// 	_, err := db.Exec(stmt,
// 		task.UserID,
// 		task.Title,
// 		time.Now(),
// 		time.Now(),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil, nil
// }

// // GetALl returns all tasks in db
// func GetAll(db *sql.DB) ([]*domain.Task, error) {
// 	query := `select id, user_id, title, created_at, updated_at from tasks`

// 	var tasks []*domain.Task
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var task domain.Task
// 		err := rows.Scan(
// 			&task.ID,
// 			&task.UserID,
// 			&task.Title,
// 			&task.CreatedAt,
// 			&task.UpdatedAt,
// 		)
// 		if err != nil {
// 			log.Println("Error scanning", err)
// 			return nil, err
// 		}

// 		tasks = append(tasks, &task)
// 	}

// 	return tasks, nil
// }
