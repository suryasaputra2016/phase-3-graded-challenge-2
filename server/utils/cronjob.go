package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/configs"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type Server struct{}

func (s *Server) StartSchedulerJob() error {
	c := cron.New()

	_, err := c.AddFunc("0 0 * * *", func() {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		updateStatus()
	})
	if err != nil {
		log.Printf("Error adding cron job: %v", err)
	}
	c.Start()

	if err != nil {
		return err
	}

	return nil
}

func updateStatus() error {
	// get collection
	collection, err := configs.BookDatabase(context.Background())
	if err != nil {
		return err
	}

	// data container
	var datas []entities.Book

	// get cursor
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return err
	}

	// loop Cursor and get all data
	for cursor.Next(context.Background()) {
		var data entities.Book
		if err := cursor.Decode(&data); err != nil {
			return err
		}
		datas = append(datas, data)
	}
	if err := cursor.Err(); err != nil {
		return err
	}

	// update data version
	for _, data := range datas {
		if data.Status != "overdue" {
			data.Status = "overdue"

			_, err = collection.UpdateOne(
				context.Background(),
				bson.M{"_id": data.ID},
				bson.M{"$set": data},
			)
			if err != nil {
				return err
			}
		}

	}

	fmt.Println("book status updated")
	return nil
}
