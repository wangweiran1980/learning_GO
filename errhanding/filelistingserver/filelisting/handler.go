package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandlerFileList(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	w.Write(all)
	return nil
}
