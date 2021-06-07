package actions

import (
	"inngest.com/edges"
)

// Redefining edges.#Edge based off of the existing definition fixes an "imported but not used"
// error within cue.
#Edge: edges.#Edge

#Action: {
	name:     string
	dsn:      string

	version: {
		major: >0 | *1
		minor: >0 | *1
	}

	// Settings represents settings stored for this action within a workspace.
	// These settings are configurable by end users.
	//
	// A setting category can be defined as:
	//  1. "*" - meaining it can have any number of settings (eg. custom webhooks, email templates).
	//  2. A map of setting names to their definitions, ensuring that users configure each setting correctly.
	settings?: {[Category=_]:
		// A setting category can have either any setting names, or a specific set.
		{[Name=_]: #Setting & {name: Name}}
	}

	// workflowMetadata represents workflow-specific metadata used to configure the action
	workflowMetadata?: {[Name=_]: #Metadata & {name: Name}}

	// response defines an object of key => values which the action returns.  This allows
	// you to strongly type the responses for coordination within larger workflows.
	response?: {[Name=_]: #Response & {name: Name}}

	// Edges allow you to add predefined edges in the UI.
	edges?: {[Name=_]: #Edge & {name: Name}}

	// runtime indicates the runtime that you wish to use to execute this action.
	// Currently, only the docker runtime is supported. 
	runtime: #Runtime
}

#Runtime: (#RuntimeDocker)

#RuntimeDocker: {
	type: "docker"
	image: string
	entrypoint?: string
}

// ActionType depicts the type of the response and metadata input.
#ActionType: "string" | "int" | "float" | "bool" | "json" | "datetime" | "duration"

// Response represents a single top-level key in the response object, generated by this
// action.
//
// For example, if this action performs an HTTP call, the 
#Response: {
	name:     string
	type:     #ActionType
	optional: bool | *false // whether this is an optional/nullable type.
}

// Setting is an enum of specific settings which can be set for this action.
#Setting: {
	name:    string
	form:    #Form
} & (#SettingEnum | #SettingString | #SettingWildcard)

#SettingWildcard: {
	name: string
	type: "wildcard"
	form: #Form
}

#SettingEnum: {
	name:    string
	type:    "enum"
	form:    #Form
	choices: [ ...string] | *[]
}

#SettingString: {
	name:    string
	type:    "string"
	form:    #Form
}

// Metadata represents metadata used to configure the action within a specific workflow.
#Metadata: {
	// The name of the metadata used in our setting.
	name: string

	// The default value for this metadata.
	default?: _

	// An expression that must evauluate to true if this metadata is valid.
	//
	// For example, during the "Wait" action you can configure one piece of
	// metadata (type) to choose between "wait for" and "wait until".  These two
	// other settings have an expression of `metadata.type == "for"` and
	// `metadata.type == "wait"` respectively.
	expression?: string

	// True, false, or an expression
	required: bool | *false

	// TODO: Fully define metadata types for each action.
	type: #ActionType

	// Form lets you specify the UI type witihn the workflow configuration panel, allowing\
	// non-technical users to edit action metadata.
	form: #Form
}

#Form: (#FormInput | #FormSelect | #FormTextarea | #FormDateTime)

#FormInput: {
	title:      string
	type:       "input"
	templating: bool | *false
}

#FormDateTime: {
	title:      string
	type:       "datetime"
	templating: bool | *false
}

#FormTextarea: {
	title:      string
	type:       "textarea"
	templating: bool | *false
}

#FormSelect: {
	title:   string
	type:    "select"
	choices: [ ...{value: string, name: string}] | *[]
	eval?:   string
}

