package wg

import (
	"fmt"
	"os"
	"wgconf/exx"

	"gopkg.in/yaml.v3"
)

const confFile = "conf.yml"

type Peer struct {
	Name       string `yaml:"-"`
	IP         string `yaml:"IP"`
	PublicKey  string `yaml:"PublicKey"`
	PrivateKey string `yaml:"PrivateKey"`
}

type Conf struct {
	DNS        string          `yaml:"DNS"`
	AllowedIPs string          `yaml:"AllowedIPs"`
	Endpoint   string          `yaml:"Endpoint"`
	ServerKey  string          `yaml:"ServerKey"`
	RunUp      []string        `yaml:"RunUp"`
	RunDown    []string        `yaml:"RunDown"`
	Peers      map[string]Peer `yaml:"Peers"`
}

func (c *Conf) Write() (err error) {
	defer exx.H(&err)

	bytes := exx.CA(yaml.Marshal(c))
	exx.C(os.WriteFile(confFile, bytes, 0766))

	return
}

func (c *Conf) AddPeer(p Peer) (err error) {
	defer exx.H(&err)

	if _, ok := c.Peers[p.Name]; ok {
		exx.C(fmt.Errorf("peer already exists: %s", p.Name))
	}
	c.Peers[p.Name] = p

	return
}

func (c *Conf) RemPeer(name string) (err error) {
	defer exx.H(&err)

	if _, ok := c.Peers[name]; ok {
		delete(c.Peers, name)
	} else {
		exx.C(fmt.Errorf("peer not found: %s", name))
	}

	return
}

func (c *Conf) FindPeer(name string) (peer Peer, err error) {
	defer exx.H(&err)

	if p, ok := c.Peers[name]; ok {
		peer = p
	} else {
		exx.C(fmt.Errorf("peer not found: %s", name))
	}

	return
}

func ReadConf() (cfg *Conf, err error) {
	defer exx.H(&err)

	cfg = new(Conf)
	bytes := exx.CA(os.ReadFile(confFile))
	exx.C(yaml.Unmarshal(bytes, cfg))

	return
}
