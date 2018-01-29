package main

import (
	"encoding/json"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"flag"
	"time"
	"strings"
)

var (
	//HYPER_HOST = "tcp://gcp-us-central1.hyper.sh:6443"
	HYPER_HOST = "tcp://127.0.0.1:6443"
	POD_NUM = 10
)

func main() {
	flag.Parse()
	action := flag.Arg(0)

	config := &rest.Config{
		Host: HYPER_HOST,
	}

	fmt.Printf("server: %v\n", config.Host)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	switch action {
	case "create":
		for i:=0; i< POD_NUM; i++ {
			createPod(i, clientset)
		}
	case "list":
		listPod(clientset)
	case "delete":
		for i:=0; i< POD_NUM; i++ {
			deletePod(i, clientset)
		}
	default :
		fmt.Println("Please specify --action")
	}
}

func createPod(index int, clientset *kubernetes.Clientset) error {
	var err error
	podStr := fmt.Sprintf(`{
			  "apiVersion": "v1",
			  "kind": "Pod",
			  "metadata": {
				"name": "pod-%v"
			  },
			  "spec": {
				"containers": [
				  {
					"name": "nginx",
					"image": "nginx:1.7.9",
					"ports": [
					  {
						"containerPort": %v
					  }
					]
				  }
				]
			  }
			}`, index, 8000 + index)

	name := fmt.Sprintf("pod-%v",index)
	begin := time.Now()
	defer func() {
		LogTimeConsumption(begin, err,"create pod %s", name)
	}()

	var podData apiv1.Pod
	json.Unmarshal([]byte(podStr), &podData)
	_, err = clientset.CoreV1().Pods(apiv1.NamespaceDefault).Create(&podData)
	return err
}

func listPod(clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.ListOptions{}
	podList, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).List(opts)
	if err != nil {
		fmt.Errorf("get pod list error: %v", err)
	} else {
		for i, item := range podList.Items {
			fmt.Printf("%v: %v %v\n", i, item.Name, item.Status.Message)
		}
	}
}

func deletePod(index int, clientset *kubernetes.Clientset) {
	var err error
	name := fmt.Sprintf("pod-%v",index)
	begin := time.Now()

	defer func() {
		LogTimeConsumption(begin, err,"delete pod %s", name)
	}()

	opts := meta_v1.DeleteOptions{}
	err = clientset.CoreV1().Pods(apiv1.NamespaceDefault).Delete(name, &opts)
}

//////////////////////////
// utility
//////////////////////////
func LogTimeConsumption(begin time.Time, err error, format string, args ...interface{}) {
	if err == nil {
		msg := fmt.Sprintf(format, args...)
		fmt.Printf("%s : %v\n", msg, time.Now().Sub(begin))
	} else {
		if strings.Contains(err.Error(),"could not find the requested resource") {
			err = fmt.Errorf("not found")
		}
		msg := fmt.Sprintf(format, args...)
		fmt.Printf("%s error: %v\n", msg, err)
	}
}
