package wg

import (
	"bytes"
	"fmt"
	"wgconf/exx"
	"wgconf/sys"

	"gopkg.in/ini.v1"
)

func AddPeer(cfg *Conf, name, ip string) (err error) {
	defer exx.H(&err)

	userKey := exx.CA(sys.Cmd("", "wg", "genkey"))
	userPub := exx.CA(sys.Cmd(userKey, "wg", "pubkey"))
	pub := exx.CA(sys.Cmd(cfg.ServerKey, "wg", "pubkey"))

	exx.C(cfg.AddPeer(Peer{
		Name:       name,
		IP:         ip,
		PublicKey:  userPub,
		PrivateKey: userKey,
	}))
	exx.C(cfg.Write())

	fmt.Println(string(exx.CA(
		outputPeerCfg(cfg, ip, userKey, pub),
	)))

	return
}

func outputPeerCfg(cfg *Conf, addr, key, pub string) (out []byte, err error) {
	defer exx.H(&err)

	u := ini.Empty()

	iface := u.Section("Interface")
	iface.Key("PrivateKey").SetValue(key)
	iface.Key("Address").SetValue(addr)
	iface.Key("DNS").SetValue(cfg.DNS)

	peer := u.Section("Peer")
	peer.Key("PublicKey").SetValue(pub)
	peer.Key("AllowedIPs").SetValue(cfg.AllowedIPs)
	peer.Key("Endpoint").SetValue(cfg.Endpoint)

	w := bytes.NewBuffer([]byte{})
	exx.CA(u.WriteTo(w))
	out = w.Bytes()

	return
}
