package config

import "os"

var AccessSecret = os.Getenv("ACCESS_SECRET")
var RefreshSecret = os.Getenv("REFRESH_SECRET")
