package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Server struct {
	db        *gorm.DB
	http      *http.Server
	tokenAuth *jwtauth.JWTAuth
	cacheRds  *redis.Client
}

func NewServer(db *gorm.DB, tokenAuth *jwtauth.JWTAuth, cacheRds *redis.Client) *Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Route("/api", func(r chi.Router) {
		r.Post("/register", deliveryAuth(db, cacheRds, tokenAuth).Register)
		r.Post("/login", deliveryAuth(db, cacheRds, tokenAuth).Login)

		// route protect area
		r.Group(func(rauth chi.Router) {
			rauth.Use(jwtauth.Verifier(tokenAuth))
			rauth.Use(jwtauth.Authenticator)
			rauth.Use(CheckTokenInRedis(cacheRds))

			rauth.Route("/auth", func(auth chi.Router) {
				auth.Post("/logout", deliveryAuth(db, cacheRds, tokenAuth).Logout)
				// bank account
				auth.Route("/bank-account", func(ba chi.Router) {
					ba.Post("/create", deliveryBankAccount(db).Create)
					ba.Get("/{id}", deliveryBankAccount(db).FindbyId)
					ba.Put("/{id}", deliveryBankAccount(db).Update)
				})
				// transaction
				auth.Route("/transaction", func(trf chi.Router) {
					trf.Post("/top-up", deliveryTransaction(db).TopUp)
				})

				// user
				auth.Route("/user", func(usr chi.Router) {
					usr.Get("/", deliveryUser(db).UserList)
					usr.Get("/{userId}", deliveryUser(db).UserDetail)

				})
			})

		})
	})

	server := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      20 * time.Second,
		Addr:              viper.GetString("listen_address"),
		Handler:           router,
	}

	server.SetKeepAlivesEnabled(true)

	srv := &Server{
		db:        db,
		http:      server,
		tokenAuth: tokenAuth,
		cacheRds:  cacheRds,
	}

	return srv
}

func (s *Server) Start(ctx context.Context) (err error) {
	fmt.Println("App: Starting")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err = s.http.ListenAndServe(); err != nil {
			return
		}
	}()

	<-ctx.Done()

	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	fmt.Println("App: Stopping...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err = s.http.Shutdown(ctx); err != nil {
			return
		}
	}()
	wg.Wait()

	return
}
