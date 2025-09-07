package model

type TVResponse struct {
	Adult          bool          `json:"adult"`
	BackdropPath   string        `json:"backdrop_path"`
	CreatedBy      []interface{} `json:"created_by"`
	EpisodeRunTime []int         `json:"episode_run_time"`
	FirstAirDate   string        `json:"first_air_date"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int      `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		ID             int         `json:"id"`
		Name           string      `json:"name"`
		Overview       string      `json:"overview"`
		VoteAverage    int         `json:"vote_average"`
		VoteCount      int         `json:"vote_count"`
		AirDate        string      `json:"air_date"`
		EpisodeNumber  int         `json:"episode_number"`
		EpisodeType    string      `json:"episode_type"`
		ProductionCode string      `json:"production_code"`
		Runtime        int         `json:"runtime"`
		SeasonNumber   int         `json:"season_number"`
		ShowID         int         `json:"show_id"`
		StillPath      interface{} `json:"still_path"`
	} `json:"last_episode_to_air"`
	Name             string      `json:"name"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air"`
	Networks         []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int           `json:"number_of_episodes"`
	NumberOfSeasons     int           `json:"number_of_seasons"`
	OriginCountry       []string      `json:"origin_country"`
	OriginalLanguage    string        `json:"original_language"`
	OriginalName        string        `json:"original_name"`
	Overview            string        `json:"overview"`
	Popularity          float64       `json:"popularity"`
	PosterPath          string        `json:"poster_path"`
	ProductionCompanies []interface{} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	Seasons []struct {
		AirDate      string  `json:"air_date"`
		EpisodeCount int     `json:"episode_count"`
		ID           int     `json:"id"`
		Name         string  `json:"name"`
		Overview     string  `json:"overview"`
		PosterPath   string  `json:"poster_path"`
		SeasonNumber int     `json:"season_number"`
		VoteAverage  float64 `json:"vote_average"`
	} `json:"seasons"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Type        string  `json:"type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

type TVSearchResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIDs         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalTitle    string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		Title            string   `json:"name"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int      `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}
