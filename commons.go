package commons

type BasicCommand struct {
	OwnerId  string `json:"OwnerId"`
	Command  string `json:"Command"`
	ExtraArg string `json:"ExtraArg"`
}

type RegisterCommand struct {
	OwnerId    string `json:"OwnerId"`
	PlayerName string `json:"PlayerName"`
}