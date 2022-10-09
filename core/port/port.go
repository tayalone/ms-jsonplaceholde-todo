package port

import (
	"github.com/tayalone/ms-jsonplaceholde-todo/core/domain"
	"github.com/tayalone/ms-jsonplaceholde-todo/core/dto"
)

/*ToDoRpstr define Bahavior of ToDo Repository */
type ToDoRpstr interface {
	Create(dto.NoteTodo) domain.ToDo
	UpdateByPk(dto.UpdateTodo) (domain.ToDo, error)
	DeleteByPk(dto.DeleteTodo) error
}

/*ToDoSrvc define Bahavior of ToDo Services */
type ToDoSrvc interface {
	Note(dto.NoteTodo) domain.ToDo
	UpdateByID(dto.UpdateTodo) (domain.ToDo, error)
	DeleteByID(dto.DeleteTodo) error
}
