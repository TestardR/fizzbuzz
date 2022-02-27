package redis

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -package=mock -source=redis.go -destination=$MOCK_FOLDER/redis.go Statser

const (
	maxKey  = "max_key"
	maxHits = "max_hits"
)

var (
	defaultKeyValue  string
	defaultHitsValue int
)

// Storager describes the methods available on the Storage struct.
type Storager interface {
	// IncrementCount Increments the counter associated to a key by 1.
	IncrementCount(ctx context.Context, key string) error
	// GetMaxEntries retrieves the most frequently incremented key with its value.
	GetMaxEntries(ctx context.Context) (string, int, error)
	// Reset removes all entries and set max_key an max_hits to default value.
	Reset(ctx context.Context) error
	// Health checks that storage layer is healthy.
	Health(ctx context.Context) error
}

// Storage holds a client to database.
type Storage struct {
	mu     sync.RWMutex
	client *redis.Client
}

// NewStorage creates a Storage instance.
func NewStorage(host, port string) (Storager, error) {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	s := Storage{client: redis.NewClient(opt)}

	ctx := context.Background()

	// Don't initialize if existing values
	_, err = s.client.Get(ctx, maxKey).Result()
	if err == nil {
		return &s, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	err = s.setMaxEntries(ctx, defaultKeyValue, defaultHitsValue)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *Storage) IncrementCount(ctx context.Context, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	result, err := s.client.Incr(ctx, key).Result()
	if err != nil {
		return err
	}

	currMaxHits, err := s.client.Get(ctx, maxHits).Int64()
	if err != nil {
		return err
	}

	if currMaxHits >= result {
		return nil
	}

	err = s.setMaxEntries(ctx, key, int(result))
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetMaxEntries(ctx context.Context) (string, int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	key, err := s.client.Get(ctx, maxKey).Result()
	if err != nil {
		return "", 0, err
	}

	hits, err := s.client.Get(ctx, maxHits).Int()
	if err != nil {
		return "", 0, err
	}

	return key, hits, nil
}

func (s *Storage) Reset(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.client.FlushDB(ctx).Err(); err != nil {
		return err
	}

	err := s.setMaxEntries(ctx, defaultKeyValue, defaultHitsValue)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) setMaxEntries(ctx context.Context, keyValue string, hitsValue int) error {
	pipe := s.client.TxPipeline()
	pipe.Set(ctx, maxKey, keyValue, 0)
	pipe.Set(ctx, maxHits, hitsValue, 0)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Health(ctx context.Context) error {
	return s.client.Ping(ctx).Err()
}
