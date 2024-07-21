package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func OpenSession(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://opentdb.com/api_token.php?command=request")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error requesting session: %s", err), 500)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading session token: %s", err), 500)
		return
	}
	var session SessionResp
	json.Unmarshal(body, &session)

	w.Header().Add("content-type", "text/plain")
	w.Write([]byte(session.Token))
	w.WriteHeader(200)
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request body: %s", err), 421)
		return
	}

	var body QuestionReq
	json.Unmarshal(req, &body)

	res, err := http.Get(fmt.Sprintf("https://opentdb.com/api.php?amount=1&category=%d&type=multiple&token=%s", body.Category, body.Token))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to request questions: %s", err), 500)
		return
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse question response: %s", err), 500)
		return
	}
	var questionResp QuestionResp
	json.Unmarshal(b, &questionResp)
	if questionResp.ResponseCode > 0 {
		http.Error(w, "Error with question response", 500)
		return
	}
	result, err := json.Marshal(questionResp.Results[len(questionResp.Results)-1])
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding question: %s", err), 500)
	}

	w.Header().Add("content-type", "application/json")
	w.Write(result)
	w.WriteHeader(200)
}
