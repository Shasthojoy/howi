package hcli

// ApplicationInfo object
type ApplicationInfo struct {
	Name             string   `json:"name"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"short-description"`
	LongDescription  string   `json:"long-description"`
	CopyRight        string   `json:"copyright"`
	CopySince        int      `json:"copy-since"`
	CopyMsg          string   `json:"copy-msg"`
	URL              string   `json:"url"`
	Version          string   `json:"version"`
	Authors          []string `json:"authors"`
}
