package routes

type Route struct {
	Namespace string
	Ingress   string
	Host      string
}

func (r Route) TabularRows() [][]string {
	return [][]string{{r.Namespace, r.Ingress, r.Host}}
}
