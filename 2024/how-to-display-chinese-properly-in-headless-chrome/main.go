package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

const (
	DefaultURL = "ws://127.0.0.1:9222"
)

type Client struct{}

// ScreenShot 接受一个上下文、URL和CSS选择器作为输入参数，并返回一个image.Image和一个错误。
// 并且，只有当 css 选择器的元素可见时，才会截屏。
func (c *Client) ScreenShot(ctx context.Context, url string, sel string, mobile bool, verbose bool, remote bool, remoteURL string) ([]byte, error) {
	var opts []chromedp.ContextOption
	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}

	var cCtx context.Context
	if remote {
		aCtx, _ := chromedp.NewRemoteAllocator(ctx, remoteURL)
		xCtx, cancel := chromedp.NewContext(aCtx, opts...)
		defer cancel()
		cCtx = xCtx
	} else {
		xCtx, cancel := chromedp.NewContext(ctx, opts...)
		defer cancel()
		cCtx = xCtx
	}

	var buf []byte
	action := []chromedp.Action{}
	if mobile {
		d := device.Info{
			Name:      "iPad Pro",
			UserAgent: "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
			Width:     512,
			Height:    1366,
			Scale:     2.500000,
			Landscape: false,
			Mobile:    true,
			Touch:     true,
		}
		action = append(action, chromedp.Emulate(d))
	}
	action = append(
		action,
		chromedp.Navigate(url),
		// chromedp.CaptureScreenshot(&buf),
		chromedp.Screenshot(sel, &buf, chromedp.NodeVisible),
	)
	if err := chromedp.Run(cCtx, action...); err != nil {
		return nil, fmt.Errorf("failed getting body of %s: %v", url, err)
	}
	return buf, nil
}

func main() {
	url := flag.String("url", "", "url")
	sel := flag.String("sel", "body", "css selector")
	output := flag.String("output", "screen.png", "output file")
	flag.Parse()

	ctx := context.Background()
	cli := Client{}
	bs, err := cli.ScreenShot(ctx, *url, *sel, true, false, true, DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(*output, bs, 0644); err != nil {
		log.Fatal(err)
	}
}
