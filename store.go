package main

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/joe-davidson1802/go-hotwire-site/models"
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

func (t *TodoStore) PostTodo(ctx context.Context, r *models.Todo) error {
	dsClient, err := datastore.NewClient(ctx, t.proj)
	if err != nil {
		return err
	}

	defer dsClient.Close()

	k := datastore.IncompleteKey(t.kind, nil)

	id, err := dsClient.Put(ctx, k, r)

	r.ID = id

	if err != nil {
		return err
	}

	return nil
}

func (t *TodoStore) DeleteTodo(ctx context.Context, id int64) error {
	dsClient, err := datastore.NewClient(ctx, t.proj)
	if err != nil {
		return err
	}

	defer dsClient.Close()

	k := datastore.IDKey(t.kind, id, nil)

	err = dsClient.Delete(ctx, k)

	if err != nil {
		return err
	}

	return nil
}

func (t *TodoStore) GetTodo(ctx context.Context, id int64) (*models.Todo, error) {
	dsClient, err := datastore.NewClient(ctx, t.proj)
	if err != nil {
		return nil, err
	}

	defer dsClient.Close()

	k := datastore.IDKey(t.kind, id, nil)

	var r models.Todo

	err = dsClient.Get(ctx, k, &r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (t *TodoStore) GetTodos(ctx context.Context, max int) (*[]models.Todo, error) {
	dsClient, err := datastore.NewClient(ctx, t.proj)
	if err != nil {
		return nil, err
	}

	defer dsClient.Close()

	var rs []models.Todo

	q := datastore.NewQuery(t.kind).Limit(max)

	_, err = dsClient.GetAll(ctx, q, &rs)

	if err != nil {
		return nil, err
	}

	return &rs, nil
}

func (t *TodoStore) CompleteTodo(ctx context.Context, id int64) (*models.Todo, error) {
	r, err := t.GetTodo(ctx, id)
	if err != nil {
		return nil, err
	}

	r.Complete = true

	if err = t.UpdateTodo(ctx, id, r); err != nil {
		return nil, err
	}

	return r, nil
}

func (t *TodoStore) UpdateTodo(ctx context.Context, id int64, r *models.Todo) error {
	dsClient, err := datastore.NewClient(ctx, t.proj)
	if err != nil {
		return err
	}

	defer dsClient.Close()

	k := datastore.IDKey(t.kind, id, nil)

	_, err = dsClient.Put(ctx, k, r)

	if err != nil {
		return err
	}

	return nil
}
