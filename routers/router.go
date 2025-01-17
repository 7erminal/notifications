// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Users/bedeabbe/Desktop/projects/golang/notification_service/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/account_verification",
			beego.NSInclude(
				&controllers.AccountVerificationController{},
			),
		),

		beego.NSNamespace("/accounts",
			beego.NSInclude(
				&controllers.AccountsController{},
			),
		),

		beego.NSNamespace("/categories",
			beego.NSInclude(
				&controllers.CategoriesController{},
			),
		),

		beego.NSNamespace("/countries",
			beego.NSInclude(
				&controllers.CountriesController{},
			),
		),

		beego.NSNamespace("/currencies",
			beego.NSInclude(
				&controllers.CurrenciesController{},
			),
		),

		beego.NSNamespace("/customer_categories",
			beego.NSInclude(
				&controllers.CustomerCategoriesController{},
			),
		),

		beego.NSNamespace("/customers",
			beego.NSInclude(
				&controllers.CustomersController{},
			),
		),

		beego.NSNamespace("/features",
			beego.NSInclude(
				&controllers.FeaturesController{},
			),
		),

		beego.NSNamespace("/item_features",
			beego.NSInclude(
				&controllers.ItemFeaturesController{},
			),
		),

		beego.NSNamespace("/item_images",
			beego.NSInclude(
				&controllers.ItemImagesController{},
			),
		),

		beego.NSNamespace("/item_prices",
			beego.NSInclude(
				&controllers.ItemPricesController{},
			),
		),

		beego.NSNamespace("/item_purposes",
			beego.NSInclude(
				&controllers.ItemPurposesController{},
			),
		),

		beego.NSNamespace("/item_quantity",
			beego.NSInclude(
				&controllers.ItemQuantityController{},
			),
		),

		beego.NSNamespace("/item_reviews",
			beego.NSInclude(
				&controllers.ItemReviewsController{},
			),
		),

		beego.NSNamespace("/items",
			beego.NSInclude(
				&controllers.ItemsController{},
			),
		),

		beego.NSNamespace("/migrations",
			beego.NSInclude(
				&controllers.MigrationsController{},
			),
		),

		beego.NSNamespace("/newsletter_customers",
			beego.NSInclude(
				&controllers.NewsletterCustomersController{},
			),
		),

		beego.NSNamespace("/order_items",
			beego.NSInclude(
				&controllers.OrderItemsController{},
			),
		),

		beego.NSNamespace("/orders",
			beego.NSInclude(
				&controllers.OrdersController{},
			),
		),

		beego.NSNamespace("/payment_history",
			beego.NSInclude(
				&controllers.PaymentHistoryController{},
			),
		),

		beego.NSNamespace("/payment_methods",
			beego.NSInclude(
				&controllers.PaymentMethodsController{},
			),
		),

		beego.NSNamespace("/payment_types",
			beego.NSInclude(
				&controllers.PaymentTypesController{},
			),
		),

		beego.NSNamespace("/payments",
			beego.NSInclude(
				&controllers.PaymentsController{},
			),
		),

		beego.NSNamespace("/purposes",
			beego.NSInclude(
				&controllers.PurposesController{},
			),
		),

		beego.NSNamespace("/service",
			beego.NSInclude(
				&controllers.ServiceController{},
			),
		),

		beego.NSNamespace("/shops",
			beego.NSInclude(
				&controllers.ShopsController{},
			),
		),

		beego.NSNamespace("/status",
			beego.NSInclude(
				&controllers.StatusController{},
			),
		),

		beego.NSNamespace("/status_codes",
			beego.NSInclude(
				&controllers.StatusCodesController{},
			),
		),

		beego.NSNamespace("/status_types",
			beego.NSInclude(
				&controllers.StatusTypesController{},
			),
		),

		beego.NSNamespace("/transaction_details",
			beego.NSInclude(
				&controllers.TransactionDetailsController{},
			),
		),

		beego.NSNamespace("/transactions",
			beego.NSInclude(
				&controllers.TransactionsController{},
			),
		),

		beego.NSNamespace("/user_otps",
			beego.NSInclude(
				&controllers.UserOtpsController{},
			),
		),

		beego.NSNamespace("/user_types",
			beego.NSInclude(
				&controllers.UserTypesController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
