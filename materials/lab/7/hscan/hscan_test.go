// Optional Todo
package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	GuessSingle("36cb3251dfb3d2b4a559796498a2ac29", "../main/misc-dictionary.txt") // Currently function returns only number of open ports
}

func TestGenHashMaps(t *testing.T) {
	GenHashMaps("../main/misc-dictionary.txt")
	got := md5lookup
	// test will be pass
	if got["36cb3251dfb3d2b4a559796498a2ac29"] != "drmike1" {
		t.Errorf("got %q want %q", got, "36cb3251dfb3d2b4a559796498a2ac29")
	}
	// test will be fail
	if got["36cb3251dfb3d2b4a559796498a2ac29"] != "moon1" {
		t.Errorf("got %q want %q", got, "36cb3251dfb3d2b4a559796498a2ac29")
	}

}
