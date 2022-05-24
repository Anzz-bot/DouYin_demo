/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 21:53:41
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-23 21:40:51
 */
package models

import (
	"gorm.io/gorm"
	"time"
)

// primary key ID
type ID struct {
	ID uint64 `json:"id" gorm:"primaryKey"`
}

// create and update time
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// soft delete
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
