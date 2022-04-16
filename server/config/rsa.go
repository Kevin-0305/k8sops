package config

type Rsa struct {
	Dir               string `yaml:"dir"`
	RsaPrivateKeyFile string `yaml:"RsaPrivateKeyFile"`
	RsaPublicKeyFile  string `yaml:"RsaPublicKeyFile"`
}
