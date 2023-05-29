package main

import (
	"log"
	"net/http"

	"github.com/your-username/may26th_leavesubmission/leavedata"
)

func main() {
	http.HandleFunc("/save-leave-data", leavedata.SaveLeaveDataHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
