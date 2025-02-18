package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "AddNotificationCategory",
            Router: `/add-notification-category`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "AddNotificationMessage",
            Router: `/add-notification-message`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "GetAllNoficationCategories",
            Router: `/get-all-notification-categories`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "GetAllNoficationMessages",
            Router: `/get-all-notification-messages`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "GetAllUserNotifications",
            Router: `/get-user-notifications/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["notification_service/controllers:NotificationsController"],
        beego.ControllerComments{
            Method: "UpdateNotificationReadStatus",
            Router: `/update-read-status/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
