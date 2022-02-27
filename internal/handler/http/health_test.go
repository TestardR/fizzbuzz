package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TestardR/fizz-buzz/pkg/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"gotest.tools/assert"
)

type handlerCaseHealth struct {
	h      handler
	status int
}

func TestHandler_Health(t *testing.T) {
	t.Parallel()
	mc := gomock.NewController(t)
	t.Cleanup(func() { mc.Finish() })

	tests := map[string]handlerCaseHealth{
		"health-not-ok": handlerHealthCaseFailCheck(mc),
		"health-ok":     handlerHealthCaseOk(mc),
	}

	for tn, tc := range tests {
		tn, tc := tn, tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			ts, uri := handlerHealthServer(tc.h)
			t.Cleanup(func() { ts.Close() })

			req, err := http.NewRequest(http.MethodGet, uri, nil)
			require.NoError(t, err)

			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)

		})
	}
}

func handlerHealthCaseFailCheck(mc *gomock.Controller) handlerCaseHealth {
	ms := mock.NewMockStorager(mc)
	ml := mock.NewMockLogger(mc)

	ml.EXPECT().Error(gomock.Any())
	ms.EXPECT().Health(gomock.Any()).Return(errors.New("mock"))

	return handlerCaseHealth{
		h: handler{
			log:   ml,
			store: ms,
		},
		status: http.StatusInternalServerError,
	}
}

func handlerHealthCaseOk(mc *gomock.Controller) handlerCaseHealth {
	ms := mock.NewMockStorager(mc)

	ms.EXPECT().Health(gomock.Any())

	return handlerCaseHealth{
		h: handler{
			store: ms,
		},
		status: http.StatusOK,
	}
}

func handlerHealthServer(h handler) (*httptest.Server, string) {
	router := gin.New()

	router.GET("", h.Health)
	ts := httptest.NewServer(router)

	return ts, ts.URL
}
