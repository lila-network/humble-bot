/*
 * Copyright (c) 2023 Laura Kalb <dev@lauka.net>
 * The code of this project is available under the MIT license. See the LICENSE file for more info.
 *
 */

package models

type Bundle struct {
	TileLogoInformation struct {
		Config struct {
			ImageType string `json:"image_type"`
			Gcs       string `json:"gcs"`
			Imgix     struct {
				Args struct {
				} `json:"args"`
				MasterImage struct {
					ImageType string `json:"image_type"`
					Gcs       string `json:"gcs"`
					Static    string `json:"static"`
					Imgix     struct {
					} `json:"imgix"`
				} `json:"master_image"`
			} `json:"imgix"`
		} `json:"config"`
	} `json:"tile_logo_information"`
	MachineName      string `json:"machine_name"`
	HighResTileImage string `json:"high_res_tile_image"`
	DisableHeroTile  bool   `json:"disable_hero_tile"`
	MarketingBlurb   string `json:"marketing_blurb"`
	HoverTitle       string `json:"hover_title"`
	ProductURL       string `json:"product_url"`
	TileImage        string `json:"tile_image"`
	Category         string `json:"category"`
	HeroHighlights   []struct {
		Heading string `json:"heading"`
		Tooltip string `json:"tooltip"`
	} `json:"hero_highlights"`
	HoverHighlights             []string `json:"hover_highlights"`
	Author                      string   `json:"author"`
	FallbackStoreSaleLogo       string   `json:"fallback_store_sale_logo"`
	HighResTileImageInformation struct {
		Config struct {
			ImageType string `json:"image_type"`
			Gcs       string `json:"gcs"`
			Imgix     struct {
				Args struct {
				} `json:"args"`
				MasterImage struct {
					ImageType string `json:"image_type"`
					Gcs       string `json:"gcs"`
					Static    string `json:"static"`
					Imgix     struct {
					} `json:"imgix"`
				} `json:"master_image"`
			} `json:"imgix"`
		} `json:"config"`
	} `json:"high_res_tile_image_information"`
	SupportsPartners       bool    `json:"supports_partners"`
	DetailedMarketingBlurb string  `json:"detailed_marketing_blurb"`
	TileLogo               string  `json:"tile_logo"`
	TileShortName          string  `json:"tile_short_name"`
	StartDateDatetime      string  `json:"start_date|datetime"`
	EndDateDatetime        string  `json:"end_date|datetime"`
	TileStamp              string  `json:"tile_stamp"`
	BundlesSoldDecimal     float64 `json:"bundles_sold|decimal"`
	TileName               string  `json:"tile_name"`
	ShortMarketingBlurb    string  `json:"short_marketing_blurb"`
	TileImageInformation   struct {
		Config struct {
			ImageType string `json:"image_type"`
			Gcs       string `json:"gcs"`
			Imgix     struct {
				Args struct {
				} `json:"args"`
				MasterImage struct {
					ImageType string `json:"image_type"`
					Gcs       string `json:"gcs"`
					Static    string `json:"static"`
					Imgix     struct {
					} `json:"imgix"`
				} `json:"master_image"`
			} `json:"imgix"`
		} `json:"config"`
	} `json:"tile_image_information"`
	Type       string   `json:"type"`
	Highlights []string `json:"highlights"`
}
