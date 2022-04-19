package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type dbStorage struct {
	db *pgxpool.Pool
}

type Task struct {
 	id, opened, closed, aut_id, assg_id int
 	tittle, content string
}



func NewDBStorage(conn string) (*dbStorage, error) {
	db, err := pgxpool.Connect(context.Background(),conn)
	if err != nil{
		return nil, err
	}
	if err = db.Ping(context.Background()); err != nil{
		return nil, err
	}
	return &dbStorage{db}, err
}

func (ds *dbStorage) NewTask(task Task) error{
	_, err := ds.db.Exec(
		context.Background(),
		CreateTask,
		task.opened,
		task.closed,
		task.aut_id,
		task.assg_id,
		task.tittle,
		task.content,
		)
	if err != nil{
		return err
	}
	return nil
}

func (ds dbStorage) TasksByAuthorID(id int) ([]Task, error) {
	rows, err := ds.db.Query(
		context.Background(),
		QueryTasksByAuthorID,
		id,
	)
	if err != nil{
		return nil, err
	}

	var task Task
	var tasks []Task
	for rows.Next(){
		err = rows.Scan(
			&task.id,
			&task.opened,
			&task.closed,
			&task.aut_id,
			&task.assg_id,
			&task.tittle,
			&task.content,
			)
		if err != nil{
			return nil, err
		}
		tasks = append(tasks, task)

	}
	return tasks, rows.Err()
}




func (ds dbStorage) TasksALL() ([]Task, error) {
	rows, err := ds.db.Query(
		context.Background(),
		QueryAllTasks,
	)
	if err != nil{
		return nil, err
	}

	var task Task
	var tasks []Task
	for rows.Next(){
		err = rows.Scan(
			&task.id,
			&task.opened,
			&task.closed,
			&task.aut_id,
			&task.assg_id,
			&task.tittle,
			&task.content,
		)
		if err != nil{
			return nil, err
		}
		tasks = append(tasks, task)

	}
	return tasks, rows.Err()
}
 type TaskByAuth struct {
	 id int
	 name string
	 opened int
	 closed int
	 tittle string
	 content string
 }

func (ds dbStorage) TasksByAuth(name string) ([]TaskByAuth, error) {
	rows, err := ds.db.Query(
		context.Background(),
		QueryTasksByAuthor,
		name,
	)
	if err != nil{
		return nil, err
	}

	var task TaskByAuth
	var tasks []TaskByAuth
	for rows.Next(){
		err = rows.Scan(
			&task.id,
			&task.name,
			&task.opened,
			&task.closed,
			&task.tittle,
			&task.content,
		)
		if err != nil{
			return nil, err
		}
		tasks = append(tasks, task)

	}
	return tasks, rows.Err()
}

type TasksBylabel struct{
	id_task int
	name_auth string
	opened, closed int
	tittle string
	content string
	label string
}

func (ds dbStorage) TasksByLabel(label string) ([]TasksBylabel, error) {
	rows, err := ds.db.Query(
		context.Background(),
		QueryTasksByLabel,
		label,
	)
	if err != nil{
		return nil, err
	}

	var task TasksBylabel
	var tasks []TasksBylabel
	for rows.Next(){
		err = rows.Scan(
			&task.id_task,
			&task.name_auth,
			&task.opened,
			&task.closed,
			&task.tittle,
			&task.content,
			&task.label,
		)
		if err != nil{
			return nil, err
		}
		tasks = append(tasks, task)

	}
	return tasks, rows.Err()
}

func (ds *dbStorage)UpdateTaskByID(id int, content string) error{
	_, err := ds.db.Exec(context.Background(), UpdateTaskByID, content, id)
	if err != nil{
		return err
	}
	return nil
}

func (ds *dbStorage)DeleteTaskByID(id int) error{
	_, err := ds.db.Exec(context.Background(), DeleteTaskByID, id)
	if err != nil{
		return err
	}
	return nil
}

