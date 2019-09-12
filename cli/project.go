package cli

import (
	"encoding/json"
	"fmt"

	errors "github.com/pkg/errors"
	config "github.com/wallix/awless/config"
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

/*
	url := "/customer/v1/websites/me"
		customer/v1/websites/logout
		req := cli.API().Path(url).Get()
		// cli.HandleBefore(handlerPath, params, req)
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
			if _, ok := decoded["contactEmail"]; ok {
				fmt.Printf("Username: %s, Id: %s\n", decoded["contactEmail"], decoded["accountId"])
			}
		} else {
			logger.Error("Invalid access token.")
		}
*/

type ProjectApi struct {
	Id   string `json:"projectId"`
	Name string `json:"name"`
}

func (api *APIInternal) Projects() *ProjectApi {
	return &ProjectApi{}
}

// POST TO sites/v1/databases/create using projectId as param

func (api *ProjectApi) GetProjects() (map[string]ProjectApi, error) {
	var resp map[string]ProjectApi
	itm, err := API().GetByKey("/customer/v1/websites/me", "projects")
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

func (api *ProjectApi) ListDatabase(projectId string, name string, dbType string) (map[string]interface{}, error) {
	// TODO ....
	return nil, nil
}

func (api *ProjectApi) CreatDatabase(projectId string, name string, dbType string) (map[string]interface{}, error) {
	// var resp map[string]ProjectApi

	url := "/sites/v1/databases/create"
	req := API().Path(url).Post()

	// Method to be used
	// req.Method("POST")

	data := map[string]string{"projectId": config.GetProjectId(), "type": dbType, "name": name}
	req.Use(body.JSON(data))

	// Perform the request
	res, err := req.Do()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return nil, errors.Wrap(err, fmt.Sprintf("Request error: %s\n", err))
	}
	if !res.Ok {
		// fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return nil, errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
	return nil, nil
}

/*
func (api *APIInternal) Project() (map[string]interface{}, err) {
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
		if _, ok := decoded["project"]; ok {
			fmt.Printf("Username: %s, Id: %s\n", decoded["contactEmail"], decoded["accountId"])
		}
		return decoded
	} else {
		logger.Error("Invalid access token.")
	}
}
*/
