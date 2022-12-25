// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one

INSERT INTO
    "order".orders (
        id,
        order_source,
        loyalty_member_id,
        order_status,
        updated
    )
VALUES ($1, $2, $3, $4, $5) RETURNING id, order_source, loyalty_member_id, order_status, updated
`

type CreateOrderParams struct {
	ID              uuid.UUID
	OrderSource     int32
	LoyaltyMemberID uuid.UUID
	OrderStatus     int32
	Updated         sql.NullTime
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (OrderOrder, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.ID,
		arg.OrderSource,
		arg.LoyaltyMemberID,
		arg.OrderStatus,
		arg.Updated,
	)
	var i OrderOrder
	err := row.Scan(
		&i.ID,
		&i.OrderSource,
		&i.LoyaltyMemberID,
		&i.OrderStatus,
		&i.Updated,
	)
	return i, err
}

const getAll = `-- name: GetAll :many

SELECT
    o.id,
    order_source,
    loyalty_member_id,
    order_status,
    l.id as "line_item_id",
    item_type,
    name,
    price,
    item_status,
    is_barista_order
FROM "order".orders o
    LEFT JOIN "order".line_items l ON o.id = l.order_id
`

type GetAllRow struct {
	ID              uuid.UUID
	OrderSource     int32
	LoyaltyMemberID uuid.UUID
	OrderStatus     int32
	LineItemID      uuid.NullUUID
	ItemType        int32
	Name            string
	Price           string
	ItemStatus      int32
	IsBaristaOrder  bool
}

func (q *Queries) GetAll(ctx context.Context) ([]GetAllRow, error) {
	rows, err := q.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllRow
	for rows.Next() {
		var i GetAllRow
		if err := rows.Scan(
			&i.ID,
			&i.OrderSource,
			&i.LoyaltyMemberID,
			&i.OrderStatus,
			&i.LineItemID,
			&i.ItemType,
			&i.Name,
			&i.Price,
			&i.ItemStatus,
			&i.IsBaristaOrder,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getByID = `-- name: GetByID :many

SELECT
    o.id,
    order_source,
    loyalty_member_id,
    order_status,
    l.id as "line_item_id",
    item_type,
    name,
    price,
    item_status,
    is_barista_order
FROM "order".orders o
    LEFT JOIN "order".line_items l ON o.id = l.order_id
WHERE o.id = $1
`

type GetByIDRow struct {
	ID              uuid.UUID
	OrderSource     int32
	LoyaltyMemberID uuid.UUID
	OrderStatus     int32
	LineItemID      uuid.NullUUID
	ItemType        int32
	Name            string
	Price           string
	ItemStatus      int32
	IsBaristaOrder  bool
}

func (q *Queries) GetByID(ctx context.Context, id uuid.UUID) ([]GetByIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetByIDRow
	for rows.Next() {
		var i GetByIDRow
		if err := rows.Scan(
			&i.ID,
			&i.OrderSource,
			&i.LoyaltyMemberID,
			&i.OrderStatus,
			&i.LineItemID,
			&i.ItemType,
			&i.Name,
			&i.Price,
			&i.ItemStatus,
			&i.IsBaristaOrder,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertItemLine = `-- name: InsertItemLine :one

INSERT INTO
    "order".line_items (
        id,
        item_type,
        name,
        price,
        item_status,
        is_barista_order,
        order_id,
        created,
        updated
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, item_type, name, price, item_status, is_barista_order, order_id, created, updated
`

type InsertItemLineParams struct {
	ID             uuid.UUID
	ItemType       int32
	Name           string
	Price          string
	ItemStatus     int32
	IsBaristaOrder bool
	OrderID        uuid.NullUUID
	Created        time.Time
	Updated        sql.NullTime
}

func (q *Queries) InsertItemLine(ctx context.Context, arg InsertItemLineParams) (OrderLineItem, error) {
	row := q.db.QueryRowContext(ctx, insertItemLine,
		arg.ID,
		arg.ItemType,
		arg.Name,
		arg.Price,
		arg.ItemStatus,
		arg.IsBaristaOrder,
		arg.OrderID,
		arg.Created,
		arg.Updated,
	)
	var i OrderLineItem
	err := row.Scan(
		&i.ID,
		&i.ItemType,
		&i.Name,
		&i.Price,
		&i.ItemStatus,
		&i.IsBaristaOrder,
		&i.OrderID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const updateItemLine = `-- name: UpdateItemLine :exec

UPDATE "order".line_items
SET
    item_status = $2,
    updated = $3
WHERE id = $1
`

type UpdateItemLineParams struct {
	ID         uuid.UUID
	ItemStatus int32
	Updated    sql.NullTime
}

func (q *Queries) UpdateItemLine(ctx context.Context, arg UpdateItemLineParams) error {
	_, err := q.db.ExecContext(ctx, updateItemLine, arg.ID, arg.ItemStatus, arg.Updated)
	return err
}

const updateOrder = `-- name: UpdateOrder :exec

UPDATE "order".orders
SET
    order_status = $2,
    updated = $3
WHERE id = $1
`

type UpdateOrderParams struct {
	ID          uuid.UUID
	OrderStatus int32
	Updated     sql.NullTime
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateOrder, arg.ID, arg.OrderStatus, arg.Updated)
	return err
}
