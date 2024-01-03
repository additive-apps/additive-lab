package ssh

import (
	"additive-apps/additive-lab/src/util"
	"path/filepath"
)

type SSHPaths struct {
	Base       string
	KeyName    string
	PrivateKey string
	PublicKey  string
	Config     string
}

func NewSSHPaths() SSHPaths {
	keyName := "additive"
	base := filepath.Join(util.HomePath(), ".ssh/")

	return SSHPaths{
		KeyName:    keyName,
		Base:       base,
		PrivateKey: filepath.Join(base, keyName),
		Config:     filepath.Join(base, "config"),
		PublicKey:  filepath.Join(base, keyName, ".pub"),
	}
}
