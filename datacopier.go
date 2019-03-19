package main

import "github.com/httpreserve/httpreserve"

// This structure is used to communicate with the server
// we may also use some static storage in the form of Bolt DB
// the final signal to the webapp will be a empty payload
// but with the complete flag set to true so that we know there
// is no more work to be processed. ls contains a link stat
// data structure if we can recreate one from the JSON we receive
// else the js variable will contain a single JSON document to
// be processed.
type processLog struct {
	complete bool
	ls       httpreserve.LinkStats
	js       string
	lmap     map[string]interface{}
}

// Thread safe data copy from one slice to another
func tsdatacopy(copyfrom *int, copyto *int, list []string) []string {
	//protect memory by copying only what we know we've got
	*copyto = len(list)
	if *copyto > 0 && *copyto > *copyfrom {
		var res []string
		res = make([]string, *copyto-*copyfrom)
		copy(res, list[*copyfrom:*copyto])
		*copyfrom = *copyto
		return res
	}
	return []string{}
}

// Thread safe data copy from one slice to another
func pldatacopy(copyfrom *int, copyto *int, list []processLog) []processLog {
	//protect memory by copying only what we know we've got
	*copyto = len(list)
	if *copyto > 0 && *copyto > *copyfrom {
		var res []processLog
		res = make([]processLog, *copyto-*copyfrom)
		copy(res, list[*copyfrom:*copyto])
		*copyfrom = *copyto
		return res
	}
	return []processLog{}
}

// Thread safe data copy from one slice to another
// This method allows us to specify a length which may be safer for
// us in the long run...
func pldatacopylen(copyfrom *int, copyto *int, list []processLog, copylen int) []processLog {
	//protect memory by copying only what we know we've got
	*copyto = *copyto + 1
	if *copyto > 0 && *copyto > *copyfrom && *copyto <= len(list) {
		var res []processLog
		res = make([]processLog, *copyto-*copyfrom)
		copy(res, list[*copyfrom:*copyto])
		*copyfrom = *copyto
		return res
	}
	return []processLog{}
}
