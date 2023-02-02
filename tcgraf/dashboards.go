package tcgraf

import (
	"encoding/json"

	graf "github.com/grafana/grafana-api-golang-client"
)

func createDash(event Event) (dash *graf.DashboardSaveResponse, err error) {
	folder, err := folderByInstanceId(event.InstanceId)
	if err != nil {
		folder, err = client.NewFolder(event.InstanceId)
		if err != nil {
			return
		}
	}

	req := graf.Dashboard{
		FolderID:  folder.ID,
		FolderUID: folder.UID,
		Model: map[string]any{
			"title":    event.NasId,
			"tags":     []string{event.InstanceId},
			"timezone": "browser",
			"refresh":  "5s",
		},
	}

	err = assignPanelsToDash(event, &req)
	if err != nil {
		return
	}

	return client.NewDashboard(req)
}

func assignPanelsToDash(event Event, dash *graf.Dashboard) (err error) {
	panels, err := client.LibraryPanels()
	if err != nil {
		return
	}

	list := make([]PanelModel, 0)
	for _, panel := range panels {
		b, _ := json.Marshal(panel.Model)
		var p PanelModel
		err := json.Unmarshal(b, &p)
		if err != nil {
			return err
		}

		targs := make([]PanelTarget, 0)
		for _, target := range p.Targets {
			target.Expr = target.Expr + "{nas_id=\"" + event.NasId + "\"}"
			targs = append(targs, target)
		}
		p.Targets = targs

		list = append(list, p)
	}

	dash.Model["panels"] = list

	return
}
