package trait

import "github.com/GoodHot/TinyCMS/core"

type VisibilityLevel int

const (
	VisibilityPrivate VisibilityLevel = 0
	VisibilityPaid    VisibilityLevel = 1
	VisibilityMember  VisibilityLevel = 2
	VisibilityPublic  VisibilityLevel = 3
)

type Channel struct {
	BaseORM
	Avatar          string          `json:"avatar"`
	AvatarSVG       string          `json:"avatar_svg"`
	Name            string          `json:"name"`
	Path            string          `json:"path"`
	Visibility      VisibilityLevel `json:"visibility"`
	MetaTitle       string          `json:"meta_title"`
	MetaDescription string          `json:"meta_description"`
	Sort            int             `json:"sort"`
	ParentID        int             `json:"parent_id"`
}

type ChannelORM interface {
	Save(channel *Channel) *core.Err
	GetByPath(path string) (*Channel, *core.Err)
	GetAll() (*[]Channel, *core.Err)
	UpdateSort(data []Channel) *core.Err
}
