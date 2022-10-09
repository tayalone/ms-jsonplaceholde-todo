package dto

/*NoteTodo Define Input Value for Create New Todo*/
type NoteTodo struct {
	UserID uint
	Title  string
}

/*UpdateTodo Define Input Value for Update existing Todo*/
type UpdateTodo struct {
	ID        uint
	Title     *string
	Completed *bool
}

/*DeleteTodo Define Input Value for Update existing Todo*/
type DeleteTodo struct {
	ID uint
}
