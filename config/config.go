package config

import (
	"os"
)

var (
	LocalVendorPath = "abc"

	MirrorCarrier string
)

func init() {
	_ = os.Getenv("DLMAN_ENV")
	LocalVendorPath = os.Getenv("DLMAN_LOCAL_VENDOR_PATH")
	MirrorCarrier = "LOCAL"
}