package utils

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

func PoweroffVM(ctx context.Context, c *vim25.Client, vmName string) {
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
	var objectvm *object.VirtualMachine
	for _, vm := range vms {
		// 判断虚拟机名称是否相同，相同的话，vm 就是查找到主机
		if vm.Summary.Config.Name == vmName {
			objectvm = object.NewVirtualMachine(c, vm.Reference())
			fmt.Println(objectvm)
			break
		}
	}
	powerstatus, err := objectvm.PowerState(ctx)
	if err != nil {
		fmt.Println("get powerstatus error", err)
		return
	}
	if powerstatus == "poweredOn" {
		//err := objectvm.ShutdownGuest(ctx)
		//task, err := objectvm.Suspend(ctx)
		//fmt.Println(task)
		err = objectvm.ShutdownGuest(ctx)
		if err != nil {
			panic(err)
		}
	}
}
