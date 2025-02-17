package odoo

import (
	"fmt"
)

// AccountAccountTag represents account.account.tag model.
type AccountAccountTag struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	Active        *Bool      `xmlrpc:"active,omptempty"`
	Applicability *Selection `xmlrpc:"applicability,omptempty"`
	Color         *Int       `xmlrpc:"color,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAccountTags represents array of account.account.tag model.
type AccountAccountTags []AccountAccountTag

// AccountAccountTagModel is the odoo model name.
const AccountAccountTagModel = "account.account.tag"

// Many2One convert AccountAccountTag to *Many2One.
func (aat *AccountAccountTag) Many2One() *Many2One {
	return NewMany2One(aat.Id.Get(), "")
}

// CreateAccountAccountTag creates a new account.account.tag model and returns its id.
func (c *Client) CreateAccountAccountTag(aat *AccountAccountTag) (int64, error) {
	return c.Create(AccountAccountTagModel, aat)
}

// UpdateAccountAccountTag updates an existing account.account.tag record.
func (c *Client) UpdateAccountAccountTag(aat *AccountAccountTag) error {
	return c.UpdateAccountAccountTags([]int64{aat.Id.Get()}, aat)
}

// UpdateAccountAccountTags updates existing account.account.tag records.
// All records (represented by ids) will be updated by aat values.
func (c *Client) UpdateAccountAccountTags(ids []int64, aat *AccountAccountTag) error {
	return c.Update(AccountAccountTagModel, ids, aat)
}

// DeleteAccountAccountTag deletes an existing account.account.tag record.
func (c *Client) DeleteAccountAccountTag(id int64) error {
	return c.DeleteAccountAccountTags([]int64{id})
}

// DeleteAccountAccountTags deletes existing account.account.tag records.
func (c *Client) DeleteAccountAccountTags(ids []int64) error {
	return c.Delete(AccountAccountTagModel, ids)
}

// GetAccountAccountTag gets account.account.tag existing record.
func (c *Client) GetAccountAccountTag(id int64) (*AccountAccountTag, error) {
	aats, err := c.GetAccountAccountTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.account.tag not found", id)
}

// GetAccountAccountTags gets account.account.tag existing records.
func (c *Client) GetAccountAccountTags(ids []int64) (*AccountAccountTags, error) {
	aats := &AccountAccountTags{}
	if err := c.Read(AccountAccountTagModel, ids, nil, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTag finds account.account.tag record by querying it with criteria.
func (c *Client) FindAccountAccountTag(criteria *Criteria) (*AccountAccountTag, error) {
	aats := &AccountAccountTags{}
	if err := c.SearchRead(AccountAccountTagModel, criteria, NewOptions().Limit(1), aats); err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("account.account.tag was not found with criteria %v", criteria)
}

// FindAccountAccountTags finds account.account.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTags(criteria *Criteria, options *Options) (*AccountAccountTags, error) {
	aats := &AccountAccountTags{}
	if err := c.SearchRead(AccountAccountTagModel, criteria, options, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAccountTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAccountTagId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.account.tag was not found with criteria %v and options %v", criteria, options)
}
