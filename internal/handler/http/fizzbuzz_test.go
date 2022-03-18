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

type handlerCaseFizzBuzz struct {
	h           handler
	queryString string
	status      int
}

func TestHandler_GetFizzBuzz(t *testing.T) {
	t.Parallel()
	mc := gomock.NewController(t)
	t.Cleanup(func() { mc.Finish() })

	tests := map[string]handlerCaseFizzBuzz{
		// TODO: improve unit tests for query string parameters
		"fail-invalid-query-parameter": handlerFizzBuzzCaseFailInvalidQueryParameter(mc),
		"fail-validation":              handlerFizzBuzzCaseFailValidation(mc),
		"fail-to-store-increment":      handlerFizzBuzzCaseFailStoreIncrement(mc),
		"success":                      handlerFizzBuzzCaseSuccess(mc),
	}

	for tn, tc := range tests {
		tn, tc := tn, tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			router := gin.New()
			router.GET("", tc.h.GetFizzbuzz)
			ts := httptest.NewServer(router)
			t.Cleanup(func() { ts.Close() })

			req, err := http.NewRequest(http.MethodGet, ts.URL+tc.queryString, nil)
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

func handlerFizzBuzzCaseFailInvalidQueryParameter(mc *gomock.Controller) handlerCaseFizzBuzz {
	ml := mock.NewMockLogger(mc)

	ml.EXPECT().Error(gomock.Any())

	return handlerCaseFizzBuzz{
		h: handler{
			log: ml,
		},
		queryString: `?int1="0"`,
		status:      http.StatusBadRequest,
	}
}

func handlerFizzBuzzCaseFailValidation(mc *gomock.Controller) handlerCaseFizzBuzz {
	ml := mock.NewMockLogger(mc)

	ml.EXPECT().Info(gomock.Any())
	ml.EXPECT().Error(gomock.Any())

	return handlerCaseFizzBuzz{
		h: handler{
			log: ml,
		},
		queryString: "?int1=0",
		status:      http.StatusBadRequest,
	}
}

func handlerFizzBuzzCaseFailStoreIncrement(mc *gomock.Controller) handlerCaseFizzBuzz {
	ml := mock.NewMockLogger(mc)
	ms := mock.NewMockStorager(mc)

	ml.EXPECT().Info("request received with payload: {3 5 10 foo bar}")
	ms.EXPECT().IncrementCount(gomock.Any(), "3,5,10,foo,bar").Return(errors.New("mock"))
	ml.EXPECT().Error(gomock.Any())

	return handlerCaseFizzBuzz{
		h: handler{
			log:   ml,
			store: ms,
		},
		queryString: `?int1=3&int2=5&limit=10&str1=foo&str2=bar`,
		status:      http.StatusOK,
	}
}

func handlerFizzBuzzCaseSuccess(mc *gomock.Controller) handlerCaseFizzBuzz {
	ml := mock.NewMockLogger(mc)
	ms := mock.NewMockStorager(mc)

	ml.EXPECT().Info("request received with payload: {3 5 10 foo bar}")
	ms.EXPECT().IncrementCount(gomock.Any(), "3,5,10,foo,bar")

	return handlerCaseFizzBuzz{
		h: handler{
			log:   ml,
			store: ms,
		},
		queryString: `?int1=3&int2=5&limit=10&str1=foo&str2=bar`,
		status:      http.StatusOK,
	}
}
