package todos

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{ID: "1", Title: "Clean Room", Completed: false},
	{ID: "2", Title: "Read Book", Completed: false},
	{ID: "3", Title: "Record Song", Completed: false},
	{ID: "4", Title: "Drink coffee", Completed: false},
}
