package handle

import "basic_gin/model"

type MemberData struct {
	data []model.Member
}

func NewMember() *MemberData {
	return &MemberData{}
}

func (m *MemberData) AllData() []model.Member {
	return m.data
}