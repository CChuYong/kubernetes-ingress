package configs

// appProtectDosResources holds the file names of APDosPolicy and APDosLogConf resources used in an Ingress resource.
type appProtectDosResources struct {
	AppProtectDosEnable       string
	AppProtectDosLogEnable    bool
	AppProtectDosMonitor      string
	AppProtectDosName         string
	AppProtectDosAccessLogDst string
	AppProtectDosPolicyFile   string
	AppProtectDosLogConfFile  string
}
