package crudchain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "crudchain/api/crudchain/crudchain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ResourceAll",
					Use:       "list-resource",
					Short:     "List all resource",
				},
				{
					RpcMethod:      "Resource",
					Use:            "show-resource [id]",
					Short:          "Shows a resource by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ShowResourceByName",
					Use:            "show-resource-by-name [name]",
					Short:          "Query show-resource-by-name",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},

				{
					RpcMethod: "FilmAll",
					Use:       "list-film",
					Short:     "List all film",
				},
				{
					RpcMethod:      "Film",
					Use:            "show-film [id]",
					Short:          "Shows a film by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ShowFilmByGenre",
					Use:            "show-film-by-genre [gere]",
					Short:          "Query show-film-by-genre",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "gere"}},
				},

				{
					RpcMethod:      "ShowFilmGenre",
					Use:            "show-film-genre [genre]",
					Short:          "Query show-film-genre",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "genre"}},
				},

				{
					RpcMethod:      "ShowFilmName",
					Use:            "show-film-name [name]",
					Short:          "Query show-film-name",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateResource",
					Use:            "create-resource [name] [description]",
					Short:          "Create resource",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "UpdateResource",
					Use:            "update-resource [id] [name] [description]",
					Short:          "Update resource",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "DeleteResource",
					Use:            "delete-resource [id]",
					Short:          "Delete resource",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateFilm",
					Use:            "create-film [name] [description] [genre]",
					Short:          "Create film",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "genre"}},
				},
				{
					RpcMethod:      "UpdateFilm",
					Use:            "update-film [id] [name] [description] [genre]",
					Short:          "Update film",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "genre"}},
				},
				{
					RpcMethod:      "DeleteFilm",
					Use:            "delete-film [id]",
					Short:          "Delete film",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
