package futures_test

import (
	"context"
	json "github.com/goccy/go-json"
	"fmt"
	"os"
	"testing"

	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/market"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var marketApi market.MarketAPI

func init() {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		Build()

	client := api.NewClient(option)

	marketApi = client.RestService().GetFuturesService().GetMarketAPI()
}

func TestMarketGetAllTickersReq(t *testing.T) {
	// GetAllTickers
	// Get All Tickers
	// /api/v1/allTickers

	resp, err := marketApi.GetAllTickers(context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetPrivateTokenReq(t *testing.T) {
	// GetPrivateToken
	// Get Private Token - Futures
	// /api/v1/bullet-private

	resp, err := marketApi.GetPrivateToken(context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetPublicTokenReq(t *testing.T) {
	// GetPublicToken
	// Get Public Token - Futures
	// /api/v1/bullet-public

	resp, err := marketApi.GetPublicToken(context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetAllSymbolsReq(t *testing.T) {
	// GetAllSymbols
	// Get All Symbols
	// /api/v1/contracts/active

	resp, err := marketApi.GetAllSymbols(context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetSymbolReq(t *testing.T) {
	// GetSymbol
	// Get Symbol
	// /api/v1/contracts/{symbol}

	builder := market.NewGetSymbolReqBuilder()
	builder.SetSymbol("XBTUSDM")
	req := builder.Build()

	resp, err := marketApi.GetSymbol(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetSpotIndexPriceReq(t *testing.T) {
	// GetSpotIndexPrice
	// Get Spot Index Price
	// /api/v1/index/query

	builder := market.NewGetSpotIndexPriceReqBuilder()
	builder.SetSymbol(".KXBTUSDT").SetStartAt(1732464000000).SetEndAt(1732521600000).SetMaxCount(10)
	req := builder.Build()

	resp, err := marketApi.GetSpotIndexPrice(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetInterestRateIndexReq(t *testing.T) {
	// GetInterestRateIndex
	// Get Interest Rate Index
	// /api/v1/interest/query

	builder := market.NewGetInterestRateIndexReqBuilder()
	builder.SetSymbol(".XBTINT").SetStartAt(1732464000000).SetEndAt(1732521600000)
	req := builder.Build()

	resp, err := marketApi.GetInterestRateIndex(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetKlinesReq(t *testing.T) {
	// GetKlines
	// Get Klines
	// /api/v1/kline/query

	builder := market.NewGetKlinesReqBuilder()
	builder.SetSymbol("DOGEUSDTM").SetGranularity(1).SetFrom(1746028800000).SetTo(1747989311786)
	req := builder.Build()

	resp, err := marketApi.GetKlines(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetPartOrderBookReq(t *testing.T) {
	// GetPartOrderBook
	// Get Part OrderBook
	// /api/v1/level2/depth{size}

	builder := market.NewGetPartOrderBookReqBuilder()
	builder.SetSymbol("XBTUSDTM").SetSize("100")
	req := builder.Build()

	resp, err := marketApi.GetPartOrderBook(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetFullOrderBookReq(t *testing.T) {
	// GetFullOrderBook
	// Get Full OrderBook
	// /api/v1/level2/snapshot

	builder := market.NewGetFullOrderBookReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := marketApi.GetFullOrderBook(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetMarkPriceReq(t *testing.T) {
	// GetMarkPrice
	// Get Mark Price
	// /api/v1/mark-price/{symbol}/current

	builder := market.NewGetMarkPriceReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := marketApi.GetMarkPrice(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetPremiumIndexReq(t *testing.T) {
	// GetPremiumIndex
	// Get Premium Index
	// /api/v1/premium/query

	builder := market.NewGetPremiumIndexReqBuilder()
	builder.SetSymbol(".XBTUSDTMPI").SetStartAt(1732464000000).SetEndAt(1732521600000)
	req := builder.Build()

	resp, err := marketApi.GetPremiumIndex(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetServiceStatusReq(t *testing.T) {
	// GetServiceStatus
	// Get Service Status
	// /api/v1/status

	resp, err := marketApi.GetServiceStatus(context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetTickerReq(t *testing.T) {
	// GetTicker
	// Get Ticker
	// /api/v1/ticker

	builder := market.NewGetTickerReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := marketApi.GetTicker(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetServerTimeReq(t *testing.T) {
	// GetServerTime
	// Get Server Time
	// /api/v1/timestamp

	resp, err := marketApi.GetServerTime(context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGetTradeHistoryReq(t *testing.T) {
	// GetTradeHistory
	// Get Trade History
	// /api/v1/trade/history

	builder := market.NewGetTradeHistoryReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := marketApi.GetTradeHistory(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestMarketGet24hrStatsReq(t *testing.T) {
	// Get24hrStats
	// Get 24hr Stats
	// /api/v1/trade-statistics

	resp, err := marketApi.Get24hrStats(context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}
