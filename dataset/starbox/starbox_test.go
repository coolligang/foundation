package starbox

import (
	"github.com/coolligang/foundation/definitions"
	"os"
	"testing"
)

func TestStarBoxLoopWithEnvPara(t *testing.T) {
	var err error
	err = os.Setenv(definitions.DLhrDatasetStarBoxAccessId, "CCbDeheb")
	err = os.Setenv(definitions.DLhrDatasetStarBoxAccessKey, "4hL7ZKn3BrjO")
	api := NewStarBox("/Public/CICD/livetest/liveCICD/1_attack_3D/1_iMac_dark_whitebg")
	for filename, storageId := api.GetFirst(); !api.IsLast(); filename, storageId = api.GetNext() {
		t.Log(filename)
		t.Log(storageId)
	}
	err = os.Unsetenv(definitions.DLhrDatasetStarBoxAccessId)
	err = os.Unsetenv(definitions.DLhrDatasetStarBoxAccessKey)
	t.Log(err)
}

func TestStarBoxLoopWithoutEnvPara(t *testing.T) {
	api := NewStarBox("/Public/CICD/livetest/liveCICD/1_attack_3D/1_iMac_dark_whitebg")
	for filename, storageId := api.GetFirst(); !api.IsLast(); filename, storageId = api.GetNext() {
		t.Log(filename)
		t.Log(storageId)
	}
}

func TestStarBoxGetFile(t *testing.T) {
	api := NewStarBox("/Public/CICD/livetest/liveCICD/1_attack_3D/1_iMac_dark_whitebg")
	_, storageId := api.GetFirst()
	t.Log(api.GetFile(storageId))
}
