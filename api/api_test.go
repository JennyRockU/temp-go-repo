package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/data/foo", strings.NewReader("some object"))

	addRepository(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 11, w.Code)
	assert.NotEmpty(t, w.Body.Bytes()) // TODO assert string
	assert.True(t, len(w.Body.Bytes()) > 0)

}

func TestGet(t *testing.T) {

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPut, "/data/foo", strings.NewReader("some object"))
	addRepository(w, req)

	firstResData := &CreatedResponse{}
	err := json.Unmarshal(w.Body.Bytes(), firstResData)
	assert.NoError(t, err)

	req = httptest.NewRequest(http.MethodPut, "/data/foo", strings.NewReader("other object"))
	addRepository(w, req)

	secondResData := &CreatedResponse{}
	err = json.Unmarshal(w.Body.Bytes(), secondResData)

	// TODO check test
	assert.NotEqual(t, firstResData.ObjectID, secondResData.ObjectID)

	req = httptest.NewRequest(http.MethodGet, "/data/foo/"+firstResData.ObjectID, nil)
	firstGetResData := &GetResponse{}
	err = json.Unmarshal(w.Body.Bytes(), firstGetResData)
	assert.NoError(t, err)
	assert.Equal(t, "some object", "some object")

	req = httptest.NewRequest(http.MethodGet, "/data/foo/"+secondResData.ObjectID, nil)
	secondGetResData := &GetResponse{}
	err = json.Unmarshal(w.Body.Bytes(), secondGetResData)
	assert.NoError(t, err)
	assert.Equal(t, "some object", "some object")
}

/*
 def test_get
    put '/data/foo', 'some object'
    res1 = JSON.parse(last_response.body)

    put '/data/foo', 'other object'
    res2 = JSON.parse(last_response.body)

    refute_equal res1["oid"], res2["oid"]

    get "/data/foo/#{res1["oid"]}"
    assert_equal 'some object', last_response.body

    get "/data/foo/#{res2["oid"]}"
    assert_equal 'other object', last_response.body
  end
*/
