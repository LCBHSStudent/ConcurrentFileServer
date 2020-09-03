package service

import (
	"ConcurrentImageFileServer/src/config"
	"ConcurrentImageFileServer/src/utils/dateUtil"
	"ConcurrentImageFileServer/src/utils/logUtil"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	_ "github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"strings"
)

const utilName = "service"

var fileServerInfo 	fasthttp.FS
var router     		*fasthttprouter.Router
var fileServer		fasthttp.Server

func StartImageFileService() {
	fileServerInfo = fasthttp.FS{
		Root:                 config.ServerConfig.UploadPath,
		GenerateIndexPages:   config.ServerConfig.GenerateIndexPages,
		Compress:             true,

		// default structure values
		AcceptByteRange:      false,
		PathRewrite:          nil,
		PathNotFound:         nil,
		CacheDuration:        0,
		CompressedFileSuffix: "",
	}

	go func() {
		if err := fasthttp.ListenAndServe(
			config.ServerConfig.VisitPort,
			fileServerInfo.NewRequestHandler(),
		); err != nil {
			panic(err)
		}
	}()
}

func StartOperationServer() {
	router = fasthttprouter.New()

	router.PanicHandler = func(ctx *fasthttp.RequestCtx, err interface{}) {
		info := fmt.Sprintf("%v", err)
		printUtilMessage(info)

		sendResponse(ctx, -1, "Undefined error", err)
	}

	router.POST("/upload-image", uploadImageHandler)
	router.POST("/remove-image", removeImageHandler)

	fileServer = fasthttp.Server{
		Handler:                            router.Handler,
		MaxRequestBodySize:                 config.ServerConfig.MaxRequestBodySize,

		ErrorHandler:                       nil,
		HeaderReceived:                     nil,
		ContinueHandler:                    nil,
		Name:                               "",
		Concurrency:                        0,
		DisableKeepalive:                   false,
		ReadBufferSize:                     0,
		WriteBufferSize:                    0,
		ReadTimeout:                        0,
		WriteTimeout:                       0,
		IdleTimeout:                        0,
		MaxConnsPerIP:                      0,
		MaxRequestsPerConn:                 0,
		MaxKeepaliveDuration:               0,
		TCPKeepalive:                       false,
		TCPKeepalivePeriod:                 0,
		ReduceMemoryUsage:                  false,
		GetOnly:                            false,
		DisablePreParseMultipartForm:       false,
		LogAllErrors:                       false,
		DisableHeaderNamesNormalizing:      false,
		SleepWhenConcurrencyLimitsExceeded: 0,
		NoDefaultServerHeader:              false,
		NoDefaultDate:                      false,
		NoDefaultContentType:               false,
		ConnState:                          nil,
		Logger:                             nil,
		KeepHijackedConns:                  false,
	}

	go func() {
		if 	err := fileServer.ListenAndServe(config.ServerConfig.OperationPort);
			err != nil {
			panic(err)
		}
	}()
}

func printUtilMessage(msg string) {
	logMsg := strings.Join(
		[]string {
			logUtil.Red(dateUtil.GetCurFormatTime()),
			" [Module ", logUtil.Blue(utilName),  "] ", msg,
		}, "",
	)
	fmt.Println(logMsg)
}