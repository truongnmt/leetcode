func findItinerary(tickets [][]string) []string {
    adj := make(map[string][]string)
    itineraries := []string{}
    stack := []string{"JFK"}
    
    for _,t := range(tickets) {
        adj[t[0]] = append(adj[t[0]], t[1])
    }
    for _,a := range(adj) {
        sort.Strings(a)
    }
    
    for len(stack)>0{
        src := stack[len(stack)-1]
        if _,ok := adj[src]; ok && len(adj[src])>0 {
            dest := adj[src][0]
            stack = append(stack, dest)
            adj[src] = adj[src][1:]
        } else {
            itineraries = append(itineraries, src)
            stack = stack[:(len(stack)-1)]
            
        }
    }
    
    for i:=0;i<len(itineraries)/2;i++ {
        itineraries[i], itineraries[len(itineraries)-i-1] = itineraries[len(itineraries)-i-1], itineraries[i]
    }
    
    return itineraries
}
