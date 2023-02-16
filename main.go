package main

import (
	"flag"
	"fmt"
	"os"
	"wgconf/exx"
	"wgconf/wg"
)

func main() {
	defer exx.HL()

	argName := flag.String("name", "", "Unique name for the new peer")
	argIP := flag.String("ip", "", "Ip address for the new peer")
	argAdd := flag.Bool("add", false, "Add a peer to config with new keys")
	argRem := flag.Bool("rem", false, "Remove a peer from config")
	argList := flag.Bool("list", false, "List peers from config file")
	argUp := flag.Bool("up", false, "Run up script gor given peer")
	argDown := flag.Bool("down", false, "Run down script for given peer")
	flag.Bool("update-all", false, "Tries to run add script for all peers in config")

	flag.Parse()
	cfg := exx.CA(wg.ReadConf())

	if *argAdd {
		exx.CI(*argName == "" || *argIP == "", "usage: wgconf -add -name PeerName -ip PeerIP")
		exx.C(wg.AddPeer(cfg, *argName, *argIP))
		os.Exit(0)
	}

	if *argRem {
		exx.CI(*argName == "", "usage: wgconf -rem -name PeerName")
		exx.C(cfg.RemPeer(*argName))
		exx.C(cfg.Write())
		os.Exit(0)
	}

	if *argList {
		for name, peer := range cfg.Peers {
			fmt.Printf("%s: %s\n", name, peer.IP)
		}
		os.Exit(0)
	}

	if *argUp {
		exx.CI(*argName == "", "usage: wgconf -up -name PeerName")
		exx.C(wg.RunUp(cfg, *argName))
		os.Exit(0)
	}

	if *argDown {
		exx.CI(*argName == "", "usage: wgconf -down -name PeerName")
		exx.C(wg.RunDown(cfg, *argName))
		os.Exit(0)
	}

	flag.Usage()
}
