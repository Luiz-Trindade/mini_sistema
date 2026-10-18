package models

import "time"

type User struct {
	ID           uint         `gorm:"primaryKey" json:"id"`
	Nome         string       `gorm:"size:150;not null" json:"nome"`
	Email        string       `gorm:"size:150;uniqueIndex;not null" json:"email"`
	Senha        string       `gorm:"size:255;not null" json:"-"`
	CriadoEm     time.Time    `gorm:"autoCreateTime" json:"criado_em"`
	AtualizEm    time.Time    `gorm:"autoUpdateTime" json:"atualizado_em"`
	UltimoLogin  *time.Time   `json:"ultimo_login"`
	GroupID      *uint        `gorm:"index" json:"group_id"`
	Group        *Group       `gorm:"foreignKey:GroupID" json:"group,omitempty"`
	Permissions  []Permission `gorm:"many2many:user_permissions;" json:"permissions,omitempty"`
	IsAdmin      bool         `gorm:"default:false" json:"is_admin"`
	IsStaff      bool         `gorm:"default:false" json:"is_staff"`
	TokenVersion int          `gorm:"default:1" json:"-"` // Versão para invalidação de tokens
}
