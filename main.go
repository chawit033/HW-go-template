package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       int
}

type Song struct {
	Id       int
	SongName string
	Band     string
	Artist   string
	Album    string
	Genre    string
}

var menus = []Menu{
	{
		Id:          1,
		Name:        "BB Corn",
		Description: "Giant breed of rare corn that was eaten by Gourmet Nobility as a snack long ago",
		Price:       4000,
	},
	{
		Id:          2,
		Name:        "Century Soup",
		Description: "A soup cooked with hundreds or even thousands of ingredients",
		Price:       10000,
	},
	{
		Id:          3,
		Name:        "Jewel Meat",
		Description: "Incandescent lamp-like radiance that dulls jewels and lights up a night sky",
		Price:       8000,
	},
}

var songs = []Song{
	{
		Id:       1,
		SongName: "Whiplash",
		Band:     "Metallica",
		Artist:   "James Hetfield",
		Album:    "KILL'EM ALL",
		Genre:    "Metal",
	},

	{
		Id:       2,
		SongName: "She Wolf",
		Band:     "Megadeth",
		Artist:   "Dave Mustaine",
		Album:    "Wryptic Writings",
		Genre:    "Metal",
	},

	{
		Id:       3,
		SongName: "Half Awake",
		Band:     "Concrete Castles",
		Artist:   "Audra Miller",
		Album:    "Half Awake",
		Genre:    "Pop Rock",
	},

	{
		Id:       4,
		SongName: "Church",
		Band:     "Fall Out Boy",
		Artist:   "Patrick Stump",
		Album:    "M A N I A",
		Genre:    "Pop punk",
	},
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/hello.html")
}

func AllMenusHandler(w http.ResponseWriter, r *http.Request) {
	menuTemplate, err := template.ParseFiles("views/menus/index.html")
	if err != nil {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	menuTemplate.Execute(w, menus)
}

func MenusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuIndex, _ := strconv.ParseInt(params["id"], 0, 64)
	menuIndex -= 1

	menuTemplate, err := template.ParseFiles("views/menus/show.html")
	if err != nil || isOutOfRange(menuIndex) {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	menuTemplate.Execute(w, menus[menuIndex])
}

func AllSongsHandler(w http.ResponseWriter, r *http.Request) {
	songTemplate, err := template.ParseFiles("views/songs/index.html")
	if err != nil {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	songTemplate.Execute(w, songs)
}

func SongsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	songIndex, _ := strconv.ParseInt(params["id"], 0, 64)
	songIndex -= 1

	songTemplate, err := template.ParseFiles("views/songs/show.html")
	if err != nil || SongOutOfRange(songIndex) {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	songTemplate.Execute(w, songs[songIndex])
}

func isOutOfRange(index int64) bool {
	return (index < 0 || index >= int64(len(menus)))
}

func SongOutOfRange(index int64) bool {
	return (index < 0 || index >= int64(len(songs)))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/home", HomeHandler)
	router.HandleFunc("/menus", AllMenusHandler)
	router.HandleFunc("/menus/{id:[0-9]+}", MenusHandler)
	router.HandleFunc("/songs", AllSongsHandler)
	router.HandleFunc("/songs/{id:[0-9]+}", SongsHandler)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.Handle("/", router)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.ListenAndServe(":8000", nil)
}
