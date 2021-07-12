package main

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var mainDB *mongo.Database

func findExperimentByID(experiments *ExperimentSlice, experimentID string) (*Experiment, int, error) {
	log.Println("Seaching for the id", experimentID)
	for index, experiment := range experiments.Experiments {
		if experiment.ID == experimentID {
			return &experiment, index, nil
		}
	}
	return nil, -1, errors.New("No experiment found with id")
}

func addMetricsToExperiment(experiments *ExperimentSlice,
	experimentID, metricName string, metricValue float32) {
	exp, _, err := findExperimentByID(experiments, experimentID)
	if err != nil {
		log.Println(err)
	}
	exp.Metrics[metricName] = append(exp.Metrics[metricName], metricValue)
}

func deleteExperiment(experiments *ExperimentSlice, experimentID string) error {
	_, index, err := findExperimentByID(experiments, experimentID)
	if err != nil {
		log.Println(err)
		return err
	}
	l := experiments.ExperimentCount
	experiments.Experiments[index] = experiments.Experiments[l-1]
	// We do not need to put experiment[i] at the end, as it will be discarded anyway
	experiments.Experiments = experiments.Experiments[:l-1]
	experiments.ExperimentCount = l - 1
	return nil
}

func connectMongo() {
	// cred := options.Credential{
	// 	Username: "root",
	// 	Password: "test",
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()
	// defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panicln(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	mainDB = client.Database("experiments")

	log.Println("Successfully connected and pinged.")
}

func addExperimentToMongo(ex Experiment) error {
	if client == nil && mainDB == nil {
		log.Println("EMPTY CLIENT")
		return nil
	}
	// db := client.Database("experiments")
	collection := mainDB.Collection("new-experiment")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, err := collection.InsertOne(ctx, ex)
	return err
}
