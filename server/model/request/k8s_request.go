package request

type K8sRequest struct {
	ClusterID  string `json:"clusterId" form:"clusterId"`   //集群ID
	Namespace  string `json:"namespace" form:"namespace"`   //集群ID
	Name       string `json:"name" form:"name"`             // 页码
	ConfigData string `json:"configData" form:"configData"` // 每页大小
}

type K8sRequestSearch struct {
	ClusterID string `json:"clusterId" form:"clusterId"` //集群ID
	Namespace string `json:"namespace" form:"namespace"` //集群ID
	PageInfo
}

type K8sSecretRequest struct {
	ClusterID  string            `json:"clusterId" form:"clusterId"`   //集群ID
	Namespace  string            `json:"namespace" form:"namespace"`   //集群ID
	Name       string            `json:"name" form:"name"`             // 页码
	ConfigData string            `json:"configData" form:"configData"` //
	Data       map[string]string `json:"data" form:"data"`             //
}
