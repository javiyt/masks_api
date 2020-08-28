package domain

import "github.com/google/uuid"

type SaveOwnerService interface {
	New(name string) error
}

type GetOwnersService interface {
	GetAll() Owners
}

type SaveOwnerServiceImpl struct {
	r OwnerRepository
}

func NewSaveOwnerService(r OwnerRepository) SaveOwnerService {
	return &SaveOwnerServiceImpl{r: r}
}

func (s *SaveOwnerServiceImpl) New(name string) error {
	return s.r.Save(NewOwner(uuid.New(), name))
}

type GetOwnersServiceImpl struct {
	r OwnerRepository
}

func NewGetOwnersServiceImpl(r OwnerRepository) GetOwnersService {
	return &GetOwnersServiceImpl{r: r}
}

func (g *GetOwnersServiceImpl) GetAll() Owners {
	return g.r.GetAll()
}





