package servicemessengertelegram

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"sync"
)

//https://gitlab.com/Athamaxy/telegram-bot-tutorial/-/blob/main/TutorialBot.go

type Handler func(
	update tgbotapi.Update,
	bot *TelegramDriver,
) error

type TelegramDriver struct {
	token  string
	logger *zap.Logger
	botAPI *tgbotapi.BotAPI
	lockMe *sync.RWMutex
}

func (r *TelegramDriver) GetChatID(message servicemessengermessage.Message) (string, error) {
	if message.ChatLink.Telegram.ChatID == 0 {
		return "", errors.New("chat ID undefined")
	}

	return strconv.FormatInt(message.ChatLink.Telegram.ChatID, 10), nil
}

func (r *TelegramDriver) SendMessage(message servicemessengermessage.Message) error {
	err := r.setMe()
	if err != nil {
		return errors.Wrap(err, "botAPI.setMe")
	}

	msg := tgbotapi.NewMessage(message.ChatLink.Telegram.ChatID, message.Text)

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
		r.logger.Error("servicemessengertelegram.SendMessage", zap.Error(err))
		return errors.Wrap(err, "servicemessengertelegram.SendMessage")
	}

	return nil
}

func (r *TelegramDriver) AskPhone(link servicemessengermessage.ChatLink) error {
	msg := tgbotapi.NewMessage(link.Telegram.ChatID, "tst1")
	kb := tgbotapi.NewReplyKeyboard()
	//kb.Keyboard = [][]tgbotapi.KeyboardButton{}
	kb.Keyboard = [][]tgbotapi.KeyboardButton{
		{
			tgbotapi.NewKeyboardButtonContact("tst2"),
		},
	}

	msg.ReplyMarkup = kb
	//msg.ReplyToMessageID = update.Message.MessageID
	msg.AllowSendingWithoutReply = true
	//msg.DisableNotification = true

	_, err := r.botAPI.Send(msg)

	return errors.Wrap(err, "botAPI.Send")
}

func NewTelegramDriver(
	token string,
	logger *zap.Logger,
) *TelegramDriver {
	b := &tgbotapi.BotAPI{
		Token:  token,
		Client: &http.Client{},
		Buffer: 100,
	}

	b.SetAPIEndpoint(tgbotapi.APIEndpoint)

	return &TelegramDriver{
		token:  token,
		logger: logger,
		botAPI: b,
		lockMe: &sync.RWMutex{},
	}
}

func (r *TelegramDriver) setMe() error {
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

func (r *TelegramDriver) ReceiveUpdates(
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
				r.logger.Error("servicemessengertelegram.ReceiveUpdates", zap.Error(err))
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
			//r.logger.Info(
			//	"telegram update received",
			//	zap.String("message", update.Message.Text),
			//	zap.String("user", update.Message.From.UserName),
			//	zap.Int64("user_id", update.Message.From.ID),
			//	zap.String("contact", update.Message.Contact.PhoneNumber),
			//	zap.Int64("contact_id", update.Message.Contact.UserID),
			//)
			//
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "tst1")
			//kb := tgbotapi.NewReplyKeyboard()
			////kb.Keyboard = [][]tgbotapi.KeyboardButton{}
			//kb.Keyboard = [][]tgbotapi.KeyboardButton{
			//	{
			//		tgbotapi.NewKeyboardButtonContact("tst2"),
			//	},
			//}
			//
			//msg.ReplyMarkup = kb
			//msg.ReplyToMessageID = update.Message.MessageID
			//msg.AllowSendingWithoutReply = true
			////msg.DisableNotification = true
			//
			//r.botAPI.Send(msg)

			// <<
		}
	}

	//return nil
}

//func (r *TelegramDriver) SendMessage(chatId int64, message string) error {
//	err := r.setMe()
//	if err != nil {
//		return errors.Wrap(err, "botAPI.setMe")
//	}
//
//	msg := tgbotapi.NewMessage(chatId, message)
//
//	//msg.Entities = append(
//	//	msg.Entities,
//	//	tgbotapi.MessageEntity{
//	//		Type: "mention", //(@username),
//	//		Offset: 0,
//	//		Length: len(message),
//	//	},
//	//)
//
//	_, err = r.botAPI.Send(msg)
//
//	if err != nil {
//		return errors.Wrap(err, "bot.SendMessage")
//	}
//
//	return nil
//}
