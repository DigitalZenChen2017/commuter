package cli

import (
	"flag"

	"github.com/KyleBanks/commuter/cmd"
)

// ArgParser parses input arguments from the command line.
type ArgParser struct {
	Args []string
}

// NewArgParser initializes and returns an ArgParser.
func NewArgParser(args []string) *ArgParser {
	return &ArgParser{
		Args: args,
	}
}

// Parse attempts to determine which command is being executed,
// parse its flags, and return it.
func (a *ArgParser) Parse(conf *cmd.Configuration, s cmd.StorageProvider) (cmd.RunnerValidator, error) {
	if conf == nil || len(a.Args) == 0 {
		return a.parseConfigureCmd(s)
	}

	switch a.Args[0] {
	case cmdAdd:
		return a.parseAddCmd(s, a.Args[1:])
	}

	return a.parseCommuteCmd(a.Args)
}

// parseConfigureCmd parses and returns a ConfigureCmd.
func (a *ArgParser) parseConfigureCmd(s cmd.StorageProvider) (*cmd.ConfigureCmd, error) {
	return &cmd.ConfigureCmd{
		Input: NewStdin(),
		Store: s,
	}, nil
}

// parseCommuteCmd parses and returns a CommuteCmd from user supplied flags.
func (a *ArgParser) parseCommuteCmd(args []string) (*cmd.CommuteCmd, error) {
	var c cmd.CommuteCmd

	f := flag.NewFlagSet(cmdDefault, flag.ExitOnError)
	f.StringVar(&c.From, defaultFromParam, cmd.DefaultLocationAlias, defaultFromUsage)
	f.StringVar(&c.To, defaultToParam, cmd.DefaultLocationAlias, defaultToUsage)
	f.Parse(args)

	return &c, nil
}

// parseAddCmd parses and returns an AddCmd from user supplied flags.
func (a *ArgParser) parseAddCmd(s cmd.StorageProvider, args []string) (*cmd.AddCmd, error) {
	c := cmd.AddCmd{Store: s}

	f := flag.NewFlagSet(cmdAdd, flag.ExitOnError)
	f.StringVar(&c.Name, addNameParam, "", addNameUsage)
	f.StringVar(&c.Value, addLocationParam, "", addLocationUsage)
	f.Parse(args)

	return &c, nil
}