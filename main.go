package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//GetOptimalUpgradePlans(GetCase1())
	//GetOptimalUpgradePlans(GetCase2())
	//GetOptimalUpgradePlans(GetCase3())
	//GetOptimalUpgradePlans(GetCase4())
	//GetOptimalUpgradePlans(GetCase5())
	GetOptimalUpgradePlans(GetCase6())
	//GetOptimalUpgradePlans(GetCase7())
	//GetOptimalUpgradePlans(GetCase8())
	//GetOptimalUpgradePlans(generateCase(1000, 800, -1))
	//GetOptimalUpgradePlans(generateCase(2000, 1600, -1))
	//GetOptimalUpgradePlans(generateCase(5000, 4000, 160))
}

func GetOptimalUpgradePlans(caseDemo CaseDemo){
	apps := caseDemo.Apps
	nodes := caseDemo.Nodes
	budgets := caseDemo.DisruptionBudgets
	timer :=time.NewTicker(time.Minute * 10)
	startTime := time.Now()
	restartCount := 0
	for {
		select {
		case <-timer.C:
			// 超时退出
			fmt.Println("time out, failed to GetOptimalUpgradePlans ")
			break
		default:

		}
		err , cu := NewClusterUpgrade(CaseDemo{nodes, apps,  budgets, caseDemo.MinNodeCountToRestart})
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if len(nodes) == 0 { // 无节点就退出
			break
		}
		if len(apps) == 0 { // 有节点，但是无app了，那么重启方案就是所有节点
			restartCount++
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
		// 打印结果即这一批次需要重启nodes
		fmt.Print("group #:[ ")
		for k, _ := range resultMap {
			fmt.Print(k+",")
		}
		fmt.Printf("], total number of nodes: %d\n", len(resultMap))
		restartCount++
		// 清理已重启node
		nodes, apps = cu.CleanNodes(resultMap)
		if len(resultMap) == 1 { // 剪枝，当resultMap长度为1时，说明剩下的节点都只能单独启动，为减少计算，直接给出结果
			for i:= 0; i < len(nodes); i++ {
				restartCount++
				fmt.Print("group #:[ ")
				fmt.Printf("%s], total number of nodes: %d\n", nodes[i].NodeName, 1)
			}
			break
		}
	}
	fmt.Printf("the total restart time: %d \n", restartCount)
	fmt.Println("start time is " + startTime.Format("2006-01-02 15:04:05"))
	fmt.Println("end time is " + time.Now().Format("2006-01-02 15:04:05"))
}

type ClusterUpgrade struct {
	NodeMap map[string]struct{} // 节点集合
	Node2Apps map[string][]string // 节点映射到的apps存储在map结构中
	AppsDisruptionAllowedNum map[string]int // 每个app保证可用性下最大可是失去的应用数
	DynamicMap map[int][]Result // 动态规划的中间量保存,<第i个节点,重启节点集>, // 缓存以i结尾的node最大可以重启的节点数
	NodeIResult Result // 记录从节点0到节点i之间可重启的最大节点集
	Nodes []Node // 所有节点
	Apps []Application // 所有app
	MinNodeCountToRestart int //当可重启节点数达到MinNodeCountToRestart是。立即返回当前最优结果
}

func NewClusterUpgrade(caseDemo CaseDemo) (error, *ClusterUpgrade) {
	if len(caseDemo.Apps) == 0 { // 没有app
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
		AppsDisruptionAllowedNum: appsAllowedCountMap,
		Nodes: caseDemo.Nodes,
		DynamicMap: make(map[int][]Result),
		Apps: caseDemo.Apps,
		MinNodeCountToRestart: caseDemo.MinNodeCountToRestart,
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
	if len(apps) == 0 { // 说明节点上没有app，那该节点随意可重启
		return true
	}
	SelectNodeAppsCountMapTemp := SelectNodeAppsCountMap // 已选择的节点中app的计数<appi, count>
	for i := 0; i < len(apps); i++ {
		if _, ok := SelectNodeAppsCountMapTemp[apps[i]]; !ok {
			SelectNodeAppsCountMapTemp[apps[i]] = 1
			continue
		}
		if (SelectNodeAppsCountMapTemp[apps[i]] + 1) <= cu.AppsDisruptionAllowedNum[apps[i]] { // 未超出app预算
			SelectNodeAppsCountMapTemp[apps[i]]++ //+1
			continue
		} else {// 超出预算
			return false
		}
	}
	return true
}

// 加入一个节点，修改SelectNodeAppsCountMap
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

// 获取剩余节点中满足可用性可重启的最长节点串
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
		nodeiMaxNodeCount := 1 // 最小是1
		for j := 0; j < i; j++ {
			results := cu.DynamicMap[j]
			for index_result := 0; index_result < len(results); index_result++ { // 遍历i之前[0, i-1]与node[i]一起重启，检查是否满足可用性
				tempMap := CopySelectNodeAppsCountMap(results[index_result].SelectNodeAppsCountMap)
				if cu.CheckNodeAvailability(cu.Nodes[i].NodeName, cu.Node2Apps[cu.Nodes[i].NodeName], tempMap) {
					// 只考虑满足可用性的results[j]与nodes[i]一起重启
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
					if newNodeCount > nodeiMaxNodeCount {
						nodeiMaxNodeCount = newNodeCount
					}
					if cu.MinNodeCountToRestart <= resultTemp.NodeCount && cu.MinNodeCountToRestart > 0{ // 剪枝，节省时间
						return nil, maxResult.SelectNodeMap
					}
					cu.DynamicMap[i] = []Result{resultTemp} // 每个位置保证最大的子串
				}
			}

		}
	}
	return nil, maxResult.SelectNodeMap
}