package v1

type Images struct {
	OperatorImage                 string `json:"operatorImage,omitempty"`
	ConfigReloaderImage           string `json:"configReloaderImage,omitempty"`
	PrometheusConfigReloaderImage string `json:"prometheusConfigReloaderImage,omitempty"`
	KubeRbacProxyImage            string `json:"kubeRbacProxyImage,omitempty"`
}
type PrometheusOperator struct {
	Enabled string `json:"enabled,omitempty"`
	Images  Images `json:"images,omitempty"`
}
type Prometheus struct {
	Enabled             string `json:"enabled,omitempty"`
	Image               string `json:"image,omitempty"`
	BasicAuthProxyImage string `json:"basicAuthProxyImage,omitempty"`
	CPURequest          string `json:"cpuRequest,omitempty"`
	MemoryRequest       string `json:"memoryRequest,omitempty"`
	SvcName             string `json:"svcName,omitempty"`
	Port                int    `json:"port,omitempty"`
	NodePort            int    `json:"nodePort,omitempty"`
	StorageSize         string `json:"storageSize,omitempty"`
	StorageClass        string `json:"storageClass,omitempty"`
	CredsRef            string `json:"credsRef,omitempty"`
}

type NodeExporter struct {
	Enabled string `json:"enabled,omitempty"`
	Image   string `json:"image,omitempty"`
}
type KubeStateMetrics struct {
	Enabled string `json:"enabled,omitempty"`
	Image   string `json:"image,omitempty"`
}
type Grafana struct {
	Enabled    string                `json:"enabled,omitempty"`
	Image      string                `json:"image,omitempty"`
	SvcName    string                `json:"svcName,omitempty"`
	Port       int                   `json:"port,omitempty"`
	NodePort   int                   `json:"nodePort,omitempty"`
	OauthProxy OauthProxyServiceConf `json:"oauthProxy,omitempty"`
}
type DefaultServiceMonitors struct {
	Enabled string `json:"enabled,omitempty"`
}
type SidekiqExporter struct {
	Enabled string `json:"enabled,omitempty"`
	Image   string `json:"image,omitempty"`
}
type MinioExporter struct {
	Enabled string `json:"enabled,omitempty"`
	Image   string `json:"image,omitempty"`
}
type DcgmExporter struct {
	Enabled string `json:"enabled,omitempty"`
	Image   string `json:"image,omitempty"`
}
type IdleMetricsExporter struct {
	Enabled string `json:"enabled,omitempty"`
}
type MetricsServer struct {
	Enabled string `json:"enabled,omitempty"`
	Image   string `json:"image,omitempty"`
}

type CnvrgInfraMonitoring struct {
	Enabled               string             `json:"enabled,omitempty"`
	PrometheusOperator    PrometheusOperator `json:"prometheusOperator,omitempty"`
	Prometheus            Prometheus         `json:"prometheus,omitempty"`
	KubeletServiceMonitor string             `json:"kubeletServiceMonitor,omitempty"`
	NodeExporter          NodeExporter       `json:"nodeExporter,omitempty"`
	KubeStateMetrics      KubeStateMetrics   `json:"kubeStateMetrics,omitempty"`
	Grafana               Grafana            `json:"grafana,omitempty"`
	DcgmExporter          DcgmExporter       `json:"dcgmExporter,omitempty"`
	//DefaultServiceMonitors DefaultServiceMonitors `json:"defaultServiceMonitors,omitempty"`
	//SidekiqExporter        SidekiqExporter        `json:"sidekiqExporter,omitempty"`
	//MinioExporter          MinioExporter          `json:"minioExporter,omitempty"`

	//IdleMetricsExporter    IdleMetricsExporter    `json:"idleMetricsExporter,omitempty"`
	//MetricsServer          MetricsServer          `json:"metricsServer,omitempty"`
}

type CnvrgAppMonitoring struct {
	Enabled            string     `json:"enabled,omitempty"`
	UpstreamPrometheus string     `json:"upstreamPrometheus,omitempty"`
	Prometheus         Prometheus `json:"prometheus,omitempty"`
	Grafana            Grafana    `json:"grafana,omitempty"`
}

var grafanaDefault = Grafana{
	Enabled:    "true",
	Image:      "grafana/grafana:7.3.4",
	SvcName:    "grafana",
	Port:       8080,
	NodePort:   30012,
	OauthProxy: OauthProxyServiceConf{SkipAuthRegex: []string{`\/api\/health`}},
}

var prometheusDefault = Prometheus{
	Enabled:             "true",
	Image:               "quay.io/prometheus/prometheus:v2.22.1",
	BasicAuthProxyImage: "docker.io/nginx:1.20",
	CPURequest:          "200m",
	MemoryRequest:       "500Mi",
	SvcName:             "prometheus",
	Port:                9091, // basic auth nginx proxy is enabled by default
	NodePort:            30909,
	StorageSize:         "50Gi",
	StorageClass:        "",
	CredsRef:            "prom-creds",
}

var cnvrgAppMonitoringDefault = CnvrgAppMonitoring{
	Enabled:            "true",
	UpstreamPrometheus: "prometheus-operated.cnvrg-infra.svc.cluster.local:9090",
	Prometheus:         prometheusDefault,
	Grafana:            grafanaDefault,
}

var infraMonitoringDefault = CnvrgInfraMonitoring{
	Enabled:               "true",
	KubeletServiceMonitor: "true",
	Prometheus:            prometheusDefault,
	Grafana:               grafanaDefault,
	PrometheusOperator: PrometheusOperator{
		Enabled: "true",
		Images: Images{
			OperatorImage:                 "quay.io/prometheus-operator/prometheus-operator:v0.44.1",
			PrometheusConfigReloaderImage: "quay.io/prometheus-operator/prometheus-config-reloader:v0.44.1",
			KubeRbacProxyImage:            "quay.io/brancz/kube-rbac-proxy:v0.8.0",
		},
	},
	KubeStateMetrics: KubeStateMetrics{
		Enabled: "true",
		Image:   "quay.io/coreos/kube-state-metrics:v1.9.7",
	},
	NodeExporter: NodeExporter{
		Enabled: "true",
		Image:   "quay.io/prometheus/node-exporter:v1.0.1",
	},
	DcgmExporter: DcgmExporter{
		Enabled: "true",
		Image:   "nvcr.io/nvidia/k8s/dcgm-exporter:2.1.4-2.3.1-ubuntu18.04",
	},
	//KubeStateMetrics: KubeStateMetrics{
	//	Enabled: "true",
	//	Image:   "quay.io/coreos/kube-state-metrics:v1.9.5",
	//},

	//DefaultServiceMonitors: DefaultServiceMonitors{Enabled: "true"},
	//SidekiqExporter: SidekiqExporter{
	//	Enabled: "true",
	//	Image:   "docker.io/strech/sidekiq-prometheus-exporter:0.1.13",
	//},
	//MinioExporter: MinioExporter{
	//	Enabled: "true",
	//	Image:   "docker.io/cnvrg/cnvrg-boot:v0.24",
	//},

	//IdleMetricsExporter: IdleMetricsExporter{Enabled: "true"},
	//MetricsServer: MetricsServer{
	//	Enabled: "true",
	//	Image:   "k8s.gcr.io/metrics-server/metrics-server:v0.3.7",
	//},
}
