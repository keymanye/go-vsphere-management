package utils

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

// 这里的 c 就是上面登录后 client 的 Client 属性
func FindVMByName(ctx context.Context, c *vim25.Client, vmName string) {
	m := view.NewManager(c)

	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		panic(err)
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all machines
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		panic(err)
	}

	for _, vm := range vms {
		// 判断虚拟机名称是否相同，相同的话，vm 就是查找到主机
		if vm.Summary.Config.Name == vmName {
			fmt.Println(vm)
			break
		}
	}
}
