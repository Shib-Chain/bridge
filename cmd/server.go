package cmd

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/Shib-Chain/bridge/pkg/bridge"
)

var (
	httpPort = flag.String("httpport", "8080", "HTTP Server port")
)

func init() {
	flag.Parse()
}

type Config struct {
}

type Server struct {
	http   *http.Server
	bridge *bridge.Bridge
}

func NewServer() *Server {
	b, err := bridge.New()
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		bridge: b,
	}
}

func (s *Server) Run() {
	e := gin.Default()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, os.Interrupt)

	// eth
	e.POST("/move-asset-eth", func(c *gin.Context) {
		txHash := c.PostForm("tx_hash")
		if txHash == "" {
			errMsg := "Invalid request, please try again."
			http.Error(c.Writer, errMsg, http.StatusBadRequest)
			return
		}

		newTxHash, err := s.bridge.ETHMove(c.Request.Context(), txHash)
		if err != nil {
			handleErrResponse(c, err)
			return
		}

		resp := struct {
			TxHash string `json:"tx_hash"`
		}{TxHash: newTxHash}
		c.JSON(http.StatusOK, resp)
	})

	// shibc
	e.POST("/move-asset-shibc", func(c *gin.Context) {
		txHash := c.PostForm("tx_hash")
		if txHash == "" {
			errMsg := "Invalid request, please try again."
			http.Error(c.Writer, errMsg, http.StatusBadRequest)
			return
		}

		newTxHash, err := s.bridge.ShibcMove(c.Request.Context(), txHash)
		if err != nil {
			handleErrResponse(c, err)
			return
		}

		resp := struct {
			TxHash string `json:"tx_hash"`
		}{TxHash: newTxHash}
		c.JSON(http.StatusOK, resp)
	})

	s.http = &http.Server{
		Addr:    ":" + *httpPort,
		Handler: e,
	}

	go func() {
		log.Infof("HTTP server is starting on %s", s.http.Addr)
		if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Start HTTP server: %v", err)
		}
	}()

	<-interrupt

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	log.Infof("HTTP server shutting down...")
	if err := s.http.Shutdown(ctx); err != nil {
		log.Errorf("Shutdown http server failed: %v", err)
		return
	}
}

func handleErrResponse(c *gin.Context, err error) {
	if err == nil {
		return
	}

	log.WithField("error", err).Error("System error")
	http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
}
