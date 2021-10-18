package repo

import "github.com/mahjadan/go/grpc-demo/pkg/grpc"

type Store interface {
	GetAll() grpc.TaskList
	Save(task grpc.Task)
}

type InMemoryStore struct {
	m map[string]grpc.Task
}

func NewInMemoryStore() Store {
	return &InMemoryStore{
		m: make(map[string]grpc.Task),
	}
}

func (i *InMemoryStore) GetAll() grpc.TaskList {
	var list []*grpc.Task
	for _, task := range i.m {
		list = append(list, &task)
	}
	return grpc.TaskList{
		Tasks: list,
	}
}

func (i *InMemoryStore) Save(task grpc.Task) {
	i.m[task.Name] = grpc.Task{
		Name: task.GetName(),
		Done: task.GetDone(),
	}
}
