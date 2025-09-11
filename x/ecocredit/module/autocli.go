package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
)

func (am Module) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: ecocreditv1beta1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateClass",
					Use:       "create-class [issuers] [credit-type-abbrev] [metadata]",
					Short:     "Creates a new credit class with transaction author (--from) as admin",
					Long: `Creates a new credit class with transaction author (--from) as admin.

The transaction author must have permission to create a new credit class by
being on the list of allowed class creators. This is a governance parameter.

They must also pay a fee (separate from the transaction fee) to create a credit
class. The list of accepted fees is defined by the credit_class_fee parameter.

Parameters:
- issuers:             comma separated (no spaces) list of issuer account addresses
- credit-type-abbrev:  the abbreviation of a credit type
- metadata:            arbitrary data attached to the credit class info`,
					Example: `regen tx ecocredit create-class regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw C regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf --class-fee 20000000uregen`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "issuers"},
						{ProtoField: "credit_type_abbrev"},
						{ProtoField: "metadata"},
					},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"fee": {
							Name:  "class-fee",
							Usage: "the fee that the class creator will pay to create the credit class (e.g. \"20regen\")",
						},
					},
				},
				{
					RpcMethod: "CreateProject",
					Use:       "create-project [class-id] [metadata] [jurisdiction]",
					Short:     "Create a new project within a credit class",
					Long: `Create a new project within a credit class.

Parameters:
- class-id:     the ID of the credit class
- metadata:     any arbitrary metadata to attach to the project
- jurisdiction: the jurisdiction of the project`,
					Example: `regen tx ecocredit create-project C01 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf "US-WA 98225" --reference-id VCS-001`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
						{ProtoField: "metadata"},
						{ProtoField: "jurisdiction"},
					},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"reference_id": {
							Name:  "reference-id",
							Usage: "a reference ID for the project",
						},
						"fee": {
							Name:  "project-fee",
							Usage: "the fee that the project creator will pay to create the project",
						},
					},
				},
				{
					RpcMethod: "CreateUnregisteredProject",
					Use:       "create-unregistered-project [metadata] [jurisdiction]",
					Short:     "Create a new project without registering it under a credit class",
					Long: `Create a new project without registering it under a credit class.

Parameters:
- metadata:     any arbitrary metadata to attach to the project
- jurisdiction: the jurisdiction of the project`,
					Example: `regen tx ecocredit create-unregistered-project regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf "US-WA 98225" --reference-id VCS-001`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "metadata"},
						{ProtoField: "jurisdiction"},
					},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"reference_id": {
							Name:  "reference-id",
							Usage: "a reference ID for the project",
						},
						"fee": {
							Name:  "project-fee",
							Usage: "the fee that the project creator will pay to create the project",
						},
					},
				},
				{
					RpcMethod: "CreateOrUpdateApplication",
					Use:       "create-or-update-application [project-id] [class-id] [metadata]",
					Short:     "Create or update a project credit class application",
					Long: `Create or update a project credit class application.

Parameters:
- project-id: the identifier of the project applying to the credit class
- class-id:   the identifier of the credit class
- metadata:   optional metadata relevant to the application`,
					Example: `regen tx ecocredit create-or-update-application P01-001 C01 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "project_id"},
						{ProtoField: "class_id"},
						{ProtoField: "metadata"},
					},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"withdraw": {
							Name:  "withdraw",
							Usage: "withdraw the application rather than update it",
						},
					},
				},
				{
					RpcMethod: "UpdateProjectEnrollment",
					Use:       "update-project-enrollment [project-id] [class-id] [new-status] [metadata]",
					Short:     "Update a project enrollment status",
					Long: `Update a project enrollment status.

Parameters:
- project-id: the identifier of the project
- class-id:   the identifier of the credit class
- new-status: the new enrollment status (accepted, changes-requested, rejected, terminated)
- metadata:   optional metadata explaining the decision`,
					Example: `regen tx ecocredit update-project-enrollment P01-001 C01 accepted "Project meets all requirements"`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "project_id"},
						{ProtoField: "class_id"},
						{ProtoField: "new_status"},
						{ProtoField: "metadata"},
					},
				},
				{
					RpcMethod: "CreateBatch",
					Use:       "create-batch [project-id] [metadata]",
					Short:     "Issues a new credit batch",
					Long: `Issues a new credit batch.

Note: This command requires complex JSON for issuance data. Consider using the existing manual CLI for full functionality.

Parameters:
- project-id: the unique identifier of the project
- metadata:   arbitrary metadata to attach to the credit batch`,
					Example: `regen tx ecocredit create-batch P01-001 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf`,
					Skip:    true,
				},
				{
					RpcMethod: "MintBatchCredits",
					Use:       "mint-batch-credits [batch-denom]",
					Short:     "Mint additional credits to an open batch",
					Long: `Mint additional credits to an open credit batch.

Parameters:
- batch-denom: the unique identifier of the credit batch`,
					Example: `regen tx ecocredit mint-batch-credits C01-001-20200101-20210101-001`,
					Skip:    true,
				},
				{
					RpcMethod: "SealBatch",
					Use:       "seal-batch [batch-denom]",
					Short:     "Seal an open credit batch",
					Long: `Seal an open credit batch to prevent further minting.

Parameters:
- batch-denom: the unique identifier of the credit batch`,
					Example: `regen tx ecocredit seal-batch C01-001-20200101-20210101-001`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "batch_denom"},
					},
				},
				{
					RpcMethod: "Send",
					Use:       "send [recipient]",
					Short:     "Send credits to another account",
					Long: `Send credits from the transaction author (--from) to the recipient.

Note: This command requires complex JSON for credits data. Use the existing manual CLI for full functionality.

Parameters:
- recipient: the recipient account address`,
					Example: `regen tx ecocredit send regen18xvpj53vaupyfejpws5sktv5lnas5xj2phm3cf`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "recipient"},
					},
					Skip: true, // Complex credits structure requires custom handling
				},
				{
					RpcMethod: "Retire",
					Use:       "retire [jurisdiction]",
					Short:     "Retire credits",
					Long: `Retire a specified amount of credits.

Note: This command requires complex JSON for credits data. Use the existing manual CLI for full functionality.

Parameters:
- jurisdiction: the jurisdiction in which credits will be retired`,
					Example: `regen tx ecocredit retire "US-WA 98225"`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "jurisdiction"},
					},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"reason": {
							Name:  "reason",
							Usage: "the reason for retiring the credits",
						},
					},
					Skip: true, // Complex credits structure requires custom handling
				},
				{
					RpcMethod: "Cancel",
					Use:       "cancel [reason]",
					Short:     "Cancel credits",
					Long: `Cancel a specified amount of credits.

Note: This command requires complex JSON for credits data. Use the existing manual CLI for full functionality.

Parameters:
- reason: the reason for cancelling credits`,
					Example: `regen tx ecocredit cancel "transferring credits to another registry"`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "reason"},
					},
					Skip: true, // Complex credits structure requires custom handling
				},
				{
					RpcMethod: "UpdateClassAdmin",
					Use:       "update-class-admin [class-id] [new-admin]",
					Short:     "Update the admin for a specific credit class",
					Long: `Update the admin for a specific credit class.

The '--from' flag must equal the current credit class admin.

WARNING: Updating the admin replaces the current admin. Be sure the new admin account address is entered correctly.

Parameters:
- class-id:  the ID of the credit class to update
- new-admin: the new admin account address`,
					Example: `regen tx ecocredit update-class-admin C01 regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
						{ProtoField: "new_admin"},
					},
				},
				{
					RpcMethod: "UpdateClassIssuers",
					Use:       "update-class-issuers [class-id]",
					Short:     "Update the list of issuers for a specific credit class",
					Long: `Update the list of issuers for a specific credit class.

The '--from' flag must equal the current credit class admin.

Parameters:
- class-id: the ID of the credit class to update`,
					Example: `regen tx ecocredit update-class-issuers C01 --add-issuers addr1,addr2,addr3 --remove-issuers addr3,addr4`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
					},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"add_issuers": {
							Name:  "add-issuers",
							Usage: "comma separated list of addresses to add",
						},
						"remove_issuers": {
							Name:  "remove-issuers",
							Usage: "comma separated list of addresses to remove",
						},
					},
				},
				{
					RpcMethod: "UpdateClassMetadata",
					Use:       "update-class-metadata [class-id] [new-metadata]",
					Short:     "Update the metadata for a specific credit class",
					Long: `Update the metadata for a specific credit class.

The '--from' flag must equal the credit class admin.

Parameters:
- class-id:     the class ID to update
- new-metadata: new metadata to attach to the credit class`,
					Example: `regen tx ecocredit update-class-metadata C01 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
						{ProtoField: "new_metadata"},
					},
				},
				{
					RpcMethod: "UpdateProjectAdmin",
					Use:       "update-project-admin [project-id] [new-admin]",
					Short:     "Update the project admin address",
					Long: `Update the project admin to the provided new admin address.

The '--from' flag must equal the current project admin.

WARNING: Updating the admin replaces the current admin. Be sure the new admin account address is entered correctly.

Parameters:
- project-id: the ID of the project to update
- new-admin:  the new admin account address`,
					Example: `regen tx ecocredit update-project-admin P01-001 regen1ynugxwpp4lfpy0epvfqwqkpuzkz62htnex3op`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "project_id"},
						{ProtoField: "new_admin"},
					},
				},
				{
					RpcMethod: "UpdateProjectMetadata",
					Use:       "update-project-metadata [project-id] [new-metadata]",
					Short:     "Update the project metadata",
					Long: `Update the project metadata, overwriting the project's current metadata.

Parameters:
- project-id:   the ID of the project to update
- new-metadata: new metadata to attach to the project`,
					Example: `regen tx ecocredit update-project-metadata P01-001 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "project_id"},
						{ProtoField: "new_metadata"},
					},
				},
				{
					RpcMethod: "UpdateBatchMetadata",
					Use:       "update-batch-metadata [batch-denom] [new-metadata]",
					Short:     "Update the metadata for a specific credit batch",
					Long: `Update the metadata for a specific credit batch.

The '--from' flag must equal the credit batch issuer.

Parameters:
- batch-denom:  the batch denom of the credit batch to update
- new-metadata: new metadata to attach to the credit batch`,
					Example: `regen tx ecocredit update-batch-metadata C01-001-20200101-20210101-001 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "batch_denom"},
						{ProtoField: "new_metadata"},
					},
				},
				{
					RpcMethod: "Bridge",
					Use:       "bridge [target] [recipient]",
					Short:     "Bridge credits to another chain",
					Long: `Bridge credits to another chain.

Note: This command requires complex JSON for credits data. Use the existing manual CLI for full functionality.

Parameters:
- target:    the target chain (e.g. "polygon")
- recipient: the address of the recipient on the other chain`,
					Example: `regen tx ecocredit bridge polygon 0x0000000000000000000000000000000000000001`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "target"},
						{ProtoField: "recipient"},
					},
					Skip: true, // Complex credits structure requires custom handling
				},
				{
					RpcMethod: "BridgeReceive",
					Use:       "bridge-receive [class-id]",
					Short:     "Process credits being sent from another chain",
					Long: `Process credits being sent from another chain.

Note: This command has complex nested structures. Consider using the existing manual CLI for full functionality.

Parameters:
- class-id: the unique identifier of the credit class`,
					Example: `regen tx ecocredit bridge-receive C01`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
					},
					Skip: true, // Complex nested structures require custom handling
				},
				{
					RpcMethod: "BurnRegen",
					Use:       "burn-regen [amount] [reason]",
					Short:     "Burn REGEN tokens",
					Long: `Burn REGEN tokens to account for platform fees.

Parameters:
- amount: the integer amount of uregen tokens to burn
- reason: the reason for burning REGEN tokens`,
					Example: `regen tx ecocredit burn-regen 1000000 "platform fee payment"`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "amount"},
						{ProtoField: "reason"},
					},
				},
				// Governance commands
				{
					RpcMethod: "AddCreditType",
					Use:       "add-credit-type",
					Short:     "Add a new credit type (governance)",
					Long:      "Add a new credit type to the network. This is a governance proposal command.",
					Skip:      true, // Governance command, typically handled separately
				},
				{
					RpcMethod: "SetClassCreatorAllowlist",
					Use:       "set-class-creator-allowlist [enabled]",
					Short:     "Enable or disable the class creator allowlist (governance)",
					Long:      "Enable or disable the class creator allowlist. This is a governance proposal command.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "enabled"},
					},
					Skip: true, // Governance command
				},
				{
					RpcMethod: "AddClassCreator",
					Use:       "add-class-creator [creator]",
					Short:     "Add an address to the class creation allowlist (governance)",
					Long:      "Add an address to the class creation allowlist. This is a governance proposal command.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "creator"},
					},
					Skip: true, // Governance command
				},
				{
					RpcMethod: "RemoveClassCreator",
					Use:       "remove-class-creator [creator]",
					Short:     "Remove an address from the class creation allowlist (governance)",
					Long:      "Remove an address from the class creation allowlist. This is a governance proposal command.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "creator"},
					},
					Skip: true, // Governance command
				},
				{
					RpcMethod: "UpdateClassFee",
					Use:       "update-class-fee",
					Short:     "Update the credit class creation fee (governance)",
					Long:      "Update the credit class creation fee. This is a governance proposal command.",
					Skip:      true, // Governance command
				},
				{
					RpcMethod: "UpdateProjectFee",
					Use:       "update-project-fee",
					Short:     "Update the project creation fee (governance)",
					Long:      "Update the project creation fee. This is a governance proposal command.",
					Skip:      true, // Governance command
				},
				{
					RpcMethod: "AddAllowedBridgeChain",
					Use:       "add-allowed-bridge-chain [chain-name]",
					Short:     "Add a chain to the allowed bridge chains (governance)",
					Long:      "Add a chain to the allowed bridge chains. This is a governance proposal command.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "chain_name"},
					},
					Skip: true, // Governance command
				},
				{
					RpcMethod: "RemoveAllowedBridgeChain",
					Use:       "remove-allowed-bridge-chain [chain-name]",
					Short:     "Remove a chain from the allowed bridge chains (governance)",
					Long:      "Remove a chain from the allowed bridge chains. This is a governance proposal command.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "chain_name"},
					},
					Skip: true, // Governance command
				},
			},
		},
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: ecocreditv1beta1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Classes",
					Use:       "classes",
					Short:     "List all credit classes",
					Long:      "List all credit classes with optional pagination flags.",
					Example: `regen q ecocredit classes
regen q ecocredit classes --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "ClassesByAdmin",
					Use:       "classes [admin]",
					Short:     "List all credit classes by admin",
					Long:      "List all credit classes with pagination support.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "admin"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "Class",
					Use:       "class [class-id]",
					Short:     "Get credit class info by class ID",
					Long:      "Get credit class information by providing the class ID.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
					},
				},
				{
					RpcMethod: "ClassIssuers",
					Use:       "class-issuers [class-id]",
					Short:     "Retrieve issuer addresses for a credit class",
					Long:      "Retrieve issuer addresses for a credit class with optional pagination flags.",
					Example: `regen q ecocredit class-issuers C01
regen q ecocredit class-issuers C01 --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "Projects",
					Use:       "projects",
					Short:     "List all projects",
					Long:      "List all projects with optional pagination flags.",
					Example: `regen q ecocredit projects
regen q ecocredit projects --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "ProjectsByClass",
					Use:       "projects-by-class [class-id]",
					Short:     "List projects by credit class",
					Long:      "List projects by credit class with optional pagination flags.",
					Example: `regen q ecocredit projects-by-class C01
regen q ecocredit projects-by-class C01 --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "ProjectsByReferenceId",
					Use:       "projects-by-reference-id [reference-id]",
					Short:     "List all projects by reference ID",
					Long:      "List all projects by reference ID with optional pagination flags.",
					Example: `regen q ecocredit projects-by-reference-id VCS-001
regen q ecocredit projects-by-reference-id VCS-001 --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "reference_id"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "ProjectsByAdmin",
					Use:       "projects-by-admin [admin]",
					Short:     "List projects by admin",
					Long:      "List projects by admin with optional pagination flags.",
					Example: `regen q ecocredit projects-by-admin regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw
regen q ecocredit projects-by-admin regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "admin"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "Project",
					Use:       "project [project-id]",
					Short:     "Retrieve project information",
					Long:      "Retrieve project information.",
					Example:   `regen q ecocredit project C01-001`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "project_id"},
					},
				},
				{
					RpcMethod: "Batches",
					Use:       "batches",
					Short:     "List all credit batches",
					Long:      "List all credit batches with optional pagination flags.",
					Example: `regen q ecocredit batches
regen q ecocredit batches --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "BatchesByIssuer",
					Use:       "batches-by-issuer [issuer]",
					Short:     "List all credit batches by issuer",
					Long:      "List all credit batches by issuer with optional pagination flags.",
					Example: `regen q ecocredit batches-by-issuer regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38
regen q ecocredit batches-by-issuer regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38 --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "issuer"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "BatchesByClass",
					Use:       "batches-by-class [class-id]",
					Short:     "List all credit batches by credit class",
					Long:      "List all credit batches by credit class with optional pagination flags.",
					Example: `regen q ecocredit batches-by-class C01
regen q ecocredit batches-by-class C01 --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "class_id"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "BatchesByProject",
					Use:       "batches-by-project [project-id]",
					Short:     "List all credit batches by project",
					Long:      "List all credit batches by project with optional pagination flags.",
					Example: `regen q ecocredit batches-by-project C01-001
regen q ecocredit batches-by-project C01-001 --limit 10 --offset 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "project_id"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "Batch",
					Use:       "batch [batch-denom]",
					Short:     "Retrieve credit batch information",
					Long:      "Retrieve credit batch information.",
					Example:   "regen q ecocredit batch C01-001-20200101-20210101-001",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "batch_denom"},
					},
				},
				{
					RpcMethod: "Balance",
					Use:       "batch-balance [batch-denom] [account]",
					Short:     "Retrieve the batch balance of an account",
					Long:      "Retrieve the batch balance of an account.",
					Example:   "regen q ecocredit batch-balance C01-001-20200101-20210101-001 regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
						{ProtoField: "batch_denom"},
					},
				},
				{
					RpcMethod: "Balances",
					Use:       "balances [address]",
					Short:     "Get all credit balances for an address",
					Long:      "Get all credit balances for a specific address with pagination support.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "BalancesByBatch",
					Use:       "balances-by-batch [batch-denom]",
					Short:     "Retrieve all ecocredit balances for a given credit batch",
					Long:      "Retrieve all ecocredit balances for a given credit batch with optional pagination flags.",
					Example: `regen q ecocredit balances-by-batch C01-001-20200101-20210101-001
regen q ecocredit balances-by-batch C01-001-20200101-20210101-001 --limit 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "batch_denom"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "AllBalances",
					Use:       "all-balances",
					Short:     "Retrieve all ecocredit balances",
					Long:      "Retrieve all ecocredit balances across all addresses and credit batches with optional pagination flags.",
					Example: `regen q ecocredit all-balances
regen q ecocredit all-balances --limit 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "Supply",
					Use:       "batch-supply [batch-denom]",
					Short:     "Retrieve the supply of a credit batch",
					Long:      "Retrieve the supply of a credit batch.",
					Example:   "regen q ecocredit batch-supply C01-001-20200101-20210101-001",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "batch_denom"},
					},
				},
				{
					RpcMethod: "CreditTypes",
					Use:       "credit-types",
					Short:     "List all credit types",
					Long:      "List all credit types.",
					Example:   "regen q ecocredit credit-types",
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "List the current ecocredit module parameters",
					Long:      "List the current ecocredit module parameters.",
					Example:   "regen q ecocredit params",
				},
				{
					RpcMethod: "CreditType",
					Use:       "credit-type [abbreviation]",
					Short:     "Retrieve credit type information",
					Long:      "Retrieve credit type information.",
					Example:   "regen q ecocredit credit-type C",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "abbreviation"},
					},
				},
				{
					RpcMethod: "ClassCreatorAllowlist",
					Use:       "class-creator-allowlist",
					Short:     "Retrieve the class creator allowlist enabled setting",
					Long:      "Retrieve the class creator allowlist enabled setting",
					Example:   "regen q ecocredit class-creator-allowlist",
				},
				{
					RpcMethod: "AllowedClassCreators",
					Use:       "allowed-class-creators",
					Short:     "Retrieve the allowed credit class creators",
					Long:      "Retrieve the list of allowed credit class creators with optional pagination flags.",
					Example: `regen q ecocredit allowed-class-creators
regen q ecocredit allowed-class-creators --limit 10`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "ClassFee",
					Use:       "class-fee",
					Short:     "Retrieve the credit class creation fee",
					Long:      "Retrieve the credit class creation fee",
					Example:   "regen q ecocredit class-fee",
				},
				{
					RpcMethod: "AllowedBridgeChains",
					Use:       "allowed-bridge-chains",
					Short:     "Retrieve the list of allowed bridge chains",
					Long:      "Retrieve the list of chains that are allowed to be used in bridge operations",
					Example:   "regen q ecocredit allowed-bridge-chains",
				},
				// {
				// 	RpcMethod: "ProjectEnrollment",
				// 	Use:       "credit-type [abbreviation]",
				// 	Short:     "Retrieve credit type information",
				// 	Long:      "Retrieve credit type information.",
				// 	Example:   "regen q ecocredit credit-type C",
				// 	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				// 		{ProtoField: "project_id"},
				// 		{ProtoField: "class_id"},
				// 	},
				// },
				// {
				// 	RpcMethod: "ProjectEnrollments",
				// 	Use:       "credit-type [abbreviation]",
				// 	Short:     "Retrieve credit type information",
				// 	Long:      "Retrieve credit type information.",
				// 	Example:   "regen q ecocredit credit-type C",
				// 	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				// 		{ProtoField: "project_id"},
				// 		{ProtoField: "pagination", Optional: true},
				// 	},
				// },
			},
		},
	}
}
