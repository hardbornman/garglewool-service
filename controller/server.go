package controller

import (
	"fmt"
	m "gitee.com/gbat/utils/middleware"
	"github.com/gin-gonic/gin"
	"github.com/hardbornman/garglewool-service/controller/middleware"
	"github.com/hardbornman/garglewool-service/controller/rest/gwComment"
	"github.com/hardbornman/garglewool-service/controller/rest/gwOrder"
	"github.com/hardbornman/garglewool-service/controller/rest/gwPackage"
	"github.com/hardbornman/garglewool-service/controller/rest/gwPackagedetail"
	"github.com/hardbornman/garglewool-service/controller/rest/gwShop"
	"github.com/hardbornman/garglewool-service/controller/rest/gwUser"
	"github.com/hardbornman/garglewool-service/controller/rest/gwVoucher"
	"github.com/hardbornman/garglewool-service/controller/rest/sysDictionary"
	"github.com/hardbornman/garglewool-service/controller/rest/sysDictionarycategory"
	"github.com/hardbornman/garglewool-service/initials/config"
	"github.com/hardbornman/garglewool-service/model"
)

// 启动服务
func Start() {

	//region 初始化gin
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
	r := gin.New()
	r.Use(gin.Recovery())
	//endregion

	//region 添加路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, model.Response{Code: "not found", Message: "Page not found"})
	})
	r.POST("wechatlogin", gwUser.WeChatLogin)
	r.POST("/auth", gwUser.Auth)

	v1 := r.Group("/v1")
	v1.Use(middleware.LimitMiddleware(), middleware.CORSMiddleware())
	{

		//region REST接口
		restNode := v1.Group("rest")
		restNode.Use(m.JWTAuth())
		{

			//region 评论管理
			gwCommentNode := restNode.Group("gwComments")
			{
				gwCommentNode.GET("", gwComment.GetGwCommentList)
				gwCommentNode.GET(":id", gwComment.GetGwCommentInfo)
				gwCommentNode.POST("", gwComment.InsertGwCommentInfo)
				gwCommentNode.PUT(":id", gwComment.UpdateGwCommentInfo)
				gwCommentNode.DELETE(":id", gwComment.DeleteGwCommentInfo)
			}
			//endregion

			//region 订单管理
			gwOrderNode := restNode.Group("gwOrders")
			{
				gwOrderNode.GET("", gwOrder.GetGwOrderList)
				gwOrderNode.GET(":id", gwOrder.GetGwOrderInfo)
				gwOrderNode.POST("", gwOrder.InsertGwOrderInfo)
				gwOrderNode.PUT(":id", gwOrder.UpdateGwOrderInfo)
				gwOrderNode.DELETE(":id", gwOrder.DeleteGwOrderInfo)
			}
			//endregion

			//region 套餐管理
			gwPackageNode := restNode.Group("gwPackages")
			{
				gwPackageNode.GET("", gwPackage.GetGwPackageList)
				gwPackageNode.GET(":id", gwPackage.GetGwPackageInfo)
				gwPackageNode.POST("", gwPackage.InsertGwPackageInfo)
				gwPackageNode.PUT(":id", gwPackage.UpdateGwPackageInfo)
				gwPackageNode.DELETE(":id", gwPackage.DeleteGwPackageInfo)
			}
			//endregion

			//region 套餐明细管理
			gwPackagedetailNode := restNode.Group("gwPackagedetails")
			{
				gwPackagedetailNode.GET("", gwPackagedetail.GetGwPackagedetailList)
				gwPackagedetailNode.GET(":id", gwPackagedetail.GetGwPackagedetailInfo)
				gwPackagedetailNode.POST("", gwPackagedetail.InsertGwPackagedetailInfo)
				gwPackagedetailNode.PUT(":id", gwPackagedetail.UpdateGwPackagedetailInfo)
				gwPackagedetailNode.DELETE(":id", gwPackagedetail.DeleteGwPackagedetailInfo)
			}
			//endregion

			//region 店铺管理
			gwShopNode := restNode.Group("gwShops")
			{
				gwShopNode.GET("", gwShop.GetGwShopList)
				gwShopNode.GET(":id", gwShop.GetGwShopInfo)
				gwShopNode.POST("", gwShop.InsertGwShopInfo)
				gwShopNode.PUT(":id", gwShop.UpdateGwShopInfo)
				gwShopNode.DELETE(":id", gwShop.DeleteGwShopInfo)
			}
			//endregion

			//region 用户管理
			gwUserNode := restNode.Group("gwUsers")
			{
				gwUserNode.GET("", gwUser.GetGwUserList)
				gwUserNode.GET(":id", gwUser.GetGwUserInfo)
				gwUserNode.POST("", gwUser.InsertGwUserInfo)
				gwUserNode.PUT(":id", gwUser.UpdateGwUserInfo)
				gwUserNode.DELETE(":id", gwUser.DeleteGwUserInfo)
			}
			//endregion

			//region 抵用券管理
			gwVoucherNode := restNode.Group("gwVouchers")
			{
				gwVoucherNode.GET("", gwVoucher.GetGwVoucherList)
				gwVoucherNode.GET(":id", gwVoucher.GetGwVoucherInfo)
				gwVoucherNode.POST("", gwVoucher.InsertGwVoucherInfo)
				gwVoucherNode.PUT(":id", gwVoucher.UpdateGwVoucherInfo)
				gwVoucherNode.DELETE(":id", gwVoucher.DeleteGwVoucherInfo)
			}
			//endregion

			//region 字典表
			sysDictionaryNode := restNode.Group("sysDictionarys")
			{
				sysDictionaryNode.GET("", sysDictionary.GetSysDictionaryList)
				sysDictionaryNode.GET(":id", sysDictionary.GetSysDictionaryInfo)
				sysDictionaryNode.POST("", sysDictionary.InsertSysDictionaryInfo)
				sysDictionaryNode.PUT(":id", sysDictionary.UpdateSysDictionaryInfo)
				sysDictionaryNode.DELETE(":id", sysDictionary.DeleteSysDictionaryInfo)
			}
			//endregion

			//region 字典分类
			sysDictionarycategoryNode := restNode.Group("sysDictionarycategorys")
			{
				sysDictionarycategoryNode.GET("", sysDictionarycategory.GetSysDictionarycategoryList)
				sysDictionarycategoryNode.GET(":id", sysDictionarycategory.GetSysDictionarycategoryInfo)
				sysDictionarycategoryNode.POST("", sysDictionarycategory.InsertSysDictionarycategoryInfo)
				sysDictionarycategoryNode.PUT(":id", sysDictionarycategory.UpdateSysDictionarycategoryInfo)
				sysDictionarycategoryNode.DELETE(":id", sysDictionarycategory.DeleteSysDictionarycategoryInfo)
			}
			//endregion

			//region 其它
			{
			}
			//endregion

		}
		//endregion

	}
	//endregion

	//region 启动api服务
	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Config.Port))
	//endregion

}
