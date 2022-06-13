package utils

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
)

func FindVMByIP(ctx context.Context, c *vim25.Client, ip string) {
	searchIndex := object.NewSearchIndex(c)
	// nil 是数据中心，你如果要指定的话，需要构造数据中心的结构体，这里就不指定了
	// ip 是你虚拟机的 ip
	// true 的意思我没搞懂
	reference, err := searchIndex.FindByIp(ctx, nil, ip, true)
	// 之所以只对 reference 进行判断而非对 err 是因为没有找到不算是 error
	// 也就是说 err 为 nil 并不代表就找到了，但是没找到 reference 一定为 nil
	if reference == nil {
		panic("vm not found")
	}
	if err != nil {
		panic(err)
	}

	// 这类的查找的对象都是 object.Reference，你需要通过对应的方法将其转换为相应的对象
	// 比如虚拟机、文件夹、模板等~
	fmt.Println("---------------------------------")
	fmt.Println("Reference", reference)
	vm := object.NewVirtualMachine(c, reference.Reference())
	fmt.Println(vm)
}
