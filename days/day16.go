package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type valve struct {
	id       string
	mask     int
	flowRate int
	tunnels  []string
	isOpen   bool
}

type memoKey struct {
	id          string
	elID        string
	mask        int
	minutes     int
	curFlowRate int
}

func Day16() int {
	reader := bufio.NewReader(os.Stdin)
	valvesMap := make(map[string]*valve)
	mask := 1
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\n" {
			break
		}
		line = line[:len(line)-1]
		split := strings.Split(line, " ")
		id := split[1]
		_, flowRate, _ := strings.Cut(split[4][:len(split[4])-1], "=")
		flowRateInt, _ := strconv.Atoi(flowRate)
		tunnels := make([]string, 0)
		for i := 9; i < len(split); i++ {
			tunnels = append(tunnels, strings.TrimRight(split[i], ","))
		}
		v := &valve{
			id:       id,
			mask:     mask,
			flowRate: flowRateInt,
			tunnels:  tunnels,
		}
		valvesMap[id] = v
		mask <<= 1
		// fmt.Println(v)
	}

	for _, valve := range valvesMap {
		fmt.Println(valve)
	}
	memo := make(map[memoKey]int)
	res := getFlowValue(valvesMap["AA"], valvesMap["AA"], 30, valvesMap, memo)

	fmt.Println()
	for _, valve := range valvesMap {
		fmt.Println(valve)
	}
	return res
}

func getFlowValue(cur *valve, curEl *valve, minutes int, valvesMap map[string]*valve, memo map[memoKey]int) int {
	if minutes <= 0 {
		return 0
	}

	curFlowRate := 0
	allOpen := true

	for _, valve := range valvesMap {
		if valve.isOpen {
			curFlowRate += valve.flowRate
		} else if valve.flowRate > 0 {
			allOpen = false
		}
	}

	// fmt.Println(cur, curFlowRate)
	key := getMemoKey(cur.id, curEl.id, valvesMap, minutes, curFlowRate)
	if v, ok := memo[key]; ok {
		return v
	}

	if allOpen {
		res := curFlowRate * minutes
		memo[key] = res
		return res
	}

	var opts = make([]func() int, 0)
	// do nothing
	opts = append(opts, func() int { return curFlowRate * minutes })

	// open cur valve
	if cur.flowRate > 0 && minutes >= 1 && !valvesMap[cur.id].isOpen {
		opts = append(opts, func() int {
			valvesMap[cur.id].isOpen = true
			flowValue := curFlowRate * 1
			res := flowValue + getFlowValue(cur, curEl, minutes-1, valvesMap, memo)
			valvesMap[cur.id].isOpen = false
			return res
		})
		// for _, tunnel := range cur.tunnels {
		// 	// first minute + second minute
		// 	t := tunnel
		// 	opt := func() int {
		// 		valvesMap[cur.id].isOpen = true
		// 		flowValue := curFlowRate + (curFlowRate + cur.flowRate)
		// 		res := flowValue + getFlowValue(valvesMap[t], curEl, minutes-2, valvesMap, memo)
		// 		valvesMap[cur.id].isOpen = false
		// 		return res
		// 	}
		// 	opts = append(opts, opt)
		// }
	}

	// skip current valve
	if minutes >= 1 {
		for _, tunnel := range cur.tunnels {
			t := tunnel
			opt := func() int {
				flowValue := curFlowRate * 1
				res := flowValue + getFlowValue(valvesMap[t], curEl, minutes-1, valvesMap, memo)
				return res
			}
			opts = append(opts, opt)
		}
	}

	res := 0
	for _, opt := range opts {
		res = max(res, opt())
	}
	memo[key] = res

	return res
}

func getMemoKey(curId string, curElID string, valvesMap map[string]*valve, minutes int, curFlowRate int) memoKey {
	mask := 0
	if valvesMap[curId].isOpen {
		mask += 1
	}
	// for _, v := range valvesMap {
	// 	if v.isOpen {
	// 		mask |= v.mask
	// 	}
	// }

	return memoKey{curId, curElID, mask, minutes, curFlowRate}
}
