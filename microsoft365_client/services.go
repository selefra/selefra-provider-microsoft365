package microsoft365_client

import (
	"bytes"
	"context"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
)

// https://github.com/Azure/go-autorest/blob/3fb5326fea196cd5af02cf105ca246a0fba59021/autorest/azure/cli/token.go#L126
// NewAuthorizerFromCLIWithResource creates an Authorizer configured from Azure CLI 2.0 for local development scenarios.
func getTenantFromCLI() (string, error) {
	// This is the path that a developer can set to tell this class what the install path for Azure CLI is.
	const azureCLIPath = "AzureCLIPath"

	// The default install paths are used to find Azure CLI. This is for security, so that any path in the calling program's Path environment is not used to execute Azure CLI.
	azureCLIDefaultPathWindows := fmt.Sprintf("%s\\Microsoft SDKs\\Azure\\CLI2\\wbin; %s\\Microsoft SDKs\\Azure\\CLI2\\wbin", os.Getenv("ProgramFiles(x86)"), os.Getenv("ProgramFiles"))

	// Default path for non-Windows.
	const azureCLIDefaultPath = "/bin:/sbin:/usr/bin:/usr/local/bin"

	// Execute Azure CLI to get token
	var cliCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cliCmd = exec.Command(fmt.Sprintf("%s\\system32\\cmd.exe", os.Getenv("windir")))
		cliCmd.Env = os.Environ()
		cliCmd.Env = append(cliCmd.Env, fmt.Sprintf("PATH=%s;%s", os.Getenv(azureCLIPath), azureCLIDefaultPathWindows))
		cliCmd.Args = append(cliCmd.Args, "/c", "az")
	} else {
		cliCmd = exec.Command("az")
		cliCmd.Env = os.Environ()
		cliCmd.Env = append(cliCmd.Env, fmt.Sprintf("PATH=%s:%s", os.Getenv(azureCLIPath), azureCLIDefaultPath))
	}
	cliCmd.Args = append(cliCmd.Args, "account", "get-access-token", "--resource-type=ms-graph", "-o", "json")

	var stderr bytes.Buffer
	cliCmd.Stderr = &stderr

	output, err := cliCmd.Output()
	if err != nil {
		return "", fmt.Errorf("invoking Azure CLI failed with the following error: %v", err)
	}

	var tokenResponse struct {
		AccessToken string `json:"accessToken"`
		ExpiresOn   string `json:"expiresOn"`
		Tenant      string `json:"tenant"`
		TokenType   string `json:"tokenType"`
	}
	err = json.Unmarshal(output, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.Tenant, nil
}

func GetTenant(ctx context.Context, config *Config) (interface{}, error) {
	return config.TenantID, nil
}

func GetUserFromConfig(ctx context.Context, config *Config) string {
	return config.UserID
}

func GetUserID(ctx context.Context, config *Config) (interface{}, error) {
	userID := GetUserFromConfig(ctx, config)

	if userID != "" {
		return userID, nil
	}

	// If the user is not provided in the config,
	// get the current authenticated user from CLI
	tenantID, err := GetTenant(ctx, config)
	if err != nil {
		return nil, err
	}

	// Create client
	cred, err := azidentity.NewAzureCLICredential(
		&azidentity.AzureCLICredentialOptions{
			TenantID: tenantID.(string),
		},
	)
	if err != nil {
		return nil, err
	}

	auth, err := a.NewAzureIdentityAuthenticationProvider(cred)
	if err != nil {
		return nil, err
	}

	adapter, err := msgraphsdkgo.NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, err
	}
	client := msgraphsdkgo.NewGraphServiceClient(adapter)

	result, err := client.Me().Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	if result.GetId() != nil {
		return *result.GetId(), nil
	}

	return nil, nil
}

func GetGraphClient(ctx context.Context, config *Config) (*msgraphsdkgo.GraphServiceClient, *msgraphsdkgo.GraphRequestAdapter, error) {
	// Disable caching since it only saves ~.25ms and results in an SDK error
	// when running consecutive queries for the mail_message and my_mail_message
	// tables:
	// Error: rpc error: code = Internal desc = hydrate function listMicrosoft365MyMailMessages failed with panic runtime error: invalid memory address or nil pointer dereference (SQLSTATE HV000)
	// Have we already created and cached the session?
	/*
		sessionCacheKey := "GetGraphClient"
		if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
			return cachedData.(*msgraphsdkgo.GraphServiceClient), nil, nil
		}
	*/
	var cloudConfiguration cloud.Configuration

	switch config.Environment {
	case "AZURECHINACLOUD":
		cloudConfiguration = cloud.AzureChina
	case "AZUREUSGOVERNMENTCLOUD":
		cloudConfiguration = cloud.AzureGovernment
	default:
		cloudConfiguration = cloud.AzurePublic
	}

	var cred azcore.TokenCredential
	var err error
	if config.TenantID == "" { // CLI authentication
		cred, err = azidentity.NewAzureCLICredential(
			&azidentity.AzureCLICredentialOptions{},
		)
		if err != nil {
			return nil, nil, fmt.Errorf("error creating credentials: %w", err)
		}
	} else if config.TenantID != "" && config.ClientID != "" && config.ClientSecret != "" { // Client secret authentication
		cred, err = azidentity.NewClientSecretCredential(
			config.TenantID,
			config.ClientID,
			config.ClientSecret,
			&azidentity.ClientSecretCredentialOptions{
				ClientOptions: policy.ClientOptions{
					Cloud: cloudConfiguration,
				},
			},
		)
		if err != nil {
			return nil, nil, fmt.Errorf("error creating credentials: %w", err)
		}
	} else if config.TenantID != "" && config.ClientID != "" && config.ClientSecret != "" { // Client certificate authentication
		// Load certificate from given path
		loadFile, err := os.ReadFile(config.CertificatePath)
		if err != nil {
			return nil, nil, fmt.Errorf("error reading certificate from %s: %v", config.CertificatePath, err)
		}

		var certs []*x509.Certificate
		var key crypto.PrivateKey
		if config.CertificatePassword == "" {
			certs, key, err = azidentity.ParseCertificates(loadFile, nil)
		} else {
			certs, key, err = azidentity.ParseCertificates(loadFile, []byte(config.CertificatePassword))
		}

		if err != nil {
			return nil, nil, fmt.Errorf("error parsing certificate from %s: %v", config.CertificatePath, err)
		}

		cred, err = azidentity.NewClientCertificateCredential(
			config.TenantID,
			config.ClientID,
			certs,
			key,
			&azidentity.ClientCertificateCredentialOptions{
				ClientOptions: policy.ClientOptions{
					Cloud: cloudConfiguration,
				},
			},
		)
		if err != nil {
			return nil, nil, err
		}
	} else if config.EnableMSI { // Managed identity authentication
		cred, err = azidentity.NewManagedIdentityCredential(
			&azidentity.ManagedIdentityCredentialOptions{},
		)
		if err != nil {
			return nil, nil, err
		}
	}

	auth, err := a.NewAzureIdentityAuthenticationProvider(cred)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating authentication provider: %v", err)
	}

	adapter, err := msgraphsdkgo.NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating graph adapter: %v", err)
	}
	client := msgraphsdkgo.NewGraphServiceClient(adapter)

	return client, adapter, nil
}

// Int32 returns a pointer to the int32 value passed in.
func Int32(v int32) *int32 {
	return &v
}
