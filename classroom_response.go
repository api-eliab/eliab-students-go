package main

type ClassroomsResponse struct {
	Classrooms []Classroom `json:"classrooms"`
}

type ClassroomDetailResponse struct {
	ClassroomDetail ClassroomDetail `json:"classroom"`
}

type ClassroomDetail struct {
	ID         int64        `json:"id"`
	SectionID  int64        `json:"section_id"`
	Name       string       `json:"name"`
	Teacher    Teacher      `json:"teacher"`
	Grade      string       `json:"grade"`
	CourseDist []CourseDist `json:"course_dist"`
}

type Teacher struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type CourseDist struct {
	Perfect       bool   `json:"perfect"`
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	CurrentPoints int64  `json:"current_points"`
	Tasks         []Task `json:"tasks"`
}

type Task struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Points  string `json:"points"`
	Type    int    `json:"type"`
	Comment string `json:"comment"`
}
