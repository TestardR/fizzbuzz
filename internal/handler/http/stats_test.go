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

type handlerCaseStats struct {
	h      handler
	status int
}

func TestHandler_GetStats(t *testing.T) {
	t.Parallel()
	mc := gomock.NewController(t)
	t.Cleanup(func() { mc.Finish() })

	tests := map[string]handlerCaseStats{
		"fail-get-max-entries": handlerStatsCaseFailGetMaxEntries(mc),
		"success-no-hits":      handlerStatsCaseSuccessNoHits(mc),
		"success-with-hits":    handlerStatsCaseSuccessWithHits(mc),
	}

	for tn, tc := range tests {
		tn, tc := tn, tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			router := gin.New()
			router.GET("", tc.h.GetStats)
			ts := httptest.NewServer(router)
			t.Cleanup(func() { ts.Close() })

			req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
			if err != nil {
				t.Error(err)
			}
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func handlerStatsCaseFailGetMaxEntries(mc *gomock.Controller) handlerCaseStats {
	ml := mock.NewMockLogger(mc)
	ms := mock.NewMockStorager(mc)

	ms.EXPECT().GetMaxEntries(gomock.Any()).Return("", 0, errors.New("mock"))
	ml.EXPECT().Error(gomock.Any())

	return handlerCaseStats{
		h: handler{
			log:   ml,
			store: ms,
		},
		status: http.StatusInternalServerError,
	}
}

func handlerStatsCaseSuccessNoHits(mc *gomock.Controller) handlerCaseStats {
	ml := mock.NewMockLogger(mc)
	ms := mock.NewMockStorager(mc)

	ms.EXPECT().GetMaxEntries(gomock.Any()).Return("", 0, nil)

	return handlerCaseStats{
		h: handler{
			log:   ml,
			store: ms,
		},
		status: http.StatusNoContent,
	}
}

func handlerStatsCaseSuccessWithHits(mc *gomock.Controller) handlerCaseStats {
	ml := mock.NewMockLogger(mc)
	ms := mock.NewMockStorager(mc)

	ms.EXPECT().GetMaxEntries(gomock.Any()).Return("", 1, nil)

	return handlerCaseStats{
		h: handler{
			log:   ml,
			store: ms,
		},
		status: http.StatusOK,
	}
}

func TestHandler_DeleteStats(t *testing.T) {
	t.Parallel()
	mc := gomock.NewController(t)
	t.Cleanup(func() { mc.Finish() })

	tests := map[string]handlerCaseStats{
		"fail-flush-entries": handleStatsCaseFailReset(mc),
		"success":            handleStatsCaseSuccess(mc),
	}

	for tn, tc := range tests {
		tn, tc := tn, tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			router := gin.New()
			router.GET("", tc.h.DeleteStats)
			ts := httptest.NewServer(router)
			t.Cleanup(func() { ts.Close() })

			req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
			if err != nil {
				t.Error(err)
			}
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func handleStatsCaseFailReset(mc *gomock.Controller) handlerCaseStats {
	ml := mock.NewMockLogger(mc)
	ms := mock.NewMockStorager(mc)

	ms.EXPECT().Reset(gomock.Any()).Return(errors.New("mock"))
	ml.EXPECT().Error(gomock.Any())

	return handlerCaseStats{
		h: handler{
			log:   ml,
			store: ms,
		},
		status: http.StatusInternalServerError,
	}
}

func handleStatsCaseSuccess(mc *gomock.Controller) handlerCaseStats {
	ml := mock.NewMockLogger(mc)
	ms := mock.NewMockStorager(mc)

	ms.EXPECT().Reset(gomock.Any())

	return handlerCaseStats{
		h: handler{
			log:   ml,
			store: ms,
		},
		status: http.StatusOK,
	}
}
