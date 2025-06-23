package main

import _ "github.com/lib/pq"
import (
	"context"
	"database/sql"
	"github.com/atulantonyz/gator/internal/config"
	"github.com/atulantonyz/gator/internal/database"
	"log"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	dbUrl := cfg.Db_url
	db, err := sql.Open("postgres", dbUrl)
	dbQueries := database.New(db)
	var app_state state
	app_state.cfg = &cfg
	app_state.db = dbQueries
	app_commands := commands{
		map[string]func(*state, command) error{},
	}
	app_commands.register("login", handlerLogin)
	app_commands.register("register", handlerRegister)
	app_commands.register("reset", handlerReset)
	app_commands.register("users", handlerUsers)
	app_commands.register("agg", handlerAgg)
	app_commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	app_commands.register("feeds", handlerFeeds)
	app_commands.register("follow", middlewareLoggedIn(handlerFollow))
	app_commands.register("following", middlewareLoggedIn(handlerFollowing))
	app_commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	app_commands.register("browse", middlewareLoggedIn(handlerBrowse))
	cmd_line_args := os.Args
	if len(cmd_line_args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmdName := cmd_line_args[1]
	cmdArgs := cmd_line_args[2:]

	err = app_commands.run(&app_state, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
