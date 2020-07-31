//START STUB OMIT

// Client API for Hydration service
type HydrationClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

func NewHydrationClient(cc *grpc.ClientConn) HydrationClient {
	return &hydrationClient{cc}
}

//END STUB OMIT

//START USAGE OMIT

// Usage
import "google.golang.org/grpc"

conn, err := grpc.Dial("localhost:8989", opts...) // HL
defer conn.Close()
c := pb.NewHydrationClient(conn) // HL

//END USAGE OMIT


//START IMPL OMIT

// Stub Implementation
type GetUserRequest struct {
	UserID    string
	FieldMask *google_protobuf.FieldMask 
}

type GetUserResponse struct {
	DataField1    string
	DataField2    int
	DataField3    AnotherType
}

func (c *hydrationClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	...
	err := grpc.Invoke(ctx, "/user.Hydration/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//END IMPL OMIT
