package services

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type BaseService struct {
}

func (ctrl BaseService) Ptrint32(p int32) *int32 {
	return &p
}

func (ctrl BaseService) GetKuberClient() *kubernetes.Clientset {
	var kubeconfig string
	kubeconfig, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}
	fmt.Println(kubeconfig)
	clientconfig := clientcmd.GetConfigFromFileOrDie(kubeconfig)
	fmt.Println(clientconfig.Contexts)

	config, err := clientcmd.NewNonInteractiveClientConfig(*clientconfig, "local", &clientcmd.ConfigOverrides{}, clientcmd.NewDefaultClientConfigLoadingRules()).ClientConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
