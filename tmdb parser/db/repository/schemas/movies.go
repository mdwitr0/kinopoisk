package schemas

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	Movie struct {
		OID bson.ObjectID `bson:"_id,omitempty"`
		ID  int           `bson:"id,omitempty"`

		Title            string `bson:"name,omitempty"`
		AlternativeTitle string `bson:"alternativeName,omitempty"`

		ExternalId    ExternalId `bson:"externalId,omitempty"`
		IsTmdbChecked bool       `bson:"isTmdbChecked,omitempty"`
		IsSeries      bool       `bson:"isSeries,omitempty"`
		Year          int        `bson:"year,omitempty"`
	}

	ExternalId struct {
		KpHD string `json:"kpHD"`
		Imdb string `json:"imdb"`
		Tmdb int    `json:"tmdb"`
	}
)

func (t Movie) CollectionName() string {
	return "movies"
}
