package domain_test

import (
	"errors"
	"github.com/javiyt/masks_api/internal/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSaveOwnerServiceImpl_New(t *testing.T) {
	r := new(domain.MockOwnerRepository)
	s := domain.NewSaveOwnerService(r)

	t.Run("it should fail saving a mask owner", func(t *testing.T) {
		r.On("Save", mock.MatchedBy(func(owner domain.Owner) bool {
			return owner.Name() == "test"
		})).Times(1).
			Return(errors.New("testing error"))

		e := s.New("test")

		r.AssertExpectations(t)
		require.EqualError(t, e, "testing error")
	})

	t.Run("it should save a mask owner", func(t *testing.T) {
		r.On("Save", mock.MatchedBy(func(owner domain.Owner) bool {
			return owner.Name() == "test"
		})).Times(1).
			Return(nil)

		e := s.New("test")

		r.AssertExpectations(t)
		require.NoError(t, e)
	})
}

func TestGetOwnersServiceImpl_GetAll(t *testing.T) {
	r := new(domain.MockOwnerRepository)
	s := domain.NewGetOwnersServiceImpl(r)

	r.On("GetAll").Times(1).Return(make(domain.Owners, 0))

	owners := s.GetAll()

	require.Empty(t, owners)
	r.AssertExpectations(t)
}
