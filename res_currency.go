package odoo

import (
	"fmt"
)

// ResCurrency represents res.currency model.
type ResCurrency struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	Active               *Bool      `xmlrpc:"active,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencySubunitLabel *String    `xmlrpc:"currency_subunit_label,omptempty"`
	CurrencyUnitLabel    *String    `xmlrpc:"currency_unit_label,omptempty"`
	Date                 *Time      `xmlrpc:"date,omptempty"`
	DecimalPlaces        *Int       `xmlrpc:"decimal_places,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	Position             *Selection `xmlrpc:"position,omptempty"`
	Rate                 *Float     `xmlrpc:"rate,omptempty"`
	RateIds              *Relation  `xmlrpc:"rate_ids,omptempty"`
	Rounding             *Float     `xmlrpc:"rounding,omptempty"`
	Symbol               *String    `xmlrpc:"symbol,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResCurrencys represents array of res.currency model.
type ResCurrencys []ResCurrency

// ResCurrencyModel is the odoo model name.
const ResCurrencyModel = "res.currency"

// Many2One convert ResCurrency to *Many2One.
func (rc *ResCurrency) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCurrency creates a new res.currency model and returns its id.
func (c *Client) CreateResCurrency(rc *ResCurrency) (int64, error) {
	return c.Create(ResCurrencyModel, rc)
}

// UpdateResCurrency updates an existing res.currency record.
func (c *Client) UpdateResCurrency(rc *ResCurrency) error {
	return c.UpdateResCurrencys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCurrencys updates existing res.currency records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCurrencys(ids []int64, rc *ResCurrency) error {
	return c.Update(ResCurrencyModel, ids, rc)
}

// DeleteResCurrency deletes an existing res.currency record.
func (c *Client) DeleteResCurrency(id int64) error {
	return c.DeleteResCurrencys([]int64{id})
}

// DeleteResCurrencys deletes existing res.currency records.
func (c *Client) DeleteResCurrencys(ids []int64) error {
	return c.Delete(ResCurrencyModel, ids)
}

// GetResCurrency gets res.currency existing record.
func (c *Client) GetResCurrency(id int64) (*ResCurrency, error) {
	rcs, err := c.GetResCurrencys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.currency not found", id)
}

// GetResCurrencys gets res.currency existing records.
func (c *Client) GetResCurrencys(ids []int64) (*ResCurrencys, error) {
	rcs := &ResCurrencys{}
	if err := c.Read(ResCurrencyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCurrency finds res.currency record by querying it with criteria.
func (c *Client) FindResCurrency(criteria *Criteria) (*ResCurrency, error) {
	rcs := &ResCurrencys{}
	if err := c.SearchRead(ResCurrencyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.currency was not found with criteria %v", criteria)
}

// FindResCurrencys finds res.currency records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCurrencys(criteria *Criteria, options *Options) (*ResCurrencys, error) {
	rcs := &ResCurrencys{}
	if err := c.SearchRead(ResCurrencyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCurrencyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCurrencyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCurrencyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCurrencyId finds record id by querying it with criteria.
func (c *Client) FindResCurrencyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCurrencyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.currency was not found with criteria %v and options %v", criteria, options)
}
