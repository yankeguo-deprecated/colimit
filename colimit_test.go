package colimit

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestColimit(t *testing.T) {
	e := echo.New()
	req1 := httptest.NewRequest(http.MethodGet, "/1", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	handler := func(c echo.Context) error {
		time.Sleep(time.Second)
		return c.String(http.StatusOK, "OK")
	}
	m := New(0)
	h := m(handler)
	h(c1)

	assert.Equal(t, "OK", rec1.Body.String())
	assert.Equal(t, http.StatusOK, rec1.Code)

	req2 := httptest.NewRequest(http.MethodGet, "/2", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	h(c2)

	assert.Equal(t, "OK", rec2.Body.String())
	assert.Equal(t, http.StatusOK, rec2.Code)

	m = New(1)
	h = m(handler)

	req3 := httptest.NewRequest(http.MethodGet, "/3", nil)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	h(c3)

	assert.Equal(t, "OK", rec3.Body.String())
	assert.Equal(t, http.StatusOK, rec3.Code)

	req4 := httptest.NewRequest(http.MethodGet, "/4", nil)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	go func() {
		time.Sleep(time.Second / 2)
		h(c4)
		assert.Equal(t, "concurrency limit exceeded", rec4.Body.String())
		assert.Equal(t, http.StatusServiceUnavailable, rec4.Code)
	}()

	req5 := httptest.NewRequest(http.MethodGet, "/5", nil)
	rec5 := httptest.NewRecorder()
	c5 := e.NewContext(req5, rec5)
	h(c5)

	assert.Equal(t, "OK", rec5.Body.String())
	assert.Equal(t, http.StatusOK, rec5.Code)

	time.Sleep(time.Second)

	req6 := httptest.NewRequest(http.MethodGet, "/6", nil)
	rec6 := httptest.NewRecorder()
	c6 := e.NewContext(req6, rec6)
	h(c6)

	assert.Equal(t, "OK", rec6.Body.String())
	assert.Equal(t, http.StatusOK, rec6.Code)
}
