package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaa-aaa-aaa"
	strnew := strings.Replace(str, "-","_", -1)
	fmt.Println(strnew)
}

// ["kubepods","burstable","pod-xxxx-xxx"]
// m.Name(name) -> kubepods-burstable-pod_xxx_xxx.slice
// cgroupPaths[memory] = path.Join(val, "kubepods-burstable-pod_xxx_xxx.slice"
//                                 val ->