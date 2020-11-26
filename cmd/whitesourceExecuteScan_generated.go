// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type whitesourceExecuteScanOptions struct {
	BuildTool                            string   `json:"buildTool,omitempty"`
	BuildDescriptorFile                  string   `json:"buildDescriptorFile,omitempty"`
	VersioningModel                      string   `json:"versioningModel,omitempty"`
	CreateProductFromPipeline            bool     `json:"createProductFromPipeline,omitempty"`
	SecurityVulnerabilities              bool     `json:"securityVulnerabilities,omitempty"`
	Timeout                              int      `json:"timeout,omitempty"`
	AgentDownloadURL                     string   `json:"agentDownloadUrl,omitempty"`
	ConfigFilePath                       string   `json:"configFilePath,omitempty"`
	ReportDirectoryName                  string   `json:"reportDirectoryName,omitempty"`
	AggregateVersionWideReport           bool     `json:"aggregateVersionWideReport,omitempty"`
	VulnerabilityReportFormat            string   `json:"vulnerabilityReportFormat,omitempty"`
	ParallelLimit                        string   `json:"parallelLimit,omitempty"`
	Reporting                            bool     `json:"reporting,omitempty"`
	ServiceURL                           string   `json:"serviceUrl,omitempty"`
	BuildDescriptorExcludeList           []string `json:"buildDescriptorExcludeList,omitempty"`
	OrgToken                             string   `json:"orgToken,omitempty"`
	UserToken                            string   `json:"userToken,omitempty"`
	LicensingVulnerabilities             bool     `json:"licensingVulnerabilities,omitempty"`
	AgentFileName                        string   `json:"agentFileName,omitempty"`
	EmailAddressesOfInitialProductAdmins []string `json:"emailAddressesOfInitialProductAdmins,omitempty"`
	ProductVersion                       string   `json:"productVersion,omitempty"`
	JreDownloadURL                       string   `json:"jreDownloadUrl,omitempty"`
	ProductName                          string   `json:"productName,omitempty"`
	ProjectName                          string   `json:"projectName,omitempty"`
	ProjectToken                         string   `json:"projectToken,omitempty"`
	VulnerabilityReportTitle             string   `json:"vulnerabilityReportTitle,omitempty"`
	InstallCommand                       string   `json:"installCommand,omitempty"`
	ScanType                             string   `json:"scanType,omitempty"`
	CvssSeverityLimit                    string   `json:"cvssSeverityLimit,omitempty"`
	Includes                             string   `json:"includes,omitempty"`
	Excludes                             string   `json:"excludes,omitempty"`
	ProductToken                         string   `json:"productToken,omitempty"`
	AgentParameters                      string   `json:"agentParameters,omitempty"`
	ProjectSettingsFile                  string   `json:"projectSettingsFile,omitempty"`
	GlobalSettingsFile                   string   `json:"globalSettingsFile,omitempty"`
	M2Path                               string   `json:"m2Path,omitempty"`
	InstallArtifacts                     bool     `json:"installArtifacts,omitempty"`
	DefaultNpmRegistry                   string   `json:"defaultNpmRegistry,omitempty"`
}

type whitesourceExecuteScanCommonPipelineEnvironment struct {
	custom struct {
		whitesourceProjectNames []string
	}
}

func (p *whitesourceExecuteScanCommonPipelineEnvironment) persist(path, resourceName string) {
	content := []struct {
		category string
		name     string
		value    interface{}
	}{
		{category: "custom", name: "whitesourceProjectNames", value: p.custom.whitesourceProjectNames},
	}

	errCount := 0
	for _, param := range content {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(param.category, param.name), param.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting piper environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Piper environment")
	}
}

