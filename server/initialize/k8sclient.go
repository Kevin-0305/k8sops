package initialize

import (
	"fmt"
	"gin-vue-admin/global"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func InitK8sClient() {
	k8sConfig, err := clientcmd.BuildConfigFromFlags("", fmt.Sprintf("/home/.kube/config/admin.conf")) // 使用 kubectl 默认配置 ~/.kube/config
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// 创建一个k8s客户端
	clientSet, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	global.GVA_K8SCLIENTS = append(global.GVA_K8SCLIENTS, clientSet)
}
