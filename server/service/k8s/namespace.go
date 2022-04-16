package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// type NameSpace struct {
// }

func CreateNamespace(clientSet *kubernetes.Clientset, name string) (err error) {
	var namespace v1.Namespace
	namespace.Name = name
	_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func ListNamespaces(clientSet *kubernetes.Clientset) (result *v1.NamespaceList, err error) {
	result, err = clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil

}
