package api

type QuestionReq struct {
	Token    string `json:"token"`
	Category uint   `json:"category"`
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
	IncorrectAnswers []string `json:"incorrect_answers"`
	Category         uint     `json:"category"`
}

type SessionResp struct {
	ResponseMessage string `json:"response_message"`
	Token           string `json:"token"`
	ResponseCode    uint   `json:"response_code"`
}
