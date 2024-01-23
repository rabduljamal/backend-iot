package lib

type DataInput struct {
	Data map[string]interface{} `json:"data"`
}

type MyResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
}

func Success(d, t string) (MyResponse, error) {

	encryptedData, _ := EncryptAES([]byte(d), t)

	myResponse := MyResponse{
		Status:    "success",
		Message:   "Success get data",
		Timestamp: t,
		Data:      encryptedData,
	}

	return myResponse, nil
}

func Error(d, t string) (MyResponse, error) {
	encryptedData, _ := EncryptAES([]byte(d), t)

	myResponse := MyResponse{
		Status:    "error",
		Message:   "error get data",
		Timestamp: t,
		Data:      encryptedData,
	}

	return myResponse, nil
}
