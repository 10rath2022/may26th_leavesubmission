package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type EmployeeLeaveData struct {
	EmployeeID    string             `json:"employeeId"`
	CasualLeaves  int                `json:"casualLeaves"`
	CompOffs      int                `json:"compOffs"`
	EarnedLeaves  int                `json:"earnedLeaves"`
	SickLeaves    int                `json:"sickLeaves"`
	LeaveRequests []LeaveRequestData `json:"leaveRequests"`
}

type LeaveRequestData struct {
	RequestID string `json:"requestId"`
}

func saveLeaveDataHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the JSON data from the request body
	var leaveData EmployeeLeaveData
	err := json.NewDecoder(r.Body).Decode(&leaveData)
	if err != nil {
		log.Println("Failed to decode JSON data:", err)
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	// Handle the leave submission
	err = HandleLeaveSubmission(leaveData)
	if err != nil {
		log.Println("Failed to handle leave submission:", err)
		http.Error(w, "Failed to handle leave submission", http.StatusInternalServerError)
		return
	}

	// Send a success response
	fmt.Fprint(w, "Leave data saved successfully!")
}

func HandleLeaveSubmission(data EmployeeLeaveData) error {
	// Replace the placeholders with your SQL database connection details
	db, err := sql.Open("mysql", "username:password@tcp(hostname:port)/database")
	if err != nil {
		return err
	}
	defer db.Close()

	// Perform the necessary operations to store the data in the database or perform business logic
	// Here you would write the appropriate code to handle the leave submission based on the provided data

	// Example code to print the received data
	fmt.Println("Employee ID:", data.EmployeeID)
	fmt.Println("Casual Leaves:", data.CasualLeaves)
	fmt.Println("Comp Offs:", data.CompOffs)
	fmt.Println("Earned Leaves:", data.EarnedLeaves)
	fmt.Println("Sick Leaves:", data.SickLeaves)
	for _, request := range data.LeaveRequests {
		fmt.Println("Request ID:", request.RequestID)
	}

	return nil
}
