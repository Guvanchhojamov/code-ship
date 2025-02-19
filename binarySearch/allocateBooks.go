package binarysearch

/*
Problem Statement:
Given an array ‘arr of integer numbers,
‘ar[i]’ represents the number of pages in the ‘i-th’ book.
There are a ‘m’ number of students, and the task is to allocate all the books to the students.
Allocate books in such a way that:

 - Each student gets at least one book.
 - Each book should be allocated to only one student.
 - Book allocation should be in a contiguous manner.
 You have to allocate the book to ‘m’ students
 such that the maximum number of pages assigned to a student is minimum.

 If the allocation of books is not possible. return -1
*/

/*
  Example 1:
   Input Format:
   m = 2, arr[] = {12, 34, 67, 90}
   Result:
    113
    Explanation:
    The allocation of books will be 12, 34, 67 | 90. One student will get the first 3 books and the other will get the last one.
*/
/*
   Given books with pages, need to allocate them to the m students
    the max number of allocated pages to student, need to be minimum, of possiblitity.
    m = 4, arr[] = {25, 46, 28, 49, 24}
     Result: 71
     Explanation: The allocation of books will be 25, 46 | 28 | 49 | 24.

     Key points:
        - Each strdent need took at least 1 book.
        - Books need to be allocated in continious manner 1,2,3....n
         We cannot take 3 th book and give it to the student 1 before 1,2..
        - The maximum number of pages received by a student should be minimum possible.
        - If we can not allocate all books return -1
    Brute force:
        When we cannot allocate books? When the students > books. then return -1
        In other all ways we can allocate.
        what is the minimum can take 1 student?
            The max(arr) If each student takes 1 book, the maxumum number of pages would be max(arr).
        What is the max can take 1 student?
            The sum(arr) If we have 1 student we needd allocate all of books to this student.
        if 1 student  can take all of the books. The pages would be sum(arr)

        So the possible range is max(arr) .. sum(arr)
        We need helper function like:
         We need to check all possible nums from max(arr) ... sum(arr)
         func CountStudents(books, pages){
             pagesInStudent = a[0]
             allocatedStudentsCount = 1
             for i:=1 in books:
                if pagesInStudent+a[i] <= pages {
                       allocatedPages+=a[i]
                } else {
                    allocatedPages = a[i]
                    allocatedStudentsCount++
                }
             return allocatedStudentsCount
         }
         a = max(arr)
         b = sum(arr)
         for i:=a;a<b;a++{
             if CountStudents(books, pages) == m {
                 return i // Firs possible answer returned, and this is the minimum(max(student ca recieve))
             }
             return -1
         }
         Now we know range and what wew need we can use binary searh in there to optimze the time complexity.
*/
func allocabeBooks() {

}
