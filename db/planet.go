package db

import (
	"context"
	"encoding/json"

	"github.com/WeslleyRibeiro-1999/api-star-wars/models"
	"github.com/WeslleyRibeiro-1999/api-star-wars/swapi"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cria planeta por json
func CreatePlanet(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(DbName, CollectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var planet models.Planet

	json.Unmarshal([]byte(c.Body()), &planet)
	apperancesFilmes := swapi.GetApparances(planet.Nome)
	planet.Qtde = int32(apperancesFilmes)

	res, err := collection.InsertOne(context.Background(), planet)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

// Busca todos os planetas ou pelo ID
func GetPlanetID(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(DbName, CollectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}

// Deleta o planeta por ID
func DeletePlanet(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(DbName, CollectionName)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)
}
