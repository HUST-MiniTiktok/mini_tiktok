// Code generated by Validator v0.1.4. DO NOT EDIT.

package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *UserRegisterRequest) IsValid() error {
	if len(p.Username) < int(1) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Username) > int(32) {
		return fmt.Errorf("field Username max_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	if len(p.Password) > int(32) {
		return fmt.Errorf("field Password max_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *UserRegisterResponse) IsValid() error {
	return nil
}
func (p *UserLoginRequest) IsValid() error {
	if len(p.Username) < int(1) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Username) > int(32) {
		return fmt.Errorf("field Username max_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	if len(p.Password) > int(32) {
		return fmt.Errorf("field Password max_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *UserLoginResponse) IsValid() error {
	return nil
}
func (p *UserRequest) IsValid() error {
	return nil
}
func (p *UserResponse) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("field User not valid, %w", err)
		}
	}
	return nil
}
func (p *CheckUserIsExistRequest) IsValid() error {
	return nil
}
func (p *CheckUserIsExistResponse) IsValid() error {
	return nil
}
