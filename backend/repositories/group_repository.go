package repositories

import (
	"context"
	"mellow/models"
)

type GroupRepository interface {
	// InsertGroup crée un nouveau groupe.
	InsertGroup(ctx context.Context, group *models.Group) error

	// UpdateGroup met à jour le titre ou la description d'un groupe.
	UpdateGroup(ctx context.Context, group *models.Group) error

	// GetGroupByID retourne un groupe par son ID.
	GetGroupByID(ctx context.Context, groupID string) (*models.Group, error)

	// GetAllGroups retourne tous les groupes.
	GetAllGroups(ctx context.Context) ([]*models.Group, error)

	// DeleteGroup supprime un groupe.
	DeleteGroup(ctx context.Context, groupID string) error

	// AddMember ajoute un utilisateur à un groupe.
	AddMember(ctx context.Context, groupID, userID string) error

	// RemoveMember retire un utilisateur d’un groupe.
	RemoveMember(ctx context.Context, groupID, userID string) error

	// GetGroupMembers retourne les membres d’un groupe.
	GetGroupMembers(ctx context.Context, groupID string) ([]*models.User, error)

	// IsMember vérifie si un utilisateur est membre d’un groupe.
	IsMember(ctx context.Context, groupID, userID string) (bool, error)

	// GetGroupsJoinedByUser retourne les groupes auxquels un utilisateur a adhéré.
	GetGroupsJoinedByUser(ctx context.Context, userID string) ([]*models.Group, error)

	// IsTitleTaken vérifie si un titre de groupe est déjà utilisé.
	IsTitleTaken(ctx context.Context, title string) (bool, error)
}
