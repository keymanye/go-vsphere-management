package utils

import (
	"encoding/csv"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"os"
	"strconv"
)

// GetVms获取所有vm信息
func GetVms(client *vim25.Client, vmshosts *VmsHosts) {
	m := view.NewManager(client)
	v, err := m.CreateContainerView(CTX, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		panic(err)
	}
	defer v.Destroy(CTX)
	var vms []mo.VirtualMachine
	err = v.Retrieve(CTX, []string{"VirtualMachine"}, []string{"summary", "runtime", "datastore"}, &vms)
	if err != nil {
		panic(err)
	}
	// 输出虚拟机信息到csv
	file, _ := os.OpenFile("./vms.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	//防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"宿主机", "虚拟机", "系统", "状态", "IP地址", "内存", "CPU", "存储"})
	w.Flush()
	for _, vm := range vms {
		//虚拟机资源信息
		mem := strconv.Itoa(int(vm.Summary.Config.MemorySizeMB)) + " MB "
		cpu := strconv.Itoa(int(vm.Summary.Config.NumCpu)) + " vCPU(s) "
		//storage := units.ByteSize(vm.Summary.Storage.Committed + vm.Summary.Storage.Uncommitted).String()
		storage := units.ByteSize(vm.Summary.Storage.Committed).String()
		//fmt.Println("主机名称", vm.Summary.Runtime.Host.Value)
		w.Write([]string{
			vm.Summary.Runtime.Host.Value,
			vm.Summary.Config.Name, vm.Summary.Config.GuestFullName,
			string(vm.Summary.Runtime.PowerState), vm.Summary.Guest.IpAddress, mem, cpu, storage})
		w.Flush()
	}
	file.Close()

	// 批量插入到数据库
	/*var modelVms []*Vm
	for _, vm := range vms {
		modelVms = append(modelVms, &Vm{
			Uuid:       vm.Summary.Config.Uuid,
			Vc:         VSPHERE_IP,
			Esxi:       vm.Summary.Runtime.Host.Value,
			Name:       vm.Summary.Config.Name,
			Ip:         vm.Summary.Guest.IpAddress,
			PowerState: string(vm.Summary.Runtime.PowerState),
		})
	}*/

}
