package domain

import (
	"github.com/google/uuid"
	"time"
)

type Owner struct {
	id   uuid.UUID
	name string
}

type Owners []Owner

func NewOwner(id uuid.UUID, name string) Owner {
	return Owner{id: id, name: name}
}

func (o Owner) Id() uuid.UUID {
	return o.id
}

func (o Owner) Name() string {
	return o.name
}

type Mask struct {
	id uuid.UUID
	owner uuid.UUID
	alias string
	maxTimes uint8
	used uint8
	firstUse time.Time
}

func NewMask(id uuid.UUID, owner uuid.UUID, alias string, maxTimes uint8, used uint8, firstUse time.Time) Mask {
	return Mask{id: id, owner: owner, alias: alias, maxTimes: maxTimes, used: used, firstUse: firstUse}
}

func (m Mask) FirstUse() time.Time {
	return m.firstUse
}

func (m Mask) Id() uuid.UUID {
	return m.id
}

func (m Mask) Owner() uuid.UUID {
	return m.owner
}

func (m Mask) Alias() string {
	return m.alias
}

func (m Mask) MaxTimes() uint8 {
	return m.maxTimes
}

func (m Mask) Used() uint8 {
	return m.used
}
