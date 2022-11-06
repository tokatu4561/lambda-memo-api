package usecases

import "github.com/tokatu4561/tasks/pkg/domain"

type TaskUsecase struct {
	Repository domain.TaskRepositoryInterface
}

func NewTaskUsecase(taskRepo domain.TaskRepositoryInterface) domain.TaskUseCaseInterface {
	return &TaskUsecase{
		Repository: taskRepo,
	}
}

func (t *TaskUsecase) GetTasks() ([]*domain.Task, error) {
	tasks, err := t.Repository.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (t *TaskUsecase) AddTask(task *domain.Task) (*domain.Task, error) {
	newTask, err := t.Repository.AddTask(task)

	if err != nil {
		return nil, err
	}

	return newTask, nil
}
