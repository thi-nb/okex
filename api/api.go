package api

import (
	"context"

	"github.com/thi-nb/okex"
	"github.com/thi-nb/okex/api/rest"
	"github.com/thi-nb/okex/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okex.Destination) (*Client, error) {
	restURL := okex.RestURL
	wsPubURL := okex.PublicWsURL
	wsPriURL := okex.PrivateWsURL
	switch destination {
	case okex.AwsServer:
		restURL = okex.AwsRestURL
		wsPubURL = okex.AwsPublicWsURL
		wsPriURL = okex.AwsPrivateWsURL
	case okex.DemoServer:
		restURL = okex.DemoRestURL
		wsPubURL = okex.DemoPublicWsURL
		wsPriURL = okex.DemoPrivateWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}

// func main() {
// 	apiKey := "e4a22e10-d653-4e5c-952e-72e03eb388c0"
// 	secretKey := "DBF20D80E12D82FF9E7EDF0B82AE7C4E"
// 	passphrase := "BYD7B88dyVkTQyh@"
// 	dest := okex.NormalServer
// 	cli, err := NewClient(context.Background(), apiKey, secretKey, passphrase, dest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	r, err := cli.Rest.Market.GetCandlesticks(market.GetCandlesticks{
// 		InstID: "BTC-USDT",
// 		Bar:    okex.Bar1H,
// 		Limit:  168,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	//fmt.Printf("kline: %+v\n", r)
// 	for _, candle := range r.Candles {
// 		fmt.Printf("open %+v\n", candle)
// 	}
// 	//r,err := cli.Rest.PublicData.GetInstruments()
// 	pCh := make(chan *event_public.MarkPriceCandlesticks)
// 	cli.Ws.Public.MarkPriceCandlesticks(ws_public.MarkPriceCandlesticks{
// 		InstID:  "BTC-USDT",
// 		Channel: okex.CandleStick1H,
// 	}, pCh)

// 	for {
// 		select {
// 		case p := <-pCh:
// 			fmt.Printf("price %+v\n", p.Prices)
// 		}
// 	}
// }
