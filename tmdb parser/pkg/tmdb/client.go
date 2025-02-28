package tmdb

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
	"tmdb_parser/pkg/tmdb/model"
)

type (
	Client struct {
		client *resty.Client
	}
)

func NewClient(apiKey string) *Client {
	client := resty.New()
	client.SetBaseURL("https://api.themoviedb.org/3")
	client.SetHeader("Authorization", "Bearer "+apiKey)

	return &Client{
		client: client,
	}
}

func (c *Client) SearchMovie(ctx context.Context, query string, year int) (*model.SearchResponse, error) {
	var response model.MovieSearchResponse

	params := map[string]string{
		"query":         query,
		"language":      "ru-RU",
		"include_adult": "true",
	}

	if year > 0 {
		params["year"] = strconv.Itoa(year)
	}

	if _, err := c.client.R().
		SetQueryParams(params).
		SetResult(&response).Get("/search/movie"); err != nil {
		return nil, err
	}

	searchResponse := model.SearchResponse{
		Movies: make([]model.SearchResult, len(response.Results)),
	}

	for i, result := range response.Results {
		var year int

		if result.ReleaseDate != "" {
			releaseDate, err := time.Parse("2006-01-02", result.ReleaseDate)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse release date")
				continue
			}
			year = releaseDate.Year()
		}

		searchResponse.Movies[i] = model.SearchResult{
			ID:            result.ID,
			OriginalTitle: result.OriginalTitle,
			Title:         result.Title,
			Year:          year,
		}
	}

	return &searchResponse, nil
}

func (c *Client) SearchTV(ctx context.Context, query string, year int) (*model.SearchResponse, error) {
	var response model.TVSearchResponse

	params := map[string]string{
		"query":         query,
		"language":      "ru-RU",
		"include_adult": "true",
	}

	if year > 0 {
		params["year"] = strconv.Itoa(year)
	}

	if _, err := c.client.R().
		SetQueryParams(params).
		SetResult(&response).Get("/search/tv"); err != nil {
		return nil, err
	}

	searchResponse := model.SearchResponse{
		Movies: make([]model.SearchResult, len(response.Results)),
	}

	for i, result := range response.Results {
		var year int

		if result.FirstAirDate != "" {
			releaseDate, err := time.Parse("2006-01-02", result.FirstAirDate)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse release date")
				continue
			}

			year = releaseDate.Year()
		}

		searchResponse.Movies[i] = model.SearchResult{
			ID:            result.ID,
			OriginalTitle: result.OriginalTitle,
			Title:         result.Title,
			Year:          year,
		}
	}

	return &searchResponse, nil
}

func (c *Client) GetTV(ctx context.Context, id int) (*model.TVResponse, error) {
	var response model.TVResponse

	if _, err := c.client.R().
		SetPathParam("id", strconv.Itoa(id)).
		SetResult(&response).Get("/tv/{id}"); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetMovie(ctx context.Context, id int) (*model.MovieResponse, error) {
	var response model.MovieResponse

	if _, err := c.client.R().
		SetPathParam("id", strconv.Itoa(id)).
		SetResult(&response).Get("/movie/{id}"); err != nil {
		return nil, err
	}

	return &response, nil
}
