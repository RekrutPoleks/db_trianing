package main

import (
	"db_training/pkg/storage"
	"fmt"
)

const connect = ""

func main(){
	db, err := storage.NewDBStorage(connect)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(db.TasksALL())
	fmt.Println(db.TasksByAuthorID(1))
	fmt.Println(db.UpdateTaskByID(1, "GO_UPDATE"))
	fmt.Println(db.TasksALL())

}
