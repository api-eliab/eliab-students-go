package main

// HomeworksResponse doc ...
type HomeworksResponse struct {
	Homeworks []Homework `json:"homeworks"`
}

// Homework doc ...
type Homework struct {
	ID                    int64   `json:"id"`
	Points                float64 `json:"points"`
	Title                 string  `json:"title"`
	ShortDescription      string  `json:"short_description"`
	LongDescription       string  `json:"long_description"`
	ClassroomID           string  `json:"classroom_id"`
	DeliveryDateFormatted string  `json:"delivery_date_formatted"`
	DeliveryDate          string  `json:"delivery_date"`
	DeliveryHour          string  `json:"delivery_hour,omitempty"`
	Type                  string  `json:"type"`
}

// HomeworkDetailResponse doc ...
type HomeworkDetailResponse struct {
	HomeworkDetail HomeworkDetail `json:"homework"`
}

// HomeworkDetail doc ...
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
	Type             string `json:"type"`
}
