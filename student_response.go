package main

type HomeworksResponse struct {
	Homeworks []struct {
		ID               int    `json:"id"`
		Points           int    `json:"points"`
		Title            string `json:"title"`
		ShortDescription string `json:"short_description"`
		LongDescription string `json:"long_description"`
		ClassroomID      string `json:"classroom_id"`
		DeliveryDate     string `json:"delivery_date"`
		DeliveryHour     string `json:"delivery_hour"`
	} `json:"homeworks"`
}


type HomeworkDetailResponse struct {
		Homework struct {
			ID               int    `json:"id"`
			Points           int    `json:"points"`
			Title            string `json:"title"`
			ShortDescription string `json:"short_description"`
			LongDescription  string `json:"long_description"`
			ClassroomID      int    `json:"classroom_id"`
			ClassroomName    string `json:"classroom_name"`
			TeachersID       string `json:"teachers_id"`
			TeachersName     string `json:"teachers_name"`
			DeliveryDate     string `json:"delivery_date"`
			DeliveryHour     string `json:"delivery_hour"`
		} `json:"homework"` 
}