package handler

import (
	"Sandstorm"
	"fmt"
	"io"
	"net/http"
)

type Proxy struct {
	URI string `json:"uri"`
}

func ProxyConnHandler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Query().Get("URI")
	// proxy := new(Proxy)

	// content, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	Sandstorm.HTTPError(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = json.Unmarshal(content, &proxy)
	// if err != nil {
	// 	Sandstorm.HTTPError(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	fmt.Println("proxy --> ", uri)
	resp, err := http.Get(uri)
	if err != nil {
		Sandstorm.HTTPError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := make([]byte, 4018)

	defer func() {
		resp.Body.Close()
	}()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	// w.Header().Set("Content-Length", fmt.Sprint(r.ContentLength))
	// resp.Header.Write(w)
	rl := 0
	n := 0
	for {
		n, err = resp.Body.Read(buf)
		// fmt.Println("--", fmt.Sprint(n), "  ", fmt.Sprint(resp.ContentLength))
		if err != nil {
			fmt.Println(err.Error())
			if err == io.EOF {
				// fmt.Println("EOF")
				return
			}
			Sandstorm.HTTPError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if n > 0 {
			// Note: avoid overwrite err returned by Read.
			rl += n
			if _, err := w.Write(buf[0:n]); err != nil {
				Sandstorm.HTTPError(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// fmt.Println(rl)

	}

}
