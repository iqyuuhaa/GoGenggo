package usecases

import (
	"context"
	"fmt"
	"log"
	"time"

	"gogenggo/config"
	"gogenggo/internals/types/constants"
	"gogenggo/internals/types/payloads"
	"gogenggo/internals/types/process"
	"gogenggo/internals/types/responses"
	"gogenggo/modules/cache"
	"gogenggo/utils"
)

func (usecase *UsecaseModules) WebhookDialogflowUsecase(ctx context.Context, payload payloads.TelegramWebhookPayload) error {
	message := ""
	dialogflowProcessTime := float64(0)
	nondialogflowProcessTime := float64(0)

	userSessionID := cache.GetUserLatestSession(payload.Message.From.ID)
	if config.Configs.Main.System.IsUsingDialogflow {
		if userSessionID == "" {
			cache.SetUserRequest(payload.Message.From.ID)
			userSessionID = cache.GetUserLatestSession(payload.Message.From.ID)
		}

		dialogflowNow := time.Now()
		indentResp, err := usecase.Pkg.Dialogflow.IndentDetectText(ctx, payload.Message.Text, userSessionID)
		if err != nil {
			log.Printf("[Usecase - WebhookDialogflowUsecase] Error detect indent text with text: %s, err: %v", payload.Message.Text, err)
			return err
		}
		dialogflowProcessTime = time.Since(dialogflowNow).Seconds()

		if indentResp.Intent.EndInteraction {
			cache.DeleteUserRequest(payload.Message.From.ID)
		}

		message = indentResp.GetFulfillmentText()
	} else {
		nondialogflowNow := time.Now()
		chatKey := constants.NotFoundKey
		if userSessionID == "" {
			if utils.InArray(payload.Message.Text, constants.IntroductionChat[constants.HiIntroductionKey], true) {
				cache.SetUserRequest(payload.Message.From.ID)
				chatKey = constants.HiIntroductionKey
			}
		} else {
			cache.SetUserRequest(payload.Message.From.ID)
			for key, value := range constants.MenusChat {
				if utils.InArray(payload.Message.Text, value, false) {
					chatKey = key
					break
				}
			}

			if chatKey != constants.NotFoundKey {
				cache.DeleteUserRequest(payload.Message.From.ID)
			}
		}

		chatbotResponse, err := usecase.Platform.DB.ChatbotCommunication.GetResponse(ctx, chatKey)
		if err != nil {
			log.Printf("[Usecase - WebhookDialogflowUsecase] Error getting chatbot response with text: %s, err: %v", payload.Message.Text, err)
			return err
		}
		nondialogflowProcessTime = time.Since(nondialogflowNow).Seconds()

		message = chatbotResponse
	}

	replyMessage := utils.ReplaceStringsFormat(message, map[string]string{
		"user_id":  fmt.Sprint(payload.Message.From.ID),
		"name":     payload.Message.From.FirstName,
		"day_time": utils.GetCurrentDayTime(),
	})

	sendMessageNow := time.Now()
	if err := usecase.Pkg.Telegram.SendMessage(ctx, payload.Message.From.ID, replyMessage); err != nil {
		log.Printf("[Usecase - WebhookDialogflowUsecase] Error sending message to %s, err: %v", payload.Message.From.FirstName, err)
		return err
	}
	sendMessageProcessTime := time.Since(sendMessageNow).Seconds()

	dateTime := time.Unix(payload.Message.Date, 0)
	if err := usecase.Platform.DB.ChatbotHistory.InsertChatbotHistory(ctx, process.ChatbotHistoryProcess{
		Method:                   constants.WebhookMethod,
		Identifier:               payload.Message.From.FirstName,
		Datetime:                 dateTime,
		ProcessTime:              time.Since(dateTime).Seconds(),
		DialogflowProcessTime:    dialogflowProcessTime,
		NonDialogflowProcessTime: nondialogflowProcessTime,
		SendMessageProcessTime:   sendMessageProcessTime,
	}); err != nil {
		log.Printf("[Usecase - WebhookDialogflowUsecase] Error saving chatbot history with name %s, err: %v", payload.Message.From.FirstName, err)
	}

	return nil
}

