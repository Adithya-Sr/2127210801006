package main

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter,status int, v any)error{
w.Header().Add("Content-Type","application/json")
w.WriteHeader(status)
if err:=json.NewEncoder(w).Encode(v);err!=nil{
	return err
}
return nil
}
