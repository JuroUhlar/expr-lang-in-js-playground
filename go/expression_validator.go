package main

import (
	"time"

	"github.com/expr-lang/expr"
)

// Environment structs mirroring the JSON shape --------------------------------

type Env struct {
	Products Products `expr:"products"`
}

type Products struct {
	Identification IdentificationModule `expr:"identification"`
	Botd          BotdModule          `expr:"botd"`
	RootApps      ResultBoolModule    `expr:"rootApps"`
	Emulator      ResultBoolModule    `expr:"emulator"`
	IpInfo        IpInfoModule        `expr:"ipInfo"`
	IpBlocklist   IpBlocklistModule   `expr:"ipBlocklist"`
	Tor           ResultBoolModule    `expr:"tor"`
	Vpn           VpnModule           `expr:"vpn"`
	Proxy         ProxyModule         `expr:"proxy"`
	Incognito     ResultBoolModule    `expr:"incognito"`
	Tampering     TamperingModule     `expr:"tampering"`
	ClonedApp     ResultBoolModule    `expr:"clonedApp"`
	FactoryReset  FactoryResetModule  `expr:"factoryReset"`
	Jailbroken    ResultBoolModule    `expr:"jailbroken"`
	Frida         ResultBoolModule    `expr:"frida"`
	PrivacySettings ResultBoolModule  `expr:"privacySettings"`
	VirtualMachine ResultBoolModule   `expr:"virtualMachine"`
	HighActivity   ResultBoolModule   `expr:"highActivity"`
	LocationSpoofing ResultBoolModule `expr:"locationSpoofing"`
	SuspectScore  SuspectScoreModule  `expr:"suspectScore"`
	RemoteControl ResultBoolModule    `expr:"remoteControl"`
	DeveloperTools ResultBoolModule   `expr:"developerTools"`
	MitmAttack    ResultBoolModule    `expr:"mitmAttack"`
}

// Generic modules -------------------------------------------------------------

type ResultBoolModule struct {
	Data struct {
		Result bool `expr:"result"`
	} `expr:"data"`
}

type ResultIntModule struct {
	Data struct {
		Result int `expr:"result"`
	} `expr:"data"`
}

// Specific modules ------------------------------------------------------------

type BotdModule struct {
	Data struct {
		Bot struct {
			Result string `expr:"result"`
		} `expr:"bot"`
		LinkedId  string `expr:"linkedId"`
		URL       string `expr:"url"`
		Ip        string `expr:"ip"`
		Time      string `expr:"time"`
		UserAgent string `expr:"userAgent"`
		RequestId string `expr:"requestId"`
	} `expr:"data"`
}

type IdentificationModule struct {
	Data IdentificationData `expr:"data"`
}

type IdentificationData struct {
	VisitorId      string         `expr:"visitorId"`
	RequestId      string         `expr:"requestId"`
	BrowserDetails BrowserDetails `expr:"browserDetails"`
	Incognito      bool           `expr:"incognito"`
	Ip             string         `expr:"ip"`
	IpLocation     IpLocation     `expr:"ipLocation"`
	LinkedId       string         `expr:"linkedId"`
	Timestamp      int64          `expr:"timestamp"`
	Time           string         `expr:"time"`
	URL            string         `expr:"url"`
	Tag            map[string]any `expr:"tag"`
	Confidence     Confidence     `expr:"confidence"`
	VisitorFound   bool           `expr:"visitorFound"`
	FirstSeenAt    SeenAt         `expr:"firstSeenAt"`
	LastSeenAt     SeenAt         `expr:"lastSeenAt"`
	Replayed       bool           `expr:"replayed"`
	SDK            SDK            `expr:"sdk"`
	EnvironmentId  string         `expr:"environmentId"`
}

type BrowserDetails struct {
	BrowserName        string `expr:"browserName"`
	BrowserMajorVersion string `expr:"browserMajorVersion"`
	BrowserFullVersion string `expr:"browserFullVersion"`
	Os                 string `expr:"os"`
	OsVersion          string `expr:"osVersion"`
	Device             string `expr:"device"`
	UserAgent          string `expr:"userAgent"`
}

