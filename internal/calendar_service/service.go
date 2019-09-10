//go:generate protoc --go_out=. --proto_path=../../api/protobuf calendar.proto
package calendar_service

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"sync"
)

type CalendarService interface {
	Create(event *CalendarEvent) error
	Update(event *CalendarEvent) error
	Delete(id int64) error
	Get(id int64) (*CalendarEvent, error)
}

type MapCalendarService struct {
	db     sync.Map
	nextId int64
	idLock sync.Mutex
}

func NewCalendarService() CalendarService {
	return &MapCalendarService{
		db:     sync.Map{},
		nextId: 1,
		idLock: sync.Mutex{},
	}
}
func (s *MapCalendarService) getNextId() (id int64) {
	s.idLock.Lock()
	defer s.idLock.Unlock()
	id = s.nextId
	s.nextId++
	return
}

func (s *MapCalendarService) Create(event *CalendarEvent) error {
	if event == nil {
		return fmt.Errorf("event must not be empty")
	}
	event.Id = s.getNextId()
	now := ptypes.TimestampNow()
	event.Created = now
	event.Updated = now
	s.db.Store(event.Id, event)
	return nil
}

func (s *MapCalendarService) Update(event *CalendarEvent) error {
	if event == nil {
		return fmt.Errorf("event must not be empty")
	}
	if event.Id < 1 {
		return fmt.Errorf("id %d is invalid", event.Id)
	}

	event.Updated = ptypes.TimestampNow()
	s.db.Store(event.Id, event)
	return nil
}

func (s *MapCalendarService) Delete(id int64) error {
	_, found := s.db.Load(id)
	if found {
		s.db.Delete(id)
		return nil
	}
	return fmt.Errorf("not found")
}

func (s *MapCalendarService) Get(id int64) (*CalendarEvent, error) {
	event, found := s.db.Load(id)
	if found {
		return event.(*CalendarEvent), nil
	}
	return nil, fmt.Errorf("not found")
}
