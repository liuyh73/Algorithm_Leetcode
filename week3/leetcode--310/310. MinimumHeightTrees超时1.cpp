class Solution {
public:
    int dfs(int node, int n, vector<pair<int, int>>& edges, bool* searched){
		int maxx = 1, temp;
		for (auto edge : edges) {
			int anotherNode = node;
			if (edge.first == node)			anotherNode = edge.second;
			else if (edge.second == node)	anotherNode = edge.first;
			if (!searched[anotherNode]) {
				searched[anotherNode] = true;
				temp = dfs(anotherNode, n, edges, searched)+1;
				if(maxx < temp)	maxx = temp;
			}
		}
		return maxx;
	}
	vector<int> findMinHeightTrees(int n, vector<pair<int, int>>& edges){
		int heights[n]={0}, minHeight=INT_MAX;
		vector<int>minHeightNodes;
		for(int i=0;i<n;i++){
			bool searched[n]={0};
			searched[i]=true;
			heights[i] = dfs(i, n, edges, searched);
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