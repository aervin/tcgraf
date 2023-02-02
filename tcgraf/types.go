package tcgraf

type EventType string

func (ev EventType) String() string {
	return string(ev)
}

const (
	CreateDash EventType = "create_dash"
)

type Event struct {
	InstanceId string    `json:"instance_id"`
	NasId      string    `json:"nas_id"`
	Type       EventType `json:"type"`
}

type ResponseSuccess struct {
	Message any  `json:"message"`
	Success bool `json:"success"`
}

type PanelModel struct {
	Id           int64           `json:"id"`
	Title        string          `json:"title"`
	Type         string          `json:"type"`
	Interval     string          `json:"interval"`
	Datasource   PanelDataSource `json:"datasource"`
	Description  string          `json:"description"`
	Targets      []PanelTarget   `json:"targets"`
	FieldConfig  map[string]any  `json:"fieldConfig"`
	GridPos      map[string]any  `json:"gridPos"`
	LibraryPanel map[string]any  `json:"libraryPanel"`
	Options      map[string]any  `json:"options"`
	TimeFrom     string          `json:"timeFrom"`
}

type PanelTarget struct {
	Datasource   PanelDataSource `json:"datasource"`
	EditorMode   string          `json:"editorMode"`
	Expr         string          `json:"expr"`
	LegendFormat string          `json:"legendFormat"`
	Range        bool            `json:"range"`
	RefId        string          `json:"refId"`
}

type PanelDataSource struct {
	Type string `json:"type"`
	Uid  string `json:"uid"`
}
