package controller

import (
	"fmt"
	m "gitee.com/gbat/utils/middleware"
	"github.com/gin-gonic/gin"
	"github.com/hardbornman/garglewool-service/controller/middleware"
	"github.com/hardbornman/garglewool-service/controller/rest"
	"github.com/hardbornman/garglewool-service/controller/rest/comment"
	"github.com/hardbornman/garglewool-service/controller/rest/dictionary"
	"github.com/hardbornman/garglewool-service/controller/rest/dictionarycategory"
	"github.com/hardbornman/garglewool-service/controller/rest/guest"
	"github.com/hardbornman/garglewool-service/controller/rest/merchant"
	"github.com/hardbornman/garglewool-service/controller/rest/order"
	"github.com/hardbornman/garglewool-service/controller/rest/setmeal"
	"github.com/hardbornman/garglewool-service/controller/rest/setmealdetail"
	"github.com/hardbornman/garglewool-service/controller/rest/shop"
	"github.com/hardbornman/garglewool-service/controller/rest/userrole"
	"github.com/hardbornman/garglewool-service/controller/rest/voucher"
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
	r.POST("/auth", rest.Auth)
	v1 := r.Group("/v1")
	v1.Use(middleware.LimitMiddleware(), middleware.CORSMiddleware())
	{

		//region REST接口
		restNode := v1.Group("rest")
		restNode.Use(m.JWTAuth())
		{

			//region 评论表
			commentNode := restNode.Group("comments")
			{
				commentNode.GET("", comment.GetCommentList)
				commentNode.GET(":id", comment.GetCommentInfo)
				commentNode.POST("", comment.InsertCommentInfo)
				commentNode.PUT(":id", comment.UpdateCommentInfo)
				commentNode.DELETE(":id", comment.DeleteCommentInfo)
			}
			//endregion

			//region 字典表
			dictionaryNode := restNode.Group("dictionarys")
			{
				dictionaryNode.GET("", dictionary.GetDictionaryList)
				dictionaryNode.GET(":id", dictionary.GetDictionaryInfo)
				dictionaryNode.POST("", dictionary.InsertDictionaryInfo)
				dictionaryNode.PUT(":id", dictionary.UpdateDictionaryInfo)
				dictionaryNode.DELETE(":id", dictionary.DeleteDictionaryInfo)
			}
			//endregion

			//region 字典分类
			dictionarycategoryNode := restNode.Group("dictionarycategorys")
			{
				dictionarycategoryNode.GET("", dictionarycategory.GetDictionarycategoryList)
				dictionarycategoryNode.GET(":id", dictionarycategory.GetDictionarycategoryInfo)
				dictionarycategoryNode.POST("", dictionarycategory.InsertDictionarycategoryInfo)
				dictionarycategoryNode.PUT(":id", dictionarycategory.UpdateDictionarycategoryInfo)
				dictionarycategoryNode.DELETE(":id", dictionarycategory.DeleteDictionarycategoryInfo)
			}
			//endregion

			//region 买家客户表
			guestNode := restNode.Group("guests")
			{
				guestNode.GET("", guest.GetGuestList)
				guestNode.GET(":id", guest.GetGuestInfo)
				guestNode.POST("", guest.InsertGuestInfo)
				guestNode.PUT(":id", guest.UpdateGuestInfo)
				guestNode.DELETE(":id", guest.DeleteGuestInfo)
			}
			//endregion

			//region 商家用户表
			merchantNode := restNode.Group("merchants")
			{
				merchantNode.GET("", merchant.GetMerchantList)
				merchantNode.GET(":id", merchant.GetMerchantInfo)
				merchantNode.POST("", merchant.InsertMerchantInfo)
				merchantNode.PUT(":id", merchant.UpdateMerchantInfo)
				merchantNode.DELETE(":id", merchant.DeleteMerchantInfo)
			}
			//endregion

			//region 订单表
			orderNode := restNode.Group("orders")
			{
				orderNode.GET("", order.GetOrderList)
				orderNode.GET(":id", order.GetOrderInfo)
				orderNode.POST("", order.InsertOrderInfo)
				orderNode.PUT(":id", order.UpdateOrderInfo)
				orderNode.DELETE(":id", order.DeleteOrderInfo)
			}
			//endregion

			//region 套餐表
			setmealNode := restNode.Group("setmeals")
			{
				setmealNode.GET("", setmeal.GetSetmealList)
				setmealNode.GET(":id", setmeal.GetSetmealInfo)
				setmealNode.POST("", setmeal.InsertSetmealInfo)
				setmealNode.PUT(":id", setmeal.UpdateSetmealInfo)
				setmealNode.DELETE(":id", setmeal.DeleteSetmealInfo)
			}
			//endregion

			//region 套餐明细表
			setmealdetailNode := restNode.Group("setmealdetails")
			{
				setmealdetailNode.GET("", setmealdetail.GetSetmealdetailList)
				setmealdetailNode.GET(":id", setmealdetail.GetSetmealdetailInfo)
				setmealdetailNode.POST("", setmealdetail.InsertSetmealdetailInfo)
				setmealdetailNode.PUT(":id", setmealdetail.UpdateSetmealdetailInfo)
				setmealdetailNode.DELETE(":id", setmealdetail.DeleteSetmealdetailInfo)
			}
			//endregion

			//region 店铺表
			shopNode := restNode.Group("shops")
			{
				shopNode.GET("", shop.GetShopList)
				shopNode.GET(":id", shop.GetShopInfo)
				shopNode.POST("", shop.InsertShopInfo)
				shopNode.PUT(":id", shop.UpdateShopInfo)
				shopNode.DELETE(":id", shop.DeleteShopInfo)
			}
			//endregion

			//region 用户角色表
			userroleNode := restNode.Group("userroles")
			{
				userroleNode.GET("", userrole.GetUserroleList)
				userroleNode.GET(":id", userrole.GetUserroleInfo)
				userroleNode.POST("", userrole.InsertUserroleInfo)
				userroleNode.PUT(":id", userrole.UpdateUserroleInfo)
				userroleNode.DELETE(":id", userrole.DeleteUserroleInfo)
			}
			//endregion

			//region 抵用券管理
			voucherNode := restNode.Group("vouchers")
			{
				voucherNode.GET("", voucher.GetVoucherList)
				voucherNode.GET(":id", voucher.GetVoucherInfo)
				voucherNode.POST("", voucher.InsertVoucherInfo)
				voucherNode.PUT(":id", voucher.UpdateVoucherInfo)
				voucherNode.DELETE(":id", voucher.DeleteVoucherInfo)
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