func (usecase *UsecaseModules) TestDialogflowUsecase(ctx context.Context, payload payloads.DialogflowWebhookPayload) (response responses.DialogflowWebhookResponse, err error) {
	indentResp, err := usecase.Pkg.Dialogflow.IndentDetectText(ctx, payload.Text, "test")
	if err != nil {
		return response, err
	}

	replyMessage := utils.ReplaceStringsFormat(indentResp.GetFulfillmentText(), map[string]string{
		"name":     "nama",
		"day_time": utils.GetCurrentDayTime(),
	})

	response.Text = replyMessage

	return response, nil
}

// func (usecase *UsecaseModules) LiveChatBotUsecase(ctx context.Context, payload payloads.ChatBotPayload) (err error) {
// if config.Configs.Main.System.IsMaintenance {
// 	usecase.Pkg.API.SendTextMessage(ctx, payload.User.Phone, payload.User.Name, config.Configs.Message.Default.MaintenanceModeMessage)
// }

// now := time.Now().Unix()
// textMessage := payload.Message.Text
// msisdn := utils.NormalizeMsisdnFormat(payload.User.Phone)
// if msisdn == "" {
// 	log.Println("[Usecase - LiveChatBotUsecase] Not expected msisdn, phone: ", payload.User.Phone)
// 	return nil
// }

// if isBanned, bannedTimeCache := cache.GetBannedTimeByMsisdn(msisdn); isBanned || (!bannedTimeCache.IsZero() && now < bannedTimeCache.Unix()) {
// 	return nil
// }

// user, err := usecase.Platform.DB.User.GetUserData(ctx, msisdn)
// if err != nil && err != constants.ErrorNotFoundData {
// 	log.Printf("[Usecase - LiveChatBotUsecase] Error getting user data with msisdn: %s, err: %v", msisdn, err)
// 	return err
// }

// isRegistered := true
// if err == constants.ErrorNotFoundData {
// 	isRegistered = false
// 	textMessage = constants.NotRegisteredKey
// }

// cache.SetMsisdnRequest(msisdn)
// if isPassedRequestRule := cache.CheckRequestRule(msisdn); isPassedRequestRule {
// 	bannedData, err := usecase.Platform.DB.Banned.GetBannedDataByMsisdn(ctx, msisdn)
// 	if err != nil && err != constants.ErrorNotFoundData {
// 		log.Printf("[Usecase - LiveChatBotUsecase] Error getting banned data with msisdn: %s, err: %v", msisdn, err)
// 		return err
// 	}

// 	counter := 0
// 	untilTime := time.Now().AddDate(0, 0, 3)
// 	if err == nil {
// 		counter = bannedData.Counter + 1
// 		if bannedData.Counter >= config.Configs.Main.Banned.MaxCounter {
// 			untilTime = time.Time{}
// 		}
// 	}

// 	if err := usecase.Platform.DB.Banned.InsertBanned(ctx, msisdn, counter, untilTime, isRegistered); err != nil {
// 		log.Printf("[Usecase - LiveChatBotUsecase] Error save banned data with msisdn: %s, err: %v", msisdn, err)
// 	}

// 	cache.SetBannedMsisdn(msisdn, untilTime)
// 	if err := usecase.Pkg.API.SendTextMessage(ctx, payload.User.Phone, payload.User.Name, fmt.Sprintf(config.Configs.Message.Default.BannedMessage, utils.FormatingDateTime(bannedData.ValidUntil.Time))); err != nil {
// 		log.Printf("[Usecase - LiveChatBotUsecase] Error sending message to %s, err: %v", msisdn, err)
// 	}

// 	return nil
// }

// indentResp, err := usecase.Pkg.Dialogflow.IndentDetectText(ctx, textMessage, cache.GetMsisdnLatestSession(msisdn))
// if err != nil {
// 	log.Printf("[Usecase - LiveChatBotUsecase] Error detect indent text with text: %s, err: %v", textMessage, err)
// 	return err
// }

// replyMessage := utils.ReplaceStringsFormat(indentResp.GetFulfillmentText(), map[string]string{
// 	"name":     payload.User.Name,
// 	"day_time": utils.GetCurrentDayTime(),
// })

// if err := usecase.Pkg.API.SendTextMessage(ctx, payload.User.Phone, payload.User.Name, replyMessage); err != nil {
// 	log.Printf("[Usecase - LiveChatBotUsecase] Error sending message to %s, err: %v", msisdn, err)
// 	return err
// }

// 	return nil
// }
