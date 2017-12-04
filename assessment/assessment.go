package assessment

type Assessment struct {
	items []AssessmentItem
}

type AssessmentItem struct {
}

func NewAssessment(items []AssessmentItem) Assessment {
	assessment := Assessment{items}
	return assessment
}

func (a Assessment) itemCount() uint32 {
	return 0
}

func (a Assessment) Items() []AssessmentItem {
	return a.items
}
