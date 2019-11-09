package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"scg-api/models"

	"github.com/gin-gonic/gin"
)

type SCGController struct{}

func (s *SCGController) FindXYZ(c *gin.Context) {

	// X, 5, 9, 15, 23, Y, Z
	// a[n] = n^2 - n + 3

	var solve = func(n float64) float64 {
		return math.Pow(n, 2) - n + 3
	}

	var result = map[string]float64{
		"X": solve(1),
		"Y": solve(6),
		"Z": solve(7),
	}

	c.JSON(http.StatusOK, result)

}

func (s *SCGController) FindPlace(c *gin.Context) {

	type Query struct {
		Key string
	}

	var q Query

	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	key := "AIzaSyDgYnf1thRrNuvGae-xi6b31FOEN_83WI4"
	query := q.Key
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/textsearch/json?query=%s&key=%s", query, key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	var response models.PlaceTextSearchResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if response.Status != models.STATUS_OK {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