// WhitesourceExecuteScanCommand BETA
func WhitesourceExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "whitesourceExecuteScan"

	metadata := whitesourceExecuteScanMetadata()
	var stepConfig whitesourceExecuteScanOptions
	var startTime time.Time
	var commonPipelineEnvironment whitesourceExecuteScanCommonPipelineEnvironment

	var createWhitesourceExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "BETA",
		Long: `BETA
With this step [WhiteSource](https://www.whitesourcesoftware.com) security and license compliance scans can be executed and assessed.
WhiteSource is a Software as a Service offering based on a so called unified agent that locally determines the dependency
tree of a node.js, Java, Python, Ruby, or Scala based solution and sends it to the WhiteSource server for a policy based license compliance
check and additional Free and Open Source Software Publicly Known Vulnerabilities detection.
!!! note "Docker Images"
    The underlying Docker images are public and specific to the solution's programming language(s) and therefore may have to be exchanged
    to fit to and support the relevant scenario. The default Python environment used is i.e. Python 3 based.
!!! warn "Restrictions"
    Currently the step does contain hardened scan configurations for ` + "`" + `scanType` + "`" + ` ` + "`" + `'pip'` + "`" + ` and ` + "`" + `'go'` + "`" + `. Other environments are still being elaborated,
    so please thoroughly check your results and do not take them for granted by default.
    Also not all environments have been thoroughly tested already therefore you might need to tweak around with the default containers used or
    create your own ones to adequately support your scenario. To do so please modify ` + "`" + `dockerImage` + "`" + ` and ` + "`" + `dockerWorkspace` + "`" + ` parameters.
    The step expects an environment containing the programming language related compiler/interpreter as well as the related build tool. For a list
    of the supported build tools per environment please refer to the [WhiteSource Unified Agent Documentation](https://whitesource.atlassian.net/wiki/spaces/WD/pages/33718339/Unified+Agent).`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.OrgToken)
			log.RegisterSecret(stepConfig.UserToken)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				commonPipelineEnvironment.persist(GeneralConfig.EnvRootPath, "commonPipelineEnvironment")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			whitesourceExecuteScan(stepConfig, &telemetryData, &commonPipelineEnvironment)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addWhitesourceExecuteScanFlags(createWhitesourceExecuteScanCmd, &stepConfig)
	return createWhitesourceExecuteScanCmd
}

