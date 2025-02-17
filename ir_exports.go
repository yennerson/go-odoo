package odoo

import (
	"fmt"
)

// IrExports represents ir.exports model.
type IrExports struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	ExportFields *Relation `xmlrpc:"export_fields,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	Resource     *String   `xmlrpc:"resource,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// IrExportss represents array of ir.exports model.
type IrExportss []IrExports

// IrExportsModel is the odoo model name.
const IrExportsModel = "ir.exports"

// Many2One convert IrExports to *Many2One.
func (ie *IrExports) Many2One() *Many2One {
	return NewMany2One(ie.Id.Get(), "")
}

// CreateIrExports creates a new ir.exports model and returns its id.
func (c *Client) CreateIrExports(ie *IrExports) (int64, error) {
	return c.Create(IrExportsModel, ie)
}

// UpdateIrExports updates an existing ir.exports record.
func (c *Client) UpdateIrExports(ie *IrExports) error {
	return c.UpdateIrExportss([]int64{ie.Id.Get()}, ie)
}

// UpdateIrExportss updates existing ir.exports records.
// All records (represented by ids) will be updated by ie values.
func (c *Client) UpdateIrExportss(ids []int64, ie *IrExports) error {
	return c.Update(IrExportsModel, ids, ie)
}

// DeleteIrExports deletes an existing ir.exports record.
func (c *Client) DeleteIrExports(id int64) error {
	return c.DeleteIrExportss([]int64{id})
}

// DeleteIrExportss deletes existing ir.exports records.
func (c *Client) DeleteIrExportss(ids []int64) error {
	return c.Delete(IrExportsModel, ids)
}

// GetIrExports gets ir.exports existing record.
func (c *Client) GetIrExports(id int64) (*IrExports, error) {
	ies, err := c.GetIrExportss([]int64{id})
	if err != nil {
		return nil, err
	}
	if ies != nil && len(*ies) > 0 {
		return &((*ies)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.exports not found", id)
}

// GetIrExportss gets ir.exports existing records.
func (c *Client) GetIrExportss(ids []int64) (*IrExportss, error) {
	ies := &IrExportss{}
	if err := c.Read(IrExportsModel, ids, nil, ies); err != nil {
		return nil, err
	}
	return ies, nil
}

// FindIrExports finds ir.exports record by querying it with criteria.
func (c *Client) FindIrExports(criteria *Criteria) (*IrExports, error) {
	ies := &IrExportss{}
	if err := c.SearchRead(IrExportsModel, criteria, NewOptions().Limit(1), ies); err != nil {
		return nil, err
	}
	if ies != nil && len(*ies) > 0 {
		return &((*ies)[0]), nil
	}
	return nil, fmt.Errorf("ir.exports was not found with criteria %v", criteria)
}

// FindIrExportss finds ir.exports records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrExportss(criteria *Criteria, options *Options) (*IrExportss, error) {
	ies := &IrExportss{}
	if err := c.SearchRead(IrExportsModel, criteria, options, ies); err != nil {
		return nil, err
	}
	return ies, nil
}

// FindIrExportsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrExportsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrExportsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrExportsId finds record id by querying it with criteria.
func (c *Client) FindIrExportsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrExportsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.exports was not found with criteria %v and options %v", criteria, options)
}
