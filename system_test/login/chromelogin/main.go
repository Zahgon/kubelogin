package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s URL", os.Args[0])
		return
	}
	url := os.Args[1]
	if err := runBrowser(context.Background(), url); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func runBrowser(ctx context.Context, url string) error { _ = "STUB: not implemented"; return nil }

func openBrowser(ctx context.Context, url string) error { _ = "STUB: not implemented"; return nil }

func logInToDex(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// https://dex-server:10443/dex/auth/local

// https://dex-server:10443/dex/approval

// http://localhost:8000

func logPageMetadata() chromedp.Action { _ = "STUB: not implemented"; return *new(chromedp.Action) }
