package assessment

import "testing"

func TestAssessmentOneItem(t *testing.T) {
	var assessment = NewAssessment([]AssessmentItem{AssessmentItem{}})
	if len(assessment.Items()) != 1 {
		t.Error("Expected 1 items for new Assessment but got ", len(assessment.Items()))
	}
}

func TestAssessmentItemCountMany(t *testing.T) {
	var assessment = NewAssessment([]AssessmentItem{AssessmentItem{}, AssessmentItem{}, AssessmentItem{}})
	if len(assessment.Items()) != 3 {
		t.Error("Expected 3 items for new Assessment but got ", len(assessment.Items()))
	}
}
