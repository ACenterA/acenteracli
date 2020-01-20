package cli

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/ktrysmt/go-bitbucket"
	errors "github.com/pkg/errors"
	config "github.com/wallix/awless/config"
	logger "github.com/wallix/awless/logger"
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

type BlueprintRelease struct {
	Enabled string `json:"enabled"`
	Release string `json:"title"`
}

type BlueprintApi struct {
	Id           string                      `json:"blueprintId"`
	Name         string                      `json:"title"`
	Status       string                      `json:"type"`
	ACenterAType string                      `json:"acentera_type"`
	Release      string                      `json:"release"`
	Project      string                      `json:"projectId"`
	Releases     map[string]BlueprintRelease `json:"releases"`
}

func BlueprintApiObject() BlueprintApi {
	var tmp BlueprintApi
	return tmp
}

func (api *APIInternal) Blueprints() *BlueprintApi {
	return &BlueprintApi{}
}

func (api *BlueprintApi) GetBlueprints(projectId string) (map[string]BlueprintApi, error) {
	var resp map[string]BlueprintApi
	itm, err := API().GetByKey(fmt.Sprintf("/api/projects/v1/%s", projectId), "blueprints")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	jsonData, err := json.Marshal(itm)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, err
}

func (api *BlueprintApi) GetBlueprintsById(projectId string, blueprintId string) (*BlueprintApi, error) {
	var resp BlueprintApi
	itm, err := API().Get(fmt.Sprintf("sites/v1/blueprints/%s/%s", projectId, blueprintId))
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
	return &resp, err
}

func (api *BlueprintApi) CreateBlueprint(repoName string, uuid string, fullRepository string, sshRepo string, httpRepo string, displayName string) (map[string]BlueprintApi, error) {
	url := "/sites/v1/blueprints/create"
	req := API().Path(url).Post()

	// Method to be used
	// req.Method("POST")

	data := make(map[string]interface{}, 0)
	data["projectId"] = config.GetProjectId()
	data["type"] = "git"
	data["git_options"] = "skip_with_basicauth" // since we create it ourselves without any templates ... for now
	data["title"] = displayName
	data["branch"] = "master"
	data["acentera_type"] = "blueprint" // blueprintci.xml required in java (ie it adds ci.xml)
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

	// fmt.Printf("zs Status: %d\n", res.StatusCode)
	// fmt.Printf("sf Body: %s", res.String())

	if res.StatusCode == 200 {
		// fmt.Println("Will create stage ?")
		var decoded map[string]interface{}
		if res.StatusCode < 400 {
			if err := UnmarshalResponse(res, &decoded); err != nil {
				// fmt.Println(errors.Wrap(err, "Unmarshalling response failed"))
				logger.Error(errors.Wrap(err, "Unmarshalling response failed"))
			}
		} else {
			logger.ExtraVerbosef("HTTP %d: %s", res.StatusCode, res.String())
		}
		blueprintId := ""
		if _, ok := decoded["blueprintId"]; ok {
			blueprintId = decoded["blueprintId"].(string)
		}
		projectId := ""
		if _, ok := decoded["projectId"]; ok {
			projectId = decoded["projectId"].(string)
		}

		fmt.Println("Will create release? using prj: ", projectId, blueprintId)

		// decoded

		// url = fmt.Sprintf("/sites/v1/blueprints/%s/%s/release/create/master", projectId, blueprintId)

		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Request error: %s\n", err))
		}
		if !res.Ok {
			// fmt.Printf("Invalid server response: %d\n", res.StatusCode)
			return nil, errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
		}
		/*
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

			// fmt.Printf("01 - Status: %d\n", res.StatusCode)
			// fmt.Printf("01 - Body: %s", res.String())
		*/
	}

	return nil, nil
}

/*
func (api *APIInternal) Website() (map[string]interface{}, err) {
	url := "/customer/v1/blueprints/me"
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
		if _, ok := decoded["blueprint"]; ok {
			fmt.Printf("Username: %s, Id: %s\n", decoded["contactEmail"], decoded["accountId"])
		}
		return decoded
	} else {
		logger.Error("Invalid access token.")
	}
}
*/
