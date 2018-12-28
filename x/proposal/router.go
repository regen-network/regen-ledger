package proposal

import (
	types "github.com/cosmos/cosmos-sdk/types"
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

func (rtr *Router) CanHandle(action ProposalAction) bool {
	for _, route := range rtr.routes {
		if route.r == action.Route() {
            canHandle := route.h.CanHandle(action)
            if canHandle {
            	return true
			}
		}
	}
	return false
}

func (rtr *Router) Handle(ctx types.Context, action ProposalAction, voters []types.AccAddress) types.Result {
	for _, route := range rtr.routes {
		if route.r == action.Route() {
			return route.h.Handle(ctx, action, voters)
		}
	}
	return types.Result{
		Code:types.CodeUnknownRequest,
		Log:"can't find handler",
	}
}
