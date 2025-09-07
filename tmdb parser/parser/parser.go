package parser

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/sync/semaphore"
	"slices"
	"strings"
	"sync"
	"tmdb_parser/db/repository"
	"tmdb_parser/db/repository/schemas"
	"tmdb_parser/pkg/tmdb"
	"tmdb_parser/pkg/tmdb/model"
)

type Parser struct {
	movieRepo  *repository.MovieRepository
	tmdbClient *tmdb.Client
}

func New(movieRepo *repository.MovieRepository, tmdbClient *tmdb.Client) *Parser {
	return &Parser{
		movieRepo:  movieRepo,
		tmdbClient: tmdbClient,
	}
}

func (p *Parser) Run(ctx context.Context) {
	logger := log.With().Str("parser", "parser").Logger()
	var nextCursor bson.ObjectID
	const concurrencyLimit = 50

	sem := semaphore.NewWeighted(concurrencyLimit)
	var wg sync.WaitGroup

	for {
		movies, err := p.movieRepo.GetWithoutTmdbID(ctx, nextCursor)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to get movies without tmdb id")
			return
		}
		if len(movies) == 0 {
			break
		}

		results := make([]error, len(movies))

		for i, movie := range movies {
			if err := sem.Acquire(ctx, 1); err != nil {
				logger.Error().Err(err).Msg("Failed to acquire semaphore")
				continue
			}

			wg.Add(1)
			movieCopy := movie
			idx := i

			go func() {
				defer wg.Done()
				defer sem.Release(1)

				err := p.processMovie(ctx, &movieCopy)
				if err != nil {
					logger.Error().Err(err).Str("title", movieCopy.Title).Int("year", movieCopy.Year).Msg("Failed to process movie")
					if setErr := p.movieRepo.SetTmdbChecked(ctx, movieCopy.OID); setErr != nil {
						logger.Error().Err(setErr).Msg("failed to set tmdb checked")
					}
				}

				results[idx] = err
			}()
		}

		wg.Wait()

		for i, movie := range movies {
			if results[i] == nil {
				nextCursor = movie.OID
			}
		}
	}
}

func (p *Parser) processMovie(ctx context.Context, movie *schemas.Movie) error {
	logger := log.With().Str("title", movie.Title).Int("year", movie.Year).Logger()
	var (
		searchResult *model.SearchResponse
		title        string
		err          error
	)

	if movie.AlternativeTitle != "" {
		title = movie.AlternativeTitle
	}

	if movie.Title != "" && movie.Title != title {
		if title != "" {
			title += "+"
		}
		title += movie.Title
	}

	searchResult, err = p.search(ctx, title, movie.Year, movie.IsSeries)
	if err != nil {
		return err
	}

	if len(searchResult.Movies) == 0 {
		if movie.AlternativeTitle != "" {
			searchResult, err = p.search(ctx, movie.AlternativeTitle, movie.Year, movie.IsSeries)
			if err != nil {
				return err
			}
		} else {
			searchResult, err = p.search(ctx, movie.Title, movie.Year, movie.IsSeries)
			if err != nil {
				return err
			}
		}
	}

	if len(searchResult.Movies) == 0 {
		logger.Error().Msg("no results found for search")
		return fmt.Errorf("no results found for %s", title)
	}

	matchingMovie, err := p.findMatchingMovie(searchResult, movie)
	if err != nil {
		logger.Error().Err(err).Msg("failed to find matching movie")

		return fmt.Errorf("failed to find matching movie: %w", err)
	}

	logger.Info().Msgf("found matching movie: %s", matchingMovie.Title)

	if err := p.movieRepo.SetExternalIDs(ctx, movie.OID, matchingMovie.ID); err != nil {
		logger.Error().Err(err).Msg("failed to update tmdb id")
		return fmt.Errorf("failed to update tmdb id: %w", err)
	}

	if err := p.movieRepo.SetTmdbChecked(ctx, movie.OID); err != nil {
		logger.Error().Err(err).Msg("failed to set tmdb checked")
		return fmt.Errorf("failed to set tmdb checked: %w", err)
	}

	return nil
}

func (p *Parser) search(ctx context.Context, title string, year int, isSeries bool) (*model.SearchResponse, error) {
	logger := zerolog.Ctx(ctx)
	var (
		searchResult *model.SearchResponse
		err          error
	)

	if isSeries {
		searchResult, err = p.tmdbClient.SearchTV(ctx, title, year)
		if err != nil {
			logger.Error().Err(err).Msg("failed to search tv")
			return nil, fmt.Errorf("failed to search tv: %w", err)
		}
	} else {
		searchResult, err = p.tmdbClient.SearchMovie(ctx, title, year)
		if err != nil {
			logger.Error().Err(err).Msg("failed to search movie")
			return nil, fmt.Errorf("failed to search movie: %w", err)
		}
	}
	return searchResult, err
}

func (p *Parser) findMatchingMovie(searchResponse *model.SearchResponse, movie *schemas.Movie) (model.SearchResult, error) {
	var movieTitles []string

	if movie.Title != "" {
		movieTitles = append(movieTitles, strings.ToLower(movie.Title))
	}

	if movie.AlternativeTitle != "" {
		movieTitles = append(movieTitles, strings.ToLower(movie.AlternativeTitle))
	}

	if len(searchResponse.Movies) == 1 {
		return searchResponse.Movies[0], nil
	}

	for _, searchItem := range searchResponse.Movies {
		var foundTitles []string

		if searchItem.Title != "" {
			foundTitles = append(foundTitles, strings.ToLower(searchItem.Title))
		}

		if searchItem.OriginalTitle != "" {
			foundTitles = append(foundTitles, strings.ToLower(searchItem.OriginalTitle))
		}

		matchResult := model.SearchResult{}
		for _, movieTitle := range movieTitles {
			if slices.Contains(foundTitles, movieTitle) {
				if searchItem.Year != 0 && searchItem.Year != searchItem.Year {
					continue
				}

				matchResult = searchItem
				break
			}
		}

		if matchResult.ID != 0 {
			return matchResult, nil
		}
	}

	return model.SearchResult{}, fmt.Errorf("failed to find matching searchItem")
}
