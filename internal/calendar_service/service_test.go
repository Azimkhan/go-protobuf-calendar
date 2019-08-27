package calendar_service

import (
	"github.com/golang/protobuf/ptypes"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestCalendarService(t *testing.T) {
	service := NewCalendarService()
	var err error

	event := &CalendarEvent{
		Type: CalendarEvent_MEETING,
		Name: "Clean the rooms",
	}

	// create event
	beforeCreate := time.Now()
	time.Sleep(10 * time.Nanosecond)
	err = service.Create(event)
	time.Sleep(10 * time.Nanosecond)
	afterCreate := time.Now()

	assert.NilError(t, err)
	assert.Equal(t, int64(1), event.Id)
	createTime, _ := ptypes.Timestamp(event.Created)
	assert.Equal(t, event.Created, event.Updated)
	assert.Assert(t, createTime.After(beforeCreate))
	assert.Assert(t, createTime.Before(afterCreate))

	// update event
	event.Type = CalendarEvent_TASK
	err = service.Update(event)
	assert.NilError(t, err)

	// get event
	e2, err := service.Get(event.Id)
	assert.NilError(t, err)
	assert.Assert(t, e2 != nil)
	assert.Equal(t, CalendarEvent_TASK, e2.Type)
	assert.Equal(t, "Clean the rooms", e2.Name)

	// delete event
	err = service.Delete(e2.Id)
	assert.NilError(t, err)

	// check that delete fails now
	err = service.Delete(e2.Id)
	assert.Error(t, err, "not found")

}
