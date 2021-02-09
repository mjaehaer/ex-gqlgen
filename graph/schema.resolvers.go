package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/mjaehaer/hackernews/graph/generated"
	"github.com/mjaehaer/hackernews/graph/model"
)

func (r *mutationResolver) CreateTweet(ctx context.Context, input model.NewTweet) (*model.Tweet, error) {
	var tweet model.Tweet
	tweet.CotentText = input.CotentText
	tweet.AuthoID = input.AuthoID
	tweet.PublicationDate = time.Now().Local().String()

	return &tweet, nil
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	var tweets []*model.Tweet
	tweets = append(tweets, &model.Tweet{CotentText: "started content 1", PublicationDate: "2021-02-09 00:00:01", AuthoID: "admin"})

	tweets = append(tweets, &model.Tweet{CotentText: "our dummy content", PublicationDate: "2021-02-09 01:31:01", AuthoID: "user"})

	return tweets, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *subscriptionResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	var tweets []*model.Tweet

	tweets = append(tweets, &model.Tweet{CotentText: "started content 1", PublicationDate: "2021-02-09 00:00:01", AuthoID: "admin"})

	tweets = append(tweets, &model.Tweet{CotentText: "our dummy content", PublicationDate: "2021-02-09 01:31:01", AuthoID: "user"})
	return tweets, nil
}
func (r *Resolver) Subscription() generated.QueryResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
