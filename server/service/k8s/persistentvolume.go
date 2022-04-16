package k8s

import (
	"context"
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPersistentVolume(clientSet *kubernetes.Clientset) (result *v1.PersistentVolumeList, err error) {
	result, err = clientSet.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetPersistentVolume(clientSet *kubernetes.Clientset, persistentVolumeName string) (result *v1.PersistentVolume, err error) {
	result, err = clientSet.CoreV1().PersistentVolumes().Get(context.TODO(), persistentVolumeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreatePersistentVolume(clientSet *kubernetes.Clientset, persistentVolumeJson string) (err error) {
	var persistentVolume v1.PersistentVolume
	err = json.Unmarshal([]byte(persistentVolumeJson), &persistentVolume)
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().PersistentVolumes().Create(context.TODO(), &persistentVolume, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func UpdatePersistentVolume(clientSet *kubernetes.Clientset, persistentVolumeJson string) (err error) {
	var persistentVolume v1.PersistentVolume
	err = json.Unmarshal([]byte(persistentVolumeJson), &persistentVolume)
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().PersistentVolumes().Update(context.TODO(), &persistentVolume, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeletePersistentVolume(clientSet *kubernetes.Clientset, persistentVolumeName string) (err error) {
	err = clientSet.CoreV1().PersistentVolumes().Delete(context.TODO(), persistentVolumeName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
