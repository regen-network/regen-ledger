package proposal

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"regexp"
)

type route struct {
	r string
	h ProposalHandler
}

type Router struct {
	// TODO change this to a map
	routes []route
}

func NewRouter() *Router {
	return &Router{
		routes: make([]route, 0),
	}
}

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

func (rtr *Router) AddRoute(r string, h ProposalHandler) *Router {
	if !isAlphaNumeric(r) {
		panic("route expressions can only contain alphanumeric characters")
	}
	rtr.routes = append(rtr.routes, route{r, h})

	return rtr
}

func (rtr *Router) CheckProposal(ctx sdk.Context, action ProposalAction) (bool, sdk.Result) {
	for _, route := range rtr.routes {
		if route.r == action.Route() {
			canHandle, res := route.h.CheckProposal(ctx, action)
			if canHandle {
				return true, res
			}
		}
	}
	return false, sdk.ErrUnknownRequest("can't handle proposal").Result()
}

func (rtr *Router) HandleProposal(ctx sdk.Context, action ProposalAction, voters []sdk.AccAddress) sdk.Result {
	for _, route := range rtr.routes {
		if route.r == action.Route() {
			return route.h.HandleProposal(ctx, action, voters)
		}
	}
	return sdk.Result{
		Code: sdk.CodeUnknownRequest,
		Log:  "can't find handler",
	}
}
