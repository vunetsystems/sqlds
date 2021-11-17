package sqlds

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

// Start listening to requests send from Grafana. This call is blocking so
// it wont finish until Grafana shutsdown the process or the plugin choose
// to exit close down by itself
func Start(pluginId string, c Driver) {
	ds := NewDatasource(c)
	if err := datasource.Manage(pluginId, ds.NewDatasource, datasource.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
