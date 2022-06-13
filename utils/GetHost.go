package utils

import (
	"fmt"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"gorm.io/gorm"
)

// VmsHost 主机结构体
type VmsHost struct {
	Name string
	Ip   string
}

// VmsHosts 主机列表结构体
type VmsHosts struct {
	VmsHosts []VmsHost
}

// NewVmsHosts 初始化结构体
func NewVmsHosts() *VmsHosts {
	return &VmsHosts{
		VmsHosts: make([]VmsHost, 10),
	}
}

// 虚拟机表
type Vm struct {
	gorm.Model
	Uuid       string `gorm:"type:varchar(40);not null;unique;comment:'虚拟机id'"`
	Vc         string `gorm:"type:varchar(30);comment:'Vcenter Ip'"`
	Esxi       string `gorm:"type:varchar(30);comment:'Esxi Id'"`
	Name       string `gorm:"type:varchar(90);comment:'Vm名字'"`
	Ip         string `gorm:"type:varchar(20);comment:'Vm ip'"`
	PowerState string `gorm:"type:varchar(20);comment:'Vm state'"`
}

// AddHost 新增主机
func (vmshosts *VmsHosts) AddHost(name string, ip string) {
	host := &VmsHost{name, ip}
	vmshosts.VmsHosts = append(vmshosts.VmsHosts, *host)
}

// SelectHost 查询主机ip
func (vmshosts *VmsHosts) SelectHost(name string) string {
	ip := "None"
	for _, hosts := range vmshosts.VmsHosts {
		if hosts.Name == name {
			ip = hosts.Ip
		}
	}
	return ip
}

// GetHosts 读取主机信息
func GetHosts(client *vim25.Client, vmshosts *VmsHosts) {
	m := view.NewManager(client)
	v, err := m.CreateContainerView(CTX, client.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	if err != nil {
		panic(err)
	}
	defer v.Destroy(CTX)
	var hss []mo.HostSystem
	err = v.Retrieve(CTX, []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		panic(err)
	}
	fmt.Printf("主机名:\t%s\n", hss[0].Summary.Host.Value)
	fmt.Printf("IP:\t%s\n", hss[0].Summary.Config.Name)
	for _, hs := range hss {
		vmshosts.AddHost(hs.Summary.Host.Value, hs.Summary.Config.Name)
	}
}
