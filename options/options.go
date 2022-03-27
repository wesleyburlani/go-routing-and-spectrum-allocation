package options

type ProgramOptions struct {
	NodesFilePath                     string `long:"nodesFilePath" description:"file path of the nodes specifications" required:"true"`
	EdgesFilePath                     string `long:"edgesFilePath" description:"file path of the edges specifications" required:"true"`
	LogType                           string `long:"logType" description:"type of the log: [file, console]" required:"true"`
	LogFilePath                       string `long:"logFilePath" description:"file path of the log" required:"false"`
	DemandsSource                     string `long:"demandsSource" description:"source of the demands [generate, file]" required:"true"`
	DemandsFilePath                   string `long:"demandsFilePath" description:"defined if demands should be loaded from the file path" required:"false"`
	PathSearchAlgorithm               string `long:"pathSearchAlgorithm" description:"algorithm to use to find the path: [djikstra]" required:"true"`
	DisjointedPathPairSearchAlgorithm string `long:"disjointedPathPairSearchAlgorithm" description:"algorithm to use to find the disjointed path pair: [suurballe]" required:"true"`
	TableFillAlgorithm                string `long:"tableFillAlgorithm" description:"algorithm to use to fill the table: [first_fit_rsa, first_fit_rmlsa]" required:"true"`
	RsaType                           string `long:"rsaType" description:"type of the rsa: [single, dedicated_protection, shared_protection]" required:"true"`
}

var Options ProgramOptions
