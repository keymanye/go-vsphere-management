package utils

import (
	"context"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25"
	"net/url"
)

var VsphereClient *vim25.Client
var CTX = context.Background()

const (
	VSPHERE_IP       = "vcenter.training.example.com"
	VSPHERE_USERNAME = "administrator@vsphere.local"
	VSPHERE_PASSWORD = "Training@321"
	Insecure         = true
)

// NewClient 链接vmware
func init() {

	u := &url.URL{
		Scheme: "https",
		Host:   VSPHERE_IP,
		Path:   "/sdk",
	}

	u.User = url.UserPassword(VSPHERE_USERNAME, VSPHERE_PASSWORD)
	c, err := govmomi.NewClient(CTX, u, Insecure)
	if err != nil {
		panic(err)
	}
	VsphereClient = c.Client
}
