package heap

/*
 given heap, push and pop elements to the heap,
 and maintain heap property.
*/
/*
 PUSH.
1. insert to after last element to heap.
2. After insert heapify from bottom to top.
	until root or parent is greater, to keep heap property.
this cals also calculate up algorythm.
tc: LogN
sc: logn
*/
/*
1. POP().
1. Take root element.
2. Copy last element to root element.
3. Remove last element.
4. Heapify algo from top to bottom.
logN from top to end. calculate bottom algorythm.
tc: LogN.
*/
