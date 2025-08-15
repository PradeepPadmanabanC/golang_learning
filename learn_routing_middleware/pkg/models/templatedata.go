package models

// Tempalte Data holds data sent from handlers to data
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Dat       map[string]interface{}
	CSRFToken string // Cross Site Request Forgery  token for security
	Flash     string
	Warning   string
	Error     string
}
