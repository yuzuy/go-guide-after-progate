package db

import (
	"errors"

	"github.com/yuzuy/go-guide-after-progate/todo"
)

type DB struct {
	m map[string]*todo.Task
}

func New() *DB {
	return &DB{
		m: make(map[string]*todo.Task),
	}
}

func (db *DB) FindTasks() []*todo.Task {
	tasks := make([]*todo.Task, 0, len(db.m))

	for _, v := range db.m {
		tasks = append(tasks, v)
	}

	return tasks
}

func (db *DB) AddTask(task *todo.Task) error {
	if task == nil {
		return errors.New("this task is nil")
	}

	if _, ok := db.m[task.ID]; ok {
		return errors.New("this task already added")
	}

	db.m[task.ID] = task

	return nil
}

func (db *DB) UpdateTask(id, name string, isDone bool) error {
	task, ok := db.m[id]
	if !ok {
		return errors.New("this task not found")
	}

	task.Name = name
	task.IsDone = isDone

	return nil
}

func (db *DB) RemoveTask(id string) error {
	if _, ok := db.m[id]; !ok {
		return errors.New("this task not found")
	}

	delete(db.m, id)

	return nil
}
