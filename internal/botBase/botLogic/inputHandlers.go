package botLogic

//func SteamTopUpBalanceHandle(c telebot.Context) error {
//	if structures.UserStates[c.Chat().ID].Step == 0 {
//		if ok, err := checkSteamLogin(keys.SteamApiKey, "mrressrus0"); ok {
//			structures.UserStates[c.Chat().ID].DataCase[0] = c.Message().Text
//			structures.UserStates[c.Chat().ID].Step = 1
//		} else if err != nil {
//			loggers.ErrorLogger.Println(err)
//			return err
//		} else {
//			fmt.Println("net")
//		}
//
//	} else if structures.UserStates[c.Chat().ID].Step == 1 {
//		structures.UserStates[c.Chat().ID].DataCase[1] = c.Message().Text
//	}
//	fmt.Println(structures.UserStates[c.Chat().ID])
//	return nil
//}

//const steamAPIURL = "http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/"
//
//func checkSteamLogin(apiKey, loginToCheck string) (bool, error) {
//	// Формируем URL для запроса к Steam API
//	url := fmt.Sprintf("%s?vanityurl=%s&key=%s", steamAPIURL, loginToCheck, apiKey)
//	url = fmt.Sprintf("https://steamcommunity.com/id/%s", loginToCheck)
//
//	// Выполняем GET-запрос к Steam API
//	response, err := http.Get(url)
//	if err != nil {
//		return false, err
//	}
//	defer response.Body.Close()
//
//	// Читаем ответ
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		return false, err
//	}
//
//	// Разбираем JSON-ответ
//	fmt.Println(string(body))
//	var result map[string]interface{}
//	if err := json.Unmarshal(body, &result); err != nil {
//		return false, err
//	}
//	fmt.Println(result)
//	// Проверяем результат
//	success, ok := result["response"].(map[string]interface{})["success"].(float64)
//	if !ok {
//		return false, fmt.Errorf("неверный формат ответа от Steam API")
//	}
//
//	return success == 1, nil
//}
