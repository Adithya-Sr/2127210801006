package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"

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
if numberid=="p"{
requestURL:= "http://20.244.56.144/test/primes"
//using the token directly due to time constraints else wouldve put it in an .env file
bearer := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzI0MDQ0MzE1LCJpYXQiOjE3MjQwNDQwMTUsImlzcyI6IkFmZm9yZG1lZCIsImp0aSI6IjNlNTU1ZjFiLTZmNWItNGE3YS05OTAwLTc4MWJlMmZlZTU1ZSIsInN1YiI6IjIwMjFpdDAzNjJAc3ZjZS5hYy5pbiJ9LCJjb21wYW55TmFtZSI6IlNyaSBWZW5rYXRlc3dhcmEgQ29sbGVnZSBPZiBFbmdpbmVlcmluZyIsImNsaWVudElEIjoiM2U1NTVmMWItNmY1Yi00YTdhLTk5MDAtNzgxYmUyZmVlNTVlIiwiY2xpZW50U2VjcmV0IjoibWNUZ29ZVnllU2FtQnVyeCIsIm93bmVyTmFtZSI6IkFkaXRoeWEgU3Jpa2FudGgiLCJvd25lckVtYWlsIjoiMjAyMWl0MDM2MkBzdmNlLmFjLmluIiwicm9sbE5vIjoiMjEyNzIxMDgwMTAwNiJ9.MCWJ8kJLAf_S7qf0tDTsrIJGu9FdvjbWQy7Q8Q3EiCo"
req, err := http.NewRequest("GET",requestURL, nil)
if err!=nil{
	return err
}
req.Header.Add("Authorization", bearer)
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
if err!=nil{
	return err
}
return WriteJSON(w,http.StatusOK,string(body))	
}

return nil
}