// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	stream_types "life-streams/internal/server/handlers/stream/types"
	"strconv"
)

func CreateTaskModal(streams []stream_types.Stream, streamID string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    dialog {\n      position: absolute;\n      top: 50%;\n      left: 50%;\n      transform: translate(-50%, -50%);\n      width: min(500px, 100%);\n      background-color: var(--tile4);\n      border: 1px solid var(--tile1);\n\n      > h2 {\n        margin: 0;\n        color: var(--text1);\n        font-size: 1.25rem;\n        font-weight: bold;\n      }\n      > form {\n        display: flex;\n        flex-direction: column;\n        gap: 10px;\n\n        > .group {\n          width: 100%;\n          display: flex;\n          flex-direction: column;\n          color: var(--text1);\n\n          > label {\n            margin-bottom: 5px;\n            color: var(--text1);\n          }\n\n          > textarea {\n            resize: none;\n            padding: 5px;\n            border: 1px solid var(--tile3);\n            background-color: var(--tile5);\n          }\n        }\n\n        > .buttons {\n          margin-top: 8px;\n          width: 100%;\n          display: flex;\n          flex-direction: row;\n          gap: 4px;\n\n          > .button {\n            padding: 5px 10px;\n            border: none;\n            border-radius: 5px;\n            cursor: pointer;\n            flex: 1;\n        }\n\n        > .task-cancel {\n            background-color: var(--tile2);\n            color: var(--red1);\n            font-weight: bold;\n        }\n      }\n    }\n  </style><dialog class=\"task-modal-dialog\"><div id=\"task-modal-form-response\"><!-- Form response will be rendered here --></div><h2>Create Task</h2><form hx-post=\"/create_task\" hx-target=\"#task-modal-form-response\" method=\"POST\" class=\"task-form\"><div class=\"group\"><label for=\"title\">Task Name</label> <input type=\"text\" id=\"title\" name=\"title\" required></div><div class=\"group\"><label for=\"description\">Description</label> <textarea id=\"description\" name=\"description\" required></textarea></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if streamID != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"hidden\" name=\"stream\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(streamID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/web/components/create_task_modal/create_task_modal.templ`, Line: 88, Col: 61}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"group\"><label for=\"stream\">Stream</label> <select id=\"stream\" name=\"stream\" required><option value=\"\" disabled selected>Select a stream</option> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, stream := range streams {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(stream.ID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/web/components/create_task_modal/create_task_modal.templ`, Line: 95, Col: 57}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(stream.Title)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/web/components/create_task_modal/create_task_modal.templ`, Line: 95, Col: 73}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"buttons\"><button value=\"cancel\" type=\"button\" class=\"button task-cancel\">Cancel</button> <button type=\"submit\" class=\"button submit\">Submit</button></div></form></dialog> <button class=\"show-task-modal-button\">Create Task</button><script>\n    (() => {\n      console.log(\"Create Task Modal Loaded\");\n      const showTaskModalButton = document.querySelector(\".show-task-modal-button\");\n      const taskModalDialog = document.querySelector(\".task-modal-dialog\");\n\n      // \"Cancel\" button closes the <dialog>\n      document.querySelector(\".task-cancel\").addEventListener(\"click\", () => {\n        document.querySelector(\".task-form\").reset();\n        document.querySelector(\"#task-modal-form-response\").innerHTML = \"\";\n        taskModalDialog.close();\n      });\n      // \"Show the dialog\" button opens the <dialog> modally\n      showTaskModalButton.addEventListener(\"click\", () => {\n        taskModalDialog.showModal();\n      });\n    })()\n  </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateTaskError(message string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    .error-wrapper {\n      display: flex;\n      justify-content: start;\n      flex-direction: column;\n      width: 100%;\n      padding: 12px;\n      text-align: left;\n      background-color: var(--red1);\n\n      > h2 {\n        font-size: 1.25rem;\n        color: var(--text1);\n      }\n\n      > p {\n        font-size: 1rem;\n        color: var(--text1);\n      }\n    }\n  </style><div class=\"error-wrapper\"><h2>Failed to create task!</h2><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(message)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/web/components/create_task_modal/create_task_modal.templ`, Line: 152, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

var helloHandle = templ.NewOnceHandle()

func CreateTaskSuccess() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    .success-wrapper {\n      display: flex;\n      justify-content: start;\n      flex-direction: column;\n      width: 100%;\n      padding: 12px;\n      text-align: left;\n      background-color: var(--green1);\n\n      > h2 {\n        font-size: 1.25rem;\n        color: var(--text1);\n      }\n    }\n  </style><script>\n    (() => {\n      let form = document.querySelector('.task-form');\n      form.reset();\n      let dialog = document.querySelector('.task-modal-dialog');\n      setTimeout(() => {\n        dialog.close();\n      }, 3000)\n    })()\n  </script><div class=\"success-wrapper\"><h2>Task created successfully</h2></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
