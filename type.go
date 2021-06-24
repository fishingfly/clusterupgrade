package main

type DisruptionBudget struct {
	DisruptionAllowed int
	AppName string
}

type Application struct {
	NodeName string
	AppName string
}

type Node struct {
	NodeName string
}

type CaseDemo struct {
	Nodes []Node
	Apps []Application
	DisruptionBudgets []DisruptionBudget
}

