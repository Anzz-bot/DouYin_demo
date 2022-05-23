/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 16:32:13
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-23 21:40:51
 */
package config

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
}
