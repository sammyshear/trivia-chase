package api

type QuestionReq struct {
	Token    string `json:"token"`
	Category string `json:"category"`
}

type AnswerReq struct {
	Answer   string `json:"answer"`
	Question string `json:"question"`
}

type QuestionResp struct {
	Results      []TextQuestion `json:"results"`
	ResponseCode uint           `json:"response_code"`
}

type TextQuestion struct {
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	Type             string   `json:"type"`
	Category         string   `json:"category"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type SessionResp struct {
	ResponseMessage string `json:"response_message"`
	Token           string `json:"token"`
	ResponseCode    uint   `json:"response_code"`
}
