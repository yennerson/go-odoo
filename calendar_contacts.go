package odoo

import (
	"fmt"
)

// CalendarContacts represents calendar.contacts model.
type CalendarContacts struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CalendarContactss represents array of calendar.contacts model.
type CalendarContactss []CalendarContacts

// CalendarContactsModel is the odoo model name.
const CalendarContactsModel = "calendar.contacts"

// Many2One convert CalendarContacts to *Many2One.
func (cc *CalendarContacts) Many2One() *Many2One {
	return NewMany2One(cc.Id.Get(), "")
}

// CreateCalendarContacts creates a new calendar.contacts model and returns its id.
func (c *Client) CreateCalendarContacts(cc *CalendarContacts) (int64, error) {
	return c.Create(CalendarContactsModel, cc)
}

// UpdateCalendarContacts updates an existing calendar.contacts record.
func (c *Client) UpdateCalendarContacts(cc *CalendarContacts) error {
	return c.UpdateCalendarContactss([]int64{cc.Id.Get()}, cc)
}

// UpdateCalendarContactss updates existing calendar.contacts records.
// All records (represented by ids) will be updated by cc values.
func (c *Client) UpdateCalendarContactss(ids []int64, cc *CalendarContacts) error {
	return c.Update(CalendarContactsModel, ids, cc)
}

// DeleteCalendarContacts deletes an existing calendar.contacts record.
func (c *Client) DeleteCalendarContacts(id int64) error {
	return c.DeleteCalendarContactss([]int64{id})
}

// DeleteCalendarContactss deletes existing calendar.contacts records.
func (c *Client) DeleteCalendarContactss(ids []int64) error {
	return c.Delete(CalendarContactsModel, ids)
}

// GetCalendarContacts gets calendar.contacts existing record.
func (c *Client) GetCalendarContacts(id int64) (*CalendarContacts, error) {
	ccs, err := c.GetCalendarContactss([]int64{id})
	if err != nil {
		return nil, err
	}
	if ccs != nil && len(*ccs) > 0 {
		return &((*ccs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.contacts not found", id)
}

// GetCalendarContactss gets calendar.contacts existing records.
func (c *Client) GetCalendarContactss(ids []int64) (*CalendarContactss, error) {
	ccs := &CalendarContactss{}
	if err := c.Read(CalendarContactsModel, ids, nil, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCalendarContacts finds calendar.contacts record by querying it with criteria.
func (c *Client) FindCalendarContacts(criteria *Criteria) (*CalendarContacts, error) {
	ccs := &CalendarContactss{}
	if err := c.SearchRead(CalendarContactsModel, criteria, NewOptions().Limit(1), ccs); err != nil {
		return nil, err
	}
	if ccs != nil && len(*ccs) > 0 {
		return &((*ccs)[0]), nil
	}
	return nil, fmt.Errorf("calendar.contacts was not found with criteria %v", criteria)
}

// FindCalendarContactss finds calendar.contacts records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarContactss(criteria *Criteria, options *Options) (*CalendarContactss, error) {
	ccs := &CalendarContactss{}
	if err := c.SearchRead(CalendarContactsModel, criteria, options, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCalendarContactsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarContactsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarContactsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarContactsId finds record id by querying it with criteria.
func (c *Client) FindCalendarContactsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarContactsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("calendar.contacts was not found with criteria %v and options %v", criteria, options)
}
