package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type EventType int64

const (
	Created EventType = 0
	Updated EventType = 1
	Removed EventType = 2
)

type Session struct {
	Id        string
	SessionId string
	UserId    string
	Timestamp int64
	EventType EventType
}

func storeSession(session Session) {
	events := DB.Collection("session-events")

	_, err := events.InsertOne(context.TODO(), bson.D{{"id", session.Id}, {"sessionId", session.SessionId},
		{"userId", session.UserId}, {"eventType", session.EventType}, {"timestamp", session.Timestamp}})
	if err != nil {
		log.Fatal(err)
	}
}

func GetSessions() ([]Session, error) {

	ctx := context.Background()

	sessions := []Session{}
	cur, err := DB.Collection("session-events").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var session Session
		cur.Decode(&session)
		sessions = append(sessions, session)
	}

	defer cur.Close(ctx)

	return sessions, err
}
