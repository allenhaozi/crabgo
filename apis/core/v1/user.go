package v1

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age,omitempty"`
	UserId string `json:"userID"`
}
