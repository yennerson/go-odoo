package odoo

import (
	"fmt"
)

// IrModelRelation represents ir.model.relation model.
type IrModelRelation struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DateInit    *Time     `xmlrpc:"date_init,omptempty"`
	DateUpdate  *Time     `xmlrpc:"date_update,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Model       *Many2One `xmlrpc:"model,omptempty"`
	Module      *Many2One `xmlrpc:"module,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// IrModelRelations represents array of ir.model.relation model.
type IrModelRelations []IrModelRelation

// IrModelRelationModel is the odoo model name.
const IrModelRelationModel = "ir.model.relation"

// Many2One convert IrModelRelation to *Many2One.
func (imr *IrModelRelation) Many2One() *Many2One {
	return NewMany2One(imr.Id.Get(), "")
}

// CreateIrModelRelation creates a new ir.model.relation model and returns its id.
func (c *Client) CreateIrModelRelation(imr *IrModelRelation) (int64, error) {
	return c.Create(IrModelRelationModel, imr)
}

// UpdateIrModelRelation updates an existing ir.model.relation record.
func (c *Client) UpdateIrModelRelation(imr *IrModelRelation) error {
	return c.UpdateIrModelRelations([]int64{imr.Id.Get()}, imr)
}

// UpdateIrModelRelations updates existing ir.model.relation records.
// All records (represented by ids) will be updated by imr values.
func (c *Client) UpdateIrModelRelations(ids []int64, imr *IrModelRelation) error {
	return c.Update(IrModelRelationModel, ids, imr)
}

// DeleteIrModelRelation deletes an existing ir.model.relation record.
func (c *Client) DeleteIrModelRelation(id int64) error {
	return c.DeleteIrModelRelations([]int64{id})
}

// DeleteIrModelRelations deletes existing ir.model.relation records.
func (c *Client) DeleteIrModelRelations(ids []int64) error {
	return c.Delete(IrModelRelationModel, ids)
}

// GetIrModelRelation gets ir.model.relation existing record.
func (c *Client) GetIrModelRelation(id int64) (*IrModelRelation, error) {
	imrs, err := c.GetIrModelRelations([]int64{id})
	if err != nil {
		return nil, err
	}
	if imrs != nil && len(*imrs) > 0 {
		return &((*imrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.model.relation not found", id)
}

// GetIrModelRelations gets ir.model.relation existing records.
func (c *Client) GetIrModelRelations(ids []int64) (*IrModelRelations, error) {
	imrs := &IrModelRelations{}
	if err := c.Read(IrModelRelationModel, ids, nil, imrs); err != nil {
		return nil, err
	}
	return imrs, nil
}

// FindIrModelRelation finds ir.model.relation record by querying it with criteria.
func (c *Client) FindIrModelRelation(criteria *Criteria) (*IrModelRelation, error) {
	imrs := &IrModelRelations{}
	if err := c.SearchRead(IrModelRelationModel, criteria, NewOptions().Limit(1), imrs); err != nil {
		return nil, err
	}
	if imrs != nil && len(*imrs) > 0 {
		return &((*imrs)[0]), nil
	}
	return nil, fmt.Errorf("ir.model.relation was not found with criteria %v", criteria)
}

// FindIrModelRelations finds ir.model.relation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelRelations(criteria *Criteria, options *Options) (*IrModelRelations, error) {
	imrs := &IrModelRelations{}
	if err := c.SearchRead(IrModelRelationModel, criteria, options, imrs); err != nil {
		return nil, err
	}
	return imrs, nil
}

// FindIrModelRelationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelRelationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModelRelationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModelRelationId finds record id by querying it with criteria.
func (c *Client) FindIrModelRelationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelRelationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.model.relation was not found with criteria %v and options %v", criteria, options)
}
