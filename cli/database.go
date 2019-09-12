package cli

import (
	"encoding/json"
	"fmt"

	errors "github.com/pkg/errors"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

/*	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
        errors "github.com/pkg/errors"

	"github.com/alecthomas/chroma/quick"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/auth"
	config "github.com/wallix/awless/config"
	global "github.com/wallix/awless/global"
	logger "github.com/wallix/awless/logger"
	"gopkg.in/h2non/gentleman.v2/context"
*/

type DatabaseApi struct {
	Id           string `json:"id"`
	Name         string `json:"title"`
	Status       string `json:"type"`
	ACenterAType string `json:"acentera_type"`
	Code         int    `json:"code"`
	Database     string `json:"dbname"`
	Pass         string `json:"pass"`
	User         string `json:"user"`
}

func DatabaseApiObject() DatabaseApi {
	var tmp DatabaseApi
	return tmp
}

func (api *APIInternal) Databases() *DatabaseApi {
	return &DatabaseApi{}
}

func (api *DatabaseApi) ListDatabases(dbserverId string) ([]DatabaseApi, error) {
	var resp []DatabaseApi
	itm, err := API().GetArr(fmt.Sprintf("/sites/v1/databases/%s/databases", dbserverId))
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(itm)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return nil, err
	}
	return resp, err
}

func (api *DatabaseApi) GetDatabases(projectId string) ([]DatabaseApi, error) {
	var resp []DatabaseApi
	itm, err := API().GetArr(fmt.Sprintf("/sites/v1/projects/%s/databases", projectId))
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(itm)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return nil, err
	}
	return resp, err
}

func (api *DatabaseApi) CreateDatabase(projectId string, websiteId string, stageId string, dbName string, dbServerId string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/sites/v1/websites/%s/%s/stage/%s/database/create", projectId, websiteId, stageId)

	req := API().Path(url).Post()

	// Method to be used
	// req.Method("POST")

	data := make(map[string]interface{}, 0)
	data["projectId"] = projectId
	data["dbserverId"] = dbServerId
	data["title"] = dbName
	data["stageId"] = stageId
	data["websiteId"] = websiteId

	req.Use(body.JSON(data))

	// Perform the request
	res, err := req.Do()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Request error: %s\n", err))
	}
	if !res.Ok {
		// fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return nil, errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}

	// fmt.Printf("zs Status: %d\n", res.StatusCode)
	// fmt.Printf("sf Body: %s", res.String())
	// var resp map[string]DatabaseApi
	return nil, err
}

/*
func (api *DatabaseApi) CreateSite(repoName string, uuid string, fullRepository string, sshRepo string, httpRepo string, displayName string) (map[string]DatabaseApi, error) {
	url := "/sites/v1/websites/create"
	req := API().Path(url).Post()

	// Method to be used
	// req.Method("POST")

	data := make(map[string]interface{}, 0)
	data["projectId"] = config.GetProjectId()
	data["type"] = "git"
	data["git_options"] = "skip_with_basicauth" // since we create it ourselves without any templates ... for now
	data["title"] = displayName
	data["branch"] = "master"
	data["acentera_type"] = "docker-simple-wp-legacy"
	data["ssh_repository"] = sshRepo
	data["http_repository"] = httpRepo
	data["repository"] = fullRepository
	repodetails := map[string]string{}
	repodetails["uuid"] = uuid
	repodetails["name"] = repoName
	data["repodetails"] = repodetails

	//todo : if bitbucket ???
	token, _ := config.Get("_bitbucket.token")
	if token != nil {
		c := bitbucket.NewOAuthbearerToken(token.(string))
		_, er := c.User.Profile()
		if er == nil {
			fmt.Println("added AUTH ? token")
			data["auth"] = token.(string)
		}
	}
	if _, ok := data["auth"]; !ok {
		auth := config.GetGitUsername("bitbucket") + ":" + config.GetGitPassword("bitbucket")
		data["auth"] = base64.StdEncoding.EncodeToString([]byte(auth))
	}
	data["token"] = "" // leave empty as we want to use basic auth

	req.Use(body.JSON(data))

	// Perform the request
	res, err := req.Do()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Request error: %s\n", err))
	}
	if !res.Ok {
		// fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return nil, errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}

	fmt.Printf("zs Status: %d\n", res.StatusCode)
	fmt.Printf("sf Body: %s", res.String())

	if res.StatusCode == 200 {
		fmt.Println("Will create stage ?")
		var decoded map[string]interface{}
		if res.StatusCode < 400 {
			if err := UnmarshalResponse(res, &decoded); err != nil {
				// fmt.Println(errors.Wrap(err, "Unmarshalling response failed"))
				logger.Error(errors.Wrap(err, "Unmarshalling response failed"))
			}
		} else {
			logger.ExtraVerbosef("HTTP %d: %s", res.StatusCode, res.String())
		}
		websiteId := ""
		if _, ok := decoded["websiteId"]; ok {
			websiteId = decoded["websiteId"].(string)
		}
		projectId := ""
		if _, ok := decoded["projectId"]; ok {
			projectId = decoded["projectId"].(string)
		}

		fmt.Println("Will create stage using prj: ", projectId, websiteId)
		url = fmt.Sprintf("/sites/v1/websites/%s/%s/stage/create/master", projectId, websiteId)

		data := make(map[string]interface{}, 0)
		data["title"] = "master"

		req = API().Path(url).Post()

		req.Use(body.JSON(data))

		// Perform the request
		res, err := req.Do()
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Request error: %s\n", err))
		}
		if !res.Ok {
			// fmt.Printf("Invalid server response: %d\n", res.StatusCode)
			return nil, errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
		}

		fmt.Printf("01 - Status: %d\n", res.StatusCode)
		fmt.Printf("01 - Body: %s", res.String())
	}

	return nil, nil
}
*/

/*
func (api *APIInternal) Database() (map[string]interface{}, err) {
	url := "/customer/v1/websites/me"
	req := cli.API().Path(url).Get()
	resp, err := req.Do()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Request failed"))
	}
	var decoded map[string]interface{}
	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			fmt.Println(errors.Wrap(err, "Unmarshalling response failed"))
		}
	} else {
		fmt.Println(errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String()))
	}
	if resp.StatusCode <= 208 && resp.StatusCode >= 200 {
		logger.Verbosef("Got decoded of %s", decoded)
		if _, ok := decoded["website"]; ok {
			fmt.Printf("Username: %s, Id: %s\n", decoded["contactEmail"], decoded["accountId"])
		}
		return decoded
	} else {
		logger.Error("Invalid access token.")
	}
}
*/
