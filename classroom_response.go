package main

type ClassroomsResponse struct {
	Classrooms []Classroom `json:"classrooms"`
}

type ClassroomDetailResponse struct {
	Classroom struct {
		Name    string `json:"name"`
		Teacher struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			LastName string `json:"last_name"`
		} `json:"teacher"`
		Grade      string `json:"grade"`
		CourseDist []struct {
			Perfect       bool   `json:"perfect"`
			Name          string `json:"name"`
			ID            int    `json:"id"`
			CurrentPoints int    `json:"current_points"`
			Tasks         []struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Points string `json:"points"`
				Type   int    `json:"type"`
			} `json:"tasks"`
		} `json:"course_dist"`
	} `json:"classroom"`
}
