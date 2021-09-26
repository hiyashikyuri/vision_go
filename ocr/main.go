package main

import (
	"context"
	"fmt"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func main() {
	// クライアントの作成
	ctx := context.Background()
	client, _ := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("./cred.json"))

	// 画像の読みこみ
	// 画像URL: https://www.tanomail.com/product/3696600/
	file, _ := os.Open("sample.jpeg")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error:%s", err)
		}
	}(file)

	image, _ := vision.NewImageFromReader(file)
	annotations, _ := client.DetectTexts(ctx, image, nil, 10)

	for _, annotation := range annotations {
		fmt.Printf(annotation.Description)
	}
}
