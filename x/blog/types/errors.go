package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/blog module sentinel errors
var (
	ErrSample     = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrCommentOld = sdkerrors.Register(ModuleName, 1300, "Sorry, The comment is old now, try another comment id")
	ErrID         = sdkerrors.Register(ModuleName, 1400, "put right error ID")
)
