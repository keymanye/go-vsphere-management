package main

import (
	"VmwareManagement/utils"
)

var hosts utils.VmsHosts
var vms utils.VmsHosts

func main() {
	utils.GetHosts(utils.VsphereClient, &hosts)
	//utils.FindVMByName(context.Background(), utils.NewClient(), "ming200.44")
	//fmt.Println(utils.FindVMByIP(context.Background(), utils.NewClient(), "172.31.200.44"))
	//utils.PoweroffVM(utils.CTX, utils.VsphereClient, "DO447-08")
	//utils.PowerOnVM(context.Background(), utils.NewClient(), "DO447-08")
}
