class Solution {
public:
    int dfs(int node, int n, unordered_map<int, vector<int>>& graph, bool* searched){
		int maxx = 1, temp;
		for (auto anotherNode : graph[node]) {
			if (!searched[anotherNode]) {
				searched[anotherNode] = true;
				temp = dfs(anotherNode, n, graph, searched)+1;
				if(maxx < temp)	maxx = temp;
			}
		}
		return maxx;
	}
	vector<int> findMinHeightTrees(int n, vector<pair<int, int>>& edges){
		int heights[n]={0}, minHeight=INT_MAX;
		vector<int>minHeightNodes;
		unordered_map<int, vector<int>> graph;
		for(auto edge : edges) {
			graph[edge.first].push_back(edge.second);
			graph[edge.second].push_back(edge.first);
		}

		for(int i=0;i<n;i++){
			bool searched[n]={0};
			searched[i]=true;
			heights[i] = dfs(i, n, graph, searched);
			if(minHeight > heights[i]){
				minHeight = heights[i];
			}
		}
		for(int i=0;i<n;i++){
			if(minHeight == heights[i])
				minHeightNodes.push_back(i);
		}
		return minHeightNodes;
	}
};