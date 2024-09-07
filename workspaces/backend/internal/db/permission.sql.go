// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: permission.sql

package db

import (
	"context"
	"database/sql"
)

const deletePermission = `-- name: DeletePermission :execrows
DELETE
FROM permissions
WHERE id = ?
`

func (q *Queries) DeletePermission(ctx context.Context, id string) (int64, error) {
	result, err := q.db.ExecContext(ctx, deletePermission, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getAllPermissions = `-- name: GetAllPermissions :many
SELECT id, created_at, updated_at, service_name, resource, action, attributes, description
FROM permissions
`

func (q *Queries) GetAllPermissions(ctx context.Context) ([]Permission, error) {
	rows, err := q.db.QueryContext(ctx, getAllPermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ServiceName,
			&i.Resource,
			&i.Action,
			&i.Attributes,
			&i.Description,
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

const insertNewPermission = `-- name: InsertNewPermission :exec
INSERT INTO permissions (id, service_name, resource, action, attributes, description)
VALUES (UUID(), ?, ?, ?, ?, ?)
`

type InsertNewPermissionParams struct {
	ServiceName string
	Resource    string
	Action      string
	Attributes  string
	Description sql.NullString
}

func (q *Queries) InsertNewPermission(ctx context.Context, arg InsertNewPermissionParams) error {
	_, err := q.db.ExecContext(ctx, insertNewPermission,
		arg.ServiceName,
		arg.Resource,
		arg.Action,
		arg.Attributes,
		arg.Description,
	)
	return err
}

const updatePermission = `-- name: UpdatePermission :exec
UPDATE permissions
SET service_name = COALESCE(?, service_name),
    resource     = COALESCE(?, resource),
    action       = COALESCE(?, action),
    attributes   = COALESCE(?, attributes),
    description  = COALESCE(?, description)
WHERE id = ?
`

type UpdatePermissionParams struct {
	ServiceName sql.NullString
	Resource    sql.NullString
	Action      sql.NullString
	Attributes  sql.NullString
	Description sql.NullString
	ID          string
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) error {
	_, err := q.db.ExecContext(ctx, updatePermission,
		arg.ServiceName,
		arg.Resource,
		arg.Action,
		arg.Attributes,
		arg.Description,
		arg.ID,
	)
	return err
}
