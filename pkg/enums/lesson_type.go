package enums

type LessonType string

const (
	Lecture  LessonType = "lecture"
	Lab      LessonType = "lab"
	Exercise LessonType = "exercise"
	Test     LessonType = "test"
	Exam     LessonType = "exam"
	Other    LessonType = "other"
)
