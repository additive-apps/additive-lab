package ssh

import (
	"additive-apps/additive-lab/src/util"
	"additive-apps/additive-lab/src/util/constants"
	"additive-apps/additive-lab/src/util/dirpath"
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
	basePath := dirpath.Join(util.HomePath(), ".ssh")

	return SSHPaths{
		Base:       basePath,
		Config:     filepath.Join(basePath, "config"),
		KeyName:    constants.SSH_KEY_NAME,
		PrivateKey: filepath.Join(basePath, constants.SSH_KEY_NAME),
		PublicKey:  filepath.Join(basePath, constants.SSH_KEY_NAME, ".pub"),
	}
}
