package handlers

import (
	"gobee/internal/database"
	"gobee/internal/handlers/album"
	albumphoto "gobee/internal/handlers/album_photo"
	apiinterface "gobee/internal/handlers/api_interface"
	auth "gobee/internal/handlers/auth"
	"gobee/internal/handlers/comment"
	"gobee/internal/handlers/common"
	"gobee/internal/handlers/essay"
	"gobee/internal/handlers/file"
	"gobee/internal/handlers/flink"
	flinkgroup "gobee/internal/handlers/flink_group"
	friendcirclerecord "gobee/internal/handlers/friend_circle_record"
	friendcirclerule "gobee/internal/handlers/friend_circle_rule"
	initialize "gobee/internal/handlers/initialize"
	paychannel "gobee/internal/handlers/pay_channel"
	payorder "gobee/internal/handlers/pay_order"
	"gobee/internal/handlers/post"
	"gobee/internal/handlers/role"
	"gobee/internal/handlers/route"
	setting_handler "gobee/internal/handlers/setting"
	storagestrategy "gobee/internal/handlers/storage_strategy"
	user_handler "gobee/internal/handlers/user"
	"gobee/internal/handlers/visit"
	"gobee/pkg"
)

type HandlerMap struct {
	AlbumHandler              album.AlbumHandler
	AlbumPhotoHandler         albumphoto.AlbumPhotoHandler
	ApiInterfaceHandler       apiinterface.ApiInterfaceHandler
	AuthHandler               auth.AuthHandler
	CommentHandler            comment.CommentHandler
	CommonHandler             common.CommonHandler
	FileHandler               file.FileHandler
	FlinkHandler              flink.FlinkHandler
	FlinkGroupHandler         flinkgroup.FlinkGroupHandler
	FriendCircleRecordHandler friendcirclerecord.FriendCircleRecordHandler
	FriendCircleRuleHandler   friendcirclerule.FriendCircleRuleHandler
	InitializeHandler         initialize.InitializeHandler
	PayChannelHandler         paychannel.PayChannelHandler
	PayOrderHandler           payorder.PayOrderHandler
	PostHandler               post.PostHandler
	RoleHandler               role.RoleHandler
	RouteHandler              route.RouteHandler
	SettingHandler            setting_handler.SettingHandler
	UserHandler               user_handler.UserHandler
	EssayHandler              essay.EssayHandler
	StorageStrategyHandler    storagestrategy.StorageStrategyHandler
	VisitHandler              visit.VisitHandler
}

func InitHandler(serviceMap pkg.ServiceMap) HandlerMap {
	albumHandler := album.NewAlbumHandlerImpl(database.DB)
	albnumPhotoHandler := albumphoto.NewAlbumPhotoHandlerImpl(database.DB)
	apiInterfaceHandler := apiinterface.NewApiInterfaceHandlerImpl(database.DB, serviceMap.ApiInterfaceService)
	authHandler := auth.NewAuthHandlerImpl(database.DB)
	commentHandler := comment.NewCommentHandlerImpl(database.DB, serviceMap.CommentService)
	commonHandler := common.NewCommonHandlerImpl(serviceMap.CommonService)
	fileHandler := file.NewFileHandlerImpl(database.DB)
	flinkHandler := flink.NewFlinkHandlerImpl(database.DB, serviceMap.FlinkService)
	flinkGroupHandler := flinkgroup.NewFlinkGroupHandlerImpl(database.DB, serviceMap.FlinkService)
	friendCircleRecordHandler := friendcirclerecord.NewFriendCircleRecordHandlerImpl(database.DB)
	friendCircleRuleHandler := friendcirclerule.NewFriendCircleRuleHandlerImpl(database.DB)
	initializeHandler := initialize.NewInitializeHandlerImpl(serviceMap.UserService, serviceMap.SettingService)
	payChannelHandler := paychannel.NewPayChannelHandlerImpl(database.DB)
	payOrderHandler := payorder.NewPayOrderHandlerImpl(database.DB, serviceMap.PayOrderService)
	postHandler := post.NewPostHandlerImpl(serviceMap.PostService)
	roleHandler := role.NewRoleHandlerImpl(serviceMap.RoleService)
	routeHandler := route.NewRouteHandlerImpl()
	settingHandler := setting_handler.NewSettingHandlerImpl(serviceMap.SettingService)
	userHandler := user_handler.NewUserHandlerImpl(serviceMap.UserService, serviceMap.RoleService)
	essayHandler := essay.NewEssayHandler(serviceMap.EssayService)
	storageStrategyHandler := storagestrategy.NewStorageStrategyHandlerImpl(database.DB)
	visitHandler := visit.NewVisitHandlerImpl(serviceMap.VisitService)

	handlerMap := HandlerMap{
		AlbumHandler:              albumHandler,
		AlbumPhotoHandler:         albnumPhotoHandler,
		ApiInterfaceHandler:       apiInterfaceHandler,
		AuthHandler:               authHandler,
		CommentHandler:            commentHandler,
		CommonHandler:             commonHandler,
		FileHandler:               fileHandler,
		FlinkHandler:              flinkHandler,
		FlinkGroupHandler:         flinkGroupHandler,
		FriendCircleRecordHandler: friendCircleRecordHandler,
		FriendCircleRuleHandler:   friendCircleRuleHandler,
		InitializeHandler:         initializeHandler,
		PayChannelHandler:         payChannelHandler,
		PayOrderHandler:           payOrderHandler,
		PostHandler:               postHandler,
		RoleHandler:               roleHandler,
		RouteHandler:              routeHandler,
		SettingHandler:            settingHandler,
		UserHandler:               userHandler,
		EssayHandler:              essayHandler,
		StorageStrategyHandler:    storageStrategyHandler,
		VisitHandler:              visitHandler,
	}

	return handlerMap

}
