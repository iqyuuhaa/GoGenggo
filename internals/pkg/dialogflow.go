package pkg

import (
	"context"
	"fmt"
	"gogenggo/internals/types/constants"
	"gogenggo/modules/secret"
	"log"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
)

func (p *PkgModules) IndentDetectText(ctx context.Context, text, userSessionID string) (*dialogflowpb.QueryResult, error) {
	client, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		log.Println("[Pkg - DetectIntentText] Error creating new session, err:", err)
		return nil, err
	}
	defer client.Close()

	sessionPath := fmt.Sprintf(constants.DialogflowSessionPath, secret.SecretObjects.Dialogflow.ProjectID, userSessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: "en-US"}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	queryParams := dialogflowpb.QueryParameters{
		TimeZone: "Asia/Jakarta",
		SentimentAnalysisRequestConfig: &dialogflowpb.SentimentAnalysisRequestConfig{
			AnalyzeQueryTextSentiment: true,
		},
	}

	request := dialogflowpb.DetectIntentRequest{
		Session:     sessionPath,
		QueryInput:  &queryInput,
		QueryParams: &queryParams,
	}

	response, err := client.DetectIntent(ctx, &request)
	if err != nil {
		log.Println("[Pkg - DetectIntentText] Error detect indent text, err:", err)
		return nil, err
	}

	return response.GetQueryResult(), nil
}
