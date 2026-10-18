package models

import "time"

type Permission struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nome      string    `gorm:"size:100;not null" json:"nome"`
	Chave     string    `gorm:"size:100;uniqueIndex;not null" json:"chave"`
	CriadoEm  time.Time `json:"criado_em"`
	AtualizEm time.Time `json:"atualizado_em"`
}
