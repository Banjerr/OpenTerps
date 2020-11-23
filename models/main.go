package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

type Effect struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

type Strain struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

type Smell struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

type Taste struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

type Terpene struct {
	gorm.Model
	Name       string     `json:"name"`
	Categories []Category `json:"category,omitempty" gorm:"many2many:terpene_categories;"`
	Tastes     []Taste    `json:"tastes,omitempty" gorm:"many2many:terpene_tastes;"`
	Smells     []Smell    `json:"smells,omitempty" gorm:"many2many:terpene_smells;"`
	Strains    []Strain   `json:"strains,omitempty" gorm:"many2many:terpene_strains;"`
	Effects    []Effect   `json:"effects,omitempty" gorm:"many2many:terpene_effects;"`
}

type TerpeneResponse struct {
	gorm.Model
	Name       string             `json:"name"`
	Categories []CategoryResponse `json:"category,omitempty" gorm:"many2many:terpene_categories;"`
	Tastes     []TasteResponse    `json:"tastes,omitempty" gorm:"many2many:terpene_tastes;"`
	Smells     []SmellResponse    `json:"smells,omitempty" gorm:"many2many:terpene_smells;"`
	Strains    []StrainResponse   `json:"strains,omitempty" gorm:"many2many:terpene_strains;"`
	Effects    []EffectResponse   `json:"effects,omitempty" gorm:"many2many:terpene_effects;"`
}

type CategoryResponse struct {
	Name string `json:"name,omitempty"`
}

type EffectResponse struct {
	Name string `json:"name,omitempty"`
}

type StrainResponse struct {
	Name string `json:"name,omitempty"`
}

type SmellResponse struct {
	Name string `json:"name,omitempty"`
}

type TasteResponse struct {
	Name string `json:"name,omitempty"`
}

type AuthHeader struct {
	APIUser string `json:"apiuser"`
	APIKey  string `json:"apikey"`
}
