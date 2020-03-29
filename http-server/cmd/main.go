package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/momocomics/backend/http-server/pkg/config"
	"github.com/momocomics/backend/http-server/pkg/pb"
	"github.com/momocomics/backend/http-server/pkg/routes"
)

var (
	host string
	port string
)

const (
	defGrpcServerAddr = "localhost:8090"
	GrpcHostEnv       = "GRPC_SERVER_HOST"
	GrpcPortEnv       = "GRPC_SERVER_PORT"
)

//TODO: this is a test key. Create prod key and add it to k8s secret
// generate pk - openssl genrsa -out app.rsa 512(keysize)
// generate pub key - openssl rsa -in app.rsa -pubout > app.rsa.pub
var signBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAJVKVZSF/8PIzJ1uoVrorfcjaULg4ZCtSHeTB6b9QnLJ/x6FXNBu
DxOR5dFhqiHpPtNhpS9mqQT9V3I6qXfGgzcCAwEAAQJAXsPl2TbKKPyQriqosC1d
KLDIw5Q+evkUNBsX02+WO4h3gCZeGouhGebaMgno1pUzsSal0W9U9XWnumWnkGwr
gQIhAMR9b00nFTjLMvRoi0R87nUNRgdhk431wfVmJopYU0VBAiEAwoFYoVz4Dckp
csRfg7bsb6HfcIlxncvNV2aleN/C0ncCICE/4J+7p1mu+PZm4no6cdeY4WrKVj/F
gIbYPFlYzO6BAiEAtb5Qx65sJc2CijeNnDBvarvRYYE8BZrqSzGhinlivG8CIG9r
CwQbSEUWg/snnrIOJH4O89fY9c/OZW48MLzsdyrh
-----END RSA PRIVATE KEY-----`)
var verifyBytes = []byte(`-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJVKVZSF/8PIzJ1uoVrorfcjaULg4ZCt
SHeTB6b9QnLJ/x6FXNBuDxOR5dFhqiHpPtNhpS9mqQT9V3I6qXfGgzcCAwEAAQ==
-----END PUBLIC KEY-----`)

func main() {
	flag.StringVar(&host, "host", "", "host of grpc server")
	flag.StringVar(&port, "port", "", "port of grpc server")
	flag.Parse()

	signBytes = regexp.MustCompile(`\s*$`).ReplaceAll(signBytes, []byte{})
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal("private key:", err)
	verifyBytes = regexp.MustCompile(`\s*$`).ReplaceAll(verifyBytes, []byte{})
	pubkey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal("public key:", err)

	if host == "" {
		if host, port = os.Getenv(GrpcHostEnv), os.Getenv(GrpcPortEnv); host == "" || port == "" {
			log.Printf("host %v port %v \n", host, port)
			host = defGrpcServerAddr
			log.Printf("addr %v \n", host)

		} else {
			host = strings.Join([]string{host, port}, ":")
		}
	}
	log.Printf("GRPC server address is %v \n", host)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewTodoClient(conn)
	//todo: get host from env
	cfg := config.New(client, privKey, pubkey, "momocomic.com")
	cfg.SetDebug(true)
	s := gin.Default()
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Rest server running.")
	})
	routes.Routes(s, cfg)

	//http.ListenAndServe("8080", gin)
	s.Run(":8081")
}

func fatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%v %v", msg, err)
	}
}
