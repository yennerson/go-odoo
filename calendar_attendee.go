package odoo

import (
	"fmt"
)

// CalendarAttendee represents calendar.attendee model.
type CalendarAttendee struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken  *String    `xmlrpc:"access_token,omptempty"`
	Availability *Selection `xmlrpc:"availability,omptempty"`
	CommonName   *String    `xmlrpc:"common_name,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	Email        *String    `xmlrpc:"email,omptempty"`
	EventId      *Many2One  `xmlrpc:"event_id,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	PartnerId    *Many2One  `xmlrpc:"partner_id,omptempty"`
	State        *Selection `xmlrpc:"state,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CalendarAttendees represents array of calendar.attendee model.
type CalendarAttendees []CalendarAttendee

// CalendarAttendeeModel is the odoo model name.
const CalendarAttendeeModel = "calendar.attendee"

// Many2One convert CalendarAttendee to *Many2One.
func (ca *CalendarAttendee) Many2One() *Many2One {
	return NewMany2One(ca.Id.Get(), "")
}

// CreateCalendarAttendee creates a new calendar.attendee model and returns its id.
func (c *Client) CreateCalendarAttendee(ca *CalendarAttendee) (int64, error) {
	return c.Create(CalendarAttendeeModel, ca)
}

// UpdateCalendarAttendee updates an existing calendar.attendee record.
func (c *Client) UpdateCalendarAttendee(ca *CalendarAttendee) error {
	return c.UpdateCalendarAttendees([]int64{ca.Id.Get()}, ca)
}

// UpdateCalendarAttendees updates existing calendar.attendee records.
// All records (represented by ids) will be updated by ca values.
func (c *Client) UpdateCalendarAttendees(ids []int64, ca *CalendarAttendee) error {
	return c.Update(CalendarAttendeeModel, ids, ca)
}

// DeleteCalendarAttendee deletes an existing calendar.attendee record.
func (c *Client) DeleteCalendarAttendee(id int64) error {
	return c.DeleteCalendarAttendees([]int64{id})
}

// DeleteCalendarAttendees deletes existing calendar.attendee records.
func (c *Client) DeleteCalendarAttendees(ids []int64) error {
	return c.Delete(CalendarAttendeeModel, ids)
}

// GetCalendarAttendee gets calendar.attendee existing record.
func (c *Client) GetCalendarAttendee(id int64) (*CalendarAttendee, error) {
	cas, err := c.GetCalendarAttendees([]int64{id})
	if err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.attendee not found", id)
}

// GetCalendarAttendees gets calendar.attendee existing records.
func (c *Client) GetCalendarAttendees(ids []int64) (*CalendarAttendees, error) {
	cas := &CalendarAttendees{}
	if err := c.Read(CalendarAttendeeModel, ids, nil, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCalendarAttendee finds calendar.attendee record by querying it with criteria.
func (c *Client) FindCalendarAttendee(criteria *Criteria) (*CalendarAttendee, error) {
	cas := &CalendarAttendees{}
	if err := c.SearchRead(CalendarAttendeeModel, criteria, NewOptions().Limit(1), cas); err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("calendar.attendee was not found with criteria %v", criteria)
}

// FindCalendarAttendees finds calendar.attendee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarAttendees(criteria *Criteria, options *Options) (*CalendarAttendees, error) {
	cas := &CalendarAttendees{}
	if err := c.SearchRead(CalendarAttendeeModel, criteria, options, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCalendarAttendeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarAttendeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarAttendeeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarAttendeeId finds record id by querying it with criteria.
func (c *Client) FindCalendarAttendeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarAttendeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("calendar.attendee was not found with criteria %v and options %v", criteria, options)
}
