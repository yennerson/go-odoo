package odoo

import (
	"fmt"
)

// SmsApi represents sms.api model.
type SmsApi struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// SmsApis represents array of sms.api model.
type SmsApis []SmsApi

// SmsApiModel is the odoo model name.
const SmsApiModel = "sms.api"

// Many2One convert SmsApi to *Many2One.
func (sa *SmsApi) Many2One() *Many2One {
	return NewMany2One(sa.Id.Get(), "")
}

// CreateSmsApi creates a new sms.api model and returns its id.
func (c *Client) CreateSmsApi(sa *SmsApi) (int64, error) {
	return c.Create(SmsApiModel, sa)
}

// UpdateSmsApi updates an existing sms.api record.
func (c *Client) UpdateSmsApi(sa *SmsApi) error {
	return c.UpdateSmsApis([]int64{sa.Id.Get()}, sa)
}

// UpdateSmsApis updates existing sms.api records.
// All records (represented by ids) will be updated by sa values.
func (c *Client) UpdateSmsApis(ids []int64, sa *SmsApi) error {
	return c.Update(SmsApiModel, ids, sa)
}

// DeleteSmsApi deletes an existing sms.api record.
func (c *Client) DeleteSmsApi(id int64) error {
	return c.DeleteSmsApis([]int64{id})
}

// DeleteSmsApis deletes existing sms.api records.
func (c *Client) DeleteSmsApis(ids []int64) error {
	return c.Delete(SmsApiModel, ids)
}

// GetSmsApi gets sms.api existing record.
func (c *Client) GetSmsApi(id int64) (*SmsApi, error) {
	sas, err := c.GetSmsApis([]int64{id})
	if err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.api not found", id)
}

// GetSmsApis gets sms.api existing records.
func (c *Client) GetSmsApis(ids []int64) (*SmsApis, error) {
	sas := &SmsApis{}
	if err := c.Read(SmsApiModel, ids, nil, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSmsApi finds sms.api record by querying it with criteria.
func (c *Client) FindSmsApi(criteria *Criteria) (*SmsApi, error) {
	sas := &SmsApis{}
	if err := c.SearchRead(SmsApiModel, criteria, NewOptions().Limit(1), sas); err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("sms.api was not found with criteria %v", criteria)
}

// FindSmsApis finds sms.api records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsApis(criteria *Criteria, options *Options) (*SmsApis, error) {
	sas := &SmsApis{}
	if err := c.SearchRead(SmsApiModel, criteria, options, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSmsApiIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsApiIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsApiModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsApiId finds record id by querying it with criteria.
func (c *Client) FindSmsApiId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsApiModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.api was not found with criteria %v and options %v", criteria, options)
}
