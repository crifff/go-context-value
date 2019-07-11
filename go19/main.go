package go19

import (
	"encoding/json"
	"net/http"

	"fmt"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	v := map[string]interface{}{}

	v["AppID"] = appengine.AppID(ctx)
	v["Datacenter"] = appengine.Datacenter(ctx)
	v["DefaultVersionHostname"] = appengine.DefaultVersionHostname(ctx)
	v["ModuleName"] = appengine.ModuleName(ctx)
	v["RequestID"] = appengine.RequestID(ctx)
	v["VersionID"] = appengine.VersionID(ctx)
	v["ServerSoftware"] = appengine.ServerSoftware()
	v["InstanceID"] = appengine.InstanceID()
	v["IsSecondGen"] = appengine.IsSecondGen()
	v["IsFlex"] = appengine.IsFlex()
	v["IsStandard"] = appengine.IsStandard()
	v["IsAppEngine"] = appengine.IsAppEngine()
	v["IsDevAppServer"] = appengine.IsDevAppServer()
	v["ServiceAccount"], _ = appengine.ServiceAccount(ctx)

	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "text/json")
	fmt.Fprint(w, string(b))
}
