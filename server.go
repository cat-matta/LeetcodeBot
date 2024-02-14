package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	  "strings"
)

func main() {
	// fmt.Printf(parse("hudson-river-trading"))
	// getDaily()
    // getProblemsTest()
	// parse("hudson-river-trading")
		http.HandleFunc("/", getRoot)
		http.HandleFunc("/daily", getDaily)
		http.HandleFunc("/company", getCompany)
		http.HandleFunc("/test", getTest)
		// http.HandleFunc("/all", getAll)

		err := http.ListenAndServe(":3333", nil)
  	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)

	    }
	}
	func getRoot(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got / request\n")
		io.WriteString(w, "Usage:\n/daily: gets the leetcode daily problem\ncompany?name={name}: gets the list of problems for a company\n/test: all the problems")
	}

    func getDaily(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got /daily request\n")
		resp := getDailyProblems()
		io.WriteString(w, resp)
	}
	func getCompany(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		fmt.Printf("got /company?name=%s request\n", r)
		resp := getCompanyProblems(strings.ToLower(param))
		io.WriteString(w, resp)
	}
	func getTest(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got /test request\n")
		resp := getProblemsTest()
		io.WriteString(w, resp)
	}
	// func getAll(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("got /all request\n")
	// 	resp := getAllProblems()
	// 	io.WriteString(w, resp)
	// }		

