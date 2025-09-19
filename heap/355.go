package heap

import (
	"container/heap"
)

/*

355. Design Twitter
Design a simplified version of Twitter where users can post tweets,
follow/unfollow another user, and is able to see the 10 most recent tweets in the user's news feed.

Implement the Twitter class:

Twitter() Initializes your twitter object.
void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId.
Each call to this function will be made with a unique tweetId.
List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user themself.
Tweets must be ordered from most recent to least recent.
void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID
followeeId.
void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.


Example 1:

Input
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Output
[null, null, [5], null, null, [6, 5], null, [5]]

Explanation
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
twitter.follow(1, 2);    // User 1 follows user 2.
twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5].
// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.unfollow(1, 2);  // User 1 unfollows user 2.
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5],
//  since user 1 is no longer following user 2.


Constraints:

1 <= userId, followerId, followeeId <= 500
0 <= tweetId <= 104
All the tweets have unique IDs.
At most 3 * 104 calls will be made to postTweet, getNewsFeed, follow, and unfollow.
A user cannot follow himself.
*/
/*
 We need implement twits. need some structure to store tweets
 need store user subscriptions.
  Tweet:
   tweetID int primary key increase each time.
   userID int  - with posts the tweet.
  User:
   Following map[int]bool
 how to store tweets?
	store in tweets map key - tweetID, val: tweet structure with userID.
	getNews - get last 10 tweets by id.
		if Tweet.UserId is not in user followring then skip.
		else add to result.
	post:
		add tweet to tweets array.
		follow - add to following map.
		unfollow - remove from following map.
   How to store most recent tweet map ? and increase key?
   sotre tweets in maxHeap, use queue with tweetId.
	post: take root tweetId and increment for new tweet, and push new tweet to heap.
	get: take recent tweets from root and compare with userId. in following map. prepare result.
		until 10 or until end of tweets slice.

*/
// type Twit struct {
// 	Id        int
// 	UserId    int
// 	CreatedAt time.Time
// }
// type Twitter struct {
// 	Twits TwitHeap
// 	Users map[int]User
// }

// type User struct {
// 	Followings map[int]bool
// }

// func Constructor() Twitter {
// 	var pq = &TwitHeap{}
// 	heap.Init(pq)
// 	return Twitter{
// 		Twits: *pq,
// 		Users: make(map[int]User),
// 	}
// }

// func (this *Twitter) initUserIfNotExist(userId int) {
// 	if _, ok := this.Users[userId]; !ok {
// 		this.Users[userId] = User{Followings: make(map[int]bool, 0)}
// 	}
// }

// func (this *Twitter) PostTweet(userId int, tweetId int) {
// 	newTwit := Twit{Id: tweetId, UserId: userId, CreatedAt: time.Now()}
// 	heap.Push(&this.Twits, newTwit)
// }

// func (this *Twitter) GetNewsFeed(userId int) []int {
// 	this.initUserIfNotExist(userId)
// 	res := make([]int, 0, 10)
// 	temp := make(TwitHeap, len(this.Twits))
// 	copy(temp, this.Twits)
// 	heap.Init(&temp)

// 	for len(res) < 10 && temp.Len() > 0 {
// 		tw := heap.Pop(&temp).(Twit)
// 		if _, ok := this.Users[userId].Followings[tw.UserId]; ok || tw.UserId == userId {
// 			res = append(res, tw.Id)
// 		}
// 	}
// 	return res
// }

// func (this *Twitter) Follow(followerId int, followeeId int) {
// 	this.initUserIfNotExist(followerId)
// 	this.initUserIfNotExist(followeeId)
// 	this.Users[followerId].Followings[followeeId] = true

// }

// func (this *Twitter) Unfollow(followerId int, followeeId int) {
// 	this.initUserIfNotExist(followerId)
// 	this.initUserIfNotExist(followeeId)
// 	delete(this.Users[followerId].Followings, followeeId)
// }

// /**
//  * Your Twitter object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.PostTweet(userId,tweetId);
//  * param_2 := obj.GetNewsFeed(userId);
//  * obj.Follow(followerId,followeeId);
//  * obj.Unfollow(followerId,followeeId);
//  */

// type TwitHeap []Twit

// func (t TwitHeap) Len() int { return len(t) }
// func (t TwitHeap) Less(i, j int) bool {
// 	return t[i].CreatedAt.UTC().After(t[j].CreatedAt.UTC())
// }
// func (t TwitHeap) Swap(i, j int)       { t[i], t[j] = t[j], t[i] }
// func (t *TwitHeap) Push(x interface{}) { *t = append(*t, x.(Twit)) }
// func (t *TwitHeap) Pop() interface{} {
// 	old := *t
// 	n := len(old)
// 	x := old[n-1]
// 	*t = old[:n-1]
// 	return x
// }
/*
 This solution is almost correct but we got time limit exceeded.
*/

// This is optimal solution. Store each user twits and followers instead sotre twit user.s
var timestamp int

type Twit struct {
	Id        int
	UserId    int
	CreatedAt int
}

type Twitter struct {
	Users map[int]*User
}

type User struct {
	Followings map[int]bool
	Tweets     []Twit
}

func Constructor() Twitter {
	return Twitter{Users: make(map[int]*User)}
}

func (this *Twitter) ensureUser(userId int) {
	if _, ok := this.Users[userId]; !ok {
		this.Users[userId] = &User{
			Followings: make(map[int]bool),
			Tweets:     []Twit{},
		}
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.ensureUser(userId)
	timestamp++
	this.Users[userId].Tweets = append(
		this.Users[userId].Tweets,
		Twit{Id: tweetId, UserId: userId, CreatedAt: timestamp},
	)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.ensureUser(userId)

	pq := &TwitHeap{}
	heap.Init(pq)

	// add user tweets
	for i := len(this.Users[userId].Tweets) - 1; i >= 0 && i >= len(this.Users[userId].Tweets)-10; i-- {
		heap.Push(pq, this.Users[userId].Tweets[i])
	}

	// foolloers tweets
	for followeeId := range this.Users[userId].Followings {
		for i := len(this.Users[followeeId].Tweets) - 1; i >= 0 && i >= len(this.Users[followeeId].Tweets)-10; i-- {
			heap.Push(pq, this.Users[followeeId].Tweets[i])
		}
	}

	res := []int{}
	for pq.Len() > 0 && len(res) < 10 {
		tw := heap.Pop(pq).(Twit)
		res = append(res, tw.Id)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.ensureUser(followerId)
	this.ensureUser(followeeId)
	this.Users[followerId].Followings[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.ensureUser(followerId)
	delete(this.Users[followerId].Followings, followeeId)
}

// Heap implementation (max-heap by timestamp)
type TwitHeap []Twit

func (t TwitHeap) Len() int           { return len(t) }
func (t TwitHeap) Less(i, j int) bool { return t[i].CreatedAt > t[j].CreatedAt }
func (t TwitHeap) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t *TwitHeap) Push(x interface{}) {
	*t = append(*t, x.(Twit))
}
func (t *TwitHeap) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[:n-1]
	return x
}
