package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MessageResponse struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

var videoCollection = getSession().DB("You_Tube").C("Videos")

/*var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Al limite del deseo", 1878, "Geancarlo Listh"},
	Movie{"El interprete del mal", 1978, "JJ Ramirez"},
}*/

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola amo, soy su nuevo servidor en GO...")
}

func Contacto(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Esta es la pagina de contacto, amo.")

}

func ListaPeliculas(w http.ResponseWriter, r *http.Request) {

	var results []Movie

	err := videoCollection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
		SetHeaders(500, w, false)
	}

	fmt.Println("Resultados ", results)

	SetHeaders(200, w, true)
	json.NewEncoder(w).Encode(results)

	//json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		SetHeaders(404, w, false)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	result := Movie{}

	err := videoCollection.FindId(oid).One(&result)

	if err != nil {
		log.Fatal(err)
		SetHeaders(500, w, false)
	}

	fmt.Println("Resultados ", result)

	SetHeaders(200, w, true)
	json.NewEncoder(w).Encode(result)

	fmt.Fprintf(w, "Amo, la pelicula numero  %s ah sido cargada.", movieId)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		SetHeaders(404, w, false)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	err := videoCollection.RemoveId(oid)

	if err != nil {
		log.Fatal(err)
		SetHeaders(500, w, false)
	}

	//results := MessageResponse{"success","Amo, la pelicula con el id: " + movieId + " ha sido removida"}

	results := new(MessageResponse)

	results.SetData("success", "Amo, la pelicula con el id: "+movieId+" ha sido removida")

	SetHeaders(200, w, true)
	json.NewEncoder(w).Encode(results)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		SetHeaders(404, w, false)
		return
	}

	decoder := json.NewDecoder(r.Body)
	oid := bson.ObjectIdHex(movieId)

	movieData := Movie{}

	err := decoder.Decode(&movieData)

	if err != nil {
		log.Fatal(err)
		panic(err)
		SetHeaders(500, w, false)
	}

	defer r.Body.Close()

	fmt.Println("Resultados ", movieData)

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movieData}

	err = videoCollection.Update(document, change)

	if err != nil {

		SetHeaders(404, w, false)
	}
}

func AddMovie(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movieData)

	/*movies = append(movies, movieData)

	log.Println(movies)*/

	err = videoCollection.Insert(movieData)

	if err != nil {
		SetHeaders(500, w, false)
		return
	}

	SetHeaders(200, w, true)

	json.NewEncoder(w).Encode(movieData)

}

func SetHeaders(result int, response http.ResponseWriter, needsContentType bool) {

	if needsContentType {

		response.Header().Set("Content-Type", "application/json")
	}

	response.WriteHeader(result)
}
