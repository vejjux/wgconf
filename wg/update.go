package wg

import (
	"fmt"
	"strings"
	"wgconf/exx"
	"wgconf/sys"
)

func run(peer Peer, lines []string) (err error) {
	defer exx.H(&err)

	for _, line := range lines {
		if len(line) > 0 {
			line = strings.ReplaceAll(line, "{PeerPub}", peer.PublicKey)
			line = strings.ReplaceAll(line, "{IP}", peer.IP)
			line := strings.Split(line, " ")
			if len(line) == 1 {
				fmt.Println(exx.CA(sys.Cmd("", line[0])))
			} else {
				fmt.Println(exx.CA(sys.Cmd("", line[0], line[1:]...)))
			}
		}
	}

	return
}

func RunUp(cfg *Conf, name string) (err error) {
	defer exx.H(&err)

	peer := exx.CA(cfg.FindPeer(name))
	run(peer, cfg.RunUp)

	return
}

func RunDown(cfg *Conf, name string) (err error) {
	defer exx.H(&err)

	peer := exx.CA(cfg.FindPeer(name))
	run(peer, cfg.RunDown)

	return
}
