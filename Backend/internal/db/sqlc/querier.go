// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateType(ctx context.Context, name string) (Type, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserGroup(ctx context.Context, arg CreateUserGroupParams) (UserGroup, error)
	DeleteEntry(ctx context.Context, id uuid.UUID) error
	DeleteGroup(ctx context.Context, id uuid.UUID) error
	DeleteType(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	DeleteUserGroup(ctx context.Context, id uuid.UUID) error
	GetEntry(ctx context.Context, id uuid.UUID) (GetEntryRow, error)
	GetGroup(ctx context.Context, id uuid.UUID) (Group, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetType(ctx context.Context, id uuid.UUID) (Type, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error)
	GetUserGroup(ctx context.Context, id uuid.UUID) (UserGroup, error)
	GetUserGroupByGroup(ctx context.Context, groupID uuid.NullUUID) (UserGroup, error)
	GetUserGroupByUser(ctx context.Context, userID uuid.NullUUID) (UserGroup, error)
	ListEntryByGroup(ctx context.Context, groupID uuid.NullUUID) ([]ListEntryByGroupRow, error)
	ListEntryByUser(ctx context.Context, userID uuid.UUID) ([]ListEntryByUserRow, error)
	ListGroupsByUser(ctx context.Context, userID uuid.UUID) ([]Group, error)
	ListTypes(ctx context.Context) ([]Type, error)
	ListUserGroups(ctx context.Context, arg ListUserGroupsParams) ([]UserGroup, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error)
	UpdateGroups(ctx context.Context, arg UpdateGroupsParams) (Group, error)
	UpdateType(ctx context.Context, arg UpdateTypeParams) (Type, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
	UpdateUserGroup(ctx context.Context, arg UpdateUserGroupParams) (UserGroup, error)
}

var _ Querier = (*Queries)(nil)
