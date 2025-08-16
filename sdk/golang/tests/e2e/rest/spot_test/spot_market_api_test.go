package spot_test

import (
	"context"
	json "github.com/goccy/go-json"
	"fmt"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/market"
	"github.com/Kucoin/kucoin-universal-sdk/sdk/golang/pkg/types"
	"os"
	"testing"
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

	marketApi = client.RestService().GetSpotService().GetMarketAPI()
}

func TestMarketGetPrivateTokenReq(t *testing.T) {
	// GetPrivateToken
	// Get Private Token - Spot/Margin
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
	// Get Public Token - Spot/Margin
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

func TestMarketGetAllTickersReq(t *testing.T) {
	// GetAllTickers
	// Get All Tickers
	// /api/v1/market/allTickers

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

func TestMarketGetKlinesReq(t *testing.T) {
	// GetKlines
	// Get Klines
	// /api/v1/market/candles

	builder := market.NewGetKlinesReqBuilder()
	builder.SetSymbol("BTC-USDT").SetType("1min").SetStartAt(1566703297).SetEndAt(1566789757)
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

func TestMarketGetTradeHistoryReq(t *testing.T) {
	// GetTradeHistory
	// Get Trade History
	// /api/v1/market/histories

	builder := market.NewGetTradeHistoryReqBuilder()
	builder.SetSymbol("BTC-USDT")
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

func TestMarketGetTickerReq(t *testing.T) {
	// GetTicker
	// Get Ticker
	// /api/v1/market/orderbook/level1

	builder := market.NewGetTickerReqBuilder()
	builder.SetSymbol("BTC-USDT")
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

func TestMarketGetPartOrderBookReq(t *testing.T) {
	// GetPartOrderBook
	// Get Part OrderBook
	// /api/v1/market/orderbook/level2_{size}

	builder := market.NewGetPartOrderBookReqBuilder()
	builder.SetSymbol("BTC-USDT").SetSize("20")
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

func TestMarketGet24hrStatsReq(t *testing.T) {
	// Get24hrStats
	// Get 24hr Stats
	// /api/v1/market/stats

	builder := market.NewGet24hrStatsReqBuilder()
	builder.SetSymbol("BTC-USDT")
	req := builder.Build()

	resp, err := marketApi.Get24hrStats(req, context.TODO())
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

func TestMarketGetMarketListReq(t *testing.T) {
	// GetMarketList
	// Get Market List
	// /api/v1/markets

	resp, err := marketApi.GetMarketList(context.TODO())
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

func TestMarketGetFiatPriceReq(t *testing.T) {
	// GetFiatPrice
	// Get Fiat Price
	// /api/v1/prices

	builder := market.NewGetFiatPriceReqBuilder()
	builder.SetBase("USD").SetCurrencies("BTC,ETH")
	req := builder.Build()

	resp, err := marketApi.GetFiatPrice(req, context.TODO())
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

func TestMarketGetAllSymbolsReq(t *testing.T) {
	// GetAllSymbols
	// Get All Symbols
	// /api/v2/symbols

	builder := market.NewGetAllSymbolsReqBuilder()
	builder.SetMarket("USDS")
	req := builder.Build()

	resp, err := marketApi.GetAllSymbols(req, context.TODO())
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
	// /api/v2/symbols/{symbol}

	builder := market.NewGetSymbolReqBuilder()
	builder.SetSymbol("BTC-USDT")
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

func TestMarketGetAnnouncementsReq(t *testing.T) {
	// GetAnnouncements
	// Get Announcements
	// /api/v3/announcements

	builder := market.NewGetAnnouncementsReqBuilder()
	builder.SetAnnType("latest-announcements").SetLang("en_US")
	req := builder.Build()

	resp, err := marketApi.GetAnnouncements(req, context.TODO())
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

func TestMarketGetCurrencyReq(t *testing.T) {
	// GetCurrency
	// Get Currency
	// /api/v3/currencies/{currency}

	builder := market.NewGetCurrencyReqBuilder()
	builder.SetCurrency("BTC").SetChain("btc")
	req := builder.Build()

	resp, err := marketApi.GetCurrency(req, context.TODO())
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

func TestMarketGetAllCurrenciesReq(t *testing.T) {
	// GetAllCurrencies
	// Get All Currencies
	// /api/v3/currencies

	resp, err := marketApi.GetAllCurrencies(context.TODO())
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
	// /api/v3/market/orderbook/level2

	builder := market.NewGetFullOrderBookReqBuilder()
	builder.SetSymbol("BTC-USDT")
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

func TestMarketGetCallAuctionPartOrderBookReq(t *testing.T) {
	// GetCallAuctionPartOrderBook
	// Get Call Auction Part OrderBook
	// /api/v1/market/orderbook/callauction/level2_{size}
	builder := market.NewGetCallAuctionPartOrderBookReqBuilder()
	builder.SetSymbol("HBAR-USDC").SetSize("20")
	req := builder.Build()

	resp, err := marketApi.GetCallAuctionPartOrderBook(req, context.TODO())
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

func TestMarketGetCallAuctionInfoReq(t *testing.T) {
	// GetCallAuctionInfo
	// Get Call Auction Info
	// /api/v1/market/callauctionData

	builder := market.NewGetCallAuctionInfoReqBuilder()
	builder.SetSymbol("HBAR-USDC")
	req := builder.Build()

	resp, err := marketApi.GetCallAuctionInfo(req, context.TODO())
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

func TestMarketGetClientIPAddressReq(t *testing.T) {
	// GetClientIPAddress
	// Get Client IP Address
	// /api/v1/my-ip

	resp, err := marketApi.GetClientIPAddress(context.TODO())
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
