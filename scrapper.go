package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
    "bytes"
    "io/ioutil"
    "net/http"
    "encoding/json"
)



type problem map[string]interface{}

func getAllProblems() string{
	url := "https://leetcode.com/problemset/all/"
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	// table := doc.Find("table")
	// elements := table.FindAll("tr")
	return fmt.Sprintf("%s", doc)


}

func getProblemsTest() string{
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
		 return fmt.Sprintf("Error marshalling data: %s\n", err)
		
	}

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Sprintf("Error creating request: %s\n", err)
		
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error making request: %s\n", err)
		
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %s\n", err)
		
	}

	return fmt.Sprintf("Response:", string(body))
}

func getDailyProblems() string{
	query := `
		query questionOfToday {
			activeDailyCodingChallengeQuestion {
				date
				link
				question {
					acRate
					difficulty
					freqBar
					title
					topicTags {
						name
					}
				}
			}
		}`
	

	payload := map[string]interface{}{
		"query":     query,
		"variables": "{}",
		"operationName": "questionOfToday",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Sprintf("Error marshalling data: %s\n", err)
	}

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Sprintf("Error creating request: %s\n", err)
		
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error making request: %s\n", err)
		
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %s\n", err)
		
	}

	return fmt.Sprintf("%s",body)
}

func getCompanyProblems(company string) string{
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
		data := []string{}

		for _, thing := range row{
			url := thing.Find("a")

			if(thing.Text() != ""){
				data = append(data, thing.Text())
			}
			if(url.Error == nil){
				data = append(data, url.Attrs()["href"])
			}
		}
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

		return fmt.Sprintf("%s",string(jsonData))
}
