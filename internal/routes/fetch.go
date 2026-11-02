package routes

import (
	"context"
	"fmt"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Fetch lists ingresses in every namespace
func Fetch(
	ctx context.Context,
	client kubernetes.Interface,
) ([]networkingv1.Ingress, error) {
	list, err := client.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list ingresses: %w", err)
	}
	return list.Items, nil
}
