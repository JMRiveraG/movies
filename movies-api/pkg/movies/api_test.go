package movies

import (
	"net/http"
	"sort"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchMovies(t *testing.T) {
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:             "RegularCase",
			mockResponseBody: `{"Search":[{"Title":"Star Wars: A New Hope","Year":"1977"},{"Title":"Star Wars: The Empire Strikes Back","Year":"1980"}]}`,
			expectedMovies: []Movie{
				{Title: "Star Wars: A New Hope", Year: "1977"},
				{Title: "Star Wars: The Empire Strikes Back", Year: "1980"},
			},
			expectedErrorString: "",
		},
	}

	searcher := &APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		// register http mock
		httpmock.RegisterResponder(
			"GET",
			"http://example.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMovies("star wars")
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}

func TestSearchMoviesSort(t *testing.T) {

	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:             "RegularCaseSortMovies",
			mockResponseBody: `{"Search":[{"Title":"The Matrix","Year":"1999"},{"Title":"Making 'The Matrix'","Year":"1999"},{"Title":"The Matrix Revolutions","Year":"2003"},{"Title":"The Matrix Reloaded","Year":"2003"}]}`,
			expectedMovies: []Movie{
				{Title: "Making 'The Matrix'", Year: "1999"},
				{Title: "The Matrix", Year: "1999"},
				{Title: "The Matrix Reloaded", Year: "2003"},
				{Title: "The Matrix Revolutions", Year: "2003"},
			},
			expectedErrorString: "",
		},
	}

	searcher := &APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		// register http mock
		httpmock.RegisterResponder(
			"GET",
			"http://example.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMoviesSorted("matrix")

			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}

			isSorted := sort.SliceIsSorted(actualMovies, func(k, v int) bool {
				if actualMovies[k].Year != actualMovies[v].Year {
					return actualMovies[k].Year < actualMovies[v].Year
				}
				return actualMovies[k].Title < actualMovies[v].Title
			})

			assert.True(t, isSorted)
			assert.EqualValues(t, c.expectedMovies, actualMovies)
		})
	}
}
