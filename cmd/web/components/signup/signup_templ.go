// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "life-streams/cmd/web"

func SignupPage(isLoggedIn bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n  .wrapper {\n      display: flex;\n      justify-content: center;\n      flex-direction: column;\n      width: min(400px, 100%);\n      margin-inline: auto;\n      background-color: var(--tile4);\n      border-radius: 4px;\n      border: 1px solid var(--tile3);\n  }\n\n    .form {\n      display: flex;\n      flex-direction: column; \n      width: 100%;\n      padding: 12px;\n\n      >h1 {\n        margin-bottom: 10px;\n        color: var(--text1);\n      }\n\n      .group {\n        width: 100%;\n        display: flex;\n        flex-direction: column;\n\n        > label {\n          margin-bottom: 5px;\n          color: var(--text1);\n        }\n      }\n\n      .button {\n        margin-top: 10px;\n        padding: 5px 10px;\n        background-color: var(--tile5);\n        color: var(--text1);\n        border: none;\n        width: 100%;\n        border-radius: 4px;\n        cursor: pointer;\n        transition: background-color 0.2s;\n\n        &:hover {\n          background-color: var(--tile6);\n        }\n      }\n\n      > a {\n        margin-top: 10px;\n        color: var(--text1);\n        transition: color 0.2s;\n\n        &:hover {\n          color: var(--text2);\n        }\n      }\n    }\n\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"wrapper\"><div id=\"form-response\"><!-- Form response will be rendered here --></div><form hx-post=\"/signup_post\" method=\"POST\" hx-target=\"#form-response\" class=\"form\"><h1>Signup</h1><div class=\"group\"><label for=\"email\">Email:</label> <input type=\"email\" id=\"email\" name=\"email\" required></div><div class=\"group\"><label for=\"password\">Password:</label> <input type=\"text\" id=\"password\" name=\"password\" required minlength=\"6\" maxlength=\"10\"></div><button type=\"submit\" class=\"button\">Sign Up</button> <a href=\"/login\">Login</a></form></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = web.Base(isLoggedIn).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SignUpError(message string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    .error-wrapper {\n      display: flex;\n      justify-content: start;\n      flex-direction: column;\n      width: 100%;\n      padding: 12px;\n      text-align: left;\n      background-color: var(--red1);\n\n      > h2 {\n        font-size: 1.25rem;\n        color: var(--text1);\n      }\n\n      > p {\n        font-size: 1rem;\n        color: var(--text1);\n      }\n    }\n  </style><div class=\"error-wrapper\"><h2>Signup failed</h2><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(message)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/web/components/signup/signup.templ`, Line: 114, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
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

func SignUpSuccess() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    .success-wrapper {\n      display: flex;\n      justify-content: start;\n      flex-direction: column;\n      width: 100%;\n      padding: 12px;\n      text-align: left;\n      background-color: var(--green1);\n\n      > h2 {\n        font-size: 1.25rem;\n        color: var(--text1);\n      }\n\n      > p {\n        font-size: 1rem;\n        color: var(--text1);\n      }\n    }\n  </style><script>\n    setTimeout(() => {\n      window.location.href = '/login'\n    }, 3000)\n  </script><div class=\"success-wrapper\"><h2>Signup successful</h2><p>Redirecting to login page...</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
