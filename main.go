package main

import (
	"html/template"
	"log"
	"net/http"

	"./assessment"
)

var assessmtn = []assessment.Assessment{}

var templates = template.Must(template.ParseGlob("assessment/templates/*"))

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

	http.HandleFunc("/assessments", func(w http.ResponseWriter, req *http.Request) {
		err := templates.ExecuteTemplate(w, "assessment_list.html", assessmtn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
