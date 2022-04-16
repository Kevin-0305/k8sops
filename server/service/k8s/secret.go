package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListSecret(clientSet *kubernetes.Clientset, namespace string) (result *v1.SecretList, err error) {
	result, err = clientSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetSecret(clientSet *kubernetes.Clientset, namespace string, secretName string) (result *v1.Secret, err error) {
	result, err = clientSet.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateSecret(clientSet *kubernetes.Clientset, namespace string, secretName string, secretData map[string]string) (err error) {
	var secret v1.Secret
	secretDataBytes := make(map[string][]byte)
	for k, v := range secretData {
		secretDataBytes[k] = []byte(v)
	}
	secret.Name = secretName
	secret.Data = secretDataBytes
	_, err = clientSet.CoreV1().Secrets(namespace).Create(context.TODO(), &secret, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func UpdateSecret(clientSet *kubernetes.Clientset, namespace string, secretName string, secretData map[string]string) (err error) {
	var secret v1.Secret
	var secretDataBytes map[string][]byte
	for k, v := range secretData {
		secretDataBytes[k] = []byte(v)
	}
	secret.Name = secretName
	secret.Data = secretDataBytes
	_, err = clientSet.CoreV1().Secrets(namespace).Update(context.TODO(), &secret, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeleteSecret(clientSet *kubernetes.Clientset, namespace string, secretName string) (err error) {
	err = clientSet.CoreV1().Secrets(namespace).Delete(context.TODO(), secretName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
