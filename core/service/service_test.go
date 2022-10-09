package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/tayalone/ms-jsonplaceholde-todo/core/domain"
	"github.com/tayalone/ms-jsonplaceholde-todo/core/dto"
	"github.com/tayalone/ms-jsonplaceholde-todo/core/port"
)

var (
	createdAt = time.Now()
	updatedAt = time.Now()
)

type mockTodoRepo struct{}

var mockRepo = &mockTodoRepo{}

func (m *mockTodoRepo) Create(payload dto.NoteTodo) domain.ToDo {
	return domain.ToDo{
		ID:        1,
		UserID:    payload.UserID,
		Title:     payload.Title,
		Completed: false,
		CreatedAt: createdAt,
		UpdateAt:  updatedAt,
	}
}

func (m *mockTodoRepo) UpdateByPk(payload dto.UpdateTodo) (domain.ToDo, error) {
	return domain.ToDo{
		ID:        1,
		UserID:    1,
		Title:     *payload.Title,
		Completed: *payload.Completed,
		CreatedAt: createdAt,
		UpdateAt:  updatedAt,
	}, nil
}

func (m *mockTodoRepo) DeleteByPk(payload dto.DeleteTodo) error {
	return nil
}

func TestNew(t *testing.T) {
	type args struct {
		todoRpstr port.ToDoRpstr
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
		{
			name: "Case 1: Create New Services",
			args: args{
				todoRpstr: mockRepo,
			},
			want: New(mockRepo),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.todoRpstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Note(t *testing.T) {
	type fields struct {
		todoRpstr port.ToDoRpstr
	}
	type args struct {
		note dto.NoteTodo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.ToDo
	}{
		// TODO: Add test cases.
		{
			name: "Case 1: Create New Todo Item",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				note: dto.NoteTodo{
					UserID: 1,
					Title:  "Lorem Ipsum",
				},
			},
			want: domain.ToDo{
				ID:        1,
				UserID:    1,
				Title:     "Lorem Ipsum",
				Completed: false,
				CreatedAt: createdAt,
				UpdateAt:  updatedAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				todoRpstr: tt.fields.todoRpstr,
			}
			if got := s.Note(tt.args.note); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Note() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateByID(t *testing.T) {
	type fields struct {
		todoRpstr port.ToDoRpstr
	}
	type args struct {
		update dto.UpdateTodo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.ToDo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				todoRpstr: tt.fields.todoRpstr,
			}
			got, err := s.UpdateByID(tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UpdateByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteByID(t *testing.T) {
	type fields struct {
		todoRpstr port.ToDoRpstr
	}
	type args struct {
		del dto.DeleteTodo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				todoRpstr: tt.fields.todoRpstr,
			}
			if err := s.DeleteByID(tt.args.del); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
