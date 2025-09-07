package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"tmdb_parser/db/repository/schemas"
)

type (
	MovieRepository struct {
		collection *mongo.Collection
	}
)

func NewMovieRepository(db *mongo.Database) *MovieRepository {
	repo := &MovieRepository{collection: db.Collection(schemas.Movie{}.CollectionName())}

	return repo
}

func (r MovieRepository) GetWithoutTmdbID(ctx context.Context, nextCursor bson.ObjectID) ([]schemas.Movie, error) {
	logger := zerolog.Ctx(ctx)

	var movies []schemas.Movie

	filter := bson.M{"externalId.tmdb": nil, "isTmdbChecked": false, "votes.kp": bson.M{"$gt": 10}}
	if nextCursor != bson.NilObjectID {
		filter["_id"] = bson.M{"$gt": nextCursor}
	}

	opts := options.Find().SetSort(bson.D{{Key: "year", Value: -1}, {Key: "_id", Value: 1}}).SetLimit(10000)
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		logger.Err(err).Send()
		return nil, err
	}

	if err := cursor.All(ctx, &movies); err != nil {
		logger.Err(err).Send()
		return nil, err
	}

	return movies, nil
}

func (r MovieRepository) SetExternalIDs(ctx context.Context, id bson.ObjectID, tmdbID int) error {
	logger := zerolog.Ctx(ctx)

	update := bson.M{
		"$set": bson.M{
			"externalId.tmdb": tmdbID,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		logger.Err(err).Send()
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("movie not found")
	}

	return nil
}

func (r MovieRepository) SetTmdbChecked(ctx context.Context, id bson.ObjectID) error {
	logger := zerolog.Ctx(ctx)

	update := bson.M{
		"$set": bson.M{
			"isTmdbChecked": true,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		logger.Err(err).Send()
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("movie not found")
	}

	return nil
}
