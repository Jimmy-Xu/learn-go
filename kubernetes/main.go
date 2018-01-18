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
	pvName := flag.String("pv-name", "test-pv", "PersistentVolume name")
	pvcName := flag.String("pvc-name", "test-pvc", "PersistentVolumeClaim name")

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
	//pod
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
	//node
	case "list-node":
		listNode(clientset)
	//pv
	case "create-pv":
		createPersistentVolume(*pvName, clientset)
	case "get-pv":
		getPersistentVolume(*pvName, clientset)
	case "delete-pv":
		deletePersistentVolume(*pvName, clientset)
	case "list-pv":
		listPersistentVolume(clientset)
	//pvc
	case "create-pvc":
		createPersistentVolumeClaim(*pvcName, clientset)
	case "get-pvc":
		getPersistentVolumeClaim(*pvcName, clientset)
	case "delete-pvc":
		deletePersistentVolumeClaim(*pvcName, clientset)
	case "list-pvc":
		listPersistentVolumeClaim(clientset)
	//other
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

func createPersistentVolume(name string, clientset *kubernetes.Clientset) {
	var err error
	pvStr := `{
	  "apiVersion": "v1",
	  "kind": "PersistentVolume",
	  "metadata": {
		"name": "pv-1"
	  },
	  "spec": {
		"accessModes": [
		  "ReadWriteOnce"
		],
		"capacity": {
		  "storage": "1Gi"
		},
		"hostPath": {
		  "path": "/tmp/pv-1",
		  "type": ""
		}
	  }
	}`
	var pvData apiv1.PersistentVolume
	json.Unmarshal([]byte(pvStr), &pvData)
	pvData.Name = name
	pvData.Spec.HostPath.Path = fmt.Sprintf("/tmp/%v", name)
	pvResp, err := clientset.CoreV1().PersistentVolumes().Create(&pvData)
	if err != nil {
		fmt.Printf("create pv %v error:%v\n", name, err)
		return
	}
	printJson(pvResp)
}

func listPersistentVolume(clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.ListOptions{}
	pvList, err := clientset.CoreV1().PersistentVolumes().List(opts)
	if err != nil {
		glog.Errorf("get pv list error: %v", err)
	} else {
		printJson(*pvList)
	}
}

func getPersistentVolume(name string, clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.GetOptions{}
	pvInfo, err := clientset.CoreV1().PersistentVolumes().Get(name, opts)
	if err != nil {
		glog.Errorf("get pv %v error: %v", name, err)
	} else {
		printJson(*pvInfo)
	}
}

func deletePersistentVolume(name string, clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.DeleteOptions{}
	err = clientset.CoreV1().PersistentVolumes().Delete(name, &opts)
	if err != nil {
		glog.Errorf("delete pv error: %v", err)
	} else {
		fmt.Printf("pv %v deleted", name)
	}
}

func createPersistentVolumeClaim(name string, clientset *kubernetes.Clientset) {
	var err error
	pvcStr := `{
	  "apiVersion": "v1",
	  "kind": "PersistentVolumeClaim",
	  "metadata": {
		"name": "pvc-1"
	  },
	  "spec": {
		"accessModes": [
		  "ReadWriteOnce"
		],
		"resources": {
		  "requests": {
			"storage": "1Gi"
			}
		}
	  }
	}`
	var pvcData apiv1.PersistentVolumeClaim
	json.Unmarshal([]byte(pvcStr), &pvcData)
	pvcData.Name = name
	pvcResp, err := clientset.CoreV1().PersistentVolumeClaims(apiv1.NamespaceDefault).Create(&pvcData)
	if err != nil {
		fmt.Printf("create pvc %v error:%v\n", name, err)
		return
	}
	printJson(pvcResp)
}

func listPersistentVolumeClaim(clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.ListOptions{}
	pvcList, err := clientset.CoreV1().PersistentVolumeClaims(apiv1.NamespaceDefault).List(opts)
	if err != nil {
		glog.Errorf("get pvc list error: %v", err)
	} else {
		printJson(*pvcList)
	}
}

func getPersistentVolumeClaim(name string, clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.GetOptions{}
	pvcInfo, err := clientset.CoreV1().PersistentVolumeClaims(apiv1.NamespaceDefault).Get(name, opts)
	if err != nil {
		glog.Errorf("get pvc %v error: %v", name, err)
	} else {
		printJson(*pvcInfo)
	}
}

func deletePersistentVolumeClaim(name string, clientset *kubernetes.Clientset) {
	var err error
	opts := meta_v1.DeleteOptions{}
	err = clientset.CoreV1().PersistentVolumeClaims(apiv1.NamespaceDefault).Delete(name, &opts)
	if err != nil {
		glog.Errorf("delete pvc error: %v", err)
	} else {
		fmt.Printf("pvc %v deleted", name)
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
