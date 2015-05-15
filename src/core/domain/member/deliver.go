/**
 * Copyright 2014 @ S1N1 Team.
 * name :
 * author : jarryliu
 * date : 2013-12-23 21:09
 * description :
 * history :
 */

package member

import (
	"go2o/src/core/domain/interface/member"
)

var _ member.IDeliver = new(deliver)

type deliver struct {
	value     *member.DeliverAddress
	memberRep member.IMemberRep
}

func newDeliver(v *member.DeliverAddress, memberRep member.IMemberRep) member.IDeliver {
	return &deliver{
		value:     v,
		memberRep: memberRep,
	}
}

func (this *deliver) GetDomainId() int {
	return this.value.Id
}

func (this *deliver) GetValue() member.DeliverAddress {
	return *this.value
}

func (this *deliver) SetValue(v *member.DeliverAddress) error {
	if this.value.MemberId == v.MemberId {
		this.value = v
	}
	return nil
}

func (this *deliver) Save() (int, error) {
	return this.memberRep.SaveDeliver(this.value)
}
