package port

import (
	"github.com/tayalone/ms-jsonplaceholde-todo/core/domain"
	"github.com/tayalone/ms-jsonplaceholde-todo/core/dto"
)

/*ToDoRpstr define Bahavior of ToDo Repository */
type ToDoRpstr interface {
	Create(payload dto.NoteTodo) domain.ToDo
	UpdateByPk(payload dto.UpdateTodo) (domain.ToDo, error)
	DeleteByPk(payload dto.DeleteTodo) error
}

/*ToDoSrvc define Bahavior of ToDo Services */
type ToDoSrvc interface {
	Note(note dto.NoteTodo) domain.ToDo
	UpdateByID(update dto.UpdateTodo) (domain.ToDo, error)
	DeleteByID(del dto.DeleteTodo) error
}
