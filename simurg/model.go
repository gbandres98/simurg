package simurg

type Manifest struct {
	Context  string `json:"@context,omitempty"`
	ID       string `json:"@id,omitempty"`
	Type     string `json:"@type,omitempty"`
	Label    string `json:"label,omitempty"`
	Metadata []struct {
		Label struct {
			Es []string `json:"es,omitempty"`
		} `json:"label,omitempty"`
		Value struct {
			Es []string `json:"es,omitempty"`
		} `json:"value,omitempty"`
	} `json:"metadata,omitempty"`
	Thumbnail struct {
		ID      string `json:"@id,omitempty"`
		Service struct {
			Context string `json:"@context,omitempty"`
			ID      string `json:"@id,omitempty"`
			Profile string `json:"profile,omitempty"`
		} `json:"service,omitempty"`
	} `json:"thumbnail,omitempty"`
	Service struct {
		Context string `json:"@context,omitempty"`
		ID      string `json:"@id,omitempty"`
		Profile string `json:"profile,omitempty"`
	} `json:"service,omitempty"`
	Sequences []struct {
		Type        string `json:"@type,omitempty"`
		ViewingHint string `json:"viewingHint,omitempty"`
		Canvases    []struct {
			ID        string `json:"@id,omitempty"`
			Type      string `json:"@type,omitempty"`
			Height    int    `json:"height,omitempty"`
			Width     int    `json:"width,omitempty"`
			Label     string `json:"label,omitempty"`
			Thumbnail struct {
				ID      string `json:"@id,omitempty"`
				Service struct {
					Context string `json:"@context,omitempty"`
					ID      string `json:"@id,omitempty"`
					Profile string `json:"profile,omitempty"`
				} `json:"service,omitempty"`
			} `json:"thumbnail,omitempty"`
			Images []struct {
				Context    string `json:"@context,omitempty"`
				ID         string `json:"@id,omitempty"`
				Type       string `json:"@type,omitempty"`
				Motivation string `json:"motivation,omitempty"`
				On         string `json:"on,omitempty"`
				Resource   struct {
					ID      string `json:"@id,omitempty"`
					Type    string `json:"@type,omitempty"`
					Format  string `json:"format,omitempty"`
					Height  int    `json:"height,omitempty"`
					Width   int    `json:"width,omitempty"`
					Service struct {
						Context string `json:"@context,omitempty"`
						ID      string `json:"@id,omitempty"`
						Profile string `json:"profile,omitempty"`
					} `json:"service,omitempty"`
				} `json:"resource,omitempty"`
			} `json:"images,omitempty"`
		} `json:"canvases,omitempty"`
	} `json:"sequences,omitempty"`
}
