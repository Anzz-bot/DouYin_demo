/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 16:32:13
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 16:32:13
 */
package config

type Configuration struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
}
