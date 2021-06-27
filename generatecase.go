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
		MinNodeCountToRestart: -1,
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
		MinNodeCountToRestart: -1,
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
		MinNodeCountToRestart: -1,
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
		MinNodeCountToRestart: -1,
	}
}

// 随机产生需要的apps和nodes, 其中app1的个数在10到200之间
func generateCase(nodeNum, appKinds, minNodeCountToRestart int) CaseDemo {
	nodes := make([]Node, 0)
	node2AppCount := make(map[int]int)
	// 产生node集合
	for i := 0; i < nodeNum; i++ {
		node2AppCount[i] = 0
		nodes = append(nodes,Node{
			NodeName: "node" + strconv.Itoa(i),
		})
	}
	// 产生app集合
	apps := make([]Application, 0)
	app2Count := make(map[int]int)
	for i := 0; i < appKinds; i++ { // app 种类
		rand.Seed(time.Now().Unix())
		appNums := rand.Intn(200) // 每个app的运行的实例数
		if appNums < 10 { // 最少10个
			appNums = 10
		}
		app2Count[i] = appNums
		record := make(map[int]int) // <nodei, 1>，表示appi是否已经在node上，为保证app的每个实例都飘在不同的节点上
		nodeNo := 1
		for j := 0; j < appNums; j++ {
			rand.Seed(time.Now().Unix())
			nodeNo = rand.Intn(nodeNum) // 随机一个node节点
			for {// 为保证app的每个实例都飘在不同的节点上
				if nodeNo == 0 {
					nodeNo = 1
				}
				// 节点上app个数不能超过50个
				if node2AppCount[nodeNo] >= 50 {
					nodeNo++
					continue
				}
				// 避免重复
				if _, ok := record[nodeNo]; ok { // 已经存在了
					nodeNo++
					continue
				} else {
					break
				}
			}
			record[nodeNo] = 1
			node2AppCount[nodeNo]++
			apps = append(apps, Application{
				NodeName: "node" + strconv.Itoa(nodeNo),
				AppName: "app" + strconv.Itoa(i),
			})
		}
	}
	// 产生预算集合
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
		MinNodeCountToRestart: minNodeCountToRestart,
	}
}
