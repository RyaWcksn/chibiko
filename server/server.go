package server

import (
	"fmt"
	"net/http"

	"github.com/RyaWcksn/chibiko/api/v1/handlers"
	"github.com/RyaWcksn/chibiko/api/v1/usecases"
	"github.com/RyaWcksn/chibiko/configs"
	"github.com/RyaWcksn/chibiko/pkgs/databases"
	"github.com/RyaWcksn/chibiko/ports/databases/mysql"
	"github.com/RyaWcksn/chibiko/server/middleware"
)

type Server struct {
	cfg          *configs.Config
	UsecaseLayer usecases.IUsecase
	HandlerLayer handlers.IHandler
}

var SVR *Server

func (s *Server) Register() {
	fmt.Println(s.cfg)
	dbConn := databases.NewMysqlConnection(s.cfg.Database)
	if dbConn == nil {
		panic("cannot connect DB")
	}
	db := dbConn.DBConnect()
	if db == nil {
		panic("cannot connect DB")
	}

	sqlLayerImpl := mysql.NewSql(db)
	s.UsecaseLayer = usecases.NewUsecase(sqlLayerImpl)
	s.HandlerLayer = handlers.NewHandler(*s.cfg, s.UsecaseLayer)

}

// New initiate new server http.
func New(cfg *configs.Config) *Server {
	SVR = &Server{
		cfg: cfg,
	}

	SVR.Register()
	return SVR
}

func (s Server) Start() {
	http.Handle("/encode", middleware.ErrHandler(s.HandlerLayer.Encode))
	http.Handle("/", middleware.ErrHandler(s.HandlerLayer.Decode))
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
