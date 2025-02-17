package odoo

import (
	"fmt"
)

// MailActivity represents mail.activity model.
type MailActivity struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityCategory          *Selection `xmlrpc:"activity_category,omptempty"`
	ActivityTypeId            *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	CalendarEventId           *Many2One  `xmlrpc:"calendar_event_id,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateDeadline              *Time      `xmlrpc:"date_deadline,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	Feedback                  *String    `xmlrpc:"feedback,omptempty"`
	HasRecommendedActivities  *Bool      `xmlrpc:"has_recommended_activities,omptempty"`
	Icon                      *String    `xmlrpc:"icon,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	Note                      *String    `xmlrpc:"note,omptempty"`
	PreviousActivityTypeId    *Many2One  `xmlrpc:"previous_activity_type_id,omptempty"`
	RecommendedActivityTypeId *Many2One  `xmlrpc:"recommended_activity_type_id,omptempty"`
	ResId                     *Int       `xmlrpc:"res_id,omptempty"`
	ResModel                  *String    `xmlrpc:"res_model,omptempty"`
	ResModelId                *Many2One  `xmlrpc:"res_model_id,omptempty"`
	ResName                   *String    `xmlrpc:"res_name,omptempty"`
	State                     *Selection `xmlrpc:"state,omptempty"`
	Summary                   *String    `xmlrpc:"summary,omptempty"`
	UserId                    *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailActivitys represents array of mail.activity model.
type MailActivitys []MailActivity

// MailActivityModel is the odoo model name.
const MailActivityModel = "mail.activity"

// Many2One convert MailActivity to *Many2One.
func (ma *MailActivity) Many2One() *Many2One {
	return NewMany2One(ma.Id.Get(), "")
}

// CreateMailActivity creates a new mail.activity model and returns its id.
func (c *Client) CreateMailActivity(ma *MailActivity) (int64, error) {
	return c.Create(MailActivityModel, ma)
}

// UpdateMailActivity updates an existing mail.activity record.
func (c *Client) UpdateMailActivity(ma *MailActivity) error {
	return c.UpdateMailActivitys([]int64{ma.Id.Get()}, ma)
}

// UpdateMailActivitys updates existing mail.activity records.
// All records (represented by ids) will be updated by ma values.
func (c *Client) UpdateMailActivitys(ids []int64, ma *MailActivity) error {
	return c.Update(MailActivityModel, ids, ma)
}

// DeleteMailActivity deletes an existing mail.activity record.
func (c *Client) DeleteMailActivity(id int64) error {
	return c.DeleteMailActivitys([]int64{id})
}

// DeleteMailActivitys deletes existing mail.activity records.
func (c *Client) DeleteMailActivitys(ids []int64) error {
	return c.Delete(MailActivityModel, ids)
}

// GetMailActivity gets mail.activity existing record.
func (c *Client) GetMailActivity(id int64) (*MailActivity, error) {
	mas, err := c.GetMailActivitys([]int64{id})
	if err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.activity not found", id)
}

// GetMailActivitys gets mail.activity existing records.
func (c *Client) GetMailActivitys(ids []int64) (*MailActivitys, error) {
	mas := &MailActivitys{}
	if err := c.Read(MailActivityModel, ids, nil, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMailActivity finds mail.activity record by querying it with criteria.
func (c *Client) FindMailActivity(criteria *Criteria) (*MailActivity, error) {
	mas := &MailActivitys{}
	if err := c.SearchRead(MailActivityModel, criteria, NewOptions().Limit(1), mas); err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("mail.activity was not found with criteria %v", criteria)
}

// FindMailActivitys finds mail.activity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivitys(criteria *Criteria, options *Options) (*MailActivitys, error) {
	mas := &MailActivitys{}
	if err := c.SearchRead(MailActivityModel, criteria, options, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMailActivityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailActivityModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailActivityId finds record id by querying it with criteria.
func (c *Client) FindMailActivityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.activity was not found with criteria %v and options %v", criteria, options)
}
