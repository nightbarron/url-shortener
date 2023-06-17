package gin_template

import (
	"gin_template/helpers"
	route "gin_template/routes"
)

func main() {

	// Init log
	helpers.InitLogger()

	// Init config
	//ctl.LoadVhostsFromStorage(&vhostList)

	r := route.SetupRouter()
	r.Run(":8080")

}
