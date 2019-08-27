package calendar_service

//go:generate protoc --go_out=. calendar.proto
type CalendarService interface {
	Create(event *CalendarEvent) error
	Update(event *CalendarEvent) error
	Delete(id int64) error
	Get(id int64) (*CalendarEvent, error)
}
