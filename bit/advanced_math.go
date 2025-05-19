package bit

/*
  print all divisiors of the number.
	brute force go until 1-> n and if div i ==0 add to list return.
But we can optimize this going until sqrt(n).
for i =0 ; i*i < n ; i++:
	if n % i== 0:
		list.add(i)
	if i != n/i:
		list.add(i)
tc: sqrt(n); sc:1
*/

/*
	 Print prime factors of given number.
	 bf: take all divisors, and chech for prime:
		tc: sqrt(n) * sqrt(n) = N
*/
func primeDivisors(n int) []int {
	var res = []int{n}

	for i := 2; i*i <= n; i++ { // go until sqrt(i)- same i*i<n
		for n%i == 0 {
			res = append(res, i)
			n = n / i
		}
	}
	// add the last prime number to answer.
	if n != 1 {
		res = append(res, n)
	}
	return res
}

/*
 780/2 =390
 390/2 = 195
 195/2 - not divisible, i++
 195/3=
*/
