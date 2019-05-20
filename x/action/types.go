package action

import cosmos "github.com/cosmos/cosmos-sdk/types"

type Action interface {
	cosmos.Msg
	RequiredCapabilities() []Capability
}

type Capability interface {
	// Every capability should be have a system wide unique ID that includes
	// both the type of capability and any params associated with it
	CapabilityID() string
	// Whether the specified action is allowed by this capability
	Accept(action Action) bool
}

type Keeper interface {
	// Store capabilities under the key actor-id/capability-id
	// Grant stores a root flag, and delegate
	GrantRootCapability(ctx cosmos.Context, actor cosmos.AccAddress, capability Capability)
	RevokeRootCapability(ctx cosmos.Context, actor cosmos.AccAddress, capability Capability)
	Delegate(ctx cosmos.Context, grantor cosmos.AccAddress, actor cosmos.AccAddress, capability Capability) bool
	Undelegate(ctx cosmos.Context, grantor cosmos.AccAddress, actor cosmos.AccAddress, capability Capability)
	HasCapability(ctx cosmos.Context, actor cosmos.AccAddress, capability Capability) bool
}

type Dispatcher interface {
	DispatchAction(ctx cosmos.Context, actor cosmos.AccAddress, action Action) cosmos.Result
}
