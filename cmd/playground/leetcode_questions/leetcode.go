package leetcode_questions

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"log"
)

type TopicTag struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Typename string `json:"__typename"`
}

type Question struct {
	Difficulty         string     `json:"difficulty"`
	ID                 int        `json:"id"`
	PaidOnly           bool       `json:"paidOnly"`
	QuestionFrontendID string     `json:"questionFrontendId"`
	Status             string     `json:"status"`
	Title              string     `json:"title"`
	TitleSlug          string     `json:"titleSlug"`
	TopicTags          []TopicTag `json:"topicTags"`
	IsInMyFavorites    bool       `json:"isInMyFavorites"`
	Frequency          float64    `json:"frequency"`
	AcRate             float64    `json:"acRate"`
	Typename           string     `json:"__typename"`
}

type QuestionListResponse struct {
	Data struct {
		FavoriteQuestionList struct {
			Questions []Question `json:"questions"`
		} `json:"favoriteQuestionList"`
	} `json:"data"`
}

func RunLeetCode() {
	csrf_token := "YZ5YcKTRHFmEcQcto3EfiOOBsPWkhmotGC8qM8hN6HG0yI1qRwIHfZMVpH4yJUV8"
	session_id := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMTM3ODk3MzYiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6ImIxMDhmNTdmNjEwNDcyN2MxZTc2NTAxOGU3MDdkODA4YmE5MjlmODczOWI1NzM4NzQ5MzAwNWQwZGM5ZmM0ZjQiLCJzZXNzaW9uX3V1aWQiOiIwM2U0NDNmMSIsImlkIjoxMzc4OTczNiwiZW1haWwiOiJjb2RlbGVldDIwMjRAZ21haWwuY29tIiwidXNlcm5hbWUiOiJjb2RlbGVldDIwMjQiLCJ1c2VyX3NsdWciOiJjb2RlbGVldDIwMjQiLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jb20vdXNlcnMvZGVmYXVsdF9hdmF0YXIuanBnIiwicmVmcmVzaGVkX2F0IjoxNzM1Mzg2MjIwLCJpcCI6IjExNi4yMDQuMTQ4LjE4OCIsImlkZW50aXR5IjoiMDg0NWIzMDljN2I5Yjk1N2FmZDllY2Y3NzVhNGMyMWYiLCJkZXZpY2Vfd2l0aF9pcCI6WyI1ZmNkZTVlMmFlNjhjMThjZDdhNmQ4M2EwYzgwMzJkYyIsIjExNi4yMDQuMTQ4LjE4OCJdLCJfc2Vzc2lvbl9leHBpcnkiOjEyMDk2MDB9.8nxM8B359eHxF4ZD2BThG_qk21kxDPfhqnTBhiz3grE"
	client := graphql.NewClient("https://leetcode.com/graphql")

	req := graphql.NewRequest(`
        query favoriteQuestionList($favoriteSlug: String!, $filter: FavoriteQuestionFilterInput, $filtersV2: QuestionFilterInput, $searchKeyword: String, $sortBy: QuestionSortByInput, $limit: Int, $skip: Int, $version: String = "v2") {
            favoriteQuestionList(
                favoriteSlug: $favoriteSlug
                filter: $filter
                filtersV2: $filtersV2
                searchKeyword: $searchKeyword
                sortBy: $sortBy
                limit: $limit
                skip: $skip
                version: $version
            ) {
                questions {
                    difficulty
                    id
                    paidOnly
                    questionFrontendId
                    status
                    title
                    titleSlug
                    topicTags {
                        name
                        slug
                        __typename
                    }
                    isInMyFavorites
                    frequency
                    acRate
                    __typename
                }
            }
        }
    `)

	// Headers
	req.Header.Set("Authority", "leetcode.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", fmt.Sprintf("LEETCODE_SESSION=%v; csrftoken=%v", session_id, csrf_token))
	//req.Header.Set("Origin", "https://leetcode.com")
	//req.Header.Set("Referer", "https://leetcode.com/problemset/")
	req.Header.Set("X-Csrf-Token", fmt.Sprintf("%v", csrf_token))

	// Set all required variables
	req.Var("skip", 0)
	req.Var("limit", 100)
	req.Var("favoriteSlug", "facebook-six-months")
	req.Var("filtersV2", map[string]interface{}{
		"filterCombineType": "ALL",
		"statusFilter": map[string]interface{}{
			"questionStatuses": []string{},
			"operator":         "IS",
		},
		"difficultyFilter": map[string]interface{}{
			"difficulties": []string{},
		},
		"languageFilter": map[string]interface{}{
			"languageSlugs": []string{},
			"operator":      "IS",
		},
		"topicFilter": map[string]interface{}{
			"topicSlugs": []string{},
			"operator":   "IS",
		},
	})
	req.Var("searchKeyword", "")
	req.Var("sortBy", map[string]interface{}{
		"sortField": "CUSTOM",
		"sortOrder": "ASCENDING",
	})

	var response QuestionListResponse
	if err := client.Run(context.Background(), req, &response); err != nil {
		log.Printf("GraphQL error: %v", err)
		return
	}

	// Print raw response for debugging
	fmt.Printf("Raw response: %+v\n", response)

	// Process response
	for _, q := range response.Data.FavoriteQuestionList.Questions {
		fmt.Printf("\nTitle Slug: %s\nDifficulty: %s\nFrequency: %.1f\nTopics: ",
			q.TitleSlug, q.Difficulty, q.Frequency)

		// Print topic slugs
		for i, tag := range q.TopicTags {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(tag.Slug)
		}
		fmt.Println()
	}
}
