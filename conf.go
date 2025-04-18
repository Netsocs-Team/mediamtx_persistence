package main

type Path struct {
	Name           string `json:"name" yaml:"name"`
	Source         string `json:"source" yaml:"source"`
	SourceOnDemand bool   `json:"sourceOnDemand" yaml:"sourceOnDemand"`
}

type Config struct {
	LogLevel            string   `json:"logLevel" yaml:"logLevel"`
	LogDestinations     []string `json:"logDestinations" yaml:"logDestinations"`
	LogFile             string   `json:"logFile" yaml:"logFile"`
	ReadTimeout         string   `json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout        string   `json:"writeTimeout" yaml:"writeTimeout"`
	WriteQueueSize      int      `json:"writeQueueSize" yaml:"writeQueueSize"`
	UDPMaxPayloadSize   int      `json:"udpMaxPayloadSize" yaml:"udpMaxPayloadSize"`
	RunOnConnect        string   `json:"runOnConnect" yaml:"runOnConnect"`
	RunOnConnectRestart bool     `json:"runOnConnectRestart" yaml:"runOnConnectRestart"`
	RunOnDisconnect     string   `json:"runOnDisconnect" yaml:"runOnDisconnect"`
	AuthMethod          string   `json:"authMethod" yaml:"authMethod"`
	AuthInternalUsers   []struct {
		User        string        `json:"user" yaml:"user"`
		Pass        string        `json:"pass" yaml:"pass"`
		Ips         []interface{} `json:"ips" yaml:"ips"`
		Permissions []struct {
			Action string `json:"action" yaml:"action"`
			Path   string `json:"path" yaml:"path"`
		} `json:"permissions" yaml:"permissions"`
	} `json:"authInternalUsers" yaml:"authInternalUsers"`
	AuthHTTPAddress string `json:"authHTTPAddress" yaml:"authHTTPAddress"`
	AuthHTTPExclude []struct {
		Action string `json:"action" yaml:"action"`
		Path   string `json:"path" yaml:"path"`
	} `json:"authHTTPExclude" yaml:"authHTTPExclude"`
	AuthJWTJWKS                 string          `json:"authJWTJWKS" yaml:"authJWTJWKS"`
	AuthJWTClaimKey             string          `json:"authJWTClaimKey" yaml:"authJWTClaimKey"`
	API                         bool            `json:"api" yaml:"api"`
	APIAddress                  string          `json:"apiAddress" yaml:"apiAddress"`
	APIEncryption               bool            `json:"apiEncryption" yaml:"apiEncryption"`
	APIServerKey                string          `json:"apiServerKey" yaml:"apiServerKey"`
	APIServerCert               string          `json:"apiServerCert" yaml:"apiServerCert"`
	APIAllowOrigin              string          `json:"apiAllowOrigin" yaml:"apiAllowOrigin"`
	APITrustedProxies           []interface{}   `json:"apiTrustedProxies" yaml:"apiTrustedProxies"`
	Metrics                     bool            `json:"metrics" yaml:"metrics"`
	MetricsAddress              string          `json:"metricsAddress" yaml:"metricsAddress"`
	MetricsEncryption           bool            `json:"metricsEncryption" yaml:"metricsEncryption"`
	MetricsServerKey            string          `json:"metricsServerKey" yaml:"metricsServerKey"`
	MetricsServerCert           string          `json:"metricsServerCert" yaml:"metricsServerCert"`
	MetricsAllowOrigin          string          `json:"metricsAllowOrigin" yaml:"metricsAllowOrigin"`
	MetricsTrustedProxies       []interface{}   `json:"metricsTrustedProxies" yaml:"metricsTrustedProxies"`
	Pprof                       bool            `json:"pprof" yaml:"pprof"`
	PprofAddress                string          `json:"pprofAddress" yaml:"pprofAddress"`
	PprofEncryption             bool            `json:"pprofEncryption" yaml:"pprofEncryption"`
	PprofServerKey              string          `json:"pprofServerKey" yaml:"pprofServerKey"`
	PprofServerCert             string          `json:"pprofServerCert" yaml:"pprofServerCert"`
	PprofAllowOrigin            string          `json:"pprofAllowOrigin" yaml:"pprofAllowOrigin"`
	PprofTrustedProxies         []interface{}   `json:"pprofTrustedProxies" yaml:"pprofTrustedProxies"`
	Playback                    bool            `json:"playback" yaml:"playback"`
	PlaybackAddress             string          `json:"playbackAddress" yaml:"playbackAddress"`
	PlaybackEncryption          bool            `json:"playbackEncryption" yaml:"playbackEncryption"`
	PlaybackServerKey           string          `json:"playbackServerKey" yaml:"playbackServerKey"`
	PlaybackServerCert          string          `json:"playbackServerCert" yaml:"playbackServerCert"`
	PlaybackAllowOrigin         string          `json:"playbackAllowOrigin" yaml:"playbackAllowOrigin"`
	PlaybackTrustedProxies      []interface{}   `json:"playbackTrustedProxies" yaml:"playbackTrustedProxies"`
	Rtsp                        bool            `json:"rtsp" yaml:"rtsp"`
	RtspTransports              []string        `json:"rtspTransports" yaml:"rtspTransports"`
	RtspEncryption              string          `json:"rtspEncryption" yaml:"rtspEncryption"`
	RtspAddress                 string          `json:"rtspAddress" yaml:"rtspAddress"`
	RtspsAddress                string          `json:"rtspsAddress" yaml:"rtspsAddress"`
	RtpAddress                  string          `json:"rtpAddress" yaml:"rtpAddress"`
	RtcpAddress                 string          `json:"rtcpAddress" yaml:"rtcpAddress"`
	MulticastIPRange            string          `json:"multicastIPRange" yaml:"multicastIPRange"`
	MulticastRTPPort            int             `json:"multicastRTPPort" yaml:"multicastRTPPort"`
	MulticastRTCPPort           int             `json:"multicastRTCPPort" yaml:"multicastRTCPPort"`
	RtspServerKey               string          `json:"rtspServerKey" yaml:"rtspServerKey"`
	RtspServerCert              string          `json:"rtspServerCert" yaml:"rtspServerCert"`
	RtspAuthMethods             []string        `json:"rtspAuthMethods" yaml:"rtspAuthMethods"`
	Rtmp                        bool            `json:"rtmp" yaml:"rtmp"`
	RtmpAddress                 string          `json:"rtmpAddress" yaml:"rtmpAddress"`
	RtmpEncryption              string          `json:"rtmpEncryption" yaml:"rtmpEncryption"`
	RtmpsAddress                string          `json:"rtmpsAddress" yaml:"rtmpsAddress"`
	RtmpServerKey               string          `json:"rtmpServerKey" yaml:"rtmpServerKey"`
	RtmpServerCert              string          `json:"rtmpServerCert" yaml:"rtmpServerCert"`
	Hls                         bool            `json:"hls" yaml:"hls"`
	HlsAddress                  string          `json:"hlsAddress" yaml:"hlsAddress"`
	HlsEncryption               bool            `json:"hlsEncryption" yaml:"hlsEncryption"`
	HlsServerKey                string          `json:"hlsServerKey" yaml:"hlsServerKey"`
	HlsServerCert               string          `json:"hlsServerCert" yaml:"hlsServerCert"`
	HlsAllowOrigin              string          `json:"hlsAllowOrigin" yaml:"hlsAllowOrigin"`
	HlsTrustedProxies           []interface{}   `json:"hlsTrustedProxies" yaml:"hlsTrustedProxies"`
	HlsAlwaysRemux              bool            `json:"hlsAlwaysRemux" yaml:"hlsAlwaysRemux"`
	HlsVariant                  string          `json:"hlsVariant" yaml:"hlsVariant"`
	HlsSegmentCount             int             `json:"hlsSegmentCount" yaml:"hlsSegmentCount"`
	HlsSegmentDuration          string          `json:"hlsSegmentDuration" yaml:"hlsSegmentDuration"`
	HlsPartDuration             string          `json:"hlsPartDuration" yaml:"hlsPartDuration"`
	HlsSegmentMaxSize           string          `json:"hlsSegmentMaxSize" yaml:"hlsSegmentMaxSize"`
	HlsDirectory                string          `json:"hlsDirectory" yaml:"hlsDirectory"`
	HlsMuxerCloseAfter          string          `json:"hlsMuxerCloseAfter" yaml:"hlsMuxerCloseAfter"`
	Webrtc                      bool            `json:"webrtc" yaml:"webrtc"`
	WebrtcAddress               string          `json:"webrtcAddress" yaml:"webrtcAddress"`
	WebrtcEncryption            bool            `json:"webrtcEncryption" yaml:"webrtcEncryption"`
	WebrtcServerKey             string          `json:"webrtcServerKey" yaml:"webrtcServerKey"`
	WebrtcServerCert            string          `json:"webrtcServerCert" yaml:"webrtcServerCert"`
	WebrtcAllowOrigin           string          `json:"webrtcAllowOrigin" yaml:"webrtcAllowOrigin"`
	WebrtcTrustedProxies        []interface{}   `json:"webrtcTrustedProxies" yaml:"webrtcTrustedProxies"`
	WebrtcLocalUDPAddress       string          `json:"webrtcLocalUDPAddress" yaml:"webrtcLocalUDPAddress"`
	WebrtcLocalTCPAddress       string          `json:"webrtcLocalTCPAddress" yaml:"webrtcLocalTCPAddress"`
	WebrtcIPsFromInterfaces     bool            `json:"webrtcIPsFromInterfaces" yaml:"webrtcIPsFromInterfaces"`
	WebrtcIPsFromInterfacesList []interface{}   `json:"webrtcIPsFromInterfacesList" yaml:"webrtcIPsFromInterfacesList"`
	WebrtcAdditionalHosts       []interface{}   `json:"webrtcAdditionalHosts" yaml:"webrtcAdditionalHosts"`
	WebrtcICEServers2           []interface{}   `json:"webrtcICEServers2" yaml:"webrtcICEServers2"`
	WebrtcHandshakeTimeout      string          `json:"webrtcHandshakeTimeout" yaml:"webrtcHandshakeTimeout"`
	WebrtcTrackGatherTimeout    string          `json:"webrtcTrackGatherTimeout" yaml:"webrtcTrackGatherTimeout"`
	WebrtcSTUNGatherTimeout     string          `json:"webrtcSTUNGatherTimeout" yaml:"webrtcSTUNGatherTimeout"`
	Srt                         bool            `json:"srt" yaml:"srt"`
	SrtAddress                  string          `json:"srtAddress" yaml:"srtAddress"`
	Paths                       map[string]Path `json:"paths" yaml:"paths"`
}
