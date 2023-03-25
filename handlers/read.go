package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/JhonyAltoe/micro-service-read-service/entities"
	"github.com/JhonyAltoe/micro-service-read-service/logs"
	"github.com/JhonyAltoe/micro-service-read-service/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var Repo = repositories.New()

func GetServices(c *gin.Context) {
	company := c.Param("company")

	logrus.Warn("Reading Schedules on %v", company)
	var services entities.Company

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Repo.Client.Database(company).Collection("company")
	result := collection.FindOne(
		ctx,
		bson.M{},
	)

	if result.Err() != nil {
		logrus.Errorf("Error Reading services on MongoDB: %v %v", company, result.Err().Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Err().Error()})
		return
	}

	err := result.Decode(&services)

	if err != nil {
		logrus.Errorf("Error Reading services on MongoDB: %v %v", company, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer logrus.Info("Query Complete successfully on ", company)
	defer logs.Elapsed("UpdateSchedule")()

	c.JSON(http.StatusOK, services)
}
