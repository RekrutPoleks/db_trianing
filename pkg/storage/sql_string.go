package storage

const (
	CreateTask = `INSERT INTO tasks(opened, closed,author_id,assigned_id,tittle,content) 
 						 VALUES ($1, $2, $3, $4, $5, $6)`

	QueryAllTasks        = "SELECT id,opened,closed,author_id,assigned_id,tittle,content FROM tasks;"

	QueryTasksByAuthorID = `SELECT id,opened,closed,author_id,assigned_id,tittle,content 
							FROM tasks WHERE author_id = $1;`

	QueryTasksByAuthor = `SELECT tasks.id,name,opened,closed,tittle,content 
						 FROM users
 						 JOIN tasks ON users.id = tasks.author_id 
						 WHERE users.name = $1;`

	QueryTasksByLabel = `SELECT tasks.id,users.name,opened,closed,tittle,content, lable 
						 FROM tasks
 						 LEFT JOIN tasks_labels as tl ON tasks.id = tl.task_id
                         LEFT JOIN labels ON tl.label_id = labels.id
						 RIGHT JOIN users ON users.id = tasks.author_id
						 WHERE labels.name = $1;`

	UpdateTaskByID = `UPDATE tasks SET content = $1 where id=$2;`

	DeleteTaskByID = `DELETE FROM tasks WHERE id = $1;`
)

