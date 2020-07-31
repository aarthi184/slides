// Server API for Hydration service
type HydrationServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
}

type GetUserRequest struct {
	UserID    string
	FieldMask *google_protobuf1.FieldMask
}

type GetUserResponse struct {
	DataField1    string
	DataField2    int
	DataField3    AnotherType
}
