package services

import (
    "fmt"
    coreV1 "k8s.io/api/core/v1"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/util/homedir"
    "os"
    "path/filepath"
    "reflect"
)
func typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}
type ClusterService struct {
    BaseService
}
type K8sNode struct {
    Name string
    Ip string
    Roles string
    Status string
    Pods int64
    Cpu int64
    Memory int64
    JoinDate string
    System string
    DockerVersion string
    Labels map[string]string
}
type K8sPod struct {
    Name string
    HostIP string
    PodIP string
    Status string
    Namespace string
    CreationTime string
}
type K8sNamespace struct {
    Name string
    Status string
    CreationTime string
}
type K8sPersistentVolume struct {
    Name string
    Status string
    CreationTime string
    Storage string
}

func (ctrl ClusterService) WriteKubeConfig(content string)  {
    var kubeConfig string
    kubeConfig, ok := os.LookupEnv("KUBECONFIG")
    if !ok {
        kubeConfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
    }
    f, err := os.OpenFile(kubeConfig, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
    if err != nil {
        fmt.Println(err.Error())
    }
    _, err = f.Write([]byte(content))
    if err != nil {
        fmt.Println(err.Error())
    }
    _ = f.Close()
}

func (ctrl ClusterService)  CreateNamespace(namespace string) error {
    client := ctrl.GetKuberClient()
    _, err := client.CoreV1().Namespaces().Get(namespace, metaV1.GetOptions{})
    if err != nil {
        _, err = client.CoreV1().Namespaces().Create(&coreV1.Namespace{
            ObjectMeta: metaV1.ObjectMeta{
                Name: namespace,
            },
        })
        if err != nil {
            return err
        }
    }
    return nil
}

func (ctrl ClusterService) NodeList() ([]K8sNode){
    client := ctrl.GetKuberClient()
    nodes, _ := client.CoreV1().Nodes().List(metaV1.ListOptions{})
    items :=  []K8sNode{}
    for _,node := range nodes.Items {
        cpu,_ := node.Status.Allocatable.Cpu().AsInt64()
        memory,_ := node.Status.Allocatable.Memory().AsInt64()
        memory = memory / 1024 / 1024 / 1024
        item := K8sNode{
            Name: node.Name,
            Ip: node.Status.Addresses[0].Address,
            Status: string(node.Status.Conditions[len(node.Status.Conditions)-1].Type),
            Cpu: cpu,
            Memory: memory,
            JoinDate: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
            System: node.Status.NodeInfo.OSImage,
            Labels: node.Labels,
            DockerVersion: node.Status.NodeInfo.ContainerRuntimeVersion,
        }
        items = append(items, item)
    }
    return items
}

func (ctrl ClusterService) NamespaceList() ([]K8sNamespace){
    client := ctrl.GetKuberClient()
    ns, _ := client.CoreV1().Namespaces().List(metaV1.ListOptions{})
    items :=  []K8sNamespace{}
    for _, ns := range ns.Items {
        item := K8sNamespace{
            Name: ns.ObjectMeta.Name,
            Status: string(ns.Status.Phase),
            CreationTime: ns.CreationTimestamp.Format("2006-01-02 15:04:05"),
        }
        items = append(items, item)
    }
    return items
}

func (ctrl ClusterService) PersistentVolumeList() ([]K8sPersistentVolume){
    client := ctrl.GetKuberClient()
    pvs, _ := client.CoreV1().PersistentVolumes().List(metaV1.ListOptions{})
    items :=  []K8sPersistentVolume{}
    for _, pv := range pvs.Items {
        rsStorage := pv.Spec.Capacity[coreV1.ResourceStorage]
        fmt.Println("")
        item := K8sPersistentVolume{
            Name: pv.ObjectMeta.Name,
            Status: string(pv.Status.Phase),
            Storage: rsStorage.String(),
            CreationTime: pv.CreationTimestamp.Format("2006-01-02 15:04:05"),
        }
        items = append(items, item)
    }
    return items
}

func (ctrl ClusterService) PodList(namespace string) ([]K8sPod){
    client := ctrl.GetKuberClient()
    pods, _ := client.CoreV1().Pods(namespace).List(metaV1.ListOptions{})
    items :=  []K8sPod{}
    /*
    fmt.Println(pods.Items[1].Name)
    fmt.Println(pods.Items[1].CreationTimestamp)
    fmt.Println(pods.Items[1].Labels)
    fmt.Println(pods.Items[1].Namespace)
    fmt.Println(pods.Items[1].Status.HostIP)
    fmt.Println(pods.Items[1].Status.PodIP)
    fmt.Println(pods.Items[1].Status.StartTime)
    fmt.Println(pods.Items[1].Status.Phase)
    fmt.Println(pods.Items[1].Status.ContainerStatuses[0].RestartCount)   //重启次数
    fmt.Println(pods.Items[1].Status.ContainerStatuses[0].Image)
    */
    for _, pod := range pods.Items {
        item := K8sPod{
            Name: pod.ObjectMeta.Name,
            Status: string(pod.Status.Phase),
            HostIP: pod.Status.HostIP,
            PodIP: pod.Status.PodIP,
            Namespace: pod.Namespace,
            CreationTime: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
        }
        items = append(items, item)
    }
    return items
}
/*
func (ctrl ClusterService) DeploymentList() (namespace string){
    clientset := ctrl.GetKuberClient()
    quos, _ := clientset.AppsV1beta1().Deployments("").List(metaV1.ListOptions{})
    fmt.Println(quos)
    items :=  []K8sPersistentVolume{}
    return items
}

func (this ClusterService) ResourceQuotasList() ([]K8sPersistentVolume){
    clientset := this.GetKuberClient()
    quos, _ := clientset.CoreV1().ResourceQuotas().List(metav1.ListOptions{})
    fmt.Println(quos)
    items :=  []K8sPersistentVolume{}
    return items
}

func (this ClusterService) LimitRangesList() ([]K8sPersistentVolume){
    clientset := this.GetKuberClient()
    limits, _ := clientset.CoreV1().LimitRanges().List(metav1.ListOptions{})
    fmt.Println(quos)
    items :=  []K8sPersistentVolume{}
    return items
}

*/
