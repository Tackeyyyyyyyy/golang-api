package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"strconv"
)

var (
	tableName = "daily_weight"
	conn, _   = dbr.Open("mysql", "root:@tcp(127.0.0.1:3306)/daily_weight", nil)
	sess      = conn.NewSession(nil)
)

type (
	weightInfo struct {
		ID     int     `db:"id"`
		Weight float32 `db:"weight"`
		Day    string  `db:"measurement_day"`
	}
	weightInfoJSON struct {
		ID     int     `json:"id"`
		Weight float32 `json:"weight"`
		Day    string  `json:"measurement_day"`
	}
	responseData struct {
		Weight []weightInfo `json:"response"`
	}
)

func selectDailyWeightAll(c echo.Context) error {
	var w []weightInfo

	sess.Select("*").From(tableName).Load(&w)
	response := new(responseData)
	response.Weight = w
	return c.JSON(http.StatusOK, response)
}

func InsertWeight(c echo.Context) error {
	w := new(weightInfoJSON)
	if err := c.Bind(w); err != nil {
		return err
	}

	sess.InsertInto(tableName).Columns("weight", "measurement_day").Values(w.Weight, w.Day).Exec()

	var text = "insert is ok."
	return c.JSON(http.StatusOK, text)
}

func updateWeight(c echo.Context) error {
	w := new(weightInfoJSON)
	if err := c.Bind(w); err != nil {
		return err
	}

	attrsMap := map[string]interface{}{"weight": w.Weight}
	sess.Update(tableName).SetMap(attrsMap).Where("id = ?", w.ID).Exec()

	var text = fmt.Sprintf("id: %s is updated.", c.Param("id"))
	return c.JSON(http.StatusOK, text)
}

func deleteWeight(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	sess.DeleteFrom(tableName).
		Where("id = ?", id).
		Exec()
	var text = fmt.Sprintf("id: %s is deleted.", c.Param("id"))
	return c.JSON(http.StatusOK, text)
}

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	api := e.Group("/api/v1")
	{
		api.GET("/weight", selectDailyWeightAll)
		api.POST("/weight", InsertWeight)
		api.PUT("/weight/:id", updateWeight)
		api.DELETE("/weight/:id", deleteWeight)
	}
	// サーバー起動
	e.Start(":1234")
}
