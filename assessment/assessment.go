package assessment


type Alternative struct{
	Text string
	Correct bool
}

type Question struct {
	Text string
	Alternatives []Alternative
}

type Assessment struct{
	Title string
	Questions []Question
}


