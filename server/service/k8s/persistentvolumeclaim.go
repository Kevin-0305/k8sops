package k8s

import (
	"context"
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPersistentVolumeClaim(clientSet *kubernetes.Clientset, namespace string) (result *v1.PersistentVolumeClaimList, err error) {
	result, err = clientSet.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetPersistentVolumeClaim(clientSet *kubernetes.Clientset, namespace string, persistentVolumeClaimName string) (result *v1.PersistentVolumeClaim, err error) {
	result, err = clientSet.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), persistentVolumeClaimName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreatePersistentVolumeClaim(clientSet *kubernetes.Clientset, namespace string, statefulSetJson string) (err error) {
	var persistentVolumeClaim v1.PersistentVolumeClaim
	err = json.Unmarshal([]byte(statefulSetJson), &persistentVolumeClaim)
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), &persistentVolumeClaim, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func UpdatePersistentVolumeClaim(clientSet *kubernetes.Clientset, namespace string, statefulSetJson string) (err error) {
	var persistentVolumeClaim v1.PersistentVolumeClaim
	err = json.Unmarshal([]byte(statefulSetJson), &persistentVolumeClaim)
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().PersistentVolumeClaims(namespace).Update(context.TODO(), &persistentVolumeClaim, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return
}

func DeletePersistentVolumeClaim(clientSet *kubernetes.Clientset, namespace string, statefulSetName string) (err error) {
	err = clientSet.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), statefulSetName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return
}
