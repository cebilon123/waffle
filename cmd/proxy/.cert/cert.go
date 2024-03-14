package cert

import (
"embed"
 _ "embed"
)

//go:embed *
var Certificates embed.FS
