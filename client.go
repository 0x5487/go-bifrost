package client

import (
	"fmt"

	"github.com/jasonsoft/request"
)

type Client struct {
	address string
	token   string
}

func NewClient(address, token string) *Client {
	return &Client{
		address: address,
		token:   token,
	}
}

func (c *Client) CreateOrUpdateConsumer(consumer *Consumer) (*Consumer, error) {
	url := c.address + "/v1/consumers"
	resp, err := request.PUT(url).
		Set("Authorization", c.token).
		End()

	if err != nil {
		return nil, err
	}

	if resp.OK {
		result := Consumer{}
		err := resp.JSON(&result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	var bifrostErr AppError
	resp.JSON(&bifrostErr)
	err = AppError{
		StatusCode: resp.StatusCode,
		ErrorCode:  bifrostErr.ErrorCode,
		Message:    bifrostErr.Message,
	}
	return nil, err
}

func (c *Client) CreateToken(token *Token) (*Token, error) {
	url := c.address + "/v1/tokens"
	resp, err := request.POST(url).
		Set("Authorization", c.token).
		End()

	if err != nil {
		return nil, err
	}

	if resp.OK {
		result := Token{}
		err := resp.JSON(&result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	var bifrostErr AppError
	resp.JSON(&bifrostErr)
	err = AppError{
		StatusCode: resp.StatusCode,
		ErrorCode:  bifrostErr.ErrorCode,
		Message:    bifrostErr.Message,
	}
	return nil, err
}

func (c *Client) DeleteTokensByConsumerId(consumerId string) error {
	url := fmt.Sprintf("%s/v1/tokens?consumer_id=%s", c.address, consumerId)

	resp, err := request.DELETE(url).
		Set("Authorization", c.token).
		End()

	if err != nil {
		return err
	}

	if resp.OK {
		return nil
	}

	var bifrostErr AppError
	resp.JSON(&bifrostErr)
	err = AppError{
		StatusCode: resp.StatusCode,
		ErrorCode:  bifrostErr.ErrorCode,
		Message:    bifrostErr.Message,
	}
	return err
}
