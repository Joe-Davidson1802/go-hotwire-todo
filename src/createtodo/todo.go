package main

import "cloud.google.com/go/datastore"

type Todo struct {
	ID       *datastore.Key `datastore:"__key__"`
	Title    string         `datastore:"title"`
	Complete bool           `datastore:"complete"`
}
