package gitcommand

import "os/exec"

func git(str ...string) *exec.Cmd {
	return exec.Command("git", str...)
}
