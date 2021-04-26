package v1

type ConsistentHash struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type SharedStorage struct {
	Enabled          *bool          `json:"enabled,omitempty"`
	UseExistingClaim string         `json:"useExistingClaim,omitempty"`
	ConsistentHash   ConsistentHash `json:"consistentHash,omitempty"`
}

type Limits struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}
type Requests struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type WebApp struct {
	Replicas                int                   `json:"replicas,omitempty"`
	Enabled                 *bool                 `json:"enabled,omitempty"`
	Image                   string                `json:"image,omitempty"`
	Port                    int                   `json:"port,omitempty"`
	CPU                     string                `json:"cpu,omitempty"`
	Memory                  string                `json:"memory,omitempty"`
	SvcName                 string                `json:"svcName,omitempty"`
	NodePort                int                   `json:"nodePort,omitempty"`
	PassengerMaxPoolSize    int                   `json:"passengerMaxPoolSize,omitempty"`
	InitialDelaySeconds     int                   `json:"initialDelaySeconds,omitempty"`
	ReadinessPeriodSeconds  int                   `json:"readinessPeriodSeconds,omitempty"`
	ReadinessTimeoutSeconds int                   `json:"readinessTimeoutSeconds,omitempty"`
	FailureThreshold        int                   `json:"failureThreshold,omitempty"`
	OauthProxy              OauthProxyServiceConf `json:"oauthProxy,omitempty"`
}

