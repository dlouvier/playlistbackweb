package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/dlouvier/playlistbackweb/controllers:DeezerController"] = append(beego.GlobalControllerRouter["github.com/dlouvier/playlistbackweb/controllers:DeezerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/dlouvier/playlistbackweb/controllers:SpotifyController"] = append(beego.GlobalControllerRouter["github.com/dlouvier/playlistbackweb/controllers:SpotifyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
