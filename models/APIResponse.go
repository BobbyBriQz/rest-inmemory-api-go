package models

type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewAPISuccessResponse(m string, d interface{}) APIResponse {
	return APIResponse{
		Status:  true,
		Message: m,
		Data:    d,
	}
}

func NewAPIFailedResponse(m string) APIResponse {
	return APIResponse{
		Status:  false,
		Message: m,
	}
}

func (a *APIResponse) SetStatus(i bool) {
	a.Status = i
}
func (a *APIResponse) SetMessage(i string) {
	a.Message = i
}
func (a *APIResponse) SetData(i interface{}) {
	a.Data = i
}
