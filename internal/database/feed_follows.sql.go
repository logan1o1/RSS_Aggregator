// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, created_at, updated_at, userId, feedId)
VALUES($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, userid, feedid
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Userid    uuid.UUID
	Feedid    uuid.UUID
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Userid,
		arg.Feedid,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Userid,
		&i.Feedid,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec

DELETE FROM feed_follows WHERE id = $1 AND userid = $2
`

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	Userid uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.Userid)
	return err
}

const getFeedFollows = `-- name: GetFeedFollows :many

SELECT id, created_at, updated_at, userid, feedid FROM feed_follows WHERE userid = $1
`

func (q *Queries) GetFeedFollows(ctx context.Context, userid uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Userid,
			&i.Feedid,
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
