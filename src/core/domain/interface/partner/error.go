/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 16:53
 * description :
 * history :
 */

package partner

import (
	"go2o/src/core/infrastructure/domain"
)

var (
	ErrNoSuchPartner *domain.DomainError = domain.NewDomainError(
		"no_such_partner", "商家不存在")

	ErrNoSuchShop *domain.DomainError = domain.NewDomainError(
		"no_such_shop", "门店不存在")

	ErrPartnerNotMatch *domain.DomainError = domain.NewDomainError(
		"not_match", "商家不匹配")

	ErrRegisterMode *domain.DomainError = domain.NewDomainError(
		"err_register_mode", "注册模式异常")

	ErrSalesPercent *domain.DomainError = domain.NewDomainError(
		"err_sales_percent", "销售比例错误")
)