type IpLocation struct {
	AccuracyRadius int     `expr:"accuracyRadius"`
	Latitude       float64 `expr:"latitude"`
	Longitude      float64 `expr:"longitude"`
	PostalCode     string  `expr:"postalCode"`
	Timezone       string  `expr:"timezone"`
	City           struct {
		Name string `expr:"name"`
	} `expr:"city"`
	Country struct {
		Code string `expr:"code"`
		Name string `expr:"name"`
	} `expr:"country"`
	Continent struct {
		Code string `expr:"code"`
		Name string `expr:"name"`
	} `expr:"continent"`
	Subdivisions []struct {
		IsoCode string `expr:"isoCode"`
		Name    string `expr:"name"`
	} `expr:"subdivisions"`
}

type Confidence struct {
	Score    float64 `expr:"score"`
	Revision string  `expr:"revision"`
}

type SeenAt struct {
	Global        string `expr:"global"`
	Subscription  string `expr:"subscription"`
}

type SDK struct {
	Platform string `expr:"platform"`
	Version  string `expr:"version"`
}

// IP info ---------------------------------------------------------------------

type IpInfoModule struct {
	Data struct {
		V4 struct {
			Address     string     `expr:"address"`
			Geolocation IpLocation `expr:"geolocation"`
			Asn struct {
				Asn     string `expr:"asn"`
				Name    string `expr:"name"`
				Network string `expr:"network"`
			} `expr:"asn"`
			Datacenter struct {
				Result bool   `expr:"result"`
				Name   string `expr:"name"`
			} `expr:"datacenter"`
		} `expr:"v4"`
	} `expr:"data"`
}

// IpBlocklist -----------------------------------------------------------------

type IpBlocklistModule struct {
	Data struct {
		Result  bool `expr:"result"`
		Details struct {
			EmailSpam    bool `expr:"emailSpam"`
			AttackSource bool `expr:"attackSource"`
		} `expr:"details"`
	} `expr:"data"`
}

// VPN -------------------------------------------------------------------------

type VpnModule struct {
	Data struct {
		Result           bool   `expr:"result"`
		Confidence       string `expr:"confidence"`
		OriginTimezone   string `expr:"originTimezone"`
		OriginCountry    string `expr:"originCountry"`
		Methods struct {
			TimezoneMismatch bool `expr:"timezoneMismatch"`
			PublicVPN        bool `expr:"publicVPN"`
			AuxiliaryMobile  bool `expr:"auxiliaryMobile"`
			OsMismatch       bool `expr:"osMismatch"`
			Relay            bool `expr:"relay"`
		} `expr:"methods"`
	} `expr:"data"`
}

// Proxy similarly -------------------------------------------------------------

type ProxyModule struct {
	Data struct {
		Result     bool   `expr:"result"`
		Confidence string `expr:"confidence"`
	} `expr:"data"`
}

// Tampering -------------------------------------------------------------------

type TamperingModule struct {
	Data struct {
		Result          bool  `expr:"result"`
		AnomalyScore    int   `expr:"anomalyScore"`
		AntiDetectBrowser bool `expr:"antiDetectBrowser"`
	} `expr:"data"`
}

// FactoryReset ----------------------------------------------------------------

type FactoryResetModule struct {
	Data struct {
		Time      time.Time `expr:"time"`
		Timestamp int64     `expr:"timestamp"`
	} `expr:"data"`
}

// SuspectScore ----------------------------------------------------------------

type SuspectScoreModule struct {
	Data struct {
		Result int `expr:"result"`
	} `expr:"data"`
}

// ----------------------------------------------------------------------------

func ValidateExpression(expression string) bool {
	_, err := expr.Compile(expression, expr.Env(Env{}))
	return err == nil
}

// isExpressionValid is a thin wrapper suitable for JS/WASM exports and tests.
func isExpressionValid(expression string) bool {
	return ValidateExpression(expression)
}
