package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListConfigMap(clientSet *kubernetes.Clientset, namespace string) (result *v1.ConfigMapList, err error) {
	result, err = clientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetConfigMap(clientSet *kubernetes.Clientset, namespace string, configMapName string) (result *v1.ConfigMap, err error) {
	result, err = clientSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateConfigMap(clientSet *kubernetes.Clientset, namespace string, configMapName string, configData map[string]string) (err error) {
	var configMap v1.ConfigMap
	configMap.Name = configMapName
	configMap.Data = configData
	_, err = clientSet.CoreV1().ConfigMaps(namespace).Create(context.TODO(), &configMap, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func UpdateConfigMap(clientSet *kubernetes.Clientset, namespace string, configMapName string, configData map[string]string) (err error) {
	var configMap v1.ConfigMap
	configMap.Name = configMapName
	configMap.Data = configData
	_, err = clientSet.CoreV1().ConfigMaps(namespace).Update(context.TODO(), &configMap, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return
}

func DeleteConfigMap(clientSet *kubernetes.Clientset, namespace string, name string) (err error) {
	err = clientSet.CoreV1().ConfigMaps("k8s-test").Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return
}
