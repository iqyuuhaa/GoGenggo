package long_polling

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gogenggo/config"
	"gogenggo/internals/pkg"
	"gogenggo/internals/platforms"
	"gogenggo/internals/types/constants"
	"gogenggo/internals/types/process"
	"gogenggo/internals/types/responses"
	"gogenggo/modules/cache"
	"gogenggo/utils"

	"github.com/go-co-op/gocron"
)

var longPollingModule *LongPollingModules

func Init(pkg *pkg.PkgWrapper, platform *platforms.PlatformModule) error {
	tempLongPollingModule := &LongPollingModules{
		Pkg:      pkg,
		Platform: platform,
	}

	longPollingModule = tempLongPollingModule
	longPollingModule.RegisterCron()

	return nil
}

func (m *LongPollingModules) RegisterCron() {
	ctx := context.Background()
	updateID, _ := strconv.ParseInt(os.Getenv(constants.LongPollingLastUpdateID), 10, 64)
	if updateID == 0 {
		bodyResponse, err := m.Pkg.Telegram.GetUpdates(ctx, 0, 100, 0, []string{})
		if err != nil {
			log.Println("[LongPolling - RegisterCron] Error getting updates, err:", err)
			return
		}

		var response responses.LongPollingTelegramResponse
		if err := json.Unmarshal(bodyResponse, &response); err != nil {
			log.Println("[LongPolling - RegisterCron] Error unmarshal get updates response, err:", err)
			return
		}

		if len(response.Result) != 0 {
			updateID = response.Result[len(response.Result)-1].UpdateID + 1
		}
	}

	enableGoroutine, _ := strconv.ParseBool(os.Getenv(constants.EnableLongPollingGoroutine))
	if !enableGoroutine {
		utils.AsyncToSync()
	}

	timezone := os.Getenv(constants.Timezone)
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	loc, _ := time.LoadLocation(timezone)
	s := gocron.NewScheduler(loc)
	s.Every(config.Configs.Main.System.LongPollingPeriodTime).Seconds().Do(func() {
		bodyResponse, err := m.Pkg.Telegram.GetUpdates(ctx, updateID, 100, 0, []string{})
		if err != nil {
			log.Println("[LongPolling - RegisterCron] Error getting updates, err:", err)
			return
		}

		var response responses.LongPollingTelegramResponse
		if err := json.Unmarshal(bodyResponse, &response); err != nil {
			log.Println("[LongPolling - RegisterCron] Error unmarshal get updates response, err:", err)
			return
		}

		for key, value := range response.Result {
			utils.AsyncFunc(func() {
				if key == len(response.Result)-1 {
					updateID = value.UpdateID + 1
				}

				message := ""
				dialogflowProcessTime := float64(0)
				nondialogflowProcessTime := float64(0)

				userSessionID := cache.GetUserLatestSession(value.Message.From.ID)
				if userSessionID == "" {
					cache.SetUserRequest(value.Message.From.ID)
					userSessionID = cache.GetUserLatestSession(value.Message.From.ID)
				}

				if config.Configs.Main.System.IsUsingDialogflow {
					dialogflowNow := time.Now()
					indentResp, err := m.Pkg.Dialogflow.IndentDetectText(ctx, value.Message.Text, userSessionID)
					if err != nil {
						log.Printf("[LongPolling - RegisterCron] Error detect indent text with text: %s, err: %v", value.Message.Text, err)
						return
					}
					dialogflowProcessTime = time.Since(dialogflowNow).Seconds()

					if indentResp.Intent.EndInteraction {
						cache.DeleteUserRequest(value.Message.From.ID)
					}

					message = indentResp.GetFulfillmentText()
				} else {
					nondialogflowNow := time.Now()
					chatKey := constants.NotFoundKey
					if utils.InArray(value.Message.Text, constants.IntroductionChat[constants.HiIntroductionKey], true) {
						chatKey = constants.HiIntroductionKey
					}

					if userSessionID != "" {
						for key, v := range constants.MenusChat {
							if utils.InArray(value.Message.Text, v, false) {
								chatKey = key
								break
							}
						}

						if chatKey != constants.NotFoundKey {
							cache.DeleteUserRequest(value.Message.From.ID)
						}
					}

					chatbotResponse, err := m.Platform.DB.ChatbotCommunication.GetResponse(ctx, chatKey)
					if err != nil {
						log.Printf("[LongPolling - RegisterCron] Error getting chatbot response with text: %s, err: %v", value.Message.Text, err)
						return
					}
					nondialogflowProcessTime = time.Since(nondialogflowNow).Seconds()

					message = chatbotResponse
				}

				replyMessage := utils.ReplaceStringsFormat(message, map[string]string{
					"user_id":  fmt.Sprint(value.Message.From.ID),
					"name":     value.Message.From.FirstName,
					"day_time": utils.GetCurrentDayTime(),
				})

				sendMessageNow := time.Now()
				if err := m.Pkg.Telegram.SendMessage(ctx, value.Message.From.ID, replyMessage); err != nil {
					log.Printf("[LongPolling - RegisterCron] Error sending message to %s, err: %v", value.Message.From.FirstName, err)
					return
				}
				sendMessageProcessTime := time.Since(sendMessageNow).Seconds()
				dateTime := time.Unix(value.Message.Date, 0)

				if err := m.Platform.DB.ChatbotHistory.InsertChatbotHistory(ctx, process.ChatbotHistoryProcess{
					Method:                   constants.LongPollingMethod,
					Identifier:               value.Message.From.FirstName,
					Datetime:                 dateTime,
					ProcessTime:              time.Since(dateTime).Seconds(),
					DialogflowProcessTime:    dialogflowProcessTime,
					NonDialogflowProcessTime: nondialogflowProcessTime,
					SendMessageProcessTime:   sendMessageProcessTime,
				}); err != nil {
					log.Printf("[LongPolling - RegisterCron] Error insert bulk chatbot history, err: %v", err)
					return
				}
			})
		}
	})

	s.StartAsync()
}
