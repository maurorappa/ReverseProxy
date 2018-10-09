package main

import (
	"fmt"
	//	"golang.org/x/time/rate"
	toml "github.com/BurntSushi/toml"
)

type config struct {
	Listen            string   `toml:"listen"`
	Backend           string   `toml:"backend"`
	Path              string   `toml:"admin_path"`
	Logfile           string   `toml:"log_file"`
	Path_len          int      `toml:"path_len"`
	Verbose           bool     `toml:"verbose"`
	Forbidden_methods []string `toml:"forbidden_methods"`
	Forbidden_extensions []string `toml:"forbidden_extensions"`
	Max_rate 	  int `toml:"max_rate"`
	Smtp_server	  string   `toml:"smtp"`
	Mailbox		  string   `toml:"mailbox"`	
	Auth_user	  string   `toml:"auth_user"`
	Auth_pwd	  string   `toml:"auth_pwd"`
}

func loadConfig(path string) (*config, error) {
	conf := &config{}
	metaData, err := toml.DecodeFile(path, conf)
	if err != nil {
		return nil, err
	}
	for _, key := range metaData.Undecoded() {
		return nil, &configError{fmt.Sprintf("unknown option %q", key.String())}
	}

	if len(conf.Listen) == 0 {
		return nil, &configError{"no listen specified"}
	}

	if len(conf.Backend) == 0 {
		return nil, &configError{"no backend specified"}
	}

	if conf.Path == "" {
		return nil, &configError{"no admin directory specified"}
	}

	if conf.Logfile == "" {
		return nil, &configError{"no logfile specified"}
	}

	if conf.Path_len == 0 {
		conf.Path_len = 16
	}

	if len(conf.Forbidden_methods) == 0 {
		conf.Forbidden_methods = []string{}
	}

	if len(conf.Forbidden_extensions) == 0 {
		conf.Forbidden_extensions = []string{}
	}

	return conf, nil
}

type configError struct {
	err string
}

func (e *configError) Error() string {
	return e.err
}
