package main

// HomeworksResponse ...
type HomeworksResponse struct {
	Homeworks []HomeworkResponse `json:"homeworks"`
}

// HomeworkResponse ...
type HomeworkResponse struct {
	ID               int64  `json:"id"`
	Points           string `json:"points"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	ClasroomID       int64  `json:"clasroom_id"`
	DeliveryDate     string `json:"delivery_date"`
	DeliveryHour     string `json:"delivery_hour"`
}

// HomeworkDetailResponse ...
type HomeworkDetailResponse struct {
	HomeworkDetail HomeworkDetail `json:"homework"`
}

// HomeworkDetail ...
type HomeworkDetail struct {
	ID               int64  `json:"id"`
	Points           int64  `json:"points"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	ClassroomID      int64  `json:"classroom_id"`
	ClassroomName    string `json:"classroom_name"`
	TeachersID       string `json:"teachers_id"`
	TeachersName     string `json:"teachers_name"`
	DeliveryDate     string `json:"delivery_date"`
	DeliveryHour     string `json:"delivery_hour"`
}
