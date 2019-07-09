package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	fw "github.com/fjammes/krasher/framework"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "kind-config-qserv"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		namespace := "default"
		pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the default ns\n", len(pods.Items))
		// fmt.Printf("%v\n", pods.Items)
		running := [2]bool{false, false}
		podname := ""
		for !running[0] || !running[1] {
			// Todo remove for with count switch
			for i := 0; i < 2; i++ {
				podname = fmt.Sprintf("xrootd-mgr-%d", i)
				pod, errGet := clientset.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
				if errGet != nil {
					panic(errGet.Error())
				}
				fmt.Printf("%v %v\n", pod.GetName(), pod.Status.Phase)
				// TODO test phase,
			}
		}

		var req fw.Request

		req.Namespace = namespace
		req.Pod = podname
		req.Container = "xrootd"
		req.Command = "/bin/ls"
		fw.CheckKubeExec(&req)

		errDel := clientset.CoreV1().Pods(namespace).Delete(podname, &metav1.DeleteOptions{})
		if errDel != nil {
			panic(errDel.Error())
		}
		fmt.Printf("Killed pod %v\n", podname)

		time.Sleep(10 * time.Second)
	}
}

// func killContainer(c *Clientset, namespace string, podname string) {
// 	err := c.CoreV1().Pods(namespace).Delete(podname, &metav1.DeleteOptions{})
// }

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
