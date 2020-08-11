package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cappyzawa/golang-lab/photo-share-api/graph/generated"
	"github.com/cappyzawa/golang-lab/photo-share-api/graph/model"
)

func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	r.id++
	photo := &model.Photo{
		ID:          strconv.Itoa(r.id),
		URL:         fmt.Sprintf("http://yoursite.com/img/%d.jpg", r.id),
		Name:        input.Name,
		Description: input.Description,
		Category:    *input.Category,
	}
	r.photos = append(r.photos, photo)
	return photo, nil
}

func (r *queryResolver) TotalPhotos(ctx context.Context) (int, error) {
	return len(r.photos), nil
}

func (r *queryResolver) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	return r.photos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
