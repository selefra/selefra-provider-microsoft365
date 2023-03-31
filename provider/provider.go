package provider

import (
	"context"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "microsoft365",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var microsoft365Config microsoft365_client.Configs

				err := config.Unmarshal(&microsoft365Config.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(microsoft365Config.Providers) == 0 {
					microsoft365Config.Providers = append(microsoft365Config.Providers, microsoft365_client.Config{})
				}

				if microsoft365Config.Providers[0].TenantID == "" {
					microsoft365Config.Providers[0].TenantID = os.Getenv("MICROSOFT365_TENANT_ID")
				}

				if microsoft365Config.Providers[0].TenantID == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing TenantID in configuration")
				}

				if microsoft365Config.Providers[0].ClientID == "" {
					microsoft365Config.Providers[0].ClientID = os.Getenv("MICROSOFT365_CLIENT_ID")
				}

				if microsoft365Config.Providers[0].ClientID == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing ClientID in configuration")
				}

				if microsoft365Config.Providers[0].ClientSecret == "" {
					microsoft365Config.Providers[0].ClientSecret = os.Getenv("MICROSOFT365_CLIENT_SECRET")
				}

				if microsoft365Config.Providers[0].ClientSecret == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing ClientSecret in configuration")
				}

				if microsoft365Config.Providers[0].UserID == "" {
					microsoft365Config.Providers[0].UserID = os.Getenv("MICROSOFT365_USER_ID")
				}

				if microsoft365Config.Providers[0].UserID == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing userId in configuration")
				}

				clients, err := microsoft365_client.NewClients(microsoft365Config)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# tenant_id: "<YOUR_TENANT_ID>"
# client_id: "<YOUR_CLIENT_ID>"
# client_secret: "<YOUR_CLIENT_SECRET>"
# user_id: "<YOUR_USER_ID>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var client_config microsoft365_client.Configs
				err := config.Unmarshal(&client_config.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
