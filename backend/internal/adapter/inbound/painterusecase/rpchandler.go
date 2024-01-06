package rpchandler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"rpc-server/internal/core/domain/enum"
	"rpc-server/internal/port/inbound"

	paintv1 "rpc-server/api/proto/gen/paint/v1"
	"rpc-server/api/proto/gen/paint/v1/paintv1connect"

	"connectrpc.com/connect"

	"github.com/rs/cors"
)

func convertColorDomainToRpcEnum(color enum.Color) (paintv1.Color, error) {
	switch "COLOR_" + strings.ToUpper(color.String()) {
	case paintv1.Color_COLOR_BLUE.String():
		return paintv1.Color_COLOR_BLUE, nil
	case paintv1.Color_COLOR_RED.String():
		return paintv1.Color_COLOR_RED, nil
	case paintv1.Color_COLOR_GREEN.String():
		return paintv1.Color_COLOR_GREEN, nil
	default:
		return paintv1.Color_COLOR_UNSPECIFIED, fmt.Errorf("unknown color: %s", color.String())
	}
}

func convertColorRpcToDomainEnum(color paintv1.Color) (enum.Color, error) {
	colorStr, found := strings.CutPrefix(strings.ToUpper(color.String()), "COLOR_")
	if !found {
		return enum.Unspecified, fmt.Errorf("paintv1.Color should contain the prefix 'COLOR_': %s", color.String())
	}

	switch colorStr {
	case strings.ToUpper(enum.Blue.String()):
		return enum.Blue, nil
	case strings.ToUpper(enum.Red.String()):
		return enum.Red, nil
	case strings.ToUpper(enum.Green.String()):
		return enum.Green, nil
	default:
		return enum.Unspecified, fmt.Errorf("unknown color: %s", color.String())
	}
}

type PainterHandler struct {
	painterUseCase inbound.PainterUseCase
}

func NewRpcPainterHandler(mux *http.ServeMux, painterUseCase inbound.PainterUseCase) http.Handler {
	painter := &PainterHandler{
		painterUseCase: painterUseCase,
	}
	mux.Handle(paintv1connect.NewPaintServiceHandler(painter))

	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{
			"http://localhost:45871", // To allow incoming request from front-end
		},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Timeout",             // Used for gRPC-web
			"X-Grpc-Web",               // Used for gRPC-web
			"X-User-Agent",             // Used for gRPC-web
		},
		ExposedHeaders: []string{
			"Content-Encoding",         // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Status",              // Required for gRPC-web
			"Grpc-Message",             // Required for gRPC-web
		},
	})
	handler := corsHandler.Handler(mux)
	return handler
}

func (p *PainterHandler) GetColor(
	ctx context.Context,
	req *connect.Request[paintv1.GetColorRequest],
) (*connect.Response[paintv1.GetColorResponse], error) {
	colorStored, err := p.painterUseCase.GetColor()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	color, err := convertColorDomainToRpcEnum(colorStored)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res := connect.NewResponse(&paintv1.GetColorResponse{Color: color})
	return res, nil
}

func (p *PainterHandler) ChangeColor(
	ctx context.Context,
	req *connect.Request[paintv1.ChangeColorRequest],
) (*connect.Response[paintv1.ChangeColorResponse], error) {
	defer log.Println("Color changed to: ", req.Msg.Color)
	newColor, err := convertColorRpcToDomainEnum(req.Msg.Color)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	p.painterUseCase.ChangeColor(newColor)
	res := connect.NewResponse(&paintv1.ChangeColorResponse{
		Succeed: true,
	})
	return res, nil
}

func (p *PainterHandler) GetColorStream(
	ctx context.Context,
	req *connect.Request[paintv1.GetColorStreamRequest],
	stream *connect.ServerStream[paintv1.GetColorStreamResponse],
) error {
	messages, err := p.painterUseCase.GetColorStream()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for message := range messages {
		color, err := convertColorDomainToRpcEnum(message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = stream.Send(&paintv1.GetColorStreamResponse{Color: color})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
