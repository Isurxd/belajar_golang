package model

type Student struct {
	StudentID int64  `json:"student_id"`
	Name      string `json:"name"`
	Age       int64  `json:"age"`
	Address   string `json:"address"`
	Hobby     string `json:"hoby"`
}
