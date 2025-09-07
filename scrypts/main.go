package main

import (
	"context"
	"github.com/mattn/go-colorable"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/semaphore"
	"log"
	"scripts/config"
	"sync"
)

var (
	personCollection    = "people"
	moviesCollection    = "movies"
	movieIDsCollection  = "movies_ids"
	personIDsCollection = "person_ids"
)

type (
	MovieID struct {
		ID int `bson:"movie_id"`
	}
	PersonID struct {
		ID int `bson:"person_id"`
	}
	Person struct {
		Movies []struct {
			ID int `bson:"id"`
		} `bson:"movies"`
	}
	Movie struct {
		Persons []struct {
			ID int `bson:"id"`
		} `bson:"persons"`
	}
)

func main() {
	zapCfg := zap.NewDevelopmentEncoderConfig()
	zapCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapCfg),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))

	zap.ReplaceGlobals(logger)

	mongoConfig := config.NewMongoConfig()

	clientOptions := options.Client().ApplyURI(mongoConfig.URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	mongodb := client.Database(mongoConfig.DB)

	loadPersonIDsFromMovies(context.Background(), mongodb)
}

func loadMoviesIDsFromPersons(ctx context.Context, db *mongo.Database) {
	logger := zap.L()

	var (
		maxWorkers = 100
		sem        = semaphore.NewWeighted(int64(maxWorkers))
		counter    struct {
			sync.Mutex
			C int
		}
	)

	cursor, err := db.Collection(personCollection).Find(ctx, bson.D{})
	if err != nil {
		logger.Error(err.Error())
	}

	for cursor.Next(ctx) {
		if err := sem.Acquire(ctx, 1); err != nil {
			logger.Error(err.Error(), zap.Int("counter", counter.C))
			continue
		}

		go func(cursor *mongo.Cursor) {
			defer sem.Release(1)

			var person Person
			if err := cursor.Decode(&person); err != nil {
				logger.Error(err.Error(), zap.Int("counter", counter.C))
				return
			}

			movieIDs := make([]interface{}, len(person.Movies))
			for i, movie := range person.Movies {
				movieIDs[i] = MovieID{ID: movie.ID}
			}

			if len(movieIDs) > 0 {
				if _, err := db.Collection(movieIDsCollection).InsertMany(ctx, movieIDs); err != nil {
					logger.Error(err.Error(), zap.Int("counter", counter.C))
					return
				}

			}

			logger.Debug("success inserted", zap.Int("movieIDs", len(movieIDs)), zap.Int("counter", counter.C))

			counter.Lock()
			counter.C++
			counter.Unlock()
		}(cursor)
	}
}

func loadPersonIDsFromMovies(ctx context.Context, db *mongo.Database) {
	logger := zap.L()

	var (
		maxWorkers = 100
		sem        = semaphore.NewWeighted(int64(maxWorkers))
		counter    struct {
			sync.Mutex
			C int
		}
	)

	cursor, err := db.Collection(moviesCollection).Find(ctx, bson.D{})
	if err != nil {
		logger.Error(err.Error())
	}

	for cursor.Next(ctx) {
		if err := sem.Acquire(ctx, 1); err != nil {
			logger.Error(err.Error(), zap.Int("counter", counter.C))
			continue
		}

		go func(cursor *mongo.Cursor) {
			defer sem.Release(1)

			var movie Movie
			if err := cursor.Decode(&movie); err != nil {
				logger.Error(err.Error(), zap.Int("counter", counter.C))
				return
			}

			personIDs := make([]interface{}, len(movie.Persons))
			for i, person := range movie.Persons {
				personIDs[i] = PersonID{ID: person.ID}
			}

			if len(personIDs) > 0 {
				if _, err := db.Collection(personIDsCollection).InsertMany(ctx, personIDs); err != nil {
					logger.Error(err.Error(), zap.Int("counter", counter.C))
					return
				}

			}

			logger.Debug("success inserted", zap.Int("personIDs", len(personIDs)), zap.Int("counter", counter.C))

			counter.Lock()
			counter.C++
			counter.Unlock()
		}(cursor)
	}
}
