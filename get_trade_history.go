package cryptopia

import (
	"context"

	"github.com/pkg/errors"
)

type GetTradesHistory []GetTradeHistory

type GetTradeHistory struct {
	TradeId     int
	TradePairId int
	Market      string
	Type        string
	Rate        float64
	Amount      float64
	Total       float64
	Fee         float64
	TimeStamp   string
}

// Returns a list of trade history for all tradepairs or specified tradepair
func (c *Client) GetTradesHistory(ctx context.Context, market map[string]interface{}) (GetTradesHistory, error) {
	req, err := c.newAuthenticatedRequest(ctx, "GetTradeHistory", market)
	if err != nil {
		return GetTradesHistory{}, errors.Wrap(err, "Faild to new authenticated request")
	}

	var ret = &GetTradesHistory{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, errors.Wrap(err, "Faild to do request")
	}
	return *ret, nil
}
