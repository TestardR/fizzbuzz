package redis

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

var (
	mock *miniredis.Miniredis
	ctx  = context.TODO()
)

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return s
}

func setup() {
	mock = mockRedis()
}

func teardown() {
	mock.Close()
}

func Test_NewStorage(t *testing.T) {
	t.Run("should_return_err_if_redis.ParseUrl_fails", func(t *testing.T) {
		setup()
		defer teardown()

		_, err := NewStorage("foo", "bar")

		assert.Error(t, err)
	})

	/* 	t.Run("should_return_default_maxKey_and_maxHits_value", func(t *testing.T) {
	   		setup()
	   		defer teardown()

	   		_, err := NewStorage(mock.Host(), mock.Port())

	   		mock.CheckGet(t, maxKey, "")
	   		mock.CheckGet(t, maxHits, "0")
	   		assert.NoError(t, err)
	   	})

	   	t.Run("should_be_ok", func(t *testing.T) {
	   		setup()
	   		defer teardown()

	   		_, err := NewStorage(mock.Host(), mock.Port())
	   		assert.NoError(t, err)
	   	}) */
}

func Test_IncrementCount(t *testing.T) {
	t.Run("should_INCR_work_well_and_update_maxHits", func(t *testing.T) {
		setup()
		defer teardown()

		c, _ := NewStorage(mock.Host(), mock.Port())

		mock.Incr("foo", 123)

		err := c.IncrementCount(ctx, "foo")
		assert.NoError(t, err)

		maxHits, err := mock.Get(maxHits)
		assert.NoError(t, err)
		assert.Equal(t, "124", maxHits)

		fopKey, err := mock.Get("foo")
		assert.NoError(t, err)
		assert.Equal(t, "124", fopKey)
	})

	t.Run("should_return_err_and_not_update_anything_if_an_error_occurs", func(t *testing.T) {
		setup()
		defer teardown()

		mock.Incr("key1", 21)
		mock.Incr("key2", 42)

		c, _ := NewStorage(mock.Host(), mock.Port())
		mock.SetError("mock")

		err := c.IncrementCount(ctx, "key2")
		assert.Error(t, err)

		maxHits, err := mock.Get(maxHits)
		assert.NoError(t, err)
		assert.Equal(t, "0", maxHits)

		maxKey, err := mock.Get(maxKey)
		assert.NoError(t, err)
		assert.Equal(t, "", maxKey)

		key2, err := mock.Get("key2")
		assert.NoError(t, err)
		assert.Equal(t, "42", key2)
	})
}

func Test_GetMaxEntries(t *testing.T) {
	t.Run("should_returb_err_if_Get_fails", func(t *testing.T) {
		setup()
		defer teardown()

		c, _ := NewStorage(mock.Host(), mock.Port())

		mock.SetError("error")
		_, _, err := c.GetMaxEntries(ctx)
		assert.Error(t, err)
	})

	t.Run("should_be_ok", func(t *testing.T) {
		setup()
		defer teardown()

		mock.Incr("key1", 123)
		mock.Incr("key2", 234)

		c, _ := NewStorage(mock.Host(), mock.Port())

		c.IncrementCount(ctx, "key1")
		c.IncrementCount(ctx, "key2")

		maxKey, maxHits, err := c.GetMaxEntries(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "key2", maxKey)
		assert.Equal(t, 235, maxHits)
	})
}

func Test_Reset(t *testing.T) {
	t.Run("should_return_err_if_Flush_fails", func(t *testing.T) {
		setup()
		defer teardown()

		c, _ := NewStorage(mock.Host(), mock.Port())

		mock.SetError("error")
		err := c.Reset(ctx)
		assert.Error(t, err)
	})

	t.Run("should_be_ok", func(t *testing.T) {
		setup()
		defer teardown()

		mock.Incr("key1", 234)
		mock.Incr("key2", 123)

		c, _ := NewStorage(mock.Host(), mock.Port())

		err := c.Reset(ctx)
		assert.NoError(t, err)

		c.IncrementCount(ctx, "key2")
		c.IncrementCount(ctx, "key2")

		maxKey, maxHits, _ := c.GetMaxEntries(ctx)

		assert.Equal(t, "key2", maxKey)
		assert.Equal(t, 2, maxHits)
	})
}
