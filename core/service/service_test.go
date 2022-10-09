package service

import (
	"errors"
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
	if payload.ID == 112 {
		return domain.ToDo{}, errors.New("Todo Does not exist")
	}

	if payload.Title != nil && payload.Completed != nil {
		return domain.ToDo{
			ID:        1,
			UserID:    1,
			Title:     *payload.Title,
			Completed: *payload.Completed,
			CreatedAt: createdAt,
			UpdateAt:  updatedAt,
		}, nil
	}

	if payload.Title != nil {
		return domain.ToDo{
			ID:        1,
			UserID:    1,
			Title:     *payload.Title,
			Completed: false,
			CreatedAt: createdAt,
			UpdateAt:  updatedAt,
		}, nil
	}

	if payload.Completed != nil {
		return domain.ToDo{
			ID:        1,
			UserID:    1,
			Title:     "title not change",
			Completed: *payload.Completed,
			CreatedAt: createdAt,
			UpdateAt:  updatedAt,
		}, nil
	}

	return domain.ToDo{}, errors.New("Something went wrong")
}

func (m *mockTodoRepo) DeleteByPk(payload dto.DeleteTodo) error {
	if payload.ID == 112 {
		return errors.New("todo doesn't existing")
	}

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
	titleBody := "I wanna Change Todo Title"
	completedBody := true

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
		{
			name: "Case 1: Change Both Title & Completed",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				update: dto.UpdateTodo{
					ID:        1,
					Title:     &titleBody,
					Completed: &completedBody,
				},
			},
			want: domain.ToDo{
				ID:        1,
				UserID:    1,
				Title:     titleBody,
				Completed: completedBody,
				CreatedAt: createdAt,
				UpdateAt:  updatedAt,
			},
			wantErr: false,
		},
		{
			name: "Case 2: Change Only Title",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				update: dto.UpdateTodo{
					ID:        1,
					Title:     &titleBody,
					Completed: nil,
				},
			},
			want: domain.ToDo{
				ID:        1,
				UserID:    1,
				Title:     titleBody,
				Completed: false,
				CreatedAt: createdAt,
				UpdateAt:  updatedAt,
			},
			wantErr: false,
		},
		{
			name: "Case 3: Change Only Completed",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				update: dto.UpdateTodo{
					ID:        1,
					Title:     nil,
					Completed: &completedBody,
				},
			},
			want: domain.ToDo{
				ID:        1,
				UserID:    1,
				Title:     "title not change",
				Completed: completedBody,
				CreatedAt: createdAt,
				UpdateAt:  updatedAt,
			},
			wantErr: false,
		},
		{
			name: "Case 4: Todo Does not exist",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				update: dto.UpdateTodo{
					ID:        112,
					Title:     &titleBody,
					Completed: &completedBody,
				},
			},
			want:    domain.ToDo{},
			wantErr: true,
		},
		{
			name: "Case 5: Payload are not contain Title & Complete",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				update: dto.UpdateTodo{
					ID:        1,
					Title:     nil,
					Completed: nil,
				},
			},
			want:    domain.ToDo{},
			wantErr: true,
		},
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
		{
			name: "Case 1: Delete Exising Todo",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				del: dto.DeleteTodo{
					ID: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "Case 2: Delete Does not Exising Todo",
			fields: fields{
				todoRpstr: mockRepo,
			},
			args: args{
				del: dto.DeleteTodo{
					ID: 112,
				},
			},
			wantErr: true,
		},
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
