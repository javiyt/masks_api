//go:generate mockery --name=OwnerRepository --inpackage
//go:generate mockery --name=MaskRepository --inpackage

package domain

import "github.com/google/uuid"

type OwnerRepository interface {
	Save(owner Owner) error
	GetAll() Owners
}

type MaskRepository interface {
	Save(mask Mask) error
	GetByOwner(owner uuid.UUID) (Mask, error)
}
