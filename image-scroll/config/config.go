package config

import "time"

var PORT = ":8000"
var PATH = "./aa"

var TICKER_FOLDER_CHECK = 1 * time.Second
var TICKER_SEND_NEW_IMAGE = 5 * time.Second
