package adapters

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	apiV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type k8sAdapter struct {
	clientSet *kubernetes.Clientset
}

func New() k8sAdapter {
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
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return k8sAdapter{
		clientSet: clientset,
	}
}

func (adapter k8sAdapter) GetPods() []string {
	pods, err := adapter.clientSet.CoreV1().Pods("staging").List(apiV1.ListOptions{})

	if err != nil {
		panic(err)
	}

	list := make([]string, len(pods.Items))

	for i, pod := range pods.Items {
		list[i] = pod.Name
	}

	return list
}

func (adapter k8sAdapter) GetLogs(podName string) string {
	var tail int64 = 20

	req := adapter.clientSet.CoreV1().Pods("staging").GetLogs(podName, &v1.PodLogOptions{
		Follow:    false,
		TailLines: &tail,
	})

	fmt.Println("testing")

	readCloser, err := req.Stream()
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, readCloser)
	if err != nil {
		panic(err)
	}
	str := buf.String()

	fmt.Println(str)

	defer readCloser.Close()

	return str
}
