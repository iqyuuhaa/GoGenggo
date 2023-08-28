package secret

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gogenggo/internals/types/constants"
	"gogenggo/utils"
)

var SecretObjects *SecretObject

func Init() error {
	populateEnvironmentCredentials()
	if SecretObjects != nil {
		return nil
	}

	env := utils.GetEnvironment()
	if env == "" {
		env = constants.DevelopmentEnvironment
	}

	secretKey := utils.GetEnv(constants.GolangChatbotSecretKey)
	if env == constants.ProductionEnvironment && secretKey == "" {
		log.Fatalln("[Secret - Init] Failed init secret data, secret key is required")
		return constants.ErrorWrongCredentials
	}

	if env == constants.DevelopmentEnvironment && secretKey == "" {
		secretKey = "YrYmhhRnp0S2yb3Txb2Cyx87G14MIpvK"
	}

	secretText := utils.GetEnv(constants.GolangChatbotSecret)
	if secretText == "" {
		secret, err := ioutil.ReadFile(fmt.Sprintf("%s/golang-chatbot.%s.json", constants.TlogoSecretPath, env))
		if err != nil {
			log.Fatalln("[Secret - Init] Error open secret file, err: ", err)
			return err
		}

		secretText = string(secret)
	}

	normalizedSecret := utils.RemoveSaltingSuffix(string(secretText))
	decriptedPhaseOneSecret, err := base64.StdEncoding.DecodeString(normalizedSecret)
	if err != nil {
		log.Fatalln("[Secret - Init] Error decode secret data using base64, err: ", err)
		return err
	}

	decriptedSecret, err := decryptGCMCipher([]byte(decriptedPhaseOneSecret), []byte(secretKey))
	if err != nil {
		log.Fatalln("[Secret - Init] Error decrypt secret data, err: ", err)
		return err
	}

	if err := json.Unmarshal(decriptedSecret, &SecretObjects); err != nil {
		log.Fatalln("[Secret - Init] Error unmarshall secret data, err: ", err)
		return err
	}

	return nil
}

func populateEnvironmentCredentials() {
	if utils.GetEnvironment() == constants.ProductionEnvironment {
		if utils.GetEnv(constants.GoogleApplicationCredentials) == "" {
			log.Fatalln("[Secret - populateEnvironmentCredentials] Empty google application environment credentials")
		}

		resp, _ := http.Get(utils.GetEnv(constants.GoogleApplicationCredentials))
		defer resp.Body.Close()

		tmpFile, err := os.Create("./google_application_credentials.json")
		if err != nil {
			log.Fatalln("[Secret - populateEnvironmentCredentials] Error create json file, err:", err)
		}
		defer tmpFile.Close()

		if _, err := io.Copy(tmpFile, resp.Body); err != nil {
			log.Fatalln("[Secret - populateEnvironmentCredentials] Error copy file, err:", err)
		}

		os.Setenv(constants.GoogleApplicationCredentials, "./google_application_credentials.json")
	}

	if utils.GetEnvironment() == "" || utils.GetEnvironment() == constants.DevelopmentEnvironment {
		os.Setenv(constants.GoogleApplicationCredentials, constants.GoogleApplicationCredentialsPath)
	}
}
