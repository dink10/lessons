package storage

import (
	"sync"
)

var Storage = sync.Map{}
