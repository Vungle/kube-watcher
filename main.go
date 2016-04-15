package main

import (
	"fmt"
	datadog "github.com/zorkian/go-datadog-api"
	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"os"
	"strconv"
	"time"
)

func main() {
	client, err := client.NewInCluster()
	if err != nil {
		// handle error
	}
	interval, _ := strconv.ParseInt(os.Getenv("MONITOR_INTERVAL"), 10, 64)
	dd_client := datadog.NewClient(os.Getenv("DD_API_KEY"), os.Getenv("DD_APP_KEY"))
	for {
		pods, _ := client.Pods("default").List(api.ListOptions{})
		for _, p := range pods.Items {
			for _, cs := range p.Status.ContainerStatuses {
				switch {
				case cs.State.Waiting != nil:
					eText := fmt.Sprintf("Reason: %s", cs.State.Waiting.Reason)
					title := fmt.Sprintf("%s pod is in a waiting state", p.Name)
					e := datadog.Event{
						Title: title,
						Text:  eText,
						Host:  p.Name,
					}
					if _, err := dd_client.PostEvent(&e); err != nil {
						fmt.Println(err)
					}
					fmt.Println("Name: ", cs.Name)
					fmt.Println("Reason: ", cs.State.Waiting.Reason)
				case cs.State.Running == nil:
					fmt.Println(cs.State)
				}
			}
		}
		fmt.Println("\n\n##########################\n\n")
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
