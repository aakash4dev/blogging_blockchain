package keeper

import (
	"context"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// make data type Post in profo/blog/blog/post.proto > auto-generatex/blog/types/...
	var post = types.Post{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Title:     msg.Title,
		Body:      msg.Body,
		CreatedAt: ctx.BlockHeight(),
	}

	// make function AppendPost at x/blog/keeper/post.go
	id := k.AppendPost(ctx, post) // func(value,data-type)

	return &types.MsgCreatePostResponse{Id: id}, nil
}
