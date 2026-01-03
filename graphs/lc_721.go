package graphs

import "sort"

/*
 721.Leetcode merge accounts problem...

*/

func accountsMerge(accounts [][]string) [][]string {
	ds := NewDisjointSet(len(accounts))
	mailToAcc := map[string]int{}
	for u, accMails := range accounts {
		for j := 1; j < len(accMails); j++ {
			if v, ok := mailToAcc[accMails[j]]; ok {
				ds.unionByRank(u, v)
			} else {
				mailToAcc[accMails[j]] = u
			}
		}
	}
	// fmt.Println("mtoa", mailToAcc) // mail -> acc number.

	// iterate over prev map, take each mail, and find() parent of curr acc.
	// add to map[parent_of_prev].append(email.)

	accMails := map[int][]string{} // acc number -> [] emails.
	for email, accNum := range mailToAcc {
		accParent := ds.findParent(accNum)
		accMails[accParent] = append(accMails[accParent], email)
	}
	//  fmt.Println(accMails)

	result := [][]string{}
	for accNum, emails := range accMails {
		name := accounts[accNum][0]
		sort.Strings(emails)
		result = append(result, append([]string{name}, emails...))
	}
	return result
}

/*
accounts = [
   0: ["John","johnsmith@mail.com","john_newyork@mail.com"],
   1: ["John","johnsmith@mail.com","john00@mail.com"],
   2: ["Mary","mary@mail.com"],
   3: ["John","johnnybravo@mail.com"]
]
Output: [
    ["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],
    ["Mary","mary@mail.com"],
    ["John","johnnybravo@mail.com"]
]

 ok, if we see we have some connections, between something..
 Its looks lik graph, so we need to think in terms of graph.
 each acc - node.
 each email - edge.
 when its connection and changing dynamically, it comes up with UF solution.
 But we have also BF solution, using DFS.
 How?
    1. Each same mail and diff accounts mean we have connection between them.
    use mail_to_accs map [string][]int.
    sotore accounts it belong for each mail in map.
    like this we can build Adjecency list for graph.
    2. Create visited[] map aslo for emails.
    3. Do DFS from each node, since the accounts may/or may not be connected.
        iterate each neightbour too for this mail.
        if neigbour is not visited do DFS for him also.
 TC: for each acc we do DFS. N*M+NLOGN.
 SC: N+M +tmp emails set.
But, it can be solved easyly with UF algo.
 So, what is UF do?
    He merges some nodes under the one parent, if they are not.
    He finds ultimate parent for each node.
node - account number.
email - is edge.
How we know need we merge or not?
    using hash map - email : account_number.
Iterate mails and store mail:acc.
 if we sotre already this mail with other acc.
    Then we need merge them.
    union(currAcc, AccFromMap)
 Now the accounts belong to the one account. A->B or B->A
 Now how we create response array?
  WE need add all emails from child account to parent account.
  map each leader acc : []emails.

*/
