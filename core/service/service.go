package service

import (
	"errors"

	"github.com/tayalone/ms-jsonplaceholde-todo/core/domain"
	"github.com/tayalone/ms-jsonplaceholde-todo/core/dto"
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
func (s *Service) Note(note dto.NoteTodo) domain.ToDo {
	return s.todoRpstr.Create(note)
}

/*UpdateByID update Todo data with Id*/
func (s *Service) UpdateByID(update dto.UpdateTodo) (domain.ToDo, error) {
	if update.Title == nil && update.Completed == nil {
		return domain.ToDo{}, errors.New("title and completed do not nil @ the sametime")
	}
	updatedTodo, err := s.todoRpstr.UpdateByPk(update)
	if err != nil {
		return domain.ToDo{}, err
	}
	return updatedTodo, nil
}

/*DeleteByID do Deleted By Id*/
func (s *Service) DeleteByID(del dto.DeleteTodo) error {
	err := s.todoRpstr.DeleteByPk(del)
	if err != nil {
		return err
	}
	return nil
}
