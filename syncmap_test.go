package syncmap

import (
	"testing"
)

var syncMap SyncMap[string, string]

func TestSyncMap_Make(t *testing.T) {
	syncMap.Make()
	t.Log("created a map")
}

func TestSyncMap_Add(t *testing.T) {
	t.Log("past")
	syncMap.Add("key", "value")
	t.Log("added key")
}

func TestSyncMap_Get(t *testing.T) {
	t.Logf("found in key: %s\n", syncMap.Get("key"))
}

func TestSyncMap_Del(t *testing.T) {
	syncMap.Del("key")
	t.Log("deleted key")
}

func TestSyncMap_Push(t *testing.T) {
	syncMap.Push("key", "value")
	t.Log("pushed key")
}

func TestSyncMap_Get2(t *testing.T) {
	t.Logf("found in key: %s\n", syncMap.Get("key"))
}
