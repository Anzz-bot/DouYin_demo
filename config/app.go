/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 16:29:31
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 16:29:31
 */
package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	port    string `mapstrusture:"port" json:"port" yaml:"port"`
	AppName string `mapstrusture:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}
