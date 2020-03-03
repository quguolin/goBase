package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"goBase/httpBind/binding"
)

// User contains user information
type User struct {
	Age   uint8     `form:"age" validate:"required,gte=0,lte=130" json:"age"`
	RID   []string  `form:"rid,split" validate:"required" json:"rid"`
	Email string    `form:"email" validate:"required,email"`
	Begin time.Time `form:"begin" time_format:"2006-01-02 15:04:05" json:"begin"`
	End   time.Time `form:"end" time_format:"2006-01-02 15:04:05" json:"end"`
	PN    int       `form:"pn" default:"1" json:"pn"`
	PS    int       `form:"ps" default:"20" json:"ps"`
}

func newServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bindValid(w, r)
	})
	log.Println("Starting  server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func bindValid(w http.ResponseWriter, req *http.Request) {
	user := &User{}
	val := &binding.Validator{}
	err := val.Bind(req, user)
	if err != nil {
		w.Write([]byte(string(err.Error())))
		return
	}
	bs, _ := json.Marshal(user)
	w.Write([]byte(bs))
}
func main() {
	newServer()
}
