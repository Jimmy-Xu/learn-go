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
	host := flag.String("host", "", "apiserver, format http://x.x.x.x:8001, tcp://x.x.x.x:6443")
	action := flag.String("action", "list", "list, create")

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

	if *host != "" {
		config.Host = *host
	}

	glog.V(4).Infof("host: %v\n", config.Host)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	switch *action {
	case "list":
		listPod(err, clientset)
	case "create":
		createPod(err, clientset)
	default:
		fmt.Printf("unsupport action:%v\n", *action)
	}
}
func createPod(err error, clientset *kubernetes.Clientset) {
	podStr := `{
			  "apiVersion": "v1",
			  "kind": "Pod",
			  "metadata": {
				"name": "nginx5"
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
	podResp, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Create(&podData)
	if err != nil {
		fmt.Printf("create pod error:%v\n", err)
		return
	}
	buf, err := json.MarshalIndent(podResp, "", "  ")
	fmt.Printf("create pod result:\n%v\n", string(buf))
}

func listPod(err error, clientset *kubernetes.Clientset) {
	opts := meta_v1.ListOptions{}
	podList, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).List(opts)
	if err != nil {
		glog.Errorf("get pod list error: %v", err)
	} else {
		buf, err := json.MarshalIndent(*podList, "", " ")
		if err != nil {
			glog.Errorf("result format error: %v", err)
		}
		fmt.Printf("podList:\n%v\n", string(buf))
	}
}
