package config

// URLs (Replaced map with constants)
const (
	API_URL              = "https://api.icicidirect.com/breezeapi/api/v1/"
	BREEZE_NEW_URL       = "https://breezeapi.icicidirect.com/api/v2/"
	LIVE_FEEDS_URL       = "https://livefeeds.icicidirect.com"
	LIVE_STREAM_URL      = "https://livestream.icicidirect.com"
	LIVE_OHLC_STREAM_URL = "https://breezeapi.icicidirect.com"
	SECURITY_MASTER_URL  = "https://directlink.icicidirect.com/NewSecurityMaster/SecurityMaster.zip"
	STOCK_SCRIPT_CSV_URL = "https://traderweb.icicidirect.com/Content/File/txtFile/ScripFile/StockScriptNew.csv"
)

// API Request Methods (Enum-style constants)
type APIRequestMethod string

const (
	HTTP_POST   APIRequestMethod = "POST"
	HTTP_GET    APIRequestMethod = "GET"
	HTTP_PUT    APIRequestMethod = "PUT"
	HTTP_DELETE APIRequestMethod = "DELETE"
)

// API Endpoints
type APIEndpoint string

const (
	ENDPOINT_CUST_DETAILS       APIEndpoint = "customerdetails"
	ENDPOINT_DEMAT_HOLDING      APIEndpoint = "dematholdings"
	ENDPOINT_FUND               APIEndpoint = "funds"
	ENDPOINT_HIST_CHART         APIEndpoint = "historicalcharts"
	ENDPOINT_MARGIN             APIEndpoint = "margin"
	ENDPOINT_ORDER              APIEndpoint = "order"
	ENDPOINT_PORTFOLIO_HOLDING  APIEndpoint = "portfolioholdings"
	ENDPOINT_PORTFOLIO_POSITION APIEndpoint = "portfoliopositions"
	ENDPOINT_QUOTE              APIEndpoint = "quotes"
	ENDPOINT_TRADE              APIEndpoint = "trades"
	ENDPOINT_OPT_CHAIN          APIEndpoint = "optionchain"
	ENDPOINT_SQUARE_OFF         APIEndpoint = "squareoff"
	ENDPOINT_PREVIEW_ORDER      APIEndpoint = "preview_order"
	ENDPOINT_LIMITCALCULATOR    APIEndpoint = "fnolmtpriceandqtycal"
	ENDPOINT_MARGINCALCULATOR   APIEndpoint = "margincalculator"
)

// Exception Messages
type ExceptionMessage string

const (
	AUTHENTICATION_EXCEPTION      ExceptionMessage = "could not authenticate credentials. please check token and keys"
	QUOTE_DEPTH_EXCEPTION         ExceptionMessage = "either getExchangeQuotes must be true or getMarketDepth must be true"
	EXCHANGE_CODE_EXCEPTION       ExceptionMessage = "exchange code allowed are 'bse', 'nse', 'ndx', 'mcx' or 'nfo'."
	EMPTY_STOCK_CODE_EXCEPTION    ExceptionMessage = "stock-code cannot be empty."
	EXPIRY_DATE_EXCEPTION         ExceptionMessage = "expiry-date cannot be empty for given exchange-code."
	PRODUCT_TYPE_EXCEPTION        ExceptionMessage = "product-type should either be futures or options for given exchange-code."
	STRIKE_PRICE_EXCEPTION        ExceptionMessage = "strike price cannot be empty for product-type 'options'."
	RIGHT_EXCEPTION               ExceptionMessage = "rights should either be put or call for product-type 'options'."
	STOCK_INVALID_EXCEPTION       ExceptionMessage = "stock-code not found."
	WRONG_EXCHANGE_CODE_EXCEPTION ExceptionMessage = "stock-token cannot be found due to wrong exchange-code."
	STOCK_NOT_EXIST_EXCEPTION     ExceptionMessage = "stock-data does not exist in exchange-code {0} for stock-token {1}."
	ISEC_NSE_STOCK_MAP_EXCEPTION  ExceptionMessage = "result not found"
	STREAM_OHLC_INTERVAL_ERROR    ExceptionMessage = "interval should be either '1second','1minute', '5minute', '30minute'"
	API_REQUEST_EXCEPTION         ExceptionMessage = "error while trying to make request {0} {1}"
)

// Response Messages
type ResponseMessage string

const (
	BLANK_EXCHANGE_CODE ResponseMessage = "Exchange-Code cannot be empty"
	BLANK_STOCK_CODE    ResponseMessage = "Stock-Code cannot be empty"
	BLANK_PRODUCT_TYPE  ResponseMessage = "Product cannot be empty"
	BLANK_ACTION        ResponseMessage = "Action cannot be empty"
	BLANK_ORDER_TYPE    ResponseMessage = "Order-type cannot be empty"
	BLANK_QUANTITY      ResponseMessage = "Quantity cannot be empty"
	BLANK_VALIDITY      ResponseMessage = "Validity cannot be empty"
	ZERO_AMOUNT_ERROR   ResponseMessage = "Amount should be more than 0"
	EXCHANGE_CODE_ERROR ResponseMessage = "Exchange-Code should be either 'nse', or 'nfo'"
	API_SESSION_ERROR   ResponseMessage = "API Session cannot be empty"
)

// Room Names
type RoomName string

const (
	ONE_CLICK_ROOM RoomName = "one_click_fno"
	I_CLICK_2_GAIN RoomName = "i_click_2_gain"
)

// Feed Interval Map
type FeedInterval string

const (
	INTERVAL_1MIN  FeedInterval = "1minute"
	INTERVAL_5MIN  FeedInterval = "5minute"
	INTERVAL_30MIN FeedInterval = "30minute"
	INTERVAL_1SEC  FeedInterval = "1second"
)

// Channel Interval Map
type ChannelInterval string

const (
	CHANNEL_1MIN  ChannelInterval = "1MIN"
	CHANNEL_5MIN  ChannelInterval = "5MIN"
	CHANNEL_30MIN ChannelInterval = "30MIN"
	CHANNEL_1SEC  ChannelInterval = "1SEC"
)
