package starbox

import (
	"lhr/foundation/dataset/starbox/api"
)

func NewStarBox(path string) *StarBox {
	item := new(StarBox)
	item.path = path
	item.last = false
	return item
}

type StarBox struct {
	path     string
	scrollId string
	last     bool
}

func (dataset *StarBox) GetFirst() (filename string, storageId string) {
	starBoxApi := api.NewStarBoxApi()
	dirId, err := starBoxApi.GetDirId(dataset.path)
	if err == nil {
		dataset.scrollId, _, filename, storageId, err = starBoxApi.GetScrollInDir(dirId)
		if err != nil {
			dataset.last = true
		}
	} else {
		dataset.last = true
	}
	return filename, storageId
}

func (dataset *StarBox) GetNext() (filename string, storageId string) {
	starBoxApi := api.NewStarBoxApi()
	_, filename, storageId, err := starBoxApi.GetNextScroll(dataset.scrollId)
	if err != nil {
		dataset.last = true
	}
	return filename, storageId
}

func (dataset *StarBox) IsLast() (res bool) {
	return dataset.last
}

func (dataset *StarBox) GetFile(id string) (content []byte, err error) {
	starBoxApi := api.NewStarBoxApi()
	content, err = starBoxApi.GetFileByStorageId(id)
	return content, err
}

