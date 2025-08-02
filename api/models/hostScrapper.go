package models

type HostScrapper int

// HostScrapper list of the host that can be scraped
const (
	RemoteOk HostScrapper = iota
	WorkRemotely
)

func (h HostScrapper) String() string {
	switch h {
	case RemoteOk:
		return "RemoteOk"
	case WorkRemotely:
		return "WorkRemotely"
	default:
		return "Unknown"
	}
}
