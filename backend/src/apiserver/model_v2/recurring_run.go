// Copyright 2022 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model_v2

import "fmt"

type mode string

const (
	mode_UNSPECIFIED mode = "MODE_UNSPECIFIED"
	mode_ENABLE      mode = "ENABLE"
	mode_DISABLE     mode = "DISABLE"
)

type status string

const (
	status_UNSPECIFIED status = "UNKNOWN_UNSPECIFIED"
	status_ENABLED     status = "ENABLED"
	status_DISABLED    status = "DISABLED"
	status_ERROR       status = "ERROR"
)

type RecurringRun struct {
	UUID        string `gorm:"column:UUID; not null; primary_key"`
	DisplayName string `gorm:"column:DisplayName; not null;"`
	Description string `gorm:"column:Description; not null"`
	PipelineSource
	RuntimeConfig
	ServiceAccount string `gorm:"column:ServiceAccount; not null;"`
	MaxConcurrency int64  `gorm:"column:MaxConcurrency;not null"`
	Trigger
	Mode           mode   `gorm:"type:mode; not null"`
	CreatedAtInSec int64  `gorm:"column:CreatedAtInSec; not null"` /* The time this record is stored in DB*/
	UpdatedAtInSec int64  `gorm:"column:UpdatedAtInSec; not null"`
	Status         status `gorm:"type:status; not null"`
	Error          string `gorm:"column:Error; not null"`
	NoCatchup      bool   `gorm:"column:NoCatchup; not null"`
	Namespace      string `gorm:"column:Namespace; not null;"`
	ExperimentID   string `gorm:"column:ExperimentID; not null;`
}

// Source of the pipeline spec. Can be either a pipeline ID, or a pipeline spec.
type PipelineSource struct {
	// ID of the pipeline.
	PipelineId string `gorm:"column:PipelineId;"`
	// Pipeline Spec.
	PipelineSpec string `gorm:"column:pipelineSpec;"`
}

// Trigger specifies when to create a new workflow.
type Trigger struct {
	// Create workflows according to a cron schedule.
	CronSchedule
	// Create workflows periodically.
	PeriodicSchedule
}

type CronSchedule struct {
	// Time at which scheduling starts.
	// If no start time is specified, the StartTime is the creation time of the schedule.
	CronScheduleStartTimeInSec *int64 `gorm:"column:CronScheduleStartTimeInSec;"`

	// Time at which scheduling ends.
	// If no end time is specified, the EndTime is the end of time.
	CronScheduleEndTimeInSec *int64 `gorm:"column:CronScheduleEndTimeInSec;"`

	// Cron string describing when a workflow should be created within the
	// time interval defined by StartTime and EndTime.
	Cron *string `gorm:"column:Schedule;"`
}

type PeriodicSchedule struct {
	// Time at which scheduling starts.
	// If no start time is specified, the StartTime is the creation time of the schedule.
	PeriodicScheduleStartTimeInSec *int64 `gorm:"column:PeriodicScheduleStartTimeInSec;"`

	// Time at which scheduling ends.
	// If no end time is specified, the EndTime is the end of time.
	PeriodicScheduleEndTimeInSec *int64 `gorm:"column:PeriodicScheduleEndTimeInSec;"`

	// Interval describing when a workflow should be created within the
	// time interval defined by StartTime and EndTime.
	IntervalSecond *int64 `gorm:"column:IntervalSecond;"`
}

func (r RecurringRun) GetValueOfPrimaryKey() string {
	return fmt.Sprint(r.UUID)
}

func GetRecurringRunTablePrimaryKeyColumn() string {
	return "UUID"
}

// PrimaryKeyColumnName returns the primary key for model RecurringRun.
func (r *RecurringRun) PrimaryKeyColumnName() string {
	return "UUID"
}

// DefaultSortField returns the default sorting field for model RecurringRun.
func (r *RecurringRun) DefaultSortField() string {
	return "CreatedAtInSec"
}

var recurringRunAPIToModelFieldMap = map[string]string{
	"id":         "UUID",
	"name":       "DisplayName",
	"created_at": "CreatedAtInSec",
	"package_id": "PipelineId",
}

// APIToModelFieldMap returns a map from API names to field names for model RecurringRun.
func (r *RecurringRun) APIToModelFieldMap() map[string]string {
	return recurringRunAPIToModelFieldMap
}

// GetModelName returns table name used as sort field prefix
func (r *RecurringRun) GetModelName() string {
	return "recurring_runs"
}

func (r *RecurringRun) GetField(name string) (string, bool) {
	if field, ok := recurringRunAPIToModelFieldMap[name]; ok {
		return field, true
	}
	return "", false
}

func (r *RecurringRun) GetFieldValue(name string) interface{} {
	switch name {
	case "UUID":
		return r.UUID
	case "DisplayName":
		return r.DisplayName
	case "CreatedAtInSec":
		return r.CreatedAtInSec
	case "PipelineId":
		return r.PipelineId
	default:
		return nil
	}
}

func (r *RecurringRun) GetSortByFieldPrefix(name string) string {
	return "recurring_runs."
}

func (r *RecurringRun) GetKeyFieldPrefix() string {
	return "recurring_runs."
}
