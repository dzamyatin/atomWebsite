package messengertelegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

//https://gitlab.com/Athamaxy/telegram-bot-tutorial/-/blob/main/TutorialBot.go

type Handler func(
	update tgbotapi.Update,
	bot *Bot,
) error

type Bot struct {
	token  string
	logger *zap.Logger
	botAPI *tgbotapi.BotAPI
	lockMe *sync.RWMutex
}

func NewBot(
	token string,
	logger *zap.Logger,
) *Bot {
	b := &tgbotapi.BotAPI{
		Token:  token,
		Client: &http.Client{},
		Buffer: 100,
	}

	b.SetAPIEndpoint(tgbotapi.APIEndpoint)

	return &Bot{
		token:  token,
		logger: logger,
		botAPI: b,
		lockMe: &sync.RWMutex{},
	}
}

func (r *Bot) setMe() error {
	r.lockMe.Lock()
	defer r.lockMe.Unlock()

	if r.botAPI.Self.ID != 0 {
		return nil
	}

	u, err := r.botAPI.GetMe()
	if err != nil {
		return errors.Wrap(err, "botAPI.GetMe")
	}

	r.botAPI.Self = u

	return nil
}

func (r *Bot) ReceiveUpdates(
	ctx context.Context,
	offset int,
	handler Handler,
) error {
	err := r.setMe()
	if err != nil {
		return errors.Wrap(err, "botAPI.setMe")
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.Offset = offset

	updates := r.botAPI.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update := <-updates:
			err = handler(update, r)
			if err != nil {
				return errors.Wrap(err, "handler")
			}

			// >>
			//msg := tgbotapi.NewMessage(
			//	update.Message.Chat.ID,
			//	"hello",
			//)
			//
			//msg := tgbotapi.NewEditMessageTextAndMarkup(
			//	update.Message.Chat.ID,
			//	update.Message.MessageID,
			//	"test",
			//	tgbotapi.InlineKeyboardMarkup{
			//		InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0),
			//	},
			//)
			//

			//msg := tgbotapi.EditMessageReplyMarkupConfig{
			//
			//}
			//msg.ReplyMarkup = tgbotapi.NewReplyKeyboard().Keyboard

			//r.botAPI.Send(msg)

			//kb := tgbotapi.NewReplyKeyboard().Keyboard
			////
			//msg := tgbotapi.NewEditMessageTextAndMarkup(
			//	update.Message.Chat.ID,
			//	update.Message.MessageID,
			//	"test",
			//	tgbotapi.InlineKeyboardMarkup{
			//		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			//			{
			//				tgbotapi.InlineKeyboardButton{
			//					Text:                         "<UNK>",
			//					URL:                          nil,
			//					LoginURL:                     nil,
			//					CallbackData:                 nil,
			//					SwitchInlineQuery:            nil,
			//					SwitchInlineQueryCurrentChat: nil,
			//					CallbackGame:                 nil,
			//					Pay:                          false,
			//				},
			//			},
			//		},
			//		//InlineKeyboard: kb.Keyboard,
			//		//
			//
			//	},
			//)

			//tgbotapi.NewReplyKeyboard()
			r.logger.Info(
				"telegram update received",
				zap.String("message", update.Message.Text),
				zap.String("user", update.Message.From.UserName),
				zap.Int64("user_id", update.Message.From.ID),
				zap.String("contact", update.Message.Contact.PhoneNumber),
				zap.Int64("contact_id", update.Message.Contact.UserID),
			)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "tst1")
			kb := tgbotapi.NewReplyKeyboard()
			//kb.Keyboard = [][]tgbotapi.KeyboardButton{}
			kb.Keyboard = [][]tgbotapi.KeyboardButton{
				{
					tgbotapi.NewKeyboardButtonContact("tst2"),
				},
			}

			msg.ReplyMarkup = kb
			msg.ReplyToMessageID = update.Message.MessageID
			msg.AllowSendingWithoutReply = true
			//msg.DisableNotification = true

			r.botAPI.Send(msg)

			// <<
		}
	}

	//return nil
}

func (r *Bot) SendMessage(chatId int64, message string) error {
	err := r.setMe()
	if err != nil {
		return errors.Wrap(err, "botAPI.setMe")
	}

	msg := tgbotapi.NewMessage(chatId, message)

	//msg.Entities = append(
	//	msg.Entities,
	//	tgbotapi.MessageEntity{
	//		Type: "mention", //(@username),
	//		Offset: 0,
	//		Length: len(message),
	//	},
	//)

	_, err = r.botAPI.Send(msg)

	if err != nil {
		return errors.Wrap(err, "bot.SendMessage")
	}

	return nil
}
