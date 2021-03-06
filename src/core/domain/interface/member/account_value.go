/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-15 13:58
 * description :
 * history :
 */

package member

// 账户值对象
type AccountValue struct {
	MemberId int `db:"member_id" pk:"yes" json:"memberId"`
	// 积分
	Integral int `db:"integral"`
	// 账户余额
	Balance float32 `db:"balance" json:"balance"`
	// 冻结金额
	FreezesFee float32 `db:"freezes_fee" json:"freezesFee"`
	// 奖金账户余额
	PresentBalance float32 `db:"present_balance" json:"presentBalance"`
	FlowBalance    float32 `db:"flow_balance" json:"flowBalance"` // 流动账户余额
	GrowAmount     float32 `db:"grow_amount" json:"growAmount"`   // 增利账户金额

	// 冻结赠送额
	FreezesPresent float32 `db:"freezes_present" json:"freezesPresent"`
	// 总赠送金额
	TotalPresentFee float32 `db:"total_present_fee" json:"totalPresentFee"`
	// 总消费额
	TotalFee float32 `db:"total_fee" json:"totalFee"`
	// 总充值额
	TotalCharge float32 `db:"total_charge" json:"totalCharge"`
	// 总支付额
	TotalPay float32 `db:"total_pay" json:"totalPay"`
	// 更新时间
	UpdateTime int64 `db:"update_time" json:"updateTime"`
}
