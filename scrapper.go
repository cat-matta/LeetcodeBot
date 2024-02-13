package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
    // "html"
    // "bytes"
    // "log"
    // "io/ioutil"

    // "net/http"
    // "net/url"
    "encoding/json"
)


// func main() {

// 	parse("hudson-river-trading")
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         // fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//     // })
// }


type problem map[string]interface{}

func getDaily(){
	url := "https://leetcode.com/problemset/all/"
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	// table := doc.Find("table")
	// elements := table.FindAll("tr")
	fmt.Println(doc)


}

func getProblemsTest() {
	query := `
		query dailyCodingQuestionRecords($year: Int!, $month: Int!) {
			dailyCodingChallengeV2(year: $year, month: $month) {
				challenges {
					date
					userStatus
					link
					question {
						questionFrontendId
						title
						titleSlug
					}
				}
				weeklyChallenges {
					date
					userStatus
					link
					question {
						questionFrontendId
						title
						titleSlug
						isPaidOnly
					}
				}
			}
		}
	`
	variables := map[string]interface{}{
		"year":  2024,
		"month": 2,
	}

	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshalling data: %s\n", err)
		return
	}

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	fmt.Println("Response:", string(body))
}

// func getDaily(){
// 	data := `{"query":"query questionOfToday {\n\tactiveDailyCodingChallengeQuestion {\n\t\tdate\n\t\tlink\n\t\tquestion {\n\t\t\tacRate\n\t\t\tdifficulty\n\t\t\tfreqBar\n\t\t\ttitle\n\t\t\ttopicTags {\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n}\n","operationName":"questionOfToday"}`
// 	json_str, _ := json.Marshal(data)
// 	b := bytes.NewBuffer(json_str)
// 	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", b)
// 	req.Header.Set("Referer", "https://leetcode.com/")
// 	req.Header.Set("Cookie", "csrftoken=Brk3S7hrQzzqD2hsXJp53WxWmoUkKLYfAOSJAzsgeOXfuvpP68nPjCr4sXixg0RT; __stripe_mid=ba437637-6583-4bf0-bc60-dc5eddd64056da29b9; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjYxNzQ3MCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImFsbGF1dGguYWNjb3VudC5hdXRoX2JhY2tlbmRzLkF1dGhlbnRpY2F0aW9uQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjY3MmUxZDgzNDMxYzg1NDUxNDQ5NTA4OGQ5M2ZiNjVmYjgzM2JlZmIiLCJpZCI6MjYxNzQ3MCwiZW1haWwiOiJjYXRoZXJpbmUubWF0dGEyMjM5OUBnbWFpbC5jb20iLCJ1c2VybmFtZSI6ImNhdG1hdHRhIiwidXNlcl9zbHVnIjoiY2F0bWF0dGEiLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jb20vdXNlcnMvY2F0bWF0dGEvYXZhdGFyXzE2MjIyMTY1NTIucG5nIiwicmVmcmVzaGVkX2F0IjoxNjc1MjkwMjM0LCJpcCI6IjY3Ljg2LjExNC41NSIsImlkZW50aXR5IjoiMWU5ZGUyZWJhNjQwZThlZWE5ODUxYzE0MzRjNzU4ODYiLCJzZXNzaW9uX2lkIjozMDI4MzgxMH0.YGCSMxE2EN5MPgHtnosLe4Jz_-4uW-BOwY4SCbrxXrU; NEW_PROBLEMLIST_PAGE=1; _dd_s=rum=1&id=0fe9fee6-abe5-4ab5-bbf2-8f403fe1a59f&created=1675375375076&expire=1675376932353")
// 	req.Header.Set("Content-Type", "application/json")

// 	// var res map[string]interface{}

//     // json.NewDecoder(resp.Body).Decode(&res)
//     client := &http.Client{}
//     resp, err := client.Do(req)
// 	if err != nil {
//         panic(err)
//     }
//     defer resp.Body.Close()

//     fmt.Println("response Status:", resp.Status)
//     fmt.Println("response Headers:", resp.Header)
//     body, _ := ioutil.ReadAll(resp.Body)
//     fmt.Println("response Body:", string(body))
// }

func parse(company string) string{
	link := fmt.Sprintf("https://leetcode-company-tagged.vercel.app/%s/", company)
	resp, err := soup.Get(link)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	table := doc.Find("table")
	elements := table.FindAll("tr")

	var problems []problem


	for  i := 1; i < len(elements); i++ {
		item := elements[i]
		row := item.FindAll("td")
		// out.cut(0,4)
		data := []string{}

		for _, thing := range row{
			url := thing.Find("a")

			if(thing.Text() != ""){
				// fmt.Println(thing.Text())
				data = append(data, thing.Text())

			}
			if(url.Error == nil){
				// fmt.Println(url.Attrs()["href"])
				data = append(data, url.Attrs()["href"])
			}
		}
		// fmt.Println(data)
		individual := problem{
			"Number": data[0],
			"Name": data[1],
			"Link": data[2],
			"Difficulty": data[3],
			"Frequency": data[4],
		
		}
		problems = append(problems, individual)

	}
		jsonData, err := json.Marshal(problems)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return "Error parsing"
		}

		return fmt.Sprintf("%s",jsonData)
}
