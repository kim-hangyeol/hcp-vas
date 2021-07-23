package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

type aksData struct {
	ResourceGroupName string
	ResourceName      string
	ApiVersion        string
}


func joinFunc(w http.ResponseWriter, req *http.Request){
	fmt.Println("---connected VAS---")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func aksGetCredential(w http.ResponseWriter, req *http.Request){
	fmt.Println("---request aks Data---")  
	
	aksData := &aksData{
		ResourceGroupName: "keti",
		ResourceName:      "hcp-cluster",
		ApiVersion:        "2021-03-01",
	}
	jsonData, _ := json.Marshal(aksData)
	w.Write(jsonData)	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main(){
	fmt.Println("main Start")
	http.HandleFunc("/join", joinFunc)
	http.HandleFunc("/aksGetCredential", aksGetCredential)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("main End")
}
