package ssh

import (
	"additive-apps/additive-lab/src/print"
	"additive-apps/additive-lab/src/util/file"
	"fmt"
	"os/exec"
)

var sshPaths = NewSSHPaths()

func Generate() {
	file.NotExists(sshPaths.PrivateKey, func() {
		command := fmt.Sprintf(`ssh-keygen -t rsa -b 2048 -f %s -N ""`, sshPaths.PrivateKey)
		_, err := exec.Command("bash", "-c", command).CombinedOutput()

		if err != nil {
			print.Fatal(err)
		}

		print.Create(sshPaths.PrivateKey)
	})
}

func GenerateConfigFile() {
	file.NotExists(sshPaths.Config, func() {
		content := `# Personal GitHub SSH key
Host github.com
    HostName github.com
    User git
    AddKeysToAgent yes
    IdentityFile ~/.ssh/id_rsa

# Additive GitHub SSH key
Host additive.github.com
    HostName github.com
    User git
    AddKeysToAgent yes
    IdentityFile ~/.ssh/additive
`

		file.Write(sshPaths.Config, content)
		print.Create(sshPaths.Config)
	})
}
