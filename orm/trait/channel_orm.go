package trait

import "github.com/GoodHot/TinyCMS/core"

type VisibleLevel int

const (
	VisiblePrivate VisibleLevel = 0
	VisiblePaid    VisibleLevel = 1
	VisibleMember  VisibleLevel = 2
	VisiblePublic  VisibleLevel = 3
)

func (VisibleLevel) GetByInt(t int) VisibleLevel {
	switch t {
	case 0:
		return VisiblePrivate
	case 1:
		return VisiblePaid
	case 2:
		return VisibleMember
	case 3:
		return VisiblePublic
	}
	return VisiblePrivate
}

type Channel struct {
	BaseORM
	Avatar          string       `json:"avatar"`
	AvatarSVG       string       `json:"avatar_svg"`
	Name            string       `json:"name"`
	Path            string       `json:"path"`
	Visible         VisibleLevel `json:"visible"`
	MetaTitle       string       `json:"meta_title"`
	MetaDescription string       `json:"meta_description"`
	Sort            int          `json:"sort"`
	ParentID        int          `json:"parent_id"`
}

type ChannelORM interface {
	Save(channel *Channel) *core.Err
	GetByPath(path string) (*Channel, *core.Err)
	GetAll() ([]Channel, *core.Err)
	UpdateSort(data []Channel) *core.Err
}
