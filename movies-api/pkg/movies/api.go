package movies

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	Search []Movie `json:"Search"`
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query string) ([]Movie, error) {

	respBody, err := getResponse(s, query)
	if err != nil {
		return nil, err
	}

	var respStruct omdbapiResponse
	json.Unmarshal(respBody, &respStruct)

	// return result
	return respStruct.Search, nil
}

func (s *APIMovieSearcher) SearchMoviesSorted(query string) ([]Movie, error) {
	//Getting http response
	respBody, err := getResponse(s, query)

	if err != nil {
		return nil, err
	}

	// return result
	return sortMovies(respBody)
}

func getResponse(s *APIMovieSearcher, query string) ([]byte, error) {
	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	resp, err := http.Get(s.URL + "?" + params.Encode())

	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// return respBody
	return respBody, nil
}

func sortMovies(respBody []byte) ([]Movie, error) {
	var respStruct omdbapiResponse
	var ss []Movie

	// unmarshall response
	json.Unmarshal(respBody, &respStruct)
	// sorting response body by year and title
	for _, v := range respStruct.Search {
		ss = append(ss, v)
		sort.Slice(ss, func(i, j int) bool {
			if ss[i].Year != ss[j].Year {
				return ss[i].Year < ss[j].Year
			}
			return ss[i].Title < ss[j].Title
		})
	}
	// return sorted movies
	return ss, nil
}
