package types

import "encoding/json"

func MarshalStringArray(vals []string) string {
	out, err := json.Marshal(vals)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func MarshalInterfaceArray(vals []interface{}) string {
	data := make([]string, len(vals))
	for i, v := range vals {
		dat, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		data[i] = string(dat)
	}
	return MarshalStringArray(data)
}

func MarshalBool(val bool) string {
	v, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return string(v)
}
