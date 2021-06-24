package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	GetSolution(GetCase3())
}

func GetSolution(caseDemo CaseDemo){
	apps := caseDemo.Apps
	nodes := caseDemo.Nodes
	budgets := caseDemo.DisruptionBudgets
	timer :=time.NewTicker(time.Minute * 2)
	for {
		select {
		case <-timer.C:
			fmt.Println("time out")
			break
		default:

		}
		err , cu := NewClusterUpgrade(CaseDemo{nodes, apps,  budgets})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if len(nodes) == 0 {
			break
		}
		err, resultMap := cu.GetMaxNodesToRestart()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// 打印一次结果即重启nodes
		fmt.Print("start to restart nodes: ")
		for k, _ := range resultMap {
			fmt.Print(k+",")
		}
		fmt.Println()
		time.Sleep(2 * time.Second)
		// 清理已重启node
		nodes, apps = cu.CleanNodes(resultMap)
	}
}

type ClusterUpgrade struct {
	NodeMap map[string]struct{} // 节点集合
	Node2Apps map[string][]string // 节点映射到的apps存储在map结构中
	//SelectNodeAppsCountMap map[string]int // 已选择的节点， app的数量记录
	AppsDisruptionAllowedNum map[string]int // 每个app保证可用性下最大可是失去的应用数
	// 缓存以i结尾的node最大可以重启的节点数
	DynamicMap map[int]Result // 动态规划的中间量保存
	Nodes []Node
	Apps []Application
}

func NewClusterUpgrade(caseDemo CaseDemo) (error, *ClusterUpgrade) {
	if len(caseDemo.Apps) == 0 {
		return nil, nil
	}
	if len(caseDemo.Nodes) == 0 {
		return errors.New("no node, so can not start"), nil
	}
	node2Apps := make(map[string][]string)
	nodeMap := make(map[string]struct{})
	for i := 0; i < len(caseDemo.Nodes); i++ {
		nodeMap[caseDemo.Nodes[i].NodeName] = struct{}{}
	}
	for i := 0; i < len(caseDemo.Apps); i++ {
		if _, ok := node2Apps[caseDemo.Apps[i].NodeName]; !ok {
			arr := make([]string, 0)
			node2Apps[caseDemo.Apps[i].NodeName] = arr
		}
		node2Apps[caseDemo.Apps[i].NodeName] = append(node2Apps[caseDemo.Apps[i].NodeName], caseDemo.Apps[i].AppName)
	}
	appsAllowedCountMap := make(map[string]int)
	for i := 0; i < len(caseDemo.DisruptionBudgets); i++ {
		appsAllowedCountMap[caseDemo.DisruptionBudgets[i].AppName] = caseDemo.DisruptionBudgets[i].DisruptionAllowed
	}
	return nil, &ClusterUpgrade{
		NodeMap: nodeMap,
		Node2Apps: node2Apps,
		//SelectNodeAppsCountMap : make(map[string]int),
		AppsDisruptionAllowedNum: appsAllowedCountMap,
		Nodes: caseDemo.Nodes,
		DynamicMap: make(map[int]Result),
		Apps: caseDemo.Apps,
	}
}

func (cu *ClusterUpgrade)CleanNodes(selectedNodesMap map[string]struct{}) ([]Node, []Application){
	nodes := make([]Node, 0)
	apps := make([]Application, 0)
	for i := 0; i < len(cu.Nodes); i++ {
		if _, ok := selectedNodesMap[cu.Nodes[i].NodeName]; !ok { // 不存在
			nodes = append(nodes, cu.Nodes[i])
		}
	}
	for i := 0; i < len(cu.Apps); i++ {
		if _, ok := selectedNodesMap[cu.Apps[i].NodeName]; !ok { // 不存在
			apps = append(apps, cu.Apps[i])
		}
	}
	return nodes, apps
}

func (cu *ClusterUpgrade)CheckNodeAvailability(nodeName string, apps []string, SelectNodeAppsCountMap map[string]int) bool {
	if len(nodeName) == 0 {
		return false
	}
	if len(apps) == 0 { // 说明节点上没有app，那该节点随时可重启
		return true
	}
	for i := 0; i < len(apps); i++ {
		if _, ok := SelectNodeAppsCountMap[apps[i]]; !ok {
			SelectNodeAppsCountMap[apps[i]] = 1
			continue
		}
		if (SelectNodeAppsCountMap[apps[i]] + 1) <= cu.AppsDisruptionAllowedNum[apps[i]] { // 未超出app预算
			SelectNodeAppsCountMap[apps[i]]++
		} else {// 超出预算
			//先回退再返回
			for j := 0; j < i; j++ {
				SelectNodeAppsCountMap[apps[j]]--
			}
			return false
		}
	}
	return true
}

// 运行一次，获取剩余节点中满足可用性可重启的最长节点串
func (cu *ClusterUpgrade)GetMaxNodesToRestart() (error, map[string]struct{}) {
	if len(cu.Node2Apps) == 0 { // 说明没有节点了
		return nil, nil
	}
	// 先塞入当前节点
	for i := 0; i < len(cu.Nodes); i++ {
		map1 := make(map[string]int)
		cu.CheckNodeAvailability(cu.Nodes[i].NodeName, cu.Node2Apps[cu.Nodes[i].NodeName], map1)
		cu.DynamicMap[i] = Result{NodeCount: 1, SelectNodeMap: map[string]struct{}{cu.Nodes[i].NodeName: struct {}{}}, SelectNodeAppsCountMap: map1}
	}
	for i := 0; i < len(cu.Nodes); i++ { // 遍历所有节点
		for j := 0; j < i; j++ {
			//if  _, ok := cu.DynamicMap[j].SelectNodeMap[cu.Nodes[j].NodeName]; ok {
			//	continue
			//}
			SelectNodeAppsCountMap := cu.DynamicMap[j].SelectNodeAppsCountMap // 第j个节点为结尾最长可用节点串的计数
			// 不存在
			if cu.CheckNodeAvailability(cu.Nodes[i].NodeName, cu.Node2Apps[cu.Nodes[i].NodeName], SelectNodeAppsCountMap) {
				nodesMap := cu.DynamicMap[j].SelectNodeMap // 以第j个节点结尾已经选择的可用节点串
				nodesMap[cu.Nodes[i].NodeName] = struct{}{} // 添加新节点
				if len(nodesMap) >= len(cu.DynamicMap[i].SelectNodeMap) {
					cu.DynamicMap[i] = Result{
						NodeCount: len(nodesMap),
						SelectNodeMap: nodesMap,
						SelectNodeAppsCountMap: SelectNodeAppsCountMap,
					}
				}
			}
		}
	}
	var result map[string]struct{}
	maxCount := 0
	for i := 0; i < len(cu.DynamicMap); i++ {
		if cu.DynamicMap[i].NodeCount > maxCount {
			maxCount = cu.DynamicMap[i].NodeCount
			result = cu.DynamicMap[i].SelectNodeMap
		}
	}
	return nil, result
}

type Result struct{
	NodeCount int
	SelectNodeMap map[string]struct{}
	SelectNodeAppsCountMap map[string]int // 已选择的节点， app的数量记录 appName->count

}