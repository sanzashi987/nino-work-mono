package enums

import (
	"fmt"

	"github.com/cza14h/nino-work/apps/canvas-pro/utils"
)

type AssetPrefix struct{}

const (
	PREFIX        = "cVs"
	PROJECT       = "A"
	BLOCK         = "B"
	DESIGN        = "C"
	FONT          = "D"
	COMPONENT     = "E"
	DATASOURCE    = "F"
	STATIC_SOURCE = "H"
)

func CreateCode(cat string) string {
	return fmt.Sprintf("%s%s%s", PREFIX, cat, utils.GetRandomId())
}
