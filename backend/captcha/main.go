package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/McaxDev/CaptchaGo/rpc"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	rpc.UnimplementedCaptchaServer
}

func (srv *server) Verify(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	valid := captcha.VerifyString(in.CaptchaId, in.CaptchaValue)
	return &rpc.Response{Valid: valid}, nil
}

func main() {

	grpcPort, exist := os.LookupEnv("GRPC_PORT")
	if !exist {
		grpcPort = "50051"
	}
	httpPort, exist := os.LookupEnv("HTTP_PORT")
	if !exist {
		httpPort = "8080"
	}

	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalln("无法监听RPC端口："+ err.Error())          
		}
		srv := grpc.NewServer()
		rpc.RegisterCaptchaServer(srv, &server{})
		reflection.Register(srv)
		if err := srv.Serve(lis); err != nil {
			log.Fatalln("开启RPC服务器失败："+err.Error())    
		}
	}()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		id := captcha.New()
		ctx.Header("Content-Type", "image/png")
		ctx.Header("X-Captcha-Id", id)
		if err := captcha.WriteImage(
			ctx.Writer, id, captcha.StdWidth, captcha.StdHeight,
		); err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"msg":  "验证码绘制失败",
				"data": nil,
			})
			return
		}
	})
	r.Run(":" + httpPort)
}
