package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"gitlab.lila.network/adora-kalb/humble-bot/models"
	_ "modernc.org/sqlite"
)

func TestE2E(t *testing.T) {
	// 1. Mock Humble Bundle Server
	humbleServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		category := ""
		if r.URL.Path == "/games" {
			category = "games"
		} else if r.URL.Path == "/books" {
			category = "books"
		} else if r.URL.Path == "/software" {
			category = "software"
		}

		if category == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		jsonData := fmt.Sprintf(`{
			"data": {
				"%s": {
					"mosaic": [
						{
							"products": [
								{
									"machine_name": "test_bundle_%s",
									"tile_name": "Test %s Bundle",
									"marketing_blurb": "This is a test %s bundle",
									"product_url": "/test-%s-bundle"
								}
							]
						}
					]
				}
			}
		}`, category, category, category, category, category)

		_, _ = fmt.Fprintf(w, `<html><body><script id="landingPage-json-data" type="application/json">%s</script></body></html>`, jsonData)
	}))
	defer humbleServer.Close()

	// 2. Mock Mastodon Server
	var mastodonReceived bool
	mastodonServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/api/v1/statuses" {
			mastodonReceived = true
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{
				"url": "https://mock-mastodon.social/@bot/123",
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer mastodonServer.Close()

	// 3. Initialize SQLite Database
	sqldb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	models.DB = bun.NewDB(sqldb, sqlitedialect.New())

	ctx := context.Background()
	_, err = models.DB.NewCreateTable().Model((*models.QueueItem)(nil)).Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = models.DB.NewCreateTable().Model((*models.SeenBundle)(nil)).Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// 4. Configure Viper
	viper.Set("humblebundle.url", humbleServer.URL)
	viper.Set("mastodon.url", mastodonServer.URL)
	viper.Set("mastodon.token", "test-token")
	viper.Set("mastodon.visibility", "private")

	// 5. Run UpdateBundles
	// This should scrape the mock Humble Bundle server and add items to the queue
	UpdateBundles()

	// Check if items were added to DB
	var items []models.QueueItem
	err = models.DB.NewSelect().Model(&items).Scan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) == 0 {
		t.Fatal("Expected items in queue, got 0")
	}

	// 6. Run RunSingleQueueItem
	// This should take one item from the queue and post it to the mock Mastodon server
	RunSingleQueueItem()

	// 7. Verify Mastodon received the post
	if !mastodonReceived {
		t.Error("Mastodon mock server did not receive a post request")
	}

	// Check if item was dequeued
	var remainingItems []models.QueueItem
	err = models.DB.NewSelect().Model(&remainingItems).Scan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(remainingItems) != len(items)-1 {
		t.Errorf("Expected %d items remaining in queue, got %d", len(items)-1, len(remainingItems))
	}

	// 8. Run UpdateBundles again
	// It should NOT add new items because they are already in "seen" table
	UpdateBundles()

	var itemsAfterSecondRun []models.QueueItem
	err = models.DB.NewSelect().Model(&itemsAfterSecondRun).Scan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(itemsAfterSecondRun) != len(remainingItems) {
		t.Errorf("Expected no new items in queue, but got %d (before) -> %d (after)", len(remainingItems), len(itemsAfterSecondRun))
	}

	// 9. Run RunSingleQueueItem until queue is empty
	for range itemsAfterSecondRun {
		RunSingleQueueItem()
	}

	var finalItems []models.QueueItem
	err = models.DB.NewSelect().Model(&finalItems).Scan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(finalItems) != 0 {
		t.Errorf("Expected empty queue, got %d items", len(finalItems))
	}

	// 10. Run RunSingleQueueItem when queue is empty
	// It should not crash or call Mastodon mock again (in a way that causes error)
	mastodonReceived = false
	RunSingleQueueItem()
	if mastodonReceived {
		t.Error("Mastodon mock received a request but queue was supposed to be empty")
	}
}
