// Code generated by svcdec; DO NOT EDIT.

package fleet

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedInventory struct {
	// Service is the service to decorate.
	Service InventoryServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(ctx context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedInventory) GetStableVersion(ctx context.Context, req *GetStableVersionRequest) (rsp *GetStableVersionResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetStableVersion", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetStableVersion(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetStableVersion", rsp, err)
	}
	return
}

func (s *DecoratedInventory) SetSatlabStableVersion(ctx context.Context, req *SetSatlabStableVersionRequest) (rsp *SetSatlabStableVersionResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "SetSatlabStableVersion", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.SetSatlabStableVersion(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "SetSatlabStableVersion", rsp, err)
	}
	return
}

func (s *DecoratedInventory) DeleteSatlabStableVersion(ctx context.Context, req *DeleteSatlabStableVersionRequest) (rsp *DeleteSatlabStableVersionResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteSatlabStableVersion", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteSatlabStableVersion(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteSatlabStableVersion", rsp, err)
	}
	return
}

func (s *DecoratedInventory) DumpStableVersionToDatastore(ctx context.Context, req *DumpStableVersionToDatastoreRequest) (rsp *DumpStableVersionToDatastoreResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DumpStableVersionToDatastore", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DumpStableVersionToDatastore(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DumpStableVersionToDatastore", rsp, err)
	}
	return
}
