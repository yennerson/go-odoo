package odoo

import (
	"fmt"
)

// AccountFullReconcile represents account.full.reconcile model.
type AccountFullReconcile struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	ExchangeMoveId      *Many2One `xmlrpc:"exchange_move_id,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	PartialReconcileIds *Relation `xmlrpc:"partial_reconcile_ids,omptempty"`
	ReconciledLineIds   *Relation `xmlrpc:"reconciled_line_ids,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountFullReconciles represents array of account.full.reconcile model.
type AccountFullReconciles []AccountFullReconcile

// AccountFullReconcileModel is the odoo model name.
const AccountFullReconcileModel = "account.full.reconcile"

// Many2One convert AccountFullReconcile to *Many2One.
func (afr *AccountFullReconcile) Many2One() *Many2One {
	return NewMany2One(afr.Id.Get(), "")
}

// CreateAccountFullReconcile creates a new account.full.reconcile model and returns its id.
func (c *Client) CreateAccountFullReconcile(afr *AccountFullReconcile) (int64, error) {
	return c.Create(AccountFullReconcileModel, afr)
}

// UpdateAccountFullReconcile updates an existing account.full.reconcile record.
func (c *Client) UpdateAccountFullReconcile(afr *AccountFullReconcile) error {
	return c.UpdateAccountFullReconciles([]int64{afr.Id.Get()}, afr)
}

// UpdateAccountFullReconciles updates existing account.full.reconcile records.
// All records (represented by ids) will be updated by afr values.
func (c *Client) UpdateAccountFullReconciles(ids []int64, afr *AccountFullReconcile) error {
	return c.Update(AccountFullReconcileModel, ids, afr)
}

// DeleteAccountFullReconcile deletes an existing account.full.reconcile record.
func (c *Client) DeleteAccountFullReconcile(id int64) error {
	return c.DeleteAccountFullReconciles([]int64{id})
}

// DeleteAccountFullReconciles deletes existing account.full.reconcile records.
func (c *Client) DeleteAccountFullReconciles(ids []int64) error {
	return c.Delete(AccountFullReconcileModel, ids)
}

// GetAccountFullReconcile gets account.full.reconcile existing record.
func (c *Client) GetAccountFullReconcile(id int64) (*AccountFullReconcile, error) {
	afrs, err := c.GetAccountFullReconciles([]int64{id})
	if err != nil {
		return nil, err
	}
	if afrs != nil && len(*afrs) > 0 {
		return &((*afrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.full.reconcile not found", id)
}

// GetAccountFullReconciles gets account.full.reconcile existing records.
func (c *Client) GetAccountFullReconciles(ids []int64) (*AccountFullReconciles, error) {
	afrs := &AccountFullReconciles{}
	if err := c.Read(AccountFullReconcileModel, ids, nil, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFullReconcile finds account.full.reconcile record by querying it with criteria.
func (c *Client) FindAccountFullReconcile(criteria *Criteria) (*AccountFullReconcile, error) {
	afrs := &AccountFullReconciles{}
	if err := c.SearchRead(AccountFullReconcileModel, criteria, NewOptions().Limit(1), afrs); err != nil {
		return nil, err
	}
	if afrs != nil && len(*afrs) > 0 {
		return &((*afrs)[0]), nil
	}
	return nil, fmt.Errorf("account.full.reconcile was not found with criteria %v", criteria)
}

// FindAccountFullReconciles finds account.full.reconcile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFullReconciles(criteria *Criteria, options *Options) (*AccountFullReconciles, error) {
	afrs := &AccountFullReconciles{}
	if err := c.SearchRead(AccountFullReconcileModel, criteria, options, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFullReconcileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFullReconcileIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFullReconcileModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFullReconcileId finds record id by querying it with criteria.
func (c *Client) FindAccountFullReconcileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFullReconcileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.full.reconcile was not found with criteria %v and options %v", criteria, options)
}
