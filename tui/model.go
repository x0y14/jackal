package tui

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	chatv1 "github.com/x0y14/jackal/gen/chat/v1"
	"github.com/x0y14/jackal/gen/chat/v1/chatv1connect"
	notifyv1 "github.com/x0y14/jackal/gen/notify/v1"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"log"
)

type ReceiveRespMsg struct {
	Resp *notifyv1.FetchMessageResponse
}
type errMsg error

type Model struct {
	viewport viewport.Model
	messages []*typesv1.Message
	//responses   []*typesv1.Message
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
	chatClient  chatv1connect.ChatServiceClient
	sender      string
	receiver    string
}

func InitialModel(client chatv1connect.ChatServiceClient, sender, receiver string) Model {
	ta := textarea.New()
	ta.Placeholder = "write msg here"
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 10)
	vp.SetContent(`enter to send`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return Model{
		textarea:    ta,
		messages:    []*typesv1.Message{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		chatClient:  client,
		sender:      sender,
		receiver:    receiver,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			if m.textarea.Value() == "" {
				return m, nil
			}
			req := connect.NewRequest(&chatv1.SendMessageRequest{Message: &typesv1.Message{
				From:     m.sender,
				To:       m.receiver,
				Text:     m.textarea.Value(),
				Metadata: "{}",
				Kind:     0,
			}})
			req.Header().Set("X-User-ID", m.sender)
			_, err := m.chatClient.SendMessage(context.Background(), req)
			if err != nil {
				log.Fatal(err)
			}
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}
	case ReceiveRespMsg:
		m.messages = append(m.messages, msg.Resp.Message)
		var screen string
		if len(m.messages) < 11 {
			for _, message := range m.messages {
				if message.From == m.sender {
					screen += fmt.Sprintf("%v: %s\n", m.senderStyle.Render(message.From), message.Text)
				} else {
					screen += fmt.Sprintf("%v: %s\n", message.From, message.Text)
				}
			}
		} else {
			for _, message := range m.messages[len(m.messages)-10 : len(m.messages)-1] {
				if message.From == m.sender {
					screen += fmt.Sprintf("%v: %s\n", m.senderStyle.Render(message.From), message.Text)
				} else {
					screen += fmt.Sprintf("%v: %s\n", message.From, message.Text)
				}
			}
		}
		m.viewport.SetContent(screen)
		m.textarea.Reset()
		m.viewport.GotoBottom()
		return m, nil

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
