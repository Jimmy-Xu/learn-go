package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/golang/glog"
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	server := flag.String("server", "", "apiserver, format http://x.x.x.x:8001, tcp://x.x.x.x:6443")
	action := flag.String("action", "list-pod", "create-pod,list-pod,get-pod,update-pod,delete-pod,list-node")
	podName := flag.String("pod-name", "test-nginx", "pod name")

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	if *server != "" {
		config.Host = *server
	}

	glog.V(4).Infof("server: %v\n", config.Host)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	switch *action {
	case "create-pod":
		createPod(*podName, clientset)
	case "get-pod":
		getPod(*podName, clientset)
	case "delete-pod":
		deletePod(*podName, clientset)
	case "update-pod":
		updatePod(*podName, clientset)
	case "list-pod":
		listPod(clientset)
	case "list-node":
		listNode(clientset)
	default:
		fmt.Printf("unsupport action:%v\n", *action)
	}
}
func createPod(name string, clientset *kubernetes.Clientset) {
	var err error
	podStr := `{
			  "apiVersion": "v1",
			  "kind": "Pod",
			  "metadata": {
				"name": ""
			  },
			  "spec": {
				"containers": [
				  {
					"name": "nginx",
					"image": "nginx:1.7.9",
					"ports": [
					  {
						"containerPort": 8000
					  }
					]
				  }
				]
			  }
			}`
	var podData apiv1.Pod
	json.Unmarshal([]byte(podStr), &podData)
	podData.Name = name
	podResp, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Create(&podData)
	if err != nil {
		fmt.Printf("create pod %v error:%v\n", name, err)
		return
	}
	printJson(podResp)
}

func listPod(clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.ListOptions{}
	podList, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).List(opts)
	if err != nil {
		glog.Errorf("get pod list error: %v", err)
	} else {
		printJson(*podList)
	}
}

func getPod(name string, clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.GetOptions{}
	podInfo, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Get(name, opts)
	if err != nil {
		glog.Errorf("get pod %v error: %v", name, err)
	} else {
		printJson(*podInfo)
	}
}

func deletePod(name string, clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.DeleteOptions{}
	err = clientset.CoreV1().Pods(apiv1.NamespaceDefault).Delete(name, &opts)
	if err != nil {
		glog.Errorf("delete pod error: %v", err)
	} else {
		fmt.Printf("pod %v deleted", name)
	}
}

func updatePod(name string, clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.GetOptions{}
	podFound, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Get(name, opts)
	if err != nil {
		glog.Errorf("get pod %v error: %v", name, err)
	} else {
		podFound.Spec.Containers[0].Image = "nginx:latest"
		podUpdated, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Update(podFound)
		if err != nil {
			glog.Errorf("update pod %v error: %v", name, err)
		} else {
			printJson(*podUpdated)
		}
	}
}

func listNode(clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.ListOptions{}
	nodeList, err := clientset.CoreV1().Nodes().List(opts)
	if err != nil {
		glog.Errorf("get node list error: %v", err)
	} else {
		printJson(*nodeList)
	}
}

//////////////////////////
// utility
//////////////////////////
func printJson(data interface{}) {
	buf, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		glog.Errorf("convert data to json error: %v", err)
	}
	fmt.Printf("%v", string(buf))
}
