package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// node1, node3
// node2
func GetCase1() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node2", AppName: "app1"},
		{NodeName: "node1", AppName: "app2"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node2", AppName: "app3"},
		{NodeName: "node3", AppName: "app3"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 1,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app3",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}

// node2,node3
// node1
// node4
func GetCase2() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node3", AppName: "app1"},
		{NodeName: "node1", AppName: "app2"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node4", AppName: "app2"},
		{NodeName: "node4", AppName: "app3"},
		{NodeName: "node3", AppName: "app3"},
		{NodeName: "node1", AppName: "app4"},
		{NodeName: "node3", AppName: "app4"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
		{"node4"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 1,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app3",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app4",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}

/*
group #:[ node2,node5,]
group #:[ node1,]
group #:[ node3,]
group #:[ node4,]
*/
func GetCase3() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node2", AppName: "app1"},
		{NodeName: "node3", AppName: "app1"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node3", AppName: "app2"},
		{NodeName: "node4", AppName: "app2"},
		{NodeName: "node4", AppName: "app3"},
		{NodeName: "node5", AppName: "app3"},
		{NodeName: "node1", AppName: "app3"},
		{NodeName: "node2", AppName: "app4"},
		{NodeName: "node3", AppName: "app4"},
		{NodeName: "node4", AppName: "app4"},
		{NodeName: "node3", AppName: "app5"},
		{NodeName: "node4", AppName: "app5"},
		{NodeName: "node5", AppName: "app5"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
		{"node4"},
		{"node5"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 1,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app3",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app4",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app5",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}

/*
group #:[ node2,node5,]
group #:[ node1,]
group #:[ node3,]
group #:[ node4,]
*/
func GetCase4() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node2", AppName: "app1"},
		{NodeName: "node3", AppName: "app1"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node3", AppName: "app2"},
		{NodeName: "node4", AppName: "app3"},
		{NodeName: "node5", AppName: "app3"},
		{NodeName: "node1", AppName: "app3"},
		{NodeName: "node2", AppName: "app4"},
		{NodeName: "node3", AppName: "app4"},
		{NodeName: "node4", AppName: "app4"},
		{NodeName: "node3", AppName: "app5"},
		{NodeName: "node5", AppName: "app5"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
		{"node4"},
		{"node5"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 1,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app3",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app4",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app5",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}


/*
group #:[ node1,node4,]
group #:[ node2,node5,]
group #:[ node3,]
*/
func GetCase5() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node2", AppName: "app1"},
		{NodeName: "node3", AppName: "app1"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node3", AppName: "app2"},
		{NodeName: "node4", AppName: "app3"},
		{NodeName: "node5", AppName: "app3"},
		{NodeName: "node2", AppName: "app4"},
		{NodeName: "node3", AppName: "app4"},
		{NodeName: "node4", AppName: "app4"},
		{NodeName: "node3", AppName: "app5"},
		{NodeName: "node5", AppName: "app5"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
		{"node4"},
		{"node5"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 1,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app3",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app4",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app5",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}

/*
group #:[ node1,node2,node3,]
group #:[ node4,node5,node6,]

*/
func GetCase6() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node3", AppName: "app1"},
		{NodeName: "node6", AppName: "app1"},
		{NodeName: "node5", AppName: "app1"},
		{NodeName: "node1", AppName: "app2"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node4", AppName: "app2"},
		{NodeName: "node6", AppName: "app2"},
		{NodeName: "node1", AppName: "app3"},
		{NodeName: "node3", AppName: "app3"},
		{NodeName: "node5", AppName: "app3"},
		{NodeName: "node6", AppName: "app3"},
		{NodeName: "node1", AppName: "app4"},
		{NodeName: "node2", AppName: "app4"},
		{NodeName: "node3", AppName: "app4"},
		{NodeName: "node4", AppName: "app4"},
		{NodeName: "node5", AppName: "app4"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
		{"node4"},
		{"node5"},
		{"node6"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 2,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 2,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 2,
			AppName: "app3",
		},
		{
			DisruptionAllowed: 3,
			AppName: "app4",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}

/*
group #:[ node2,node1,]
group #:[ node4,node5,]
group #:[ node6,node3,]

*/
func GetCase7() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node3", AppName: "app1"},
		{NodeName: "node4", AppName: "app1"},
		{NodeName: "node5", AppName: "app1"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node4", AppName: "app2"},
		{NodeName: "node6", AppName: "app2"},
		{NodeName: "node1", AppName: "app3"},
		{NodeName: "node3", AppName: "app3"},
		{NodeName: "node5", AppName: "app3"},
		{NodeName: "node2", AppName: "app4"},
		{NodeName: "node6", AppName: "app4"},
		{NodeName: "node3", AppName: "app5"},
		{NodeName: "node4", AppName: "app5"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
		{"node4"},
		{"node5"},
		{"node6"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 2,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app3",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app4",
		},
		{
			DisruptionAllowed: 1,
			AppName: "app5",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}

/*
node7上无app

group #:[ node1,node2,node3,node7,]
group #:[ node6,node4,node5,]

*/
func GetCase8() CaseDemo {
	apps := []Application{
		{NodeName: "node1", AppName: "app1"},
		{NodeName: "node3", AppName: "app1"},
		{NodeName: "node6", AppName: "app1"},
		{NodeName: "node5", AppName: "app1"},
		{NodeName: "node1", AppName: "app2"},
		{NodeName: "node2", AppName: "app2"},
		{NodeName: "node4", AppName: "app2"},
		{NodeName: "node6", AppName: "app2"},
		{NodeName: "node1", AppName: "app3"},
		{NodeName: "node3", AppName: "app3"},
		{NodeName: "node5", AppName: "app3"},
		{NodeName: "node6", AppName: "app3"},
		{NodeName: "node1", AppName: "app4"},
		{NodeName: "node2", AppName: "app4"},
		{NodeName: "node3", AppName: "app4"},
		{NodeName: "node4", AppName: "app4"},
		{NodeName: "node5", AppName: "app4"},
	}
	nodes := []Node{
		{"node1"},
		{"node2"},
		{"node3"},
		{"node4"},
		{"node5"},
		{"node6"},
		{"node7"},
	}
	budgets := []DisruptionBudget{
		{
			DisruptionAllowed: 2,
			AppName: "app1",
		},
		{
			DisruptionAllowed: 2,
			AppName: "app2",
		},
		{
			DisruptionAllowed: 2,
			AppName: "app3",
		},
		{
			DisruptionAllowed: 3,
			AppName: "app4",
		},
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}

func generateCase(nodeNum, appKinds int) CaseDemo {
	nodes := make([]Node, 0)
	for i := 0; i < nodeNum; i++ {
		nodes = append(nodes,Node{
			NodeName: "node" + strconv.Itoa(i),
		})
	}
	apps := make([]Application, 0)
	app2Count := make(map[int]int)
	for i := 0; i < appKinds; i++ {
		rand.Seed(time.Now().Unix())
		appNums := rand.Intn(appKinds - 1)
		app2Count[i] = appNums
		if appNums <= 0 { // 最少10个
			appNums = 1
		}
		record := make(map[int]int)
		for j := 0; j < appNums; j++ {
			nodeNo := 1
			for {
				rand.Seed(time.Now().Unix())
				nodeNo = rand.Intn(nodeNum)
				if nodeNo == 0 {
					nodeNo = 1
				}
				// 避免重复
				if _, ok := record[nodeNo]; ok {
					continue
				} else {
					break
				}
			}
			record[nodeNo] = 1
			apps = append(apps, Application{
				NodeName: "node" + strconv.Itoa(nodeNo),
				AppName: "app" + strconv.Itoa(j),
			})
		}
	}

	budgets := make([]DisruptionBudget, 0)

	for i := 0; i < appKinds; i++ {
		budgets = append(budgets, DisruptionBudget{
			DisruptionAllowed: int(math.Ceil(float64(app2Count[i]) * 0.2)),
			AppName:  "app" + strconv.Itoa(i),
		})
	}
	return CaseDemo{
		Nodes: nodes,
		Apps: apps,
		DisruptionBudgets: budgets,
	}
}
