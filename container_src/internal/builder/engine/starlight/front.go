package starlight

type Front struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Template    string `yaml:"template,omitempty"`
	Hero        *Hero  `yaml:"hero,omitempty"`
}

type Hero struct {
	Tagline string       `yaml:"tagline,omitempty"`
	Image   *HeroImage   `yaml:"image,omitempty"`
	Actions []HeroAction `yaml:"actions,omitempty"`
}

type HeroImage struct {
	File string `yaml:"file"`
}

type HeroAction struct {
	Text    string `yaml:"text"`
	Link    string `yaml:"link"`
	Icon    string `yaml:"icon,omitempty"`
	Variant string `yaml:"variant,omitempty"`
}
