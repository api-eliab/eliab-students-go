package main

type NotificationsResponse struct {
	Notifications []Notification `json:"messages"`
}

type Notification struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedAt   string `json:"createdAt"`
	TeacherName string `json:"teacherName"`
	SectionID   int64  `json:"sectionId"`
	SectionName string `json:"sectionName"`
	CourseID    int64  `json:"courseId"`
	CourseName  string `json:"courseName"`
	GradeID     int64  `json:"gradeId"`
	GradeName   string `json:"gradeName"`
	File        string `json:"file"`
	Type        string `json:"type"`
}

type MessageRequest struct {
	Message string `json:"message"`
}

type MessagesResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Text        string `json:"text"`
	Type        string `json:"type"`
	TeacherName string `json:"teacher_name"`
	OwnerName   string `json:"owner_name"`
	Time        string `json:"time"`
	Aproved     int    `json:"approved"`
	SectionID   int64  `json:"sectionId"`
	SectionName string `json:"sectionName"`
	CourseID    int64  `json:"courseId"`
	CourseName  string `json:"courseName"`
	GradeID     int64  `json:"gradeId"`
	GradeName   string `json:"gradeName"`
}
