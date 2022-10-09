package port

import "github.com/tayalone/ms-jsonplaceholde-todo/core/domain"

/*ToDoRpstr define Bahavior of ToDo Repository */
type ToDoRpstr interface {
	Create(userID uint, title string) domain.ToDo
	UpdateByPk(id uint, title string, completed bool) (domain.ToDo, error)
	DeleteByPk(id uint) error
}

/*ToDoSrvc define Bahavior of ToDo Services */
type ToDoSrvc interface {
	Note(userID uint, title string) domain.ToDo
	UpdateByID(id uint, title string, completed bool) (domain.ToDo, error)
	DeleteByID(id uint) error
}
