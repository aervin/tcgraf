package tcgraf

import (
	"errors"
	
	graf "github.com/grafana/grafana-api-golang-client"
)

func folderByInstanceId(instanceId string) (folder graf.Folder, err error) {
	folders, err := client.Folders()
	if err != nil {
		return
	}

	for _, f := range folders {
		if f.Title == instanceId {
			folder = f
			return
		}
	}

	err = errors.New("folder not found for instance id " + instanceId)

	return
}