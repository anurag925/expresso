package configurations

type secrets struct {
	Aws struct {
		SecretKeyID     string `yaml:"secret_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
	} `yaml:"aws"`
}
