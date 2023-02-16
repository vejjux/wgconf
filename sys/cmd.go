package sys

import (
	"os/exec"
	"strings"
	"wgconf/exx"
)

func Cmd(in string, name string, args ...string) (out string, err error) {
	defer exx.H(&err)

	cmd := exec.Command(name, args...)

	stdIn := exx.CA(cmd.StdinPipe())
	if len(in) > 0 {
		exx.CA(stdIn.Write([]byte(in)))
	}
	exx.C(stdIn.Close())

	out = strings.Trim(
		string(exx.CA(cmd.Output())),
		"\n\t ",
	)

	return
}
