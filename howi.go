package howi

import "github.com/okramlabs/howi/lib/metadata"

// New creates new howi application instance
func New() *HOWI {
	return &HOWI{}
}

// HOWI Application wrapper
type HOWI struct {
	metadata *metadata.Basic
}

// Meta [creates] returns metadata pointer
func (h *HOWI) Meta() *metadata.Basic {
	if h.metadata == nil {
		h.metadata = &metadata.Basic{}
	}
	return h.metadata
}

// CLI [creates] returns application command-line interface instance
func (h *HOWI) CLI() {

}
