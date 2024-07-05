package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"ociswrapper/ocis"
)

type BasicResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type CommandResponse struct {
	*BasicResponse
	ExitCode int `json:"exitCode"`
}

func parseJsonBody(reqBody io.ReadCloser) (map[string]any, error) {
	body, _ := io.ReadAll(reqBody)

	if len(body) == 0 || !json.Valid(body) {
		return nil, errors.New("invalid json data")
	}

	var bodyMap map[string]any
	json.Unmarshal(body, &bodyMap)

	return bodyMap, nil
}

func sendResponse(res http.ResponseWriter, success bool, message string) {
	res.Header().Set("Content-Type", "application/json")

	var status string
	if success {
		status = "OK"
		res.WriteHeader(http.StatusOK)
	} else {
		status = "ERROR"
		res.WriteHeader(http.StatusInternalServerError)
	}

	resBody := BasicResponse{
		Status:  status,
		Message: message,
	}

	jsonResponse, _ := json.Marshal(resBody)
	res.Write(jsonResponse)
}

func sendCmdResponse(res http.ResponseWriter, exitCode int, message string) {
	var resBody CommandResponse

	res.WriteHeader(http.StatusOK)
	if exitCode == 0 {
		resBody.Status = "OK"
		resBody.ExitCode = exitCode
		resBody.Message = message
	} else {
		resBody.Status = "ERROR"
		resBody.ExitCode = exitCode
		resBody.Message = message
	}
	res.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(resBody)
	res.Write(jsonResponse)
}

func SetEnvHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	environments, err := parseJsonBody(req.Body)
	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	var message string

	success, _ := ocis.Restart(environments)
	if success {
		message = "oCIS configured successfully"
	} else {
		message = "Failed to restart oCIS with new configuration"
	}

	sendResponse(res, success, message)
}

func RollbackHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var message string

	success, _ := ocis.Restart(nil)
	if success {
		message = "oCIS configuration rolled back successfully"
	} else {
		message = "Failed to restart oCIS with initial configuration"
	}

	sendResponse(res, success, message)
}

func StopOcisHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	success, message := ocis.Stop()
	sendResponse(res, success, message)
}

func CommandHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := parseJsonBody(req.Body)
	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}
	if body["command"] == nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	command, ok := body["command"].(string)
	if !ok || command == "" {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	exitCode, out := ocis.RunCommand(command)

	sendCmdResponse(res, exitCode, out)
}
