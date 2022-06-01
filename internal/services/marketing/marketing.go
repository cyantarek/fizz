package marketing

import (
	"context"
	"github.com/cyantarek/fizz/internal/domains"
)

type marketing struct {
	emailStats EmailStats
}

type EmailStats interface {
	GetStats(ctx context.Context) ([]domains.Stats, error)
}

func NewMarketingService(emailStats EmailStats) *marketing {
	return &marketing{emailStats: emailStats}
}

func (m marketing) GetCompleteStats(ctx context.Context) ([]Stats, error) {
	stats, err := m.emailStats.GetStats(ctx)
	if err != nil {
		return nil, err
	}

	var out []Stats

	for _, st := range stats {
		out = append(out, Stats{
			Accepted:     st.Accepted,
			Delivered:    st.Delivered,
			Failed:       st.Failed,
			Stored:       st.Stored,
			Opened:       st.Opened,
			Clicked:      st.Clicked,
			Unsubscribed: st.Unsubscribed,
			Complained:   st.Complained,
		})
	}

	return out, nil
}

type Stats struct {
	Accepted     int
	Delivered    int
	Failed       int
	Stored       int
	Opened       int
	Clicked      int
	Unsubscribed int
	Complained   int
}
