/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-12 16:21
 * description :
 * history :
 */
package delivery

type IDelivery interface {
	// 返回聚合编号
	GetAggregateRootId() int

	// 等同于GetAggregateRootId()
	GetPartnerId() int

	// 获取最近的配送区域
	GetNearestCoverage(lng, lat float64) ICoverageArea

	// 根据地址获取地区(可能会有重复的区名)
	GetArea(addr string) ([]*AreaValue, error)

	//　获取覆盖区域
	GetCoverageArea(id int) ICoverageArea

	// 查看单个所在的区域
	FindSingleCoverageArea(lng, lat float64) ICoverageArea

	// 查找所有所在的区域
	FindCoverageAreas(lng, lat float64) []ICoverageArea

	// 获取配送信息
	GetDeliveryInfo(coverageId int) (shopId, deliverUsrId int, err error)
}
