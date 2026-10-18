package models

import "time"

type Group struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Nome        string       `gorm:"size:100;uniqueIndex;not null" json:"nome"`
	Permissions []Permission `gorm:"many2many:group_permissions;" json:"permissions,omitempty"`
	CriadoEm    time.Time    `json:"criado_em"`
	AtualizEm   time.Time    `json:"atualizado_em"`
}
