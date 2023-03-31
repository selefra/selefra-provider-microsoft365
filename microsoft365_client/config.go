package microsoft365_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	TenantID            string `yaml:"tenant_id,omitempty" mapstructure:"tenant_id"`
	ClientID            string `yaml:"client_id,omitempty" mapstructure:"client_id"`
	ClientSecret        string `yaml:"client_secret,omitempty" mapstructure:"client_secret"`
	CertificatePath     string `yaml:"certificate_path,omitempty" mapstructure:"certificate_path"`
	CertificatePassword string `yaml:"certificate_password,omitempty" mapstructure:"certificate_password"`
	EnableMSI           bool   `yaml:"enable_msi,omitempty" mapstructure:"enable_msi"`
	MSIEndpoint         string `yaml:"msi_endpoint,omitempty" mapstructure:"msi_endpoint"`
	Environment         string `yaml:"environment,omitempty" mapstructure:"environment"`
	UserID              string `yaml:"user_id,omitempty" mapstructure:"user_id"`
}
