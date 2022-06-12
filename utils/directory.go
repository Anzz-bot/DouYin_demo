/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 19:10:11
 * @Last Modified by: alexander.huang hyperyyyy
 * @Last Modified time: 2022-05-18 19:10:11
 */
package utils

import "os"

// judge path exist
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
