package routes

import (
	"net/http"
	"os"
	"testing"

	"github.com/fredrik-hjarner/ztorage/diskv"
)

var HTTPHandler http.Handler

func TestMain(m *testing.M) {
	diskv.Diskv.EraseAll()
	HTTPHandler = CreateHTTPHandler()
	code := m.Run()
	os.Exit(code)
}

func SetupFixture() {
	diskv.Diskv.EraseAll()
}
