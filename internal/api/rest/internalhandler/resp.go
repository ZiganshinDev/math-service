package internalhandler

type Problems struct {
	Problems []Problem `json:"problems"`
}

type Problem struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Solution string `json:"solution"`
	Answer   string `json:"answer"`
}
