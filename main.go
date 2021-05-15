package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"mocha/problem"
	"mocha/script"

	"github.com/bwmarrin/discordgo"
)

const help_mocha_command = "help, mocha"

const run_js_command = "run js"
const estimate_js_command = "estimate js"
const run_go_command = "run go"
const estimate_go_command = "estimate go"

const enroll_problem_command = "enroll problem"
const random_problem_command = "random problem"
const view_problem_command = "view problem"
const submit_solution_command = "submit solution"

//go:embed bot.txt
var botKey string

func main() {
	bot, err := discordgo.New("Bot " + botKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		strs := strings.SplitN(m.Content, "\n", 2)
		resp := ""
		switch strs[0] {
		case help_mocha_command:
			c, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				return
			}
			s.ChannelMessageSend(c.ID, helpPhrase)
			return
		case run_js_command:
			resp, _ = script.RunJS(strs[len(strs)-1])
		case estimate_js_command:
			rs, time := script.RunJS(strs[len(strs)-1])
			resp = rs + "\n" + time
		case run_go_command:
			resp, _ = script.RunGo(strs[len(strs)-1])
		case estimate_go_command:
			rs, time := script.RunGo(strs[len(strs)-1])
			resp = rs + "\n" + time
		case enroll_problem_command:
			if err := problem.Enroll(strs[len(strs)-1]); err != nil {
				resp = "등록에 실패했습니다.\n"
			} else {
				resp = "등록하였습니다."
			}
		case random_problem_command:
			if p := problem.SelectOne(); p == nil {
				resp = "선택할 수 있는 문제가 없습니다."
			} else {
				resp = fmt.Sprintf("```제목: %s\n설명: %s```", p.Name, p.Conetent)
			}
		case view_problem_command:
			if p := problem.Select(strs[len(strs)-1]); p == nil {
				resp = "해당하는 문제가 없습니다."
			} else {
				resp = fmt.Sprintf("```제목: %s\n설명: %s```", p.Name, p.Conetent)
			}
		case submit_solution_command:
			dt, err := problem.Submit(strs[len(strs)-1])
			if err != nil {
				resp = "실패했습니다.\n" + err.Error()
			} else {
				resp = "성공했습니다.\n평균 실행 시간: " + dt
			}
		default:
			return
		}

		_, err := s.ChannelMessageSendReply(m.ChannelID, resp, m.Reference())
		if err != nil {
			log.Println(err)
			return
		}
	})

	if err := bot.Open(); err != nil {
		log.Fatal(err)
	}
	defer bot.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
