package user

//type GrpcControllers struct {
//	types.UserCrud
//
//}
//
//func NewGrpcControllers(userCrud types.UserCrud) *GrpcControllers {
//	return &GrpcControllers{UserCrud: userCrud}
//}
//
//func (g GrpcControllers) CreateUser(ctx context.Context, in *ws.UserData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
//	//user := models.User{
//	//	ID:      1,
//	//	Name:    "nameInCrud",
//	//	Address: "addressInCrud",
//	//}
//	//g.Create(user)
//	fmt.Println("CreateUser")
//	//panic("implement")
//}
//
//func (g GrpcControllers) GetUserByID(ctx context.Context, in *ws.Id, opts ...grpc.CallOption) (*ws.UserDataResponse, error) {
//	panic("implement me")
//}
//
//func (g GrpcControllers) UpdateUserByID(ctx context.Context, in *ws.UserDataUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
//	panic("implement me")
//}
//
//func (g GrpcControllers) DeleteUserByID(ctx context.Context, in *ws.Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
//	panic("implement me")
//}
