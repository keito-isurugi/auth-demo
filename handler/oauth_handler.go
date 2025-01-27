package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/keito-isurugi/auth-demo/helper"
)

type RequestData struct {
	Code string `json:"code"`
}

type ResponseData struct {
	Message string `json:"message"`
	Code string `json:code`
}

func GetAuthCode(w http.ResponseWriter, r *http.Request) {
	// リクエストのBodyからJSONを読み取る
	body, err := helper.GetJsonWithBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// JSONをパース
	var requestData RequestData
	err = helper.GetPearseJson(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// レスポンスデータを作成
	responseData := ResponseData{
		Message: fmt.Sprintf("Received code: %s", requestData.Code),
		Code: "code",
	}

	// JSONレスポンスを返す
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
