package routes

import networkingv1 "k8s.io/api/networking/v1"

// Build creates and returns a list of Routes from ingresses
func Build(ingresses []networkingv1.Ingress) []Route {
	var routes []Route
	for _, ingress := range ingresses {
		for _, rule := range ingress.Spec.Rules {
			routes = append(routes, Route{
				Namespace: ingress.Namespace,
				Ingress:   ingress.Name,
				Host:      rule.Host,
			})
		}
	}
	return routes
}
