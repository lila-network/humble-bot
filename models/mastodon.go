/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package models

import "time"

type MastodonPost struct {
	ID                 string    `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	InReplyToID        string    `json:"in_reply_to_id"`
	InReplyToAccountID string    `json:"in_reply_to_account_id"`
	Sensitive          bool      `json:"sensitive"`
	SpoilerText        string    `json:"spoiler_text"`
	Visibility         string    `json:"visibility"`
	Language           string    `json:"language"`
	URI                string    `json:"uri"`
	URL                string    `json:"url"`
	RepliesCount       int       `json:"replies_count"`
	ReblogsCount       int       `json:"reblogs_count"`
	FavouritesCount    int       `json:"favourites_count"`
	Favourited         bool      `json:"favourited"`
	Reblogged          bool      `json:"reblogged"`
	Muted              bool      `json:"muted"`
	Bookmarked         bool      `json:"bookmarked"`
	Content            string    `json:"content"`
	Reblog             any       `json:"reblog,omitempty"`
	Application        struct {
		Name    string `json:"name"`
		Website string `json:"website,omitempty"`
	} `json:"application"`
	Account struct {
		ID             string    `json:"id"`
		Username       string    `json:"username"`
		Acct           string    `json:"acct"`
		DisplayName    string    `json:"display_name"`
		Locked         bool      `json:"locked"`
		Bot            bool      `json:"bot"`
		Discoverable   bool      `json:"discoverable"`
		Group          bool      `json:"group"`
		CreatedAt      time.Time `json:"created_at"`
		Note           string    `json:"note"`
		URL            string    `json:"url"`
		Avatar         string    `json:"avatar"`
		AvatarStatic   string    `json:"avatar_static"`
		Header         string    `json:"header"`
		HeaderStatic   string    `json:"header_static"`
		FollowersCount int       `json:"followers_count"`
		FollowingCount int       `json:"following_count"`
		StatusesCount  int       `json:"statuses_count"`
		LastStatusAt   time.Time `json:"last_status_at"`
		Emojis         []any     `json:"emojis,omitempty"`
		Fields         []struct {
			Name       string `json:"name"`
			Value      string `json:"value"`
			VerifiedAt any    `json:"verified_at,omitempty"`
		} `json:"fields"`
	} `json:"account"`
	MediaAttachments []any `json:"media_attachments,omitempty"`
	Mentions         []any `json:"mentions,omitempty"`
	Tags             []any `json:"tags,omitempty"`
	Emojis           []any `json:"emojis,omitempty"`
	Card             struct {
		URL          string `json:"url"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		Type         string `json:"type"`
		AuthorName   string `json:"author_name"`
		AuthorURL    string `json:"author_url"`
		ProviderName string `json:"provider_name"`
		ProviderURL  string `json:"provider_url"`
		HTML         string `json:"html"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		Image        any    `json:"image,omitempty"`
		EmbedURL     string `json:"embed_url"`
	} `json:"card"`
	Poll any `json:"poll,omitempty"`
}
