package botLogic

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

type StoredMessage struct {
	MessageID string `sql:"message_id" json:"message_id"`
	ChatID    int64  `sql:"chat_id" json:"chat_id"`
	telebot.Editable
}

func (x StoredMessage) MessageSig() (string, int64) {
	return x.MessageID, x.ChatID
}

func Menu(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/Hydra.webp"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	commands := []*[]structures.Command{
		{
			{Text: "Магазин", Command: "showShop"},
			{Text: "Кабинет", Command: "showPersonalArea"}},
		{
			{Text: "Поддержка", Command: "showSupport"},
			{Text: "FAQ", Command: "showFAQ"},
		}}
	data := helpingMethods.ParseData(c.Callback().Data)
	// =========PARAMS=========

	keyboard := helpingMethods.CreateInline(data, commands...)
	//stringMsgID := strconv.Itoa(data.MessageID)
	//message := StoredMessage{MessageID: stringMsgID, ChatID: data.ChatID}
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	//editConf := helpingMethods.EditMessageWithPhotoAndReplyMarkup(data, commands, messageContent, picPath, positions)
	//_, err := bot.Send(editConf)
	//if err != nil {
	//	loggers.ErrorLogger.Println(err.Error(), "mediaError")
	//}
	return nil
}

func ShowShop(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	data := helpingMethods.ParseData(c.Callback().Data)
	commands := []*[]structures.Command{
		{
			{Text: "Игровые сервисы", Command: ""},
			{Text: "Сервисы", Command: ""}},
		{
			{Text: "Pepega", Command: ""},
		}}
	// =========PARAMS=========

	keyboard := helpingMethods.CreateInline(data, commands...)
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}
