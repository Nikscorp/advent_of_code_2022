package days

import (
	"bufio"
	"fmt"
	"os"
)

type robotType int

const (
	oreType robotType = iota
	clayType
	obsidianType
	geodeType
)

type oreRobot struct {
	oreCost int
}

type clayRobot struct {
	oreCost int
}

type obsidianRobot struct {
	oreCost  int
	clayCost int
}

type geodeRobot struct {
	oreCost      int
	obsidianCost int
}

type blueprint struct {
	oreRobot      oreRobot
	clayRobot     clayRobot
	obsidianRobot obsidianRobot
	geodeRobot    geodeRobot
}

type resources struct {
	ore, clay, obsidian, geode int
}

type robotsMemoKey struct {
	minutes  int
	robots   string
	resource string
}

func Day19() int {
	reader := bufio.NewReader(os.Stdin)
	blueprints := make([]blueprint, 0)
	for {
		var id, oreCostInOre, clayCostInOre, obsidianCostInOre, obsidianCostInClay, geodeCostInOre, geodeCostInObsidian int
		_, err := fmt.Fscanf(reader, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.\n", &id, &oreCostInOre, &clayCostInOre, &obsidianCostInOre, &obsidianCostInClay, &geodeCostInOre, &geodeCostInObsidian)
		if err != nil {
			break
		}

		oreRobot := oreRobot{oreCostInOre}
		clayRobot := clayRobot{clayCostInOre}
		obsidianRobot := obsidianRobot{obsidianCostInOre, obsidianCostInClay}
		geodeRobot := geodeRobot{geodeCostInOre, geodeCostInObsidian}
		b := blueprint{oreRobot, clayRobot, obsidianRobot, geodeRobot}
		blueprints = append(blueprints, b)
	}

	fmt.Println(blueprints)

	res := 1
	for i, blueprint := range blueprints {
		robots := make(map[robotType]int)
		robots[oreType] = 1
		resources := resources{}
		memo := make(map[robotsMemoKey]int)
		cur := dp(32, blueprint, robots, resources, memo)
		fmt.Println(i+1, cur)
		res *= cur
		if i == 2 {
			break
		}
	}

	return res
}

func dp(minutes int, blueprint blueprint, robots map[robotType]int, resources resources, memo map[robotsMemoKey]int) int {
	if minutes <= 0 {
		return resources.geode
	}

	key := getRobotsMemoKey(minutes, robots, resources)
	if v, ok := memo[key]; ok {
		return v
	}

	opts := make([]func() int, 0)

	// Gather resources
	newResources := resources
	newResources.clay += robots[clayType]
	newResources.ore += robots[oreType]
	newResources.obsidian += robots[obsidianType]
	newResources.geode += robots[geodeType]

	buildGeode := false
	buildObsidian := false
	// 1. Build robots of one type or skip (up to 5 options)
	if resources.obsidian >= blueprint.geodeRobot.obsidianCost && resources.ore >= blueprint.geodeRobot.oreCost {
		buildGeode = true
		opts = append(opts, func() int {
			robots[geodeType]++
			newResources.ore -= blueprint.geodeRobot.oreCost
			newResources.obsidian -= blueprint.geodeRobot.obsidianCost
			res := dp(minutes-1, blueprint, robots, newResources, memo)
			robots[geodeType]--
			newResources.ore += blueprint.geodeRobot.oreCost
			newResources.obsidian += blueprint.geodeRobot.obsidianCost
			return res
		})
	}
	if resources.ore >= blueprint.obsidianRobot.oreCost && resources.clay >= blueprint.obsidianRobot.clayCost && !buildGeode {
		buildObsidian = true
		opts = append(opts, func() int {
			robots[obsidianType]++
			newResources.ore -= blueprint.obsidianRobot.oreCost
			newResources.clay -= blueprint.obsidianRobot.clayCost
			res := dp(minutes-1, blueprint, robots, newResources, memo)
			robots[obsidianType]--
			newResources.ore += blueprint.obsidianRobot.oreCost
			newResources.clay += blueprint.obsidianRobot.clayCost
			return res
		})
	}
	if resources.ore >= blueprint.oreRobot.oreCost && !buildGeode && !buildObsidian {
		opts = append(opts, func() int {
			robots[oreType]++
			newResources.ore -= blueprint.oreRobot.oreCost
			res := dp(minutes-1, blueprint, robots, newResources, memo)
			robots[oreType]--
			newResources.ore += blueprint.oreRobot.oreCost
			return res
		})
	}
	if resources.ore >= blueprint.clayRobot.oreCost && !buildGeode && !buildObsidian {
		opts = append(opts, func() int {
			robots[clayType]++
			newResources.ore -= blueprint.clayRobot.oreCost
			res := dp(minutes-1, blueprint, robots, newResources, memo)
			robots[clayType]--
			newResources.ore += blueprint.clayRobot.oreCost
			return res
		})
	}

	if len(opts) == 0 { // for q1 replace with !buildGeode
		opts = append(opts, func() int {
			return dp(minutes-1, blueprint, robots, newResources, memo)
		})
	}

	res := 0
	for _, opt := range opts {
		res = max(res, opt())
	}

	memo[key] = res

	// if len(memo) > 1000000 {
	// 	fmt.Println(len(memo))
	// }

	return res
}

func getRobotsMemoKey(minutes int, robots map[robotType]int, resources resources) robotsMemoKey {
	robotsKey := fmt.Sprintf("%d;%d;%d;%d", robots[oreType], robots[clayType], robots[obsidianType], robots[geodeType])
	resourcesKey := fmt.Sprintf("%d;%d;%d;%d", resources.ore, resources.clay, resources.obsidian, resources.geode)

	return robotsMemoKey{
		minutes:  minutes,
		robots:   robotsKey,
		resource: resourcesKey,
	}
}
