package apiserver


import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API Server ...
type APIServer struct {
	 config *Config
	 logger *logrus.Logger
	 router *mux.Router
}



// Create New APIServer ...
func New(config *Config) *APIServer {

	// in php return new APIServer(config, logger.New(), mux.NewRouter())
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}


// Start server
func (s *APIServer) Start() error {

     if err := s.configureLogger(); err != nil {
         return err
	 }

	  
	 s.logger.Info("starting api server")


     return nil
}


// Function for configuration logger
func (s *APIServer)  configureLogger() error {
	
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	 
	// Выходим из функции
	return nil
}


// Config Router
func (s *APIServer) configureRouter() {
	
	// Routes
	s.router.HandleFunc("/hello", s.handleHello('')) 
}


// Routes Hanldes
func (s *APIServer) handleHello() http.HandlerFunc {
	
	type request struct {
		name string
	}
	
    return func (w http.ResponseWriter, r *http.Request) {

	}
}