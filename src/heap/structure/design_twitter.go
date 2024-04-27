package structure

import (
	"sort"
)

/*
Problem:
Ref: https://leetcode.com/problems/design-twitter/description/

Design:
Simple Twitter:
- void postTweet(int userId, int tweetId):
Composes a new tweet with ID tweetId by the user userId.
Each call to this function will be made with a unique tweetId.
- List<Integer> getNewsFeed(int userId):
Retrieves the 10 most recent tweet IDs in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the users themselves.
Tweets must be ordered from most recent to least recent.
- void follow(int followerId, int followeeId):
The user with ID followerId started following the user with ID followeeId.
- void unfollow(int followerId, int followeeId):
The user with ID followerId started unfollowing the user with ID followeeId.

Constraints:
- userId, followerId, followeeId: [1, 500]
- tweetId: [1, 10^4]
- all tweets have unique ids
- At most 3 * 104 calls will be made to postTweet, getNewsFeed, follow, and unfollow
*/

/*
First approach:
In order to get news feed, get all the tweets from user and his/her followees,
Sort them and return top 10
*/

type Tweet struct {
	Id       int
	PostedAt int
}

type Twitter struct {
	TweetCount int
	UserTweets [][]*Tweet
	Followings []map[int]bool
}

func TwitterConstructor() Twitter {
	return Twitter{
		TweetCount: 0,
		UserTweets: make([][]*Tweet, 500),
		Followings: make([]map[int]bool, 500),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := &Tweet{
		Id:       tweetId,
		PostedAt: this.TweetCount,
	}
	this.TweetCount++
	this.UserTweets[userId] = append(this.UserTweets[userId], tweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	var tweets []*Tweet
	for _, tweet := range this.UserTweets[userId] {
		tweets = append(tweets, tweet)
	}

	for followeeId, _ := range this.Followings[userId] {
		for _, tweet := range this.UserTweets[followeeId] {
			tweets = append(tweets, tweet)
		}
	}

	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].PostedAt > tweets[j].PostedAt
	})

	var result []int
	for i := 0; i < 10; i++ {
		if i >= len(tweets) {
			break
		}

		result = append(result, tweets[i].Id)
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	followees := this.Followings[followerId]
	if followees != nil {
		followees[followeeId] = true
	} else {
		followees = make(map[int]bool)
		followees[followeeId] = true
		this.Followings[followerId] = followees
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.Followings[followerId] != nil {
		delete(this.Followings[followerId], followeeId)
	}
}
