package application

import (
	"context"
	"fizz/internal/core/application/dto"
	"fizz/internal/core/port/outgoing"
)

type MarketingService struct {
	emailStats outgoing.EmailStats
}

func NewMarketingService(emailStats outgoing.EmailStats) *MarketingService {
	return &MarketingService{emailStats: emailStats}
}

func (m MarketingService) GetCompleteStats(ctx context.Context) ([]dto.Stats, error) {
	stats, err := m.emailStats.GetStats(ctx)
	if err != nil {
		return nil, err
	}

	var out []dto.Stats

	for _, st := range stats {
		out = append(out, dto.Stats{
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

