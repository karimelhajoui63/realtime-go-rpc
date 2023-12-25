package main

import (
	"context"
	"log"
	"net/http"

	paintv1 "rpc-server/gen/proto/paint/v1"
	"rpc-server/gen/proto/paint/v1/paintv1connect"

	"connectrpc.com/connect"

	"github.com/rs/cors"
)

type PaintServer struct{}

func (s *PaintServer) ChangeColor(
	ctx context.Context,
	req *connect.Request[paintv1.ChangeColorRequest],
) (*connect.Response[paintv1.ChangeColorResponse], error) {
	defer log.Println("Color changed to: ", req.Msg.Color)
	res := connect.NewResponse(&paintv1.ChangeColorResponse{
		Succeed: true,
	})
	res.Header().Set("Paint-Version", "v1")
	return res, nil
}

func rpc() {
	painter := &PaintServer{}
	mux := http.NewServeMux()
	mux.Handle(paintv1connect.NewPaintServiceHandler(painter))

	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{
			"http://localhost:5173", // To allow incoming request from front-end
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
	err := http.ListenAndServe("localhost:8080", handler)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
}

func main() {
	rpc()
}