func addWhitesourceExecuteScanFlags(cmd *cobra.Command, stepConfig *whitesourceExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.BuildTool, "buildTool", os.Getenv("PIPER_buildTool"), "Defines the tool which is used for building the artifact.")
	cmd.Flags().StringVar(&stepConfig.BuildDescriptorFile, "buildDescriptorFile", os.Getenv("PIPER_buildDescriptorFile"), "Explicit path to the build descriptor file.")
	cmd.Flags().StringVar(&stepConfig.VersioningModel, "versioningModel", `major`, "The default project versioning model used in case `projectVersion` parameter is empty for creating the version based on the build descriptor version to report results in Whitesource, can be one of `'major'`, `'major-minor'`, `'semantic'`, `'full'`")
	cmd.Flags().BoolVar(&stepConfig.CreateProductFromPipeline, "createProductFromPipeline", true, "Whether to create the related WhiteSource product on the fly based on the supplied pipeline configuration.")
	cmd.Flags().BoolVar(&stepConfig.SecurityVulnerabilities, "securityVulnerabilities", true, "Whether security compliance is considered and reported as part of the assessment.")
	cmd.Flags().IntVar(&stepConfig.Timeout, "timeout", 900, "Timeout in seconds until an HTTP call is forcefully terminated.")
	cmd.Flags().StringVar(&stepConfig.AgentDownloadURL, "agentDownloadUrl", `https://github.com/whitesource/unified-agent-distribution/releases/latest/download/wss-unified-agent.jar`, "URL used to download the latest version of the WhiteSource Unified Agent.")
	cmd.Flags().StringVar(&stepConfig.ConfigFilePath, "configFilePath", `./wss-generated-file.config`, "Explicit path to the WhiteSource Unified Agent configuration file.")
	cmd.Flags().StringVar(&stepConfig.ReportDirectoryName, "reportDirectoryName", `whitesource-reports`, "Name of the directory to save vulnerability/risk reports to")
	cmd.Flags().BoolVar(&stepConfig.AggregateVersionWideReport, "aggregateVersionWideReport", false, "This does not run a scan, instead just generated a report for all projects with projectVersion = config.ProductVersion")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityReportFormat, "vulnerabilityReportFormat", `xlsx`, "Format of the file the vulnerability report is written to.")
	cmd.Flags().StringVar(&stepConfig.ParallelLimit, "parallelLimit", `15`, "[NOT IMPLEMENTED] Limit of parallel jobs being run at once in case of `scanType: 'mta'` based scenarios, defaults to `15`.")
	cmd.Flags().BoolVar(&stepConfig.Reporting, "reporting", true, "Whether assessment is being done at all, defaults to `true`")
	cmd.Flags().StringVar(&stepConfig.ServiceURL, "serviceUrl", `https://saas.whitesourcesoftware.com/api`, "URL to the WhiteSource server API used for communication.")
	cmd.Flags().StringSliceVar(&stepConfig.BuildDescriptorExcludeList, "buildDescriptorExcludeList", []string{`unit-tests/pom.xml`, `integration-tests/pom.xml`}, "List of build descriptors and therefore modules to exclude from the scan and assessment activities.")
	cmd.Flags().StringVar(&stepConfig.OrgToken, "orgToken", os.Getenv("PIPER_orgToken"), "WhiteSource token identifying your organization.")
	cmd.Flags().StringVar(&stepConfig.UserToken, "userToken", os.Getenv("PIPER_userToken"), "WhiteSource token identifying the user executing the scan.")
	cmd.Flags().BoolVar(&stepConfig.LicensingVulnerabilities, "licensingVulnerabilities", true, "[NOT IMPLEMENTED] Whether license compliance is considered and reported as part of the assessment.")
	cmd.Flags().StringVar(&stepConfig.AgentFileName, "agentFileName", `wss-unified-agent.jar`, "Locally used name for the Unified Agent jar file after download.")
	cmd.Flags().StringSliceVar(&stepConfig.EmailAddressesOfInitialProductAdmins, "emailAddressesOfInitialProductAdmins", []string{}, "The list of email addresses to assign as product admins for newly created WhiteSource products.")
	cmd.Flags().StringVar(&stepConfig.ProductVersion, "productVersion", os.Getenv("PIPER_productVersion"), "Version of the WhiteSource product to be created and used for results aggregation, usually determined automatically.")
	cmd.Flags().StringVar(&stepConfig.JreDownloadURL, "jreDownloadUrl", os.Getenv("PIPER_jreDownloadUrl"), "[NOT IMPLEMENTED] URL used for downloading the Java Runtime Environment (JRE) required to run the WhiteSource Unified Agent.")
	cmd.Flags().StringVar(&stepConfig.ProductName, "productName", os.Getenv("PIPER_productName"), "Name of the WhiteSource product used for results aggregation. This parameter is mandatory if the parameter `createProductFromPipeline` is set to `true` and the WhiteSource product does not yet exist. It is also mandatory if the parameter `productToken` is not provided.")
	cmd.Flags().StringVar(&stepConfig.ProjectName, "projectName", os.Getenv("PIPER_projectName"), "The project name used for reporting results in WhiteSource. When provided, all source modules will be scanned into one aggregated WhiteSource project. For scan types `maven`, `mta`, `npm`, the default is to generate one WhiteSource project per module, whereas the project name is derived from the module's build descriptor. For NPM modules, project aggregation is not supported, the last scanned NPM module will override all previously aggregated scan results!")
	cmd.Flags().StringVar(&stepConfig.ProjectToken, "projectToken", os.Getenv("PIPER_projectToken"), "Project token to execute scan on. Ignored for scan types `maven`, `mta` and `npm`. Used for project aggregation when scanning with the Unified Agent and can be provided as an alternative to `projectName`.")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityReportTitle, "vulnerabilityReportTitle", `WhiteSource Security Vulnerability Report`, "Title of vulnerability report written during the assessment phase.")
	cmd.Flags().StringVar(&stepConfig.InstallCommand, "installCommand", os.Getenv("PIPER_installCommand"), "[NOT IMPLEMENTED] Install command that can be used to populate the default docker image for some scenarios.")
	cmd.Flags().StringVar(&stepConfig.ScanType, "scanType", os.Getenv("PIPER_scanType"), "Type of development stack used to implement the solution. For scan types other than `mta`, `maven`, and `npm`, the WhiteSource Unified Agent is downloaded and used to perform the scan. If the parameter is not provided, it is derived from the parameter `buildTool`, which is usually configured in the general section of the pipeline config file.")
	cmd.Flags().StringVar(&stepConfig.CvssSeverityLimit, "cvssSeverityLimit", `-1`, "Limit of tolerable CVSS v3 score upon assessment and in consequence fails the build, defaults to  `-1`.")
	cmd.Flags().StringVar(&stepConfig.Includes, "includes", `**\/src\/main\/**\/*.java **\/*.py **\/*.go **\/*.js **\/*.ts`, "Space separated list of file path patterns to include in the scan, slashes must be escaped for sed.")
	cmd.Flags().StringVar(&stepConfig.Excludes, "excludes", `tests/**/*.py **/src/test/**/*.java`, "Space separated list of file path patterns to exclude in the scan")
	cmd.Flags().StringVar(&stepConfig.ProductToken, "productToken", os.Getenv("PIPER_productToken"), "Token of the WhiteSource product to be created and used for results aggregation, usually determined automatically. Can optionally be provided as an alternative to `productName`.")
	cmd.Flags().StringVar(&stepConfig.AgentParameters, "agentParameters", os.Getenv("PIPER_agentParameters"), "[NOT IMPLEMENTED] Additional parameters passed to the Unified Agent command line.")
	cmd.Flags().StringVar(&stepConfig.ProjectSettingsFile, "projectSettingsFile", os.Getenv("PIPER_projectSettingsFile"), "Path to the mvn settings file that should be used as project settings file.")
	cmd.Flags().StringVar(&stepConfig.GlobalSettingsFile, "globalSettingsFile", os.Getenv("PIPER_globalSettingsFile"), "Path to the mvn settings file that should be used as global settings file.")
	cmd.Flags().StringVar(&stepConfig.M2Path, "m2Path", os.Getenv("PIPER_m2Path"), "Path to the location of the local repository that should be used.")
	cmd.Flags().BoolVar(&stepConfig.InstallArtifacts, "installArtifacts", false, "If enabled, it will install all artifacts to the local maven repository to make them available before running whitesource. This is required if any maven module has dependencies to other modules in the repository and they were not installed before.")
	cmd.Flags().StringVar(&stepConfig.DefaultNpmRegistry, "defaultNpmRegistry", os.Getenv("PIPER_defaultNpmRegistry"), "URL of the npm registry to use. Defaults to https://registry.npmjs.org/")

	cmd.MarkFlagRequired("buildTool")
	cmd.MarkFlagRequired("orgToken")
	cmd.MarkFlagRequired("userToken")
	cmd.MarkFlagRequired("productName")
}

