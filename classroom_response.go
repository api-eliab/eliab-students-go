package main

// ClassroomsResponse ...
type ClassroomsResponse struct {
	Classrooms []Classroom `json:"classrooms"`
}

type ClassroomDetailResponse struct {
	Classroom ClassroomDetail `json:"classroom"`
}