type Sidekiq struct {
	Enabled  *bool  `json:"enabled,omitempty"`
	Split    *bool  `json:"split,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	Memory   string `json:"memory,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}
type Searchkiq struct {
	Enabled  *bool  `json:"enabled,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	Memory   string `json:"memory,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}
type Systemkiq struct {
	Enabled  *bool  `json:"enabled,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	Memory   string `json:"memory,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}

type Registry struct {
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type Hyper struct {
	Enabled                 *bool  `json:"enabled,omitempty"`
	Image                   string `json:"image,omitempty"`
	Port                    int    `json:"port,omitempty"`
	Replicas                int    `json:"replicas,omitempty"`
	NodePort                int    `json:"nodePort,omitempty"`
	SvcName                 string `json:"svcName,omitempty"`
	Token                   string `json:"token,omitempty"`
	CPURequest              string `json:"cpuRequest,omitempty"`
	MemoryRequest           string `json:"memoryRequest,omitempty"`
	CPULimit                string `json:"cpuLimit,omitempty"`
	MemoryLimit             string `json:"memoryLimit,omitempty"`
	ReadinessPeriodSeconds  int    `json:"readinessPeriodSeconds,omitempty"`
	ReadinessTimeoutSeconds int    `json:"readinessTimeoutSeconds,omitempty"`
}
type Seeder struct {
	Image           string `json:"image,omitempty"`
	SeedCmd         string `json:"seedCmd,omitempty"`
	CreateBucketCmd string `json:"createBucketCmd,omitempty"`
}
type Ldap struct {
	Enabled       *bool  `json:"enabled,omitempty"`
	Host          string `json:"host,omitempty"`
	Port          string `json:"port,omitempty"`
	Account       string `json:"account,omitempty"`
	Base          string `json:"base,omitempty"`
	AdminUser     string `json:"adminUser,omitempty"`
	AdminPassword string `json:"adminPassword,omitempty"`
	Ssl           string `json:"ssl,omitempty"`
}
type Rbac struct {
	Role               string `json:"role,omitempty"`
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
	RoleBindingName    string `json:"roleBindingName,omitempty"`
}
type SMTP struct {
	Server   string `json:"server,omitempty"`
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Domain   string `json:"domain,omitempty"`
}

type ObjectStorage struct {
	CnvrgStorageType             string `json:"cnvrgStorageType,omitempty"`
	CnvrgStorageBucket           string `json:"cnvrgStorageBucket,omitempty"`
	CnvrgStorageAccessKey        string `json:"cnvrgStorageAccessKey,omitempty"`
	CnvrgStorageSecretKey        string `json:"cnvrgStorageSecretKey,omitempty"`
	CnvrgStorageEndpoint         string `json:"cnvrgStorageEndpoint,omitempty"`
	MinioSseMasterKey            string `json:"minioSseMasterKey,omitempty"`
	CnvrgStorageAzureAccessKey   string `json:"cnvrgStorageAzureAccessKey,omitempty"`
	CnvrgStorageAzureAccountName string `json:"cnvrgStorageAzureAccountName,omitempty"`
	CnvrgStorageAzureContainer   string `json:"cnvrgStorageAzureContainer,omitempty"`
	CnvrgStorageRegion           string `json:"cnvrgStorageRegion,omitempty"`
	CnvrgStorageProject          string `json:"cnvrgStorageProject,omitempty"`
	GcpStorageSecret             string `json:"gcpStorageSecret,omitempty"`
	GcpKeyfileMountPath          string `json:"gcpKeyfileMountPath,omitempty"`
	GcpKeyfileName               string `json:"gcpKeyfileName,omitempty"`
	SecretKeyBase                string `json:"secretKeyBase,omitempty"`
	StsIv                        string `json:"stsIv,omitempty"`
	StsKey                       string `json:"stsKey,omitempty"`
}

type BaseConfig struct {
	JobsStorageClass     string            `json:"jobsStorageClass,omitempty"`
	FeatureFlags         map[string]string `json:"featureFlags,omitempty"`
	SentryURL            string            `json:"sentryUrl,omitempty"`
	PassengerAppEnv      string            `json:"passengerAppEnv,omitempty"`
	RailsEnv             string            `json:"railsEnv,omitempty"`
	RunJobsOnSelfCluster string            `json:"runJobsOnSelfCluster,omitempty"`
	DefaultComputeConfig string            `json:"defaultComputeConfig,omitempty"`
	DefaultComputeName   string            `json:"defaultComputeName,omitempty"`
	UseStdout            string            `json:"useStdout,omitempty"`
	ExtractTagsFromCmd   string            `json:"extractTagsFromCmd,omitempty"`
	CheckJobExpiration   string            `json:"checkJobExpiration,omitempty"`
	AgentCustomTag       string            `json:"agentCustomTag,omitempty"`
	Intercom             string            `json:"intercom,omitempty"`
	CnvrgJobUID          string            `json:"cnvrgJobUid,omitempty"`
	CcpStorageClass      string            `json:"ccpStorageClass,omitempty"`
	HostpathNode         string            `json:"hostpathNode,omitempty"`
}

type ControlPlane struct {
	WebApp        WebApp        `json:"webapp,omitempty"`
	Sidekiq       Sidekiq       `json:"sidekiq,omitempty"`
	Searchkiq     Searchkiq     `json:"searchkiq,omitempty"`
	Systemkiq     Systemkiq     `json:"systemkiq,omitempty"`
	Hyper         Hyper         `json:"hyper,omitempty"`
	Seeder        Seeder        `json:"seeder,omitempty"`
	BaseConfig    BaseConfig    `json:"baseConfig,omitempty"`
	Ldap          Ldap          `json:"ldap,omitempty"`
	Rbac          Rbac          `json:"rbac,omitempty"`
	SMTP          SMTP          `json:"smtp,omitempty"`
	Tenancy       Tenancy       `json:"tenancy,omitempty"`
	ObjectStorage ObjectStorage `json:"objectStorage,omitempty"`
	Mpi           Mpi           `json:"mpi,omitempty"`
}

type Tenancy struct {
	Enabled        *bool  `json:"enabled,omitempty"`
	DedicatedNodes string `json:"dedicatedNodes,omitempty"`
	Key            string `json:"key,omitempty"`
	Value          string `json:"value,omitempty"`
}

type Cnvrg struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Mpi struct {
	Enabled              *bool             `json:"enabled,omitempty"`
	Image                string            `json:"image,omitempty"`
	KubectlDeliveryImage string            `json:"kubectlDeliveryImage,omitempty"`
	ExtraArgs            map[string]string `json:"extraArgs,omitempty"`
	Registry             Registry          `json:"registry,omitempty"`
}

var mpiDefault = Mpi{
	Enabled:              &defaultEnabled,
	Image:                "mpioperator/mpi-operator:v0.2.3",
	KubectlDeliveryImage: "mpioperator/kubectl-delivery:v0.2.3",
	ExtraArgs:            nil,
	Registry: Registry{

		Name:     "mpi-private-registry",
		URL:      "docker.io",
		User:     "",
		Password: "",
	},
}

var appRegistryDefault = Registry{

	Name:     "cnvrg-app-registry",
	URL:      "docker.io",
	User:     "",
	Password: "",
}

var infraRegistryDefault = Registry{

	Name:     "cnvrg-infra-registry",
	URL:      "docker.io",
	User:     "",
	Password: "",
}

var controlPlanDefault = ControlPlane{

	WebApp: WebApp{
		Enabled:                 &defaultEnabled,
		Replicas:                1,
		Image:                   "cnvrg/core:3.1.5",
		Port:                    8080,
		CPU:                     "2000m",
		Memory:                  "4Gi",
		SvcName:                 "app",
		NodePort:                30080,
		PassengerMaxPoolSize:    20,
		InitialDelaySeconds:     10,
		ReadinessPeriodSeconds:  25,
		ReadinessTimeoutSeconds: 20,
		FailureThreshold:        4,
		OauthProxy: OauthProxyServiceConf{
			SkipAuthRegex: []string{
				`^\/api`,
				`\/assets`,
				`\/healthz`,
				`\/public`,
				`\/pack`,
				`\/vscode.tar.gz`,
				`\/gitlens.vsix`,
				`\/ms-python-release.vsix`,
			},
		},
	},

	Sidekiq: Sidekiq{
		Enabled:  &defaultEnabled,
		Split:    &defaultEnabled,
		CPU:      "1000m",
		Memory:   "3750Mi",
		Replicas: 2,
	},

	Searchkiq: Searchkiq{
		Enabled:  &defaultEnabled,
		CPU:      "750m",
		Memory:   "750Mi",
		Replicas: 1,
	},

	Systemkiq: Systemkiq{
		Enabled:  &defaultEnabled,
		CPU:      "500m",
		Memory:   "500Mi",
		Replicas: 1,
	},

	Hyper: Hyper{
		Enabled:                 &defaultEnabled,
		Image:                   "cnvrg/hyper-server:latest",
		Port:                    5050,
		Replicas:                1,
		NodePort:                30050,
		SvcName:                 "hyper",
		Token:                   "token",
		CPURequest:              "100m",
		MemoryRequest:           "200Mi",
		CPULimit:                "2000m",
		MemoryLimit:             "4Gi",
		ReadinessPeriodSeconds:  100,
		ReadinessTimeoutSeconds: 60,
	},

	Seeder: Seeder{

		Image:           "docker.io/cnvrg/cnvrg-boot:v0.27-tenancy",
		SeedCmd:         "rails db:migrate && rails db:seed && rails libraries:update",
		CreateBucketCmd: "mb.sh",
	},

	BaseConfig: BaseConfig{
		JobsStorageClass:     "",
		FeatureFlags:         nil,
		SentryURL:            "https://4409141e4a204282bd1f5c021e587509:dc15f684faa9479a839cf913b98b4ee2@sentry.cnvrg.io/32",
		PassengerAppEnv:      "app",
		RailsEnv:             "app",
		RunJobsOnSelfCluster: "true",
		DefaultComputeConfig: "/opt/kube",
		DefaultComputeName:   "default",
		UseStdout:            "true",
		ExtractTagsFromCmd:   "&defaultEnabled",
		CheckJobExpiration:   "true",
		AgentCustomTag:       "latest",
		Intercom:             "true",
		CnvrgJobUID:          "1000",
		CcpStorageClass:      "",
		HostpathNode:         "",
	},

	ObjectStorage: ObjectStorage{

		CnvrgStorageType:             "minio",
		CnvrgStorageBucket:           "cnvrg-storage",
		CnvrgStorageAccessKey:        "AKIAIOSFODNN7EXAMPLE",
		CnvrgStorageSecretKey:        "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		CnvrgStorageEndpoint:         "",
		MinioSseMasterKey:            "my-minio-key:a310aadcefdb634b748ae31225f175e3f64591f955dfc66ccc20e128a6817ff9",
		CnvrgStorageAzureAccessKey:   "",
		CnvrgStorageAzureAccountName: "",
		CnvrgStorageAzureContainer:   "",
		CnvrgStorageRegion:           "eastus",
		CnvrgStorageProject:          "",
		SecretKeyBase:                "0d2b33c2cc19cfaa838d3c354354a18fcc92beaaa8e97889ef99341c8aaf963ad3afcf0f7c20454cabb5c573c3fc35b60221034e109f4fb651ed1415bf61e9d5",
		StsIv:                        "DeJ/CGz/Hkb/IbRe4t1xLg==",
		StsKey:                       "05646d3cbf8baa5be7150b4283eda07d",
		GcpStorageSecret:             "gcp-storage-secret",
		GcpKeyfileMountPath:          "/tmp/gcp_keyfile",
		GcpKeyfileName:               "key.json",
	},

	Ldap: Ldap{
		Enabled:       &defaultEnabled,
		Host:          "",
		Port:          "",
		Account:       "userPrincipalName",
		Base:          "", // dc=my-domain,dc=local
		AdminUser:     "",
		AdminPassword: "",
		Ssl:           "", // true/&defaultEnabled
	},

	Rbac: Rbac{

		Role:               "cnvrg-control-plan-role",
		ServiceAccountName: "cnvrg",
		RoleBindingName:    "cnvrg-control-plan-binding",
	},

	SMTP: SMTP{
		Server:   "",
		Port:     "",
		Username: "",
		Password: "",
		Domain:   "",
	},

	Tenancy: Tenancy{
		Enabled:        &defaultEnabled,
		DedicatedNodes: "&defaultEnabled",
		Key:            "cnvrg-taint",
		Value:          "true",
	},

	Mpi: mpiDefault,
}
