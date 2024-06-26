// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: project.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
  id,
  entity_id,
  title,
  description,
  amount,
  amount_raised
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, entity_id, title, description, amount, amount_raised, created_at, updated_at
`

type CreateProjectParams struct {
	ID           string      `json:"id"`
	EntityID     pgtype.Text `json:"entity_id"`
	Title        pgtype.Text `json:"title"`
	Description  pgtype.Text `json:"description"`
	Amount       pgtype.Int8 `json:"amount"`
	AmountRaised pgtype.Int8 `json:"amount_raised"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject,
		arg.ID,
		arg.EntityID,
		arg.Title,
		arg.Description,
		arg.Amount,
		arg.AmountRaised,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.EntityID,
		&i.Title,
		&i.Description,
		&i.Amount,
		&i.AmountRaised,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProject = `-- name: GetProject :one
SELECT id, entity_id, title, description, amount, amount_raised, created_at, updated_at FROM projects
WHERE id = $1 LIMIT $1
`

func (q *Queries) GetProject(ctx context.Context, limit int32) (Project, error) {
	row := q.db.QueryRow(ctx, getProject, limit)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.EntityID,
		&i.Title,
		&i.Description,
		&i.Amount,
		&i.AmountRaised,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
