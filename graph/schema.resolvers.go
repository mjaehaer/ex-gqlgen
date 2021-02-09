package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/mjaehaer/hackernews/graph/generated"
	"github.com/mjaehaer/hackernews/graph/model"
)

func (r *mutationResolver) CreateTweet(ctx context.Context, input model.NewTweet) (*model.Tweet, error) {
	var tweet model.Tweet
	tweet.CotentText = input.CotentText
	tweet.AuthoID = &model.User{}
	tweet.PublicationDate = time.Now().Local().String()

	return &tweet, nil
}

func (r *mutationResolver) GetTweetsOnDate(ctx context.Context, input string) ([]*model.Tweet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	var tweets []*model.Tweet
	tweets = append(tweets, &model.Tweet{CotentText: "started content 1", PublicationDate: "2021-02-09 00:00:01"})

	tweets = append(tweets, &model.Tweet{CotentText: "our dummy content", PublicationDate: "2021-02-09 01:31:01"})

	return tweets, nil
}

func (r *subscriptionResolver) SubscritTweet(ctx context.Context, user string) (<-chan []*model.Tweet, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
