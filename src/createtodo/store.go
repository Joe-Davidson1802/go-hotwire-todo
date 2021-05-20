package main

import (
	"context"

	"cloud.google.com/go/datastore"
)

type TodoStore struct {
	proj string
	kind string
}

func NewTodoStore(proj string, kind string) TodoStore {
	return TodoStore{
		proj: proj,
		kind: kind,
	}
}

func (t *TodoStore) PutTodo(ctx context.Context, r Todo) (*Todo, error) {
	dsClient, err := datastore.NewClient(ctx, t.proj)
	if err != nil {
		return nil, err
	}

	defer dsClient.Close()

	k := datastore.NameKey(t.kind, "stringID", nil)

	_, err = dsClient.Put(ctx, k, r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}
