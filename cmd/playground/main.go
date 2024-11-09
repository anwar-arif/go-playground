package main

type Tweet struct {
	userId  int
	tweetId int
}

type Twitter struct {
	follow map[int]map[int]bool
	tweets []Tweet
}

func Constructor() Twitter {
	return Twitter{
		follow: make(map[int]map[int]bool, 0),
		tweets: make([]Tweet, 0),
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.tweets = append(t.tweets, Tweet{
		userId:  userId,
		tweetId: tweetId,
	})
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	var feed []int
	for i := len(t.tweets) - 1; i >= 0; i-- {
		id := t.tweets[i].userId
		if id == userId || t.follow[userId][id] {
			feed = append(feed, id)
		}
		if len(feed) == 10 {
			break
		}
	}
	return feed
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	t.follow[followeeId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	t.follow[followeeId][followeeId] = false
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
