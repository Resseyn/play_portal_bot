package helpingMethods

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/pkg/utils/structures"
)

// TODO: ПАДРАЧИТЬ ХУЙ

/*
НОВАЯ КЛАВИАТУРА:
теперь вместо rows и columns используется слайс с количеством элементов в строке
раньше:
	3 rows 2 columns:
---------------------
|	button, button  |
|	button, button  |
|   button, button  |
---------------------
теперь: positions = []int{3,1,2}
----------------------------
|  button, button, button  |
|          button          |
|      button,button       |
----------------------------
*/

//func CreateInline(data *structures.MessageData, positions []int, commands ...structures.Command) *telebotbot.ReplyMarkup {
//	rows := len(positions)
//	su := 0
//	for _, v := range positions {
//		su += v
//	}
//
//	if su-len(commands) != 0 {
//		panic(fmt.Errorf(" ПОСЧИТАЙ СТРОЧКИ И СТОЛБЦЫ"))
//	}
//	prev := false
//
//	var resrows [][]telebotbot.InlineButton
//	if data.PrevCommand != "" {
//		positions = append(positions, 1)
//		rows = len(positions)
//		resrows = make([][]telebotbot.InlineButton, rows)
//		for k, v := range positions {
//			if k == rows && k > 0 {
//				resrows[k] = make([]telebotbot.InlineButton, 1)
//				break
//			}
//			resrows[k] = make([]telebotbot.InlineButton, v)
//		}
//		rows--
//		prev = true
//	} else {
//		resrows = make([][]telebotbot.InlineButton, rows)
//		for k, v := range positions {
//			resrows[k] = make([]telebotbot.InlineButton, v)
//		}
//	}
//	cmdcount := 0
//	for row := 0; row < rows; row++ {
//		for column := 0; column < positions[row]; column++ {
//			dataFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, commands[cmdcount].Command, data.Command)
//
//			resrows[row][column] = telebotbot.InlineButton{
//				Unique: "unique",
//				Data:   dataFormat,
//				Text:   commands[cmdcount].Text,
//			}
//
//			cmdcount++
//		}
//	}
//	if prev {
//		backFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, data.PrevCommand, "")
//		resrows[rows][0] = telebotbot.InlineButton{
//			Unique: "backButton",
//			Data:   backFormat,
//			Text:   "Назад",
//		}
//	}
//
//	kb := &telebotbot.ReplyMarkup{
//		InlineKeyboard: resrows,
//	}
//	return kb
//}

// CreateInline создает инлайн-клавиатуру с кнопками и кнопкой возврата.
// Каждый вложенный массив команд создает новую строку кнопок.
func CreateInline(data *structures.MessageData, commands ...*[]structures.Command) *telebot.ReplyMarkup {

	var rows [][]telebot.InlineButton

	for _, cmdRow := range commands {

		var row []telebot.InlineButton

		for _, cmd := range *cmdRow {

			dataFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, data.Command, data.PrevCommand)
			fmt.Println(dataFormat)
			unique := fmt.Sprintf("%v_%v_%v", data.ChatID, data.MessageID, cmd.Command)

			button := telebot.InlineButton{
				Unique: unique,
				Data:   dataFormat,
				Text:   cmd.Text,
			}
			row = append(row, button)
		}
		rows = append(rows, row)
	}

	/*
				1. Which two historical events gave birth to British Parliament?
		agreed list of 63 rules,invite for kto-to tam for first parlament in 1265
				2. Which parts does Parliament consist of? What are their purposes?
		house of commons, house of lords,and the sovereign???
				3. How many MPs are there? Must they belong to any political party?
		650, they belong to a political party or can be independent
				4. Who does the House of Lords consist of? How many Lords are there?
		bishops,life peers and hereditary peers, 800 members
				5. How can Parliament influence the work of Government?
		influences the work of the Government by questioning ministers, debating policies, and examining work through committees
				6. How does Parliament make and shape laws?
		parlament makes and shapes laws through a process of several reading and stages in both house of commons and house of lords
				7. How are MPs elected through the General Elections?
		mps are elected through
				8. How can candidates campaign to get people to vote for them?
		candidates can campaign by holding public meetings, door-to-door visits, distributing leaflets
				9. What happens when there is a “hung Parliament”?
		“hung Parliament” happens when no single political party has an absolute majority in the House of Commons after a General Election
				10. How can people participate in the political life of the UK?
		people can participate in the political life of the UK by voting, joining a political party, standing as a candidate, contacting their local MP or Lord, and participating in demonstrations or protests
	*/

	if data.PrevCommand != "" {
		backFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, data.PrevCommand, "")
		backUnique := fmt.Sprintf("%v_%v_back", data.ChatID, data.MessageID)

		backButton := telebot.InlineButton{
			Unique: backUnique,
			Data:   backFormat,
			Text:   "Назад",
		}
		rows = append(rows, []telebot.InlineButton{backButton})
	}

	kb := &telebot.ReplyMarkup{
		InlineKeyboard: rows,
	}
	return kb
}
