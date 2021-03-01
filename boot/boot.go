package boot

import (
	"nfc-tags/app/service"
	"nfc-tags/library"
)

func init() {
	if err := service.Analysis.StandingRestImageUrl(); err != nil {
		library.AsyncPrintStack(err)
	}
}
