package main

import (
 
	"log"
	"net/http"
	"slices"
   "fmt"
	"github.com/gorilla/mux"
)

//used to store the errors falling in status code range 400 alone , Server errors are responded with a generic message
type APIError struct{
	Error string `json:"error"`
}

type apiFunc func(http.ResponseWriter, *http.Request)error


func makeHttpHandler(f apiFunc)http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request){
if err:=f(w,r);err!=nil{
	if err:=WriteJSON(w,http.StatusBadRequest,APIError{Error: err.Error()});err!=nil{
		log.Fatal(err)
	}
} 
}
}


type APIServer struct{
	listenAddr string
   
}

func NewAPIServer(listenAddr string)*APIServer{
return &APIServer{
	listenAddr: listenAddr,
 
}
}


func (s *APIServer) Run()error{
router:=mux.NewRouter()
router.Handle("/numbers/{numberid}",makeHttpHandler(s.getNumbers))
err:=http.ListenAndServe(s.listenAddr,router);if err!=nil{
	return err
}
return nil
}


func (s *APIServer) getNumbers(w http.ResponseWriter, r *http.Request)error{
numberid := mux.Vars(r)["numberid"]
if !slices.Contains(Types, numberid){
return fmt.Errorf("id not allowed!:%s",numberid)
}

return nil
}