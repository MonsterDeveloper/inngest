// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type FunctionRunEvent interface {
	IsFunctionRunEvent()
}

type ActionVersionQuery struct {
	Dsn          string `json:"dsn"`
	VersionMajor *int   `json:"versionMajor,omitempty"`
	VersionMinor *int   `json:"versionMinor,omitempty"`
}

type CreateAppInput struct {
	URL string `json:"url"`
}

type Event struct {
	ID           string         `json:"id"`
	Workspace    *Workspace     `json:"workspace,omitempty"`
	Name         *string        `json:"name,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	Payload      *string        `json:"payload,omitempty"`
	Schema       *string        `json:"schema,omitempty"`
	Status       *EventStatus   `json:"status,omitempty"`
	PendingRuns  *int           `json:"pendingRuns,omitempty"`
	TotalRuns    *int           `json:"totalRuns,omitempty"`
	Raw          *string        `json:"raw,omitempty"`
	FunctionRuns []*FunctionRun `json:"functionRuns,omitempty"`
}

type EventQuery struct {
	WorkspaceID string `json:"workspaceId"`
	EventID     string `json:"eventId"`
}

type EventsQuery struct {
	WorkspaceID string  `json:"workspaceId"`
	LastEventID *string `json:"lastEventId,omitempty"`
}

type Function struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Slug        string             `json:"slug"`
	Config      string             `json:"config"`
	Concurrency int                `json:"concurrency"`
	Triggers    []*FunctionTrigger `json:"triggers,omitempty"`
	URL         string             `json:"url"`
}

type FunctionEvent struct {
	Workspace   *Workspace         `json:"workspace,omitempty"`
	FunctionRun *FunctionRun       `json:"functionRun,omitempty"`
	Type        *FunctionEventType `json:"type,omitempty"`
	Output      *string            `json:"output,omitempty"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty"`
}

func (FunctionEvent) IsFunctionRunEvent() {}

type FunctionRun struct {
	ID           string             `json:"id"`
	FunctionID   string             `json:"functionID"`
	Function     *Function          `json:"function,omitempty"`
	Workspace    *Workspace         `json:"workspace,omitempty"`
	Event        *Event             `json:"event,omitempty"`
	Status       *FunctionRunStatus `json:"status,omitempty"`
	WaitingFor   *StepEventWait     `json:"waitingFor,omitempty"`
	PendingSteps *int               `json:"pendingSteps,omitempty"`
	StartedAt    *time.Time         `json:"startedAt,omitempty"`
	FinishedAt   *time.Time         `json:"finishedAt,omitempty"`
	Output       *string            `json:"output,omitempty"`
	Timeline     []FunctionRunEvent `json:"timeline,omitempty"`
	Name         *string            `json:"name,omitempty"`
}

type FunctionRunQuery struct {
	WorkspaceID   string `json:"workspaceId"`
	FunctionRunID string `json:"functionRunId"`
}

type FunctionRunsQuery struct {
	WorkspaceID string `json:"workspaceId"`
}

type FunctionTrigger struct {
	Type  FunctionTriggerTypes `json:"type"`
	Value string               `json:"value"`
}

type StepEvent struct {
	Workspace   *Workspace     `json:"workspace,omitempty"`
	FunctionRun *FunctionRun   `json:"functionRun,omitempty"`
	StepID      *string        `json:"stepID,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Type        *StepEventType `json:"type,omitempty"`
	Output      *string        `json:"output,omitempty"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty"`
	WaitingFor  *StepEventWait `json:"waitingFor,omitempty"`
}

func (StepEvent) IsFunctionRunEvent() {}

type StepEventWait struct {
	EventName  *string   `json:"eventName,omitempty"`
	Expression *string   `json:"expression,omitempty"`
	ExpiryTime time.Time `json:"expiryTime"`
}

type StreamItem struct {
	ID        string         `json:"id"`
	Trigger   string         `json:"trigger"`
	Type      StreamType     `json:"type"`
	CreatedAt time.Time      `json:"createdAt"`
	Runs      []*FunctionRun `json:"runs,omitempty"`
}

type StreamQuery struct {
	After  *time.Time `json:"after,omitempty"`
	Before *time.Time `json:"before,omitempty"`
	Limit  int        `json:"limit"`
}

type UpdateAppInput struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Workspace struct {
	ID string `json:"id"`
}

type EventStatus string

const (
	EventStatusRunning         EventStatus = "RUNNING"
	EventStatusCompleted       EventStatus = "COMPLETED"
	EventStatusPaused          EventStatus = "PAUSED"
	EventStatusFailed          EventStatus = "FAILED"
	EventStatusPartiallyFailed EventStatus = "PARTIALLY_FAILED"
	EventStatusNoFunctions     EventStatus = "NO_FUNCTIONS"
)

var AllEventStatus = []EventStatus{
	EventStatusRunning,
	EventStatusCompleted,
	EventStatusPaused,
	EventStatusFailed,
	EventStatusPartiallyFailed,
	EventStatusNoFunctions,
}

func (e EventStatus) IsValid() bool {
	switch e {
	case EventStatusRunning, EventStatusCompleted, EventStatusPaused, EventStatusFailed, EventStatusPartiallyFailed, EventStatusNoFunctions:
		return true
	}
	return false
}

func (e EventStatus) String() string {
	return string(e)
}

func (e *EventStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventStatus", str)
	}
	return nil
}

