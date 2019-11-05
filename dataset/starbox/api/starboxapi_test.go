package api

import (
	"lhr/foundation/definitions"
	"os"
	"syscall"
	"testing"
)

func TestAllWithEnvPara(t *testing.T) {
	var err error
	err = os.Setenv(definitions.DLhrDatasetStarBoxAccessId, "CCbDeheb")
	err = os.Setenv(definitions.DLhrDatasetStarBoxAccessKey, "4hL7ZKn3BrjO")
	t.Log(syscall.Getenv(definitions.DLhrDatasetStarBoxAccessId))
	t.Log(syscall.Getenv(definitions.DLhrDatasetStarBoxAccessKey))
	api := NewStarBoxApi()
	dirId, _ := api.GetDirId("/Public/CICD/livetest/liveCICD/1_attack_3D/1_iMac_dark_whitebg")
	scrollId, _, _, _, _ := api.GetScrollInDir(dirId)
	totalNum, filename, storageId, err := api.GetNextScroll(scrollId)
	t.Log(totalNum)
	t.Log(filename)
	t.Log(storageId)
	t.Log(err)
	content, err := api.GetFileByStorageId(storageId)
	t.Log(content)
	t.Log(err)
	err = os.Unsetenv(definitions.DLhrDatasetStarBoxAccessId)
	err = os.Unsetenv(definitions.DLhrDatasetStarBoxAccessKey)
}

func TestAllWithoutEnvPara(t *testing.T) {
	api := NewStarBoxApi()
	dirId, _ := api.GetDirId("/Public/CICD/livetest/liveCICD/1_attack_3D/1_iMac_dark_whitebg")
	scrollId, _, _, _, _ := api.GetScrollInDir(dirId)
	totalNum, filename, storageId, err := api.GetNextScroll(scrollId)
	t.Log(totalNum)
	t.Log(filename)
	t.Log(storageId)
	t.Log(err)
	content, err := api.GetFileByStorageId(storageId)
	t.Log(content)
	t.Log(err)
}
