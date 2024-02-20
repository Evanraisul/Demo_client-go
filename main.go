package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeConfig := flag.String("kubeConfig", "/home/appscodepc/.kube/config", "Location of my kubeConfig File")
	//fmt.Println(kubeConfig)

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		fmt.Printf("Building config from flags %s", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Getting InCLusterConfig %s", err.Error())
		}
	}
	//fmt.Println(config)

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Creating ClientSet %s", err.Error())
	}
	//fmt.Println(clientSet)
	ctx := context.Background()
	pods, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Listening Pods %s", err.Error())
	}
	// fmt.Println(pods)
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	deployments, err := clientSet.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Listening Deployments %s", err.Error())
	}
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
}