// retrieve step metadata
func whitesourceExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "whitesourceExecuteScan",
			Aliases:     []config.Alias{},
			Description: "BETA",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name: "buildTool",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "buildTool",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "versioningModel",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "defaultVersioningModel"}},
					},
					{
						Name:        "createProductFromPipeline",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "securityVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "timeout",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentDownloadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "configFilePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "reportDirectoryName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "aggregateVersionWideReport",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityReportFormat",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "parallelLimit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "reporting",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "serviceUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorExcludeList",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "orgToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "orgAdminUserTokenCredentialsId",
								Type: "secret",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name: "userToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "userTokenCredentialsId",
								Type: "secret",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "licensingVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentFileName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "emailAddressesOfInitialProductAdmins",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "productVersion",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "jreDownloadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "productName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "whitesourceProductName"}},
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesourceProjectName"}},
					},
					{
						Name:        "projectToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityReportTitle",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "installCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "scanType",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "cvssSeverityLimit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "includes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "excludes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "productToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesourceProductToken"}},
					},
					{
						Name:        "agentParameters",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "projectSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/projectSettingsFile"}},
					},
					{
						Name:        "globalSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/globalSettingsFile"}},
					},
					{
						Name:        "m2Path",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/m2Path"}},
					},
					{
						Name:        "installArtifacts",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "defaultNpmRegistry",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "GENERAL", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "npm/defaultNpmRegistry"}},
					},
				},
			},
			Containers: []config.Container{
				{Image: "devxci/mbtci:1.0.14", WorkingDir: "/project", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "scanType", Value: "mta"}}}}},
				{Image: "maven:3.5-jdk-8", WorkingDir: "/home/java", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "scanType", Value: "maven"}}}}},
				{Image: "node:lts-stretch", WorkingDir: "/home/node", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "scanType", Value: "npm"}}}}},
				{Image: "hseeberger/scala-sbt:8u181_2.12.8_1.2.8", WorkingDir: "/home/scala", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "scanType", Value: "sbt"}}}}},
				{Image: "buildpack-deps:stretch-curl", WorkingDir: "/home/dub", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "scanType", Value: "dub"}}}}},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "commonPipelineEnvironment",
						Type: "piperEnvironment",
						Parameters: []map[string]interface{}{
							{"Name": "custom/whitesourceProjectNames"},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
