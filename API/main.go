package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type pessoa struct {
    ID     string  `json:"id"`
    Nome  string  `json:"nome"`
    Id_Profissao string  `json:"profissao"`
}

type projeto struct {
    ID     string  `json:"id"`
    Nome  string  `json:"nome"`
    Participantes []pessoa `json:"?????"`
}

// albums slice to seed record album data.
var pessoas = []pessoa{
    {ID: "1", Nome: "Bruno de Calcinha", Id_Profissao: "45"},
    {ID: "2", Nome: "Pedro Pelado", Id_Profissao: "12"},
    {ID: "3", Nome: "Caio de Sunga", Id_Profissao: "13"},
}

func main() {
    router := gin.Default()
    router.GET("/pessoas", getPessoas)

    router.POST("/pessoas", postPessoas)
    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getPessoas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, pessoas)
}

func postPessoas(c *gin.Context) {
    var newAlbum pessoa

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    pessoas = append(pessoas, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func postProjetos(c *gin.Context) {
    var newAlbum pessoa

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    pessoas = append(pessoas, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