func (e EventStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FunctionEventType string

const (
	FunctionEventTypeStarted   FunctionEventType = "STARTED"
	FunctionEventTypeCompleted FunctionEventType = "COMPLETED"
	FunctionEventTypeFailed    FunctionEventType = "FAILED"
	FunctionEventTypeCancelled FunctionEventType = "CANCELLED"
)

var AllFunctionEventType = []FunctionEventType{
	FunctionEventTypeStarted,
	FunctionEventTypeCompleted,
	FunctionEventTypeFailed,
	FunctionEventTypeCancelled,
}

func (e FunctionEventType) IsValid() bool {
	switch e {
	case FunctionEventTypeStarted, FunctionEventTypeCompleted, FunctionEventTypeFailed, FunctionEventTypeCancelled:
		return true
	}
	return false
}

func (e FunctionEventType) String() string {
	return string(e)
}

func (e *FunctionEventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FunctionEventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FunctionEventType", str)
	}
	return nil
}

func (e FunctionEventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FunctionRunStatus string

const (
	FunctionRunStatusCompleted FunctionRunStatus = "COMPLETED"
	FunctionRunStatusFailed    FunctionRunStatus = "FAILED"
	FunctionRunStatusCancelled FunctionRunStatus = "CANCELLED"
	FunctionRunStatusRunning   FunctionRunStatus = "RUNNING"
)

var AllFunctionRunStatus = []FunctionRunStatus{
	FunctionRunStatusCompleted,
	FunctionRunStatusFailed,
	FunctionRunStatusCancelled,
	FunctionRunStatusRunning,
}

func (e FunctionRunStatus) IsValid() bool {
	switch e {
	case FunctionRunStatusCompleted, FunctionRunStatusFailed, FunctionRunStatusCancelled, FunctionRunStatusRunning:
		return true
	}
	return false
}

func (e FunctionRunStatus) String() string {
	return string(e)
}

func (e *FunctionRunStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FunctionRunStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FunctionRunStatus", str)
	}
	return nil
}

func (e FunctionRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FunctionStatus string

const (
	FunctionStatusRunning   FunctionStatus = "RUNNING"
	FunctionStatusCompleted FunctionStatus = "COMPLETED"
	FunctionStatusFailed    FunctionStatus = "FAILED"
	FunctionStatusCancelled FunctionStatus = "CANCELLED"
)

var AllFunctionStatus = []FunctionStatus{
	FunctionStatusRunning,
	FunctionStatusCompleted,
	FunctionStatusFailed,
	FunctionStatusCancelled,
}

func (e FunctionStatus) IsValid() bool {
	switch e {
	case FunctionStatusRunning, FunctionStatusCompleted, FunctionStatusFailed, FunctionStatusCancelled:
		return true
	}
	return false
}

func (e FunctionStatus) String() string {
	return string(e)
}

func (e *FunctionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FunctionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FunctionStatus", str)
	}
	return nil
}

func (e FunctionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FunctionTriggerTypes string

const (
	FunctionTriggerTypesEvent FunctionTriggerTypes = "EVENT"
	FunctionTriggerTypesCron  FunctionTriggerTypes = "CRON"
)

var AllFunctionTriggerTypes = []FunctionTriggerTypes{
	FunctionTriggerTypesEvent,
	FunctionTriggerTypesCron,
}

func (e FunctionTriggerTypes) IsValid() bool {
	switch e {
	case FunctionTriggerTypesEvent, FunctionTriggerTypesCron:
		return true
	}
	return false
}

func (e FunctionTriggerTypes) String() string {
	return string(e)
}

func (e *FunctionTriggerTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FunctionTriggerTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FunctionTriggerTypes", str)
	}
	return nil
}

func (e FunctionTriggerTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StepEventType string

const (
	StepEventTypeScheduled StepEventType = "SCHEDULED"
	StepEventTypeStarted   StepEventType = "STARTED"
	StepEventTypeCompleted StepEventType = "COMPLETED"
	StepEventTypeErrored   StepEventType = "ERRORED"
	StepEventTypeFailed    StepEventType = "FAILED"
	StepEventTypeWaiting   StepEventType = "WAITING"
)

var AllStepEventType = []StepEventType{
	StepEventTypeScheduled,
	StepEventTypeStarted,
	StepEventTypeCompleted,
	StepEventTypeErrored,
	StepEventTypeFailed,
	StepEventTypeWaiting,
}

func (e StepEventType) IsValid() bool {
	switch e {
	case StepEventTypeScheduled, StepEventTypeStarted, StepEventTypeCompleted, StepEventTypeErrored, StepEventTypeFailed, StepEventTypeWaiting:
		return true
	}
	return false
}

func (e StepEventType) String() string {
	return string(e)
}

func (e *StepEventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StepEventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StepEventType", str)
	}
	return nil
}

func (e StepEventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StreamType string

const (
	StreamTypeEvent StreamType = "EVENT"
	StreamTypeCron  StreamType = "CRON"
)

var AllStreamType = []StreamType{
	StreamTypeEvent,
	StreamTypeCron,
}

func (e StreamType) IsValid() bool {
	switch e {
	case StreamTypeEvent, StreamTypeCron:
		return true
	}
	return false
}

func (e StreamType) String() string {
	return string(e)
}

func (e *StreamType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StreamType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StreamType", str)
	}
	return nil
}

func (e StreamType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
