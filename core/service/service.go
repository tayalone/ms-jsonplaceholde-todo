package service

import (
	"github.com/tayalone/ms-jsonplaceholde-todo/core/domain"
	"github.com/tayalone/ms-jsonplaceholde-todo/core/port"
)

/*Service Define variables which Services need to do business Req */
type Service struct {
	todoRpstr port.ToDoRpstr
}

var srv *Service = nil

/*New is Return *Service as Singleton */
func New(todoRpstr port.ToDoRpstr) *Service {
	srv = &Service{
		todoRpstr: todoRpstr,
	}
	return srv
}

/*Note do note somedata to database*/
func (s *Service) Note(userID uint, title string) domain.ToDo {
	return s.todoRpstr.Create(userID, title)
}

/*UpdateByID update Todo data with Id*/
func (s *Service) UpdateByID(id uint, title string, completed bool) (domain.ToDo, error) {
	updatedTodo, err := s.todoRpstr.UpdateByPk(id, title, completed)
	if err != nil {
		return domain.ToDo{}, err
	}
	return updatedTodo, nil
}

/*DeleteByID do Deleted By Id*/
func (s *Service) DeleteByID(id uint, title string, completed bool) (domain.ToDo, error) {
	updatedTodo, err := s.todoRpstr.UpdateByPk(id, title, completed)
	if err != nil {
		return domain.ToDo{}, err
	}
	return updatedTodo, nil
}
