package main

import (
	"fmt"
	"log"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {
	log.Printf("Hello")
	wf()
}

func wf() {
	var workflow = wfv1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "name",
			Annotations: map[string]string{
				"workflows.argoproj.io/description": "This demo workflow was made from the ArgoFlow SDK for Go",
			},
		},
		Spec: wfv1.WorkflowSpec{
			Entrypoint:         "root",
			ServiceAccountName: "sa",
			Volumes:            []v1.Volume{},
		},
	}

	log.Printf("WF=%v", workflow)
	connectK8s()
}
func connectK8s() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("unable to get the in cluster config : %w", err))
		return err
	}
	log.Printf("config=%s", config)
	return err
}
