package main

import (
	"go-distributed-motion-s3/dms3libs"
)

// this script will be copied to the dms3 device component platform, executed, and
// then deleted automatically
//
// NOTE: must be run with admin privileges

func main() {

	// stop existing upstart service (if running)
	dms3libs.RunCommand("service dms3server stop")

	// move binary files into /usr/local/bin
	dms3libs.CopyFile("dms3_release/go_dms3server", "/usr/local/bin/go_dms3server")
	_, err := dms3libs.RunCommand("chmod +x " + "/usr/local/bin/go_dms3server")
	dms3libs.CheckErr(err)

	// copy configuration files into /etc/distributed-motion-s3
	dms3libs.MkDir("/etc/distributed-motion-s3")
	dms3libs.CopyDir("dms3_release/dms3server", "/etc/distributed-motion-s3")
	dms3libs.CopyDir("dms3_release/dms3libs", "/etc/distributed-motion-s3")
	dms3libs.RmDir("dms3_release")

	// restart upstart service
	dms3libs.RunCommand("service dms3server start")

}
