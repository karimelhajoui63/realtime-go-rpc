package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	paintv1 "rpc-server/gen/proto/paint/v1"
	"rpc-server/gen/proto/paint/v1/paintv1connect"
	"strings"

	"connectrpc.com/connect"
)

var colorWantedStr = flag.String("color", paintv1.Color_COLOR_UNSPECIFIED.String(), "the color wanted")

func convertStringToColorEnum(colorStr string) (paintv1.Color, error) {
	switch "COLOR_" + strings.ToUpper(colorStr) {
	case paintv1.Color_COLOR_BLUE.String():
		return paintv1.Color_COLOR_BLUE, nil
	case paintv1.Color_COLOR_RED.String():
		return paintv1.Color_COLOR_RED, nil
	case paintv1.Color_COLOR_GREEN.String():
		return paintv1.Color_COLOR_GREEN, nil
	default:
		return paintv1.Color_COLOR_UNSPECIFIED, fmt.Errorf("unknown color: %s", colorStr)
	}
}

func ChangeColor() {
	flag.Parse()

	colorWanted, err := convertStringToColorEnum(*colorWantedStr)
	if err != nil {
		log.Fatalf("Error converting color: %v", err)
	}

	client := paintv1connect.NewPaintServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.ChangeColor(
		context.Background(),
		connect.NewRequest(&paintv1.ChangeColorRequest{Color: colorWanted}),
	)
	if err != nil {
		log.Println(err)
		return
	}

	if res.Msg.Succeed {
		log.Println("Call succeed! ✅")
	} else {
		log.Println("Call failed! ❌")
	}
}

func main() {
	ChangeColor()
}
