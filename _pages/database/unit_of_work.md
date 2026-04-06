---
layout: default
title: Unit of Work (Transaction Manager)
description: "Unit of Work (Transaction Manager) in Go: Maintain data consistency by grouping related database operations into a single atomic unit."
nav_order: 1
parent: Database Patterns
permalink: /database/unit-of-work
---

# Unit of Work (Transaction Manager)

**Unit of Work** pattern helps maintain data consistency by grouping related database operations into a single atomic unit, preventing partial updates and maintaining referential integrity.

## Applicability

 - You need to perform multiple related database operations that must all succeed or fail together
 - You want to avoid scattered transaction management code across your application
 - You need to coordinate operations across multiple repositories within a single transaction
 - You want to decouple business logic from transaction management details
 - You're implementing complex workflows that involve multiple database tables

## Example: Standard library (`database/sql`)

Wrap `BeginTx`, `Commit`, and `Rollback` once, then pass work as a function so callers keep business logic in one place and transactions stay consistent.

```go
package main

import (
	"context"
	"database/sql"
)

// RunInTx executes fn inside a single transaction. If fn returns an error, the
// transaction is rolled back; otherwise it is committed.
func RunInTx(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }() // no-op after successful Commit
	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit()
}

// PlaceOrder inserts a parent row and a child row atomically.
func PlaceOrder(ctx context.Context, db *sql.DB, orderID, itemID string) error {
	return RunInTx(ctx, db, func(tx *sql.Tx) error {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO orders (id) VALUES (?)`, orderID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx,
			`INSERT INTO order_items (order_id, id) VALUES (?, ?)`, orderID, itemID)
		return err
	})
}
```

## Example: Ent (code-generated client)

**More advanced:** the snippet below is a fuller, Ent-specific setup—a `UnitOfWork` interface, transactional `*ent.Client` in `context` (`GetClient`), panic-safe rollback, and a service that runs several repository calls in one transaction.

```go
package main

import (
	"context"
	"fmt"
	"log/slog"

	ent "example.com/yourapp/ent" // your generated Ent client package
)

// UnitOfWork is the transaction boundary.
type UnitOfWork interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

type txKey struct{}

// entUnitOfWork implements UnitOfWork for *ent.Client.
type entUnitOfWork struct {
	client *ent.Client
}

func NewEntUnitOfWork(c *ent.Client) UnitOfWork {
	return &entUnitOfWork{client: c}
}

func (u *entUnitOfWork) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := u.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("start transaction: %w", err)
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	txCtx := context.WithValue(ctx, txKey{}, tx.Client())

	if err := fn(txCtx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("rollback: %w (original: %v)", rerr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}

// GetClient returns the transactional *ent.Client from ctx inside Do, or defaultClient otherwise.
func GetClient(ctx context.Context, defaultClient *ent.Client) *ent.Client {
	if c, ok := ctx.Value(txKey{}).(*ent.Client); ok {
		return c
	}
	return defaultClient
}

type UserService struct {
	uow               UnitOfWork
	logger            *slog.Logger
	commentRepository CommentRepository
	postRepository    PostRepository
	sessionRepository SessionRepository
	userRepository    UserRepository
}

// Delete removes the user and related data atomically (adjust order to your schema).
func (s *UserService) Delete(ctx context.Context, userID int) error {
	return s.uow.Do(ctx, func(txCtx context.Context) error {
		err := s.commentRepository.DeleteByUserID(txCtx, userID)
		if err != nil {
			s.logger.Error("delete comments by user",
				"error", err,
				"userID", userID,
			)
			return err
		}

		err = s.postRepository.DeleteByAuthorID(txCtx, userID)
		if err != nil {
			s.logger.Error("delete posts by author",
				"error", err,
				"userID", userID,
			)
			return err
		}

		err = s.sessionRepository.DeleteByUserID(txCtx, userID)
		if err != nil {
			s.logger.Error("delete sessions for user",
				"error", err,
				"userID", userID,
			)
			return err
		}

		err = s.userRepository.Delete(txCtx, userID)
		if err != nil {
			s.logger.Error("delete user",
				"error", err,
				"userID", userID,
			)
			return err
		}

		return nil
	})
}

// Implementations use GetClient(txCtx, appClient) and Ent builders (e.g. client.Comment(), client.Post()).
type CommentRepository interface {
	DeleteByUserID(txCtx context.Context, userID int) error
}
type PostRepository interface {
	DeleteByAuthorID(txCtx context.Context, userID int) error
}
type SessionRepository interface {
	DeleteByUserID(txCtx context.Context, userID int) error
}
type UserRepository interface {
	Delete(txCtx context.Context, userID int) error
}
```
