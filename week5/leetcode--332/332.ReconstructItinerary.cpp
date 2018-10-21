class Solution {
public:
    vector<string> findItinerary(vector<pair<string, string>> tickets) {
        unordered_map<string,multiset<string>> graph;
        vector<string>itinerary;
        for(auto ticket:tickets) {
            graph[ticket.first].insert(ticket.second);
        }
        // for(auto iter=graph.begin(); iter!=graph.end(); iter++){
        //     cout<<iter.first<<": ";
        //     for(auto des:iter->second) {
        //         cout<<des<<" ";
        //     }
        //     cout<<endl;
        // }
        itinerary.push_back("JFK");
        while(graph.size()!=0) {
            int tail = itinerary.size()-1;
            mutiset<string>::iterator iter = graph[itinerary[tail]].begin();
            string destination = "";
            destination=*iter;
            if(graph[itinerary[tail]].size() == 1){
                graph.erase(itinerary[tail]);
            } else {
                graph[itinerary[tail]].erase(iter);
            }
            if(destination!="")
                itinerary.push_back(destination);
        }
        return itinerary;
    }
};