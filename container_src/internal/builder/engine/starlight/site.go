package starlight

type SiteConfig struct {
	Title   string         `json:"title"`
	Sidebar []SidebarGroup `json:"sidebar"`
}

type SidebarGroup struct {
	Label string        `json:"label"`
	Items []SidebarItem `json:"items"`
}

type SidebarItem struct {
	Autogenerate *SidebarAutogenerate `json:"autogenerate,omitempty"`
}

type SidebarAutogenerate struct {
	Directory string `json:"directory"`
}
