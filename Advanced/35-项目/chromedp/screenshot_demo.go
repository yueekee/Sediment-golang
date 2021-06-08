package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	//var b1, b2 []byte
	var b2 []byte

	if err := chromedp.Run(ctx,
		//// emulate iPhone 7 landscape
		//chromedp.Emulate(device.IPhone8Plus),
		//chromedp.Navigate(`https://www.baidu.com/`),
		//chromedp.CaptureScreenshot(&b1),
		//
		//// reset
		//chromedp.Emulate(device.Reset),

		// set really large viewport
		chromedp.EmulateViewport(1600, 900),
		//emulation.SetDeviceMetricsOverride(0, 0, 16.0/9.0, false),
		chromedp.Navigate(`https://www.baidu.com/`),
		chromedp.CaptureScreenshot(&b2),
	); err != nil {
		log.Fatal(err)
	}

	//if err := ioutil.WriteFile("baidu_IPhone8Plus.png", b1, 0777); err != nil {
	//	log.Fatal(err)
	//}
	if err := ioutil.WriteFile("baidu_PC.png", b2, 0777); err != nil {
		log.Fatal(err)
	}
}