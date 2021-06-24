package main

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

// node1, node4
// node2,node5
// node3
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
