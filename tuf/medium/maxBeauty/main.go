package main

/*
 lc - 2070.

*/

func main() {
	// call func.
}

func maximumBeauty(items [][]int, queries []int) []int {
	var pricesMap = make(map[int]int, 0)
	var maxBeauty = 0
	var result = make([]int, 0, len(queries))
	for _, item := range items {
		pricesMap[item[0]] = max(pricesMap[item[0]], item[1])
	}

	for _, query := range queries {
		maxBeauty = 0
		for price, mBeauty := range pricesMap {
			if price <= query {
				maxBeauty = max(maxBeauty, mBeauty)
			}
		}
		result = append(result, maxBeauty)
	}
	return result
}

/*
1. use mapOf nums like map[price]=max_beuaty just store max for each price
2. check each query iterating this map.
 query=key=price
 for query in range queries:
    maxBeauty=0
    if map[query]:
      maxBeauty = map[query] is max price already.
    else  {
        for i,val in range mapNums :
            if  price < query :
                maxBeauty = max(maxBeauty,map[query])
    }
    res = append(res, maxxBeauty)
end
*/
