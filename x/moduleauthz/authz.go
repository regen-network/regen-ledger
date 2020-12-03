package moduleauthz

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
)

type Permission interface {
	Allow(signer sdk.AccAddress, request sdk.MsgRequest) bool
}

type AdminPermissionMiddleware struct {
	permissions map[string]map[string]Permission
}

func (apm *AdminPermissionMiddleware) SetPermission(methodName string, moduleId module.ModuleID, permission Permission) {
	apm.permissions[methodName][moduleId.Address().String()] = permission
}

func (apm *AdminPermissionMiddleware) AsMiddleware() server.AuthorizationMiddleware {
	return func(_ sdk.Context, methodName string, req sdk.MsgRequest, signer sdk.AccAddress) bool {
		perm, found := apm.permissions[methodName][signer.String()]
		if !found {
			return false
		}

		return perm.Allow(signer, req)
	}
}
