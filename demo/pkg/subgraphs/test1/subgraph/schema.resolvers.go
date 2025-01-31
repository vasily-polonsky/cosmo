package subgraph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wundergraph/cosmo/demo/pkg/injector"
	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/test1/subgraph/generated"
	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/test1/subgraph/model"
)

// HeaderValue is the resolver for the headerValue field.
func (r *queryResolver) HeaderValue(ctx context.Context, name string) (string, error) {
	header := injector.Header(ctx)
	if header == nil {
		return "", errors.New("headers not injected into context.Context")
	}
	return header.Get(name), nil
}

// InitPayloadValue is the resolver for the initPayloadValue field.
func (r *queryResolver) InitPayloadValue(ctx context.Context, key string) (string, error) {
	payload := injector.InitPayload(ctx)
	if payload == nil {
		return "", errors.New("payload not injected into context.Context")
	}
	return fmt.Sprintf("%v", payload[key]), nil
}

// InitialPayload is the resolver for the initialPayload field.
func (r *queryResolver) InitialPayload(ctx context.Context) (map[string]interface{}, error) {
	payload := injector.InitPayload(ctx)
	if payload == nil {
		return nil, errors.New("payload not injected into context.Context")
	}
	return payload, nil
}

// Delay is the resolver for the delay field.
func (r *queryResolver) Delay(ctx context.Context, response string, ms int) (string, error) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return response, nil
}

// HeaderValue is the resolver for the headerValue field.
func (r *subscriptionResolver) HeaderValue(ctx context.Context, name string, repeat *int) (<-chan *model.TimestampedString, error) {
	header := injector.Header(ctx)
	if header == nil {
		return nil, errors.New("headers not injected into context.Context")
	}
	ch := make(chan *model.TimestampedString, 1)

	if repeat == nil {
		repeat = new(int)
		*repeat = 1
	}

	payload := injector.InitPayload(ctx)
	if payload == nil {
		payload = map[string]any{}
	}

	go func() {
		defer close(ch)

		for ii := 0; ii < *repeat; ii++ {
			// In our example we'll send the current time every second.
			time.Sleep(100 * time.Millisecond)
			select {
			case <-ctx.Done():
				return

			case ch <- &model.TimestampedString{
				Value:          header.Get(name),
				UnixTime:       int(time.Now().Unix()),
				Seq:            ii,
				Total:          *repeat,
				InitialPayload: payload,
			}:
			}
		}
	}()
	return ch, nil
}

// InitPayloadValue is the resolver for the initPayloadValue field.
func (r *subscriptionResolver) InitPayloadValue(ctx context.Context, key string, repeat *int) (<-chan *model.TimestampedString, error) {
	payload := injector.InitPayload(ctx)
	if payload == nil {
		return nil, errors.New("payload not injected into context.Context")
	}
	ch := make(chan *model.TimestampedString, 1)

	if repeat == nil {
		repeat = new(int)
		*repeat = 1
	}

	go func() {
		defer close(ch)

		for ii := 0; ii < *repeat; ii++ {
			// In our example we'll send the current time every second.
			time.Sleep(100 * time.Millisecond)
			select {
			case <-ctx.Done():
				return

			case ch <- &model.TimestampedString{
				Value:          fmt.Sprintf("%v", payload[key]),
				UnixTime:       int(time.Now().Unix()),
				Seq:            ii,
				Total:          *repeat,
				InitialPayload: payload,
			}:
			}
		}
	}()
	return ch, nil
}

// InitialPayload is the resolver for the initialPayload field.
func (r *subscriptionResolver) InitialPayload(ctx context.Context, repeat *int) (<-chan map[string]interface{}, error) {
	payload := injector.InitPayload(ctx)
	if payload == nil {
		payload = make(map[string]any)
	}
	ch := make(chan map[string]any, 1)

	if repeat == nil {
		repeat = new(int)
		*repeat = 1
	}

	go func() {
		defer close(ch)

		for ii := 0; ii < *repeat; ii++ {
			// In our example we'll send the current time every second.
			time.Sleep(100 * time.Millisecond)
			select {
			case <-ctx.Done():
				return

			case ch <- payload:

			}
		}
	}()
	return ch, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
