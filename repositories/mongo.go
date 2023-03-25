package repositories

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Client *mongo.Client
	Ctx context.Context
	CtxCancel context.CancelFunc
	ClientOptions *options.ClientOptions
}

func New() *Repository {
	logrus.Warn("Creating new repository")
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	options := options.Client().ApplyURI("mongodb+srv://omicronrpl:oIe3hKERFNXPtiyF@hsa.sqnczja.mongodb.net/?authSource=admin")
	client, err := mongo.Connect(ctx, options)

	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}

	defer logrus.Info("Repository Created")

	return &Repository{
		Client: client,
		Ctx: ctx,
		CtxCancel: cancel,
		ClientOptions: options,
	}
}

func (r *Repository) Ping() {
	logrus.Warn("Pinging MongoDB ...")
	if err := r.Client.Ping(r.Ctx, r.ClientOptions.ReadPreference); err != nil {
		logrus.Fatal("Mongo Database Unavailable")
		panic(err)
	}
}
