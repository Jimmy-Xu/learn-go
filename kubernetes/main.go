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
	host := flag.String("host", "", "kubernetes server")

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
