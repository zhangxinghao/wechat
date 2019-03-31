package oauth2

import (
	mpoauth2 "gopkg.in/zhangxinghao/wechat.v3/mp/oauth2"
	"gopkg.in/zhangxinghao/wechat.v3/oauth2"
)

var _ oauth2.Endpoint = (*Endpoint)(nil)

type Endpoint mpoauth2.Endpoint

func NewEndpoint(AppId, AppSecret string) *Endpoint {
	return (*Endpoint)(mpoauth2.NewEndpoint(AppId, AppSecret))
}

func (p *Endpoint) ExchangeTokenURL(code string) string {
	return ((*mpoauth2.Endpoint)(p)).ExchangeTokenURL(code)
}

func (p *Endpoint) RefreshTokenURL(refreshToken string) string {
	return ((*mpoauth2.Endpoint)(p)).RefreshTokenURL(refreshToken)
}
