package starbox
import (
	"archive/zip"
	"io"
	"lhr/foundation/definitions"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("testdata.zip")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}
func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()
		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}
func Starbox_Down_Unzip(path string) error {
	var err error
	err = os.Setenv(definitions.DLhrDatasetStarBoxAccessId, "UdKNpmc5")
	err = os.Setenv(definitions.DLhrDatasetStarBoxAccessKey, "OkR1A5fMaPG0")
	api_value := NewStarBox(path)
	_, storageId := api_value.GetFirst()
	url := `http://ks3.kylin.cloudwalk.work/starbox-prd-ai/` + storageId
	Download(url)
	Unzip("testdata.zip", "./")
	err = os.Unsetenv(definitions.DLhrDatasetStarBoxAccessId)
	err = os.Unsetenv(definitions.DLhrDatasetStarBoxAccessKey)
	return err
}
