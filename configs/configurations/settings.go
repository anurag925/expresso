package configurations

type settings struct {
	Aws struct {
		Region   string `yaml:"region"`
		S3Bucket string `yaml:"s3_bucket"`
	} `yaml:"aws"`
}
