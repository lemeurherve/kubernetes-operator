package plugins

const (
	configurationAsCodePlugin           = "configuration-as-code:1.54"
	gitPlugin                           = "git:4.9.0"
	jobDslPlugin                        = "job-dsl:1.77"
	kubernetesCredentialsProviderPlugin = "kubernetes-credentials-provider:0.20"
	kubernetesPlugin                    = "kubernetes:1.30.4"
	workflowAggregatorPlugin            = "workflow-aggregator:2.6"
	workflowJobPlugin                   = "workflow-job:2.42"
)

// basePluginsList contains plugins to install by operator.
var basePluginsList = []Plugin{
	Must(New(kubernetesPlugin)),
	Must(New(workflowJobPlugin)),
	Must(New(workflowAggregatorPlugin)),
	Must(New(gitPlugin)),
	Must(New(jobDslPlugin)),
	Must(New(configurationAsCodePlugin)),
	Must(New(kubernetesCredentialsProviderPlugin)),
}

// BasePlugins returns list of plugins to install by operator.
func BasePlugins() []Plugin {
	return basePluginsList
}
