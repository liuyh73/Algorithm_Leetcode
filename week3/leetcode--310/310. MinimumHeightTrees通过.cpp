class Solution {
public:
	vector<int> findMinHeightTrees(int n, vector<pair<int, int>>& edges){
		queue<int>nodes;
		int degree[n] = {0};
		bool searched[n]={0};
		unordered_map<int, vector<int>> graph;
		for(auto edge : edges) {
			degree[edge.first]++;
			degree[edge.second]++;
			graph[edge.first].push_back(edge.second);
			graph[edge.second].push_back(edge.first);
		}
		for(int i=0;i<n;i++){
			if(degree[i] ==1 || degree[i] == 0){
				searched[i] = true;
				nodes.push(i);
			}
		}

        vector<int>roots;
		while(nodes.size()){
			int size = nodes.size();
			roots.clear();
			while(size>0){
				int node = nodes.front();
				nodes.pop();
				size--;
				roots.push_back(node);
				for(auto anotherNode : graph[node]){
					if(!searched[anotherNode]){
						if(--degree[anotherNode] == 0 || degree[anotherNode] == 1){
							searched[anotherNode] = true;
							nodes.push(anotherNode);			
						}
					}
				}
			}
		}
		return roots;
	}
};