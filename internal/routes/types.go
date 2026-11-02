package routes

// Route is one ingress host and the ingress that serves it.
type Route struct {
	Namespace string
	Ingress   string
	Host      string
}

func (r Route) TabularRows() [][]string {
	return [][]string{{r.Namespace, r.Ingress, r.Host}}
}

// Headers returns the column names, in the same order as TabularRows.
func Headers() []string {
	return []string{"NAMESPACE", "INGRESS", "HOST"}
}
