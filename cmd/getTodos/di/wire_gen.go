// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/tokatu4561/tasks/pkg/application"
	"github.com/tokatu4561/tasks/pkg/infrastructure/dynamo"
	"github.com/tokatu4561/tasks/pkg/usecases"
)

// Injectors from wire.go:

func NewTaskController() *application.TaskController {
	db := dynamo.NewDynamoDatabaseHandler()
	taskRepositoryInterface := dynamo.NewTaskRepository(db)
	taskUseCaseInterface := usecases.NewTaskUsecase(taskRepositoryInterface)
	taskController := application.NewTaskController(taskUseCaseInterface)
	return taskController
}