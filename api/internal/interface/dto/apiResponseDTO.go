package dto

// APIレスポンスのステータスを表す定数
const (
	StatusSuccess = "success" // 成功
	StatusError   = "error"   // エラー（クライアント側の問題）
	StatusFail    = "fail"    // 失敗（サーバー側の問題）
)

// APIResponseDTO は全てのAPIレスポンスの共通構造
type APIResponseDTO struct {
	Status  string      `json:"status"`            // success, error, fail
	Message string      `json:"message,omitempty"` // エラーや成功メッセージ
	Data    interface{} `json:"data,omitempty"`    // レスポンスデータ
	Error   string      `json:"error,omitempty"`   // エラー詳細（開発用）
}

// SuccessResponse 成功レスポンスを作成
func SuccessResponse(data interface{}, message string) *APIResponseDTO {
	return &APIResponseDTO{
		Status:  StatusSuccess,
		Data:    data,
		Message: message,
	}
}

// ErrorResponse クライアントエラーレスポンスを作成（400系エラー）
func ErrorResponse(message string, err error) *APIResponseDTO {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	return &APIResponseDTO{
		Status:  StatusError,
		Message: message,
		Error:   errorMessage,
	}
}

// FailResponse サーバーエラーレスポンスを作成（500系エラー）
func FailResponse(message string, err error) *APIResponseDTO {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	return &APIResponseDTO{
		Status:  StatusFail,
		Message: message,
		Error:   errorMessage,
	}
}
