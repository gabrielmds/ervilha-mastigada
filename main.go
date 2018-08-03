package main

import (
	"fmt"
	"net/http"

	"./assessment"
	"log"
)

var assessmtn = []assessment.Assessment{}

func init() {
	assessmtn = []assessment.Assessment{
		assessment.Assessment{
			Title: "test",
			Questions: []assessment.Question{
				assessment.Question{
					Text: "Question 1",
					Alternatives: []assessment.Alternative{
						assessment.Alternative{
							Text:    "alt 1",
							Correct: true,
						},
						assessment.Alternative{
							Text:    "alt 2",
							Correct: false,
						},
						assessment.Alternative{
							Text:    "alt 3",
							Correct: false,
						},
						assessment.Alternative{
							Text:    "alt 4",
							Correct: false,
						},
					},
				},
				assessment.Question{
					Text: "Question 2",
					Alternatives: []assessment.Alternative{
						assessment.Alternative{
							Text:    "alt 1",
							Correct: false,
						},
						assessment.Alternative{
							Text:    "alt 2",
							Correct: false,
						},
						assessment.Alternative{
							Text:    "alt 3",
							Correct: true,
						},
						assessment.Alternative{
							Text:    "alt 4",
							Correct: false,
						},
					},
				},
			},
		},
	}
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
			for idx, assesment := range assessmtn {
				fmt.Fprint(w, fmt.Sprintf("%d) %v <br>", idx, assesment));
			}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
