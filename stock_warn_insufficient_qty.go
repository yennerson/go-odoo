package odoo

import (
	"fmt"
)

// StockWarnInsufficientQty represents stock.warn.insufficient.qty model.
type StockWarnInsufficientQty struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LocationId  *Many2One `xmlrpc:"location_id,omptempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omptempty"`
	QuantIds    *Relation `xmlrpc:"quant_ids,omptempty"`
}

// StockWarnInsufficientQtys represents array of stock.warn.insufficient.qty model.
type StockWarnInsufficientQtys []StockWarnInsufficientQty

// StockWarnInsufficientQtyModel is the odoo model name.
const StockWarnInsufficientQtyModel = "stock.warn.insufficient.qty"

// Many2One convert StockWarnInsufficientQty to *Many2One.
func (swiq *StockWarnInsufficientQty) Many2One() *Many2One {
	return NewMany2One(swiq.Id.Get(), "")
}

// CreateStockWarnInsufficientQty creates a new stock.warn.insufficient.qty model and returns its id.
func (c *Client) CreateStockWarnInsufficientQty(swiq *StockWarnInsufficientQty) (int64, error) {
	return c.Create(StockWarnInsufficientQtyModel, swiq)
}

// UpdateStockWarnInsufficientQty updates an existing stock.warn.insufficient.qty record.
func (c *Client) UpdateStockWarnInsufficientQty(swiq *StockWarnInsufficientQty) error {
	return c.UpdateStockWarnInsufficientQtys([]int64{swiq.Id.Get()}, swiq)
}

// UpdateStockWarnInsufficientQtys updates existing stock.warn.insufficient.qty records.
// All records (represented by ids) will be updated by swiq values.
func (c *Client) UpdateStockWarnInsufficientQtys(ids []int64, swiq *StockWarnInsufficientQty) error {
	return c.Update(StockWarnInsufficientQtyModel, ids, swiq)
}

// DeleteStockWarnInsufficientQty deletes an existing stock.warn.insufficient.qty record.
func (c *Client) DeleteStockWarnInsufficientQty(id int64) error {
	return c.DeleteStockWarnInsufficientQtys([]int64{id})
}

// DeleteStockWarnInsufficientQtys deletes existing stock.warn.insufficient.qty records.
func (c *Client) DeleteStockWarnInsufficientQtys(ids []int64) error {
	return c.Delete(StockWarnInsufficientQtyModel, ids)
}

// GetStockWarnInsufficientQty gets stock.warn.insufficient.qty existing record.
func (c *Client) GetStockWarnInsufficientQty(id int64) (*StockWarnInsufficientQty, error) {
	swiqs, err := c.GetStockWarnInsufficientQtys([]int64{id})
	if err != nil {
		return nil, err
	}
	if swiqs != nil && len(*swiqs) > 0 {
		return &((*swiqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.warn.insufficient.qty not found", id)
}

// GetStockWarnInsufficientQtys gets stock.warn.insufficient.qty existing records.
func (c *Client) GetStockWarnInsufficientQtys(ids []int64) (*StockWarnInsufficientQtys, error) {
	swiqs := &StockWarnInsufficientQtys{}
	if err := c.Read(StockWarnInsufficientQtyModel, ids, nil, swiqs); err != nil {
		return nil, err
	}
	return swiqs, nil
}

// FindStockWarnInsufficientQty finds stock.warn.insufficient.qty record by querying it with criteria.
func (c *Client) FindStockWarnInsufficientQty(criteria *Criteria) (*StockWarnInsufficientQty, error) {
	swiqs := &StockWarnInsufficientQtys{}
	if err := c.SearchRead(StockWarnInsufficientQtyModel, criteria, NewOptions().Limit(1), swiqs); err != nil {
		return nil, err
	}
	if swiqs != nil && len(*swiqs) > 0 {
		return &((*swiqs)[0]), nil
	}
	return nil, fmt.Errorf("stock.warn.insufficient.qty was not found with criteria %v", criteria)
}

// FindStockWarnInsufficientQtys finds stock.warn.insufficient.qty records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarnInsufficientQtys(criteria *Criteria, options *Options) (*StockWarnInsufficientQtys, error) {
	swiqs := &StockWarnInsufficientQtys{}
	if err := c.SearchRead(StockWarnInsufficientQtyModel, criteria, options, swiqs); err != nil {
		return nil, err
	}
	return swiqs, nil
}

// FindStockWarnInsufficientQtyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarnInsufficientQtyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockWarnInsufficientQtyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockWarnInsufficientQtyId finds record id by querying it with criteria.
func (c *Client) FindStockWarnInsufficientQtyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockWarnInsufficientQtyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.warn.insufficient.qty was not found with criteria %v and options %v", criteria, options)
}
