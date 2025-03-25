package main

type PathResponse struct {
	ItemCount int                `json:"itemCount"`
	PageCount int                `json:"pageCount"`
	Items     []PathResponseItem `json:"items"`
}

type PathResponseItem struct {
	Name                       string `json:"name"`
	Source                     string `json:"source"`
	SourceFingerprint          string `json:"sourceFingerprint"`
	SourceOnDemand             bool   `json:"sourceOnDemand"`
	SourceOnDemandStartTimeout string `json:"sourceOnDemandStartTimeout"`
	SourceOnDemandCloseAfter   string `json:"sourceOnDemandCloseAfter"`
	MaxReaders                 int    `json:"maxReaders"`
	SrtReadPassphrase          string `json:"srtReadPassphrase"`
	Fallback                   string `json:"fallback"`
	Record                     bool   `json:"record"`
	RecordPath                 string `json:"recordPath"`
	RecordFormat               string `json:"recordFormat"`
	RecordPartDuration         string `json:"recordPartDuration"`
	RecordSegmentDuration      string `json:"recordSegmentDuration"`
	RecordDeleteAfter          string `json:"recordDeleteAfter"`
	OverridePublisher          bool   `json:"overridePublisher"`
	SrtPublishPassphrase       string `json:"srtPublishPassphrase"`
	RtspTransport              string `json:"rtspTransport"`
	RtspAnyPort                bool   `json:"rtspAnyPort"`
	RtspRangeType              string `json:"rtspRangeType"`
	RtspRangeStart             string `json:"rtspRangeStart"`
	SourceRedirect             string `json:"sourceRedirect"`
	RpiCameraCamID             int    `json:"rpiCameraCamID"`
	RpiCameraWidth             int    `json:"rpiCameraWidth"`
	RpiCameraHeight            int    `json:"rpiCameraHeight"`
	RpiCameraHFlip             bool   `json:"rpiCameraHFlip"`
	RpiCameraVFlip             bool   `json:"rpiCameraVFlip"`
	RpiCameraBrightness        int    `json:"rpiCameraBrightness"`
	RpiCameraContrast          int    `json:"rpiCameraContrast"`
	RpiCameraSaturation        int    `json:"rpiCameraSaturation"`
	RpiCameraSharpness         int    `json:"rpiCameraSharpness"`
	RpiCameraExposure          string `json:"rpiCameraExposure"`
	RpiCameraAWB               string `json:"rpiCameraAWB"`
	RpiCameraAWBGains          []int  `json:"rpiCameraAWBGains"`
	RpiCameraDenoise           string `json:"rpiCameraDenoise"`
	RpiCameraShutter           int    `json:"rpiCameraShutter"`
	RpiCameraMetering          string `json:"rpiCameraMetering"`
	RpiCameraGain              int    `json:"rpiCameraGain"`
	RpiCameraEV                int    `json:"rpiCameraEV"`
	RpiCameraROI               string `json:"rpiCameraROI"`
	RpiCameraHDR               bool   `json:"rpiCameraHDR"`
	RpiCameraTuningFile        string `json:"rpiCameraTuningFile"`
	RpiCameraMode              string `json:"rpiCameraMode"`
	RpiCameraFPS               int    `json:"rpiCameraFPS"`
	RpiCameraAfMode            string `json:"rpiCameraAfMode"`
	RpiCameraAfRange           string `json:"rpiCameraAfRange"`
	RpiCameraAfSpeed           string `json:"rpiCameraAfSpeed"`
	RpiCameraLensPosition      int    `json:"rpiCameraLensPosition"`
	RpiCameraAfWindow          string `json:"rpiCameraAfWindow"`
	RpiCameraFlickerPeriod     int    `json:"rpiCameraFlickerPeriod"`
	RpiCameraTextOverlayEnable bool   `json:"rpiCameraTextOverlayEnable"`
	RpiCameraTextOverlay       string `json:"rpiCameraTextOverlay"`
	RpiCameraCodec             string `json:"rpiCameraCodec"`
	RpiCameraIDRPeriod         int    `json:"rpiCameraIDRPeriod"`
	RpiCameraBitrate           int    `json:"rpiCameraBitrate"`
	RpiCameraProfile           string `json:"rpiCameraProfile"`
	RpiCameraLevel             string `json:"rpiCameraLevel"`
	RunOnInit                  string `json:"runOnInit"`
	RunOnInitRestart           bool   `json:"runOnInitRestart"`
	RunOnDemand                string `json:"runOnDemand"`
	RunOnDemandRestart         bool   `json:"runOnDemandRestart"`
	RunOnDemandStartTimeout    string `json:"runOnDemandStartTimeout"`
	RunOnDemandCloseAfter      string `json:"runOnDemandCloseAfter"`
	RunOnUnDemand              string `json:"runOnUnDemand"`
	RunOnReady                 string `json:"runOnReady"`
	RunOnReadyRestart          bool   `json:"runOnReadyRestart"`
	RunOnNotReady              string `json:"runOnNotReady"`
	RunOnRead                  string `json:"runOnRead"`
	RunOnReadRestart           bool   `json:"runOnReadRestart"`
	RunOnUnread                string `json:"runOnUnread"`
	RunOnRecordSegmentCreate   string `json:"runOnRecordSegmentCreate"`
	RunOnRecordSegmentComplete string `json:"runOnRecordSegmentComplete"`
}
