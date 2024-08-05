package proxies

import "github.com/AstroStreakNet/telescope/astrometry"

type Astrometry struct {
	client *astrometry.Client
}

// NewAstrometryProxy constructor for Astrometry proxy
func NewAstrometryProxy(apiKey string) *Astrometry {
	return &Astrometry{
		astrometry.NewAstrometryClient(apiKey),
	}
}
