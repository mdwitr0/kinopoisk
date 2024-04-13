package model

type GetMoviesListRequest struct {
	Slug                      string            `json:"slug"`
	Platform                  string            `json:"platform"`
	RegionId                  int               `json:"regionId"`
	WithUserData              bool              `json:"withUserData"`
	SupportedFilterTypes      []string          `json:"supportedFilterTypes"`
	Filters                   MoviesListFilters `json:"filters"`
	SingleSelectFiltersLimit  int               `json:"singleSelectFiltersLimit"`
	SingleSelectFiltersOffset int               `json:"singleSelectFiltersOffset"`
	Limit                     int               `json:"moviesLimit"`
	Offset                    int               `json:"moviesOffset"`
	MoviesOrder               string            `json:"moviesOrder"`
	SupportedItemTypes        []string          `json:"supportedItemTypes"`
}

type MoviesListFilters struct {
	BooleanFilterValues      []string `json:"booleanFilterValues"`
	IntRangeFilterValues     []string `json:"intRangeFilterValues"`
	SingleSelectFilterValues []string `json:"singleSelectFilterValues"`
	MultiSelectFilterValues  []string `json:"multiSelectFilterValues"`
	RealRangeFilterValues    []string `json:"realRangeFilterValues"`
}

type MovieList struct {
	List MovieListBySlug `json:"movieListBySlug"`
}

type MovieListBySlug struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cover       struct {
		AvatarsUrl string `json:"avatarsUrl"`
		Typename   string `json:"__typename"`
	} `json:"cover"`
	Movies struct {
		Total    int    `json:"total"`
		Items    []Item `json:"items"`
		Typename string `json:"__typename"`
	} `json:"movies"`
}

type Item struct {
	Movie    MovieID `json:"movie"`
	Position int     `json:"position"`
	Typename string  `json:"__typename"`
}
