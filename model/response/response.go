package res

type MessageResponse struct {
	Message string `json:"message"`
}

func APIMessage(message string) MessageResponse {
	return MessageResponse{Message: message}
}

func APIBadRequest() MessageResponse {
	return APIMessage("Bad Request.")
}

func NoTokenProvider() MessageResponse {
	return APIMessage("No Token Provider.")
}
