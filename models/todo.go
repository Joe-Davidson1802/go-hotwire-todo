package models

import (
	"encoding/json"

	"cloud.google.com/go/datastore"
)

type Todo struct {
	ID       *datastore.Key `datastore:"__key__"`
	Title    string         `datastore:"title"`
	Complete bool           `datastore:"complete"`
}

type jsonTodo struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

func (t *Todo) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonTodo{
		ID:       t.ID.ID,
		Title:    t.Title,
		Complete: t.Complete,
	})
}
