package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//GetOptimalUpgradePlans(GetCase8())
	GetOptimalUpgradePlans(generateCase(100, 40))
}

func GetOptimalUpgradePlans(caseDemo CaseDemo){
	apps := caseDemo.Apps
	nodes := caseDemo.Nodes
	budgets := caseDemo.DisruptionBudgets
	startTime := time.Now()
	timer :=time.NewTicker(time.Minute * 10)
	for {
		select {
		case <-timer.C:
			// 超时退出
			fmt.Println("time out, failed to GetOptimalUpgradePlans ")
			break
		default:

		}
		err , cu := NewClusterUpgrade(CaseDemo{nodes, apps,  budgets})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if len(nodes) == 0 { // 剪枝无节点就退出
			break
		}
		if len(apps) == 0 { // 剪枝，有节点，但是无app了，那么重启方案就是所有节点
			fmt.Print("group #:[ ")
			for _, v := range nodes {
				fmt.Print(v.NodeName + ",")
			}
			fmt.Print("]\n")
		}
		err, resultMap := cu.GetMaxNodesToRestart()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// 打印一次结果即重启nodes
		fmt.Print("group #:[ ")
		for k, _ := range resultMap {
			fmt.Print(k+",")
		}
		fmt.Print("]\n")
		// 清理已重启node
		nodes, apps = cu.CleanNodes(resultMap)
	}
	fmt.Println(time.Now().Sub(startTime).Microseconds())
}

type ClusterUpgrade struct {
	NodeMap map[string]struct{} // 节点集合
	Node2Apps map[string][]string // 节点映射到的apps存储在map结构中
	//SelectNodeAppsCountMap map[string]int // 已选择的节点， app的数量记录
	AppsDisruptionAllowedNum map[string]int // 每个app保证可用性下最大可是失去的应用数
	// 缓存以i结尾的node最大可以重启的节点数
	DynamicMap map[int][]Result // 动态规划的中间量保存,<第i个节点,重启节点集>
	NodeIResult Result // 记录从节点0到节点i之间可重启的最大节点集
	Nodes []Node
	Apps []Application
	// 存放最有结果结果
	RestartPolicyResult []string
	MinRestartCount int
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
		DynamicMap: make(map[int][]Result),
		Apps: caseDemo.Apps,
		MinRestartCount: len(nodeMap),
		RestartPolicyResult: make([]string, 0),
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
	SelectNodeAppsCountMapTemp := SelectNodeAppsCountMap
	for i := 0; i < len(apps); i++ {
		if _, ok := SelectNodeAppsCountMapTemp[apps[i]]; !ok {
			SelectNodeAppsCountMapTemp[apps[i]] = 1
			continue
		}
		if (SelectNodeAppsCountMapTemp[apps[i]] + 1) <= cu.AppsDisruptionAllowedNum[apps[i]] { // 未超出app预算
			SelectNodeAppsCountMapTemp[apps[i]]++
			continue
		} else {// 超出预算
			return false
		}
	}
	return true
}

// 加入一个节点
func (cu *ClusterUpgrade)addAppCount(nodeName string, apps []string, SelectNodeAppsCountMap map[string]int) {
	for i := 0; i < len(apps); i++ {
		if _, ok := SelectNodeAppsCountMap[apps[i]]; !ok {
			SelectNodeAppsCountMap[apps[i]] = 1
			continue
		}
		if (SelectNodeAppsCountMap[apps[i]] + 1) <= cu.AppsDisruptionAllowedNum[apps[i]] { // 未超出app预算
			SelectNodeAppsCountMap[apps[i]]++
			continue
		}
	}
}

// 回退一个节点
func (cu *ClusterUpgrade)minusAppCount(nodeName string, apps []string, SelectNodeAppsCountMap map[string]int) {
	for i := 0; i < len(apps); i++ {
		SelectNodeAppsCountMap[apps[i]]--
	}
}

// 运行一次，获取剩余节点中满足可用性可重启的最长节点串
func (cu *ClusterUpgrade)GetMaxNodesToRestart() (error, map[string]struct{}) {
	if len(cu.Node2Apps) == 0 { // 说明没有节点了
		return nil, nil
	}
	if len(cu.Nodes) == 1 { // 只有一个节点
		return nil, map[string]struct{}{cu.Nodes[0].NodeName: {}}
	}
	// 先向dp table中塞入当前节点，每次只重启一个是肯定可以的
	for i := 0; i < len(cu.Nodes); i++ {
		map1 := make(map[string]int)
		cu.addAppCount(cu.Nodes[i].NodeName, cu.Node2Apps[cu.Nodes[i].NodeName], map1)
		results := []Result{{NodeCount: 1, SelectNodeMap: map[string]struct{}{cu.Nodes[i].NodeName: struct{}{}}, SelectNodeAppsCountMap: map1}}
		cu.DynamicMap[i] = results
	}
	maxRestartNodeCount := 1
	var maxResult Result = cu.DynamicMap[0][0] // 给个默认值
	for i := 0; i < len(cu.Nodes); i++ { // 遍历所有节点
		for j := 0; j < i; j++ {
			results := cu.DynamicMap[j]
			for index_result := 0; index_result < len(results); index_result++ {
				tempMap := CopySelectNodeAppsCountMap(results[index_result].SelectNodeAppsCountMap)
				// 剪枝
				if cu.CheckNodeAvailability(cu.Nodes[i].NodeName, cu.Node2Apps[cu.Nodes[i].NodeName], tempMap) {
					selectNodeAppsCountMapTemp := CopySelectNodeAppsCountMap(results[index_result].SelectNodeAppsCountMap)
					newNodeCount := results[index_result].NodeCount + 1
					newNodeMap := CopySelectNodeMap(results[index_result].SelectNodeMap)
					newNodeMap[cu.Nodes[i].NodeName] = struct{}{}
					cu.addAppCount(cu.Nodes[i].NodeName, cu.Node2Apps[cu.Nodes[i].NodeName], selectNodeAppsCountMapTemp)
					resultTemp := Result{NodeCount: newNodeCount, SelectNodeMap: newNodeMap, SelectNodeAppsCountMap: selectNodeAppsCountMapTemp}
					// 保存以node[i]为结尾的可行方案
					if maxRestartNodeCount < newNodeCount {
						maxResult = resultTemp
						maxRestartNodeCount = newNodeCount
					}
					cu.DynamicMap[i] = append(cu.DynamicMap[i], resultTemp)
				}
			}

		}
	}
	return nil, maxResult.SelectNodeMap
}

type Result struct{
	NodeCount int
	SelectNodeMap map[string]struct{} // 已经选择的节点
	SelectNodeAppsCountMap map[string]int // 已选择的节点， app的数量计数 appName->count， 便于做判断。
}

func CopySelectNodeMap(SelectNodeMap map[string]struct{} ) map[string]struct{}{
	ret := make(map[string]struct{})
	for k, v := range SelectNodeMap {
		ret[k] = v
	}
	return ret
}

func CopySelectNodeAppsCountMap(SelectNodeAppsCountMap map[string]int) map[string]int{
	ret := make(map[string]int)
	for k, v := range SelectNodeAppsCountMap {
		ret[k] = v
	}
	return ret
}