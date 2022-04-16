package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListService(clientSet *kubernetes.Clientset, namespace string) (result *v1.ServiceList, err error) {
	result, err = clientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetService(clientSet *kubernetes.Clientset, namespace string, serviceName string) (result *v1.Service, err error) {
	result, err = clientSet.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateService(clientSet *kubernetes.Clientset, namespace string, serviceJson string) (err error) {
	var service v1.Service
	fmt.Println(serviceJson)
	err = json.Unmarshal([]byte(serviceJson), &service)
	if err != nil {
		return
	}
	_, err = clientSet.CoreV1().Services(namespace).Create(context.TODO(), &service, metav1.CreateOptions{})
	if err != nil {
		return
	}
	return nil
}

func UpdateService(clientSet *kubernetes.Clientset, namespace string, serviceJson string) (err error) {
	var service v1.Service

	err = json.Unmarshal([]byte(serviceJson), &service)
	if err != nil {
		return
	}
	_, err = clientSet.CoreV1().Services(namespace).Update(context.TODO(), &service, metav1.UpdateOptions{})
	if err != nil {
		return
	}
	return nil
}

func DeleteService(clientSet *kubernetes.Clientset, namespace string, serviceName string) (err error) {
	err = clientSet.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	return
}
