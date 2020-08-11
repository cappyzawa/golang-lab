package graph

import "github.com/cappyzawa/golang-lab/photo-share-api/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	id     int
	photos []*model.Photo
}
