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
	MinNodeCountToRestart int //当可重启节点数达到MinNodeCountToRestart是。立即返回当前最优结果
}

type Result struct{
	NodeCount int
	SelectNodeMap map[string]struct{} // 已经选择的节点
	SelectNodeAppsCountMap map[string]int // 已选择的节点， app的数量计数 appName->count， 便于做判断。
}

// 工具方法拷贝map
func CopySelectNodeMap(SelectNodeMap map[string]struct{} ) map[string]struct{}{
	ret := make(map[string]struct{})
	for k, v := range SelectNodeMap {
		ret[k] = v
	}
	return ret
}

// 工具方法拷贝map
func CopySelectNodeAppsCountMap(SelectNodeAppsCountMap map[string]int) map[string]int{
	ret := make(map[string]int)
	for k, v := range SelectNodeAppsCountMap {
		ret[k] = v
	}
	return ret
}