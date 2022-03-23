package worker

import (
	"encoding/json"
	"enterpret/errors"
	"enterpret/ingester"
	"enterpret/store"
	"enterpret/types"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Wroker interface {
	Pull()
}

type disListPosts struct {
	Posts []struct {
		ID                 int       `json:"id"`
		Name               string    `json:"name"`
		Username           string    `json:"username"`
		AvatarTemplate     string    `json:"avatar_template"`
		CreatedAt          time.Time `json:"created_at"`
		LikeCount          int       `json:"like_count"`
		Blurb              string    `json:"blurb"`
		PostNumber         int       `json:"post_number"`
		TopicTitleHeadline string    `json:"topic_title_headline"`
		TopicID            int       `json:"topic_id"`
	} `json:"posts"`
}

type disPosts struct {
	PostStream struct {
		Posts []disPost `json:"posts"`
	} `json:"post_stream"`
	ID int `json:"id"`
}

type disPost struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Username          string      `json:"username"`
	AvatarTemplate    string      `json:"avatar_template"`
	CreatedAt         time.Time   `json:"created_at"`
	Cooked            string      `json:"cooked"`
	PostNumber        int         `json:"post_number"`
	PostType          int         `json:"post_type"`
	UpdatedAt         time.Time   `json:"updated_at"`
	ReplyCount        int         `json:"reply_count"`
	ReplyToPostNumber interface{} `json:"reply_to_post_number"`
	QuoteCount        int         `json:"quote_count"`
	IncomingLinkCount int         `json:"incoming_link_count"`
	Reads             int         `json:"reads"`
	ReadersCount      int         `json:"readers_count"`
	Score             float64     `json:"score"`
	Yours             bool        `json:"yours"`
	TopicID           int         `json:"topic_id"`
	TopicSlug         string      `json:"topic_slug"`
}

type Discourse struct {
	endpoint string
	slug     string
	// date
}

func NewDiscourse() Wroker {
	return &Discourse{
		slug:     "discourse",
		endpoint: "https://meta.discourse.org/search.json?page=1&q=after%3A2021-01-01+before%3A2021-02-20",
	}
}

func (d *Discourse) Pull() {
	resp, err := http.Get(d.endpoint)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Println("Status code:", resp.StatusCode)
		return
	}
	var posts disListPosts
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		log.Println(err)
		return
	}

	for _, post := range posts.Posts {
		postURL := fmt.Sprintf("https://meta.discourse.org/t/%d/posts.json?post_ids=%d", post.TopicID, +post.ID)
		resp, err := http.Get(postURL)
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode > 299 {
			log.Println("Status code:", resp.StatusCode)
			return
		}
		var post disPosts
		if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
			log.Println(err)
			continue
		}
		if len(post.PostStream.Posts) == 0 {
			continue
		}
		if err := d.ingestPost(&post.PostStream.Posts[0]); err != nil {
			log.Println(err)
			continue
		}
	}
}

func (d *Discourse) ingestPost(ds *disPost) *errors.AppError {
	store := store.Store
	alertConfig, err := store.AlertConfig().GetBySlug(d.slug, "v1")
	if err != nil {
		return err
	}
	integration, err := store.GetIntegrationByID(1)
	if err != nil {
		return err
	}
	org, err := store.GetOrgByID(integration.OrganizationID)
	if err != nil {
		return err
	}

	ingestion := ingester.NewIngestion(org, integration, alertConfig)

	b, jerr := json.Marshal(ds)
	if jerr != nil {
		return nil
	}
	var data types.JSON
	if err := json.Unmarshal(b, &data); err != nil {
		return nil
	}

	alert := ingestion.Ingest(data)
	log.Println(alert)
	return nil
}
