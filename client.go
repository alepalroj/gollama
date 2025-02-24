package gollama

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Config struct {
	APIURL             	string			`yaml:"api_url"`
	Timeout            	int				`yaml:"timeout"`
	RetryCount         	int    			`yaml:"retry_count"`
	RetryWaitTime      	int    			`yaml:"retry_wait_time"`
	RetryMaxWaitTime   	int    			`yaml:"retry_max_wait_time"`
	ContentType    		string 			`yaml:"content_type"`
	APIEndpoint			string 			`yaml:"api_endpoint"`
	Model              	string 			`yaml:"model"`
	Messages           	MessagesConfig 	`yaml:"messages"`
}

type MessagesConfig struct {
	ModelEmpty     string 	`yaml:"model_empty"`
	PromptEmpty    string 	`yaml:"prompt_empty"`
	RequestError   string 	`yaml:"request_error"`
	ResponseError  string 	`yaml:"response_error"`
	ReadConfigError string 	`yaml:"read_config_error"`
	ParseConfigError string `yaml:"parse_config_error"`
	StatusCodeError string 	`yaml:"status_code_error"`
}

type Request struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream"`
	Options map[string]interface{} `json:"options,omitempty"`
}

type Response struct {
	Model               string    `json:"model"`
	CreatedAt           time.Time `json:"created_at"`
	Response           	string    `json:"response"`
	Done               	bool      `json:"done"`
	Context            	[]int     `json:"context"`
	TotalDuration      	int64     `json:"total_duration"`
	LoadDuration       	int64     `json:"load_duration"`
	PromptEvalCount    	int       `json:"prompt_eval_count"`
	PromptEvalDuration 	int64     `json:"prompt_eval_duration"`
	EvalCount          	int       `json:"eval_count"`
	EvalDuration       	int64     `json:"eval_duration"`
}

type Client struct {
	httpClient *resty.Client
	config     Config
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf(config.Messages.ReadConfigError+": %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf(config.Messages.ParseConfigError+": %v", err)
	}
	return config, nil
}

func OllamaClient(config Config) *Client {
	client := resty.New()
	client.SetTimeout(time.Duration(config.Timeout) * time.Second)
	client.SetRetryCount(config.RetryCount)
	client.SetRetryWaitTime(time.Duration(config.RetryWaitTime) * time.Second)
	client.SetRetryMaxWaitTime(time.Duration(config.RetryMaxWaitTime) * time.Second)

	return &Client{
		httpClient: client,
		config:     config,
	}
}

func (c *Client) Generate(req Request) (*Response, error) {

	if req.Model == "" {
		return nil, fmt.Errorf(c.config.Messages.ModelEmpty)
	}

	if req.Prompt == "" {
		return nil, fmt.Errorf(c.config.Messages.PromptEmpty)
	}

	resp, err := c.httpClient.R().
		SetHeader(c.config.ContentTypeName, c.config.ContentType).
		SetBody(req).
		Post(c.config.APIURL + c.config.APIEndpoint)

	if err != nil {
		return nil, fmt.Errorf(c.config.Messages.RequestError+": %v", err)
	}

	var apiResp Response
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf(c.config.Messages.ResponseError+": %v", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf(c.config.Messages.StatusCodeError+": %s", resp.Status())
	}

	return &apiResp, nil
}

func (c *Client) GenerateAsync(req Request, callback func(*Response, error)) {
	go func() {
		resp, err := c.Generate(req)
		callback(resp, err)
	}()
}
