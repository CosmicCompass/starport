package app

// Options ...
type Options struct {
	AppName          string
	BinaryNamePrefix string
	ModulePath       string
	Denom            string
	AccountPrefix    string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
