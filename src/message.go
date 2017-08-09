package main

type Register struct {
	Type   string   `json:"type"`
	Topics []string `json:"topics"`
}

func getRegisterMessage(topics []string) Register {
	return Register{
		Type:   TypeRegister,
		Topics: topics,
	}
}
