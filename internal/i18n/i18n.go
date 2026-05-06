package i18n

// Labels returns UI label translations keyed by locale and label key.
func Labels() map[string]map[string]string {
	return labels
}

var labels = map[string]map[string]string{
	"de": {
		"app_title":      "Heidelberg Guide",
		"nav_home":       "Startseite",
		"back":           "Zurück",
		"year_built":     "Erbaut",
		"year_destroyed": "Zerstört",
		"language":       "Sprache",
		"description":    "Beschreibung",
		"history":        "Geschichte",
	},
	"en": {
		"app_title":      "Heidelberg Guide",
		"nav_home":       "Home",
		"back":           "Back",
		"year_built":     "Built",
		"year_destroyed": "Destroyed",
		"language":       "Language",
		"description":    "Description",
		"history":        "History",
	},
}
