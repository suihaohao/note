/*
Linux系统需要Install:
  yum install -y epel-release
  yum install -y chromium
  yum groupinstall "fonts"
*/

package screenshot

import (
	"context"
	"flag"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func screenshot(url string, toPath string) error {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	// create context
	// timeout 90 秒
	timeContext, cancelFunc := context.WithTimeout(context.Background(), time.Second*90)
	defer cancelFunc()
	ctx, cancel := chromedp.NewContext(timeContext)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(url, 90, &buf)); err != nil {
		return err
	}

	os.MkdirAll(filepath.Dir(toPath), os.ModeAppend)
	file, err := os.OpenFile(toPath, os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(buf)

	return nil
}

// 截取全部的viewport
// 注意会重新 viewport emulation 配置
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		// 设置浏览器窗口大小 高度设置为0 表示自适应
		chromedp.EmulateViewport(1400, 0, chromedp.EmulateScale(1)),
		// 设置等待页面加载的条件
		// chromedp.WaitReady("xxx", chromedp.BySearch),
		// chromedp.WaitVisible("show", chromedp.ByID),
		// chromedp.Sleep(time.Second * 3),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func Example() {
	url := flag.String("url", "http://www.chaipip.com", "like 'http://baidu.com'")
	toPath := flag.String("to", "./tmp/tmp.png", "like './tmp/tmp.png'")
	flag.Parse()
	err := screenshot(*url, *toPath)
	if err != nil {
		log.Fatal(err)
	}
}
