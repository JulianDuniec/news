package jsonpproxy

import(
	"net/http"
	"fmt"
	"io/ioutil"
)


type Proxy struct{}

func (f *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(req.FormValue("url"))
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
		return
	}
	fmt.Fprintf(w, "%s(%s)", req.FormValue("jsonp"), body)
}

func Start(port string) {
	http.ListenAndServe(":" + port, &Proxy{})
}