package kp

import (
	"context"
	"github.com/idoubi/goz"
	"github.com/mdwitr0/kinopoisk/crawler/external/kp/model"
	"github.com/rs/zerolog"
)

func (c *Client) GetAllMovies(ctx context.Context, limit int, offset int) (*model.Response[model.MovieList], error) {
	logger := zerolog.Ctx(ctx)

	operation := "MovieDesktopListPage"
	query := `query MovieDesktopListPage($slug: String!, $platform: WebClientPlatform!, $withUserData: Boolean!, $regionId: Int!, $supportedFilterTypes: [FilterType]!, $filters: FilterValuesInput, $singleSelectFiltersLimit: Int!, $singleSelectFiltersOffset: Int!, $moviesLimit: Int, $moviesOffset: Int, $moviesOrder: MovieListItemOrderBy, $supportedItemTypes: [MovieListItemType]) { movieListBySlug(slug: $slug, supportedFilterTypes: $supportedFilterTypes, filters: $filters) { id name description cover { avatarsUrl __typename } ...MovieListCompositeName ...MovieListAvailableFilters ...MovieList ...DescriptionLink __typename } webPage(platform: $platform) { kpMovieListPage(movieListSlug: $slug) { htmlMeta { ...OgImage __typename } footer { ...FooterConfigData __typename } featuring { ...MovieListFeaturingData __typename } __typename } __typename } } fragment MovieListCompositeName on MovieListMeta { compositeName { parts { ... on FilterReferencedMovieListNamePart { filterValue { ... on SingleSelectFilterValue { filterId __typename } __typename } name __typename } ... on StaticMovieListNamePart { name __typename } __typename } __typename } __typename } fragment MovieListAvailableFilters on MovieListMeta { availableFilters { items { ... on BooleanFilter { ...ToggleFilter __typename } ... on SingleSelectFilter { ...SingleSelectFilters __typename } __typename } __typename } __typename } fragment ToggleFilter on BooleanFilter { id enabled name { russian __typename } __typename } fragment SingleSelectFilters on SingleSelectFilter { id name { russian __typename } hint { russian __typename } values(offset: $singleSelectFiltersOffset, limit: $singleSelectFiltersLimit) { items { name { russian __typename } selectable value __typename } __typename } __typename } fragment MovieList on MovieListMeta { movies(limit: $moviesLimit, offset: $moviesOffset, orderBy: $moviesOrder, supportedItemTypes: $supportedItemTypes) { total items { movie { id title { russian original __typename } poster { avatarsUrl fallbackUrl __typename } countries { id name __typename } genres { id name __typename } cast: members(role: [ACTOR], limit: 3) { items { details person { name originalName __typename } __typename } __typename } directors: members(role: [DIRECTOR], limit: 1) { items { details person { name originalName __typename } __typename } __typename } url rating { kinopoisk { isActive count value __typename } expectation { isActive count value __typename } __typename } mainTrailer { id __typename } viewOption { buttonText originalButtonText promotionIcons { avatarsUrl fallbackUrl __typename } isAvailableOnline: isWatchable(filter: {anyDevice: false, anyRegion: false}) purchasabilityStatus subscriptionPurchaseTag type rightholderLogoUrlForPoster availabilityAnnounce { availabilityDate type groupPeriodType announcePromise __typename } __typename } isTicketsAvailable(regionId: $regionId) ... on Film { productionYear duration isShortFilm top250 __typename } ... on TvSeries { releaseYears { start end __typename } seriesDuration totalDuration top250 __typename } ... on MiniSeries { releaseYears { start end __typename } seriesDuration totalDuration top250 __typename } ... on TvShow { releaseYears { start end __typename } seriesDuration totalDuration top250 __typename } ... on Video { productionYear duration isShortFilm __typename } ...MovieListUserData @include(if: $withUserData) __typename } ... on TopMovieListItem { position positionDiff rate votes __typename } ... on MostProfitableMovieListItem { boxOffice { amount __typename } budget { amount __typename } ratio __typename } ... on MostExpensiveMovieListItem { budget { amount __typename } __typename } ... on OfflineAudienceMovieListItem { viewers __typename } ... on PopularMovieListItem { positionDiff __typename } ... on BoxOfficeMovieListItem { boxOffice { amount __typename } __typename } ... on RecommendationMovieListItem { __typename } ... on ComingSoonMovieListItem { releaseDate { date accuracy __typename } __typename } __typename } __typename } __typename } fragment MovieListUserData on Movie { userData { folders { id name public __typename } watchStatuses { notInterested { value __typename } watched { value __typename } __typename } voting { value votedAt __typename } __typename } __typename } fragment DescriptionLink on MovieListMeta { descriptionLink { title url __typename } __typename } fragment OgImage on HtmlMeta { openGraph { image { avatarsUrl __typename } __typename } __typename } fragment FooterConfigData on FooterConfiguration { socialNetworkLinks { icon { avatarsUrl __typename } url __typename } appMarketLinks { icon { avatarsUrl __typename } url __typename } links { title url __typename } __typename } fragment MovieListFeaturingData on MovieListFeaturing { items { title url __typename } __typename } `

	vars := model.GetMoviesListRequest{
		Slug:                      "",
		Platform:                  "DESKTOP",
		RegionId:                  10522,
		WithUserData:              true,
		SupportedFilterTypes:      []string{"BOOLEAN", "SINGLE_SELECT"},
		Filters:                   model.MoviesListFilters{BooleanFilterValues: []string{}, IntRangeFilterValues: []string{}, SingleSelectFilterValues: []string{}, MultiSelectFilterValues: []string{}, RealRangeFilterValues: []string{}},
		SingleSelectFiltersLimit:  250,
		SingleSelectFiltersOffset: 0,
		Limit:                     limit,
		Offset:                    offset,
		MoviesOrder:               "POSITION_ASC",
		SupportedItemTypes:        []string{"COMING_SOON_MOVIE_LIST_ITEM", "MOVIE_LIST_ITEM", "TOP_MOVIE_LIST_ITEM", "POPULAR_MOVIE_LIST_ITEM", "MOST_PROFITABLE_MOVIE_LIST_ITEM", "MOST_EXPENSIVE_MOVIE_LIST_ITEM", "BOX_OFFICE_MOVIE_LIST_ITEM", "OFFLINE_AUDIENCE_MOVIE_LIST_ITEM", "RECOMMENDATION_MOVIE_LIST_ITEM"},
	}

	options := goz.Options{
		JSON: model.Request[model.GetMoviesListRequest]{
			OperationName: operation,
			Query:         query,
			Variables:     vars,
		},
		Query: map[string]interface{}{"operationName": operation},
	}

	response, err := post[model.Response[model.MovieList]](c, ctx, "", options)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting movies list")

		return nil, err
	}

	return response, nil
}
