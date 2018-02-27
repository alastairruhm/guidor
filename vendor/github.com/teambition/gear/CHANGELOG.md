# Change Log

All notable changes to this project will be documented in this file starting from version **v1.0.0**.
This project adheres to [Semantic Versioning](http://semver.org/).

-----
## [1.9.11] - 2018-01-27

**Changed:**

- Add X-Request-ID header to log.

## [1.9.10] - 2018-01-18

**Changed:**

- Add Referer and Origin header to log.
- Add gRPC midddleware.

## [1.9.9] - 2018-01-16

**Fixed:**

- Fix logging.

## [1.9.8] - 2018-01-11

**Changed:**

- Add SetTrustedProxy setting.

## [1.9.7] - 2018-01-08

**Changed:**

- Add status code 499 ClientClosedRequest and use it for context canceled error.

## [1.9.6] - 2017-12-26

**Fixed:**

- Change and fix ctx.IP().

## [1.9.5] - 2017-12-17

**Fixed:**

- Make ctx.ParseBody support application/jrd+json, application/jose+json and so on.

## [1.9.4] - 2017-12-16

**Changed:**

- Add ctx.Send method.

## [1.9.3] - 2017-12-15

**Fixed:**

- Fix gear.ContentDisposition for IE7/8.

## [1.9.2] - 2017-11-24

**Changed:**

- Add HTTP Status code 103 Early Hints.

**Fixed:**

- Fix CORS middleware.

## [1.9.1] - 2017-11-18

**Changed:**

- Improve logging.

## [1.9.0] - 2017-11-15

**Changed:**

- Add GetRouterNodeFromCtx function.

## [1.8.10] - 2017-11-13

**Changed:**

- Add Decompress function; ctx.ParseBody support compress content.

## [1.8.9] - 2017-11-07

**Changed:**

- Some methods of App, Router, and Logger return itself for chain-style.

## [1.8.8] - 2017-10-20

**Changed:**

- Add ctx.MustAny method and logger.SetTo method.

-----

## [1.8.7] - 2017-10-01

**Changed:**

- Add ctx.LogErr method.

-----

## [1.8.6] - 2017-09-19

**Changed:**

- Return error when recalling end method.

-----

## [1.8.5] - 2017-09-07

**Changed:**

- Add Log.With method.

-----

## [1.8.4] - 2017-08-26

**Changed:**

- Add Options.OnlyFiles for static middleware.

## [1.8.3] - 2017-08-12

**Fixed:**

- Fix logging. Add more test.

## [1.8.2] - 2017-08-11

**Fixed:**

- Fix error logging.

## [1.8.1] - 2017-08-08

**Changed:**

- Add `Includes` option to static middleware.

## [1.8.0] - 2017-08-07

**Changed:**

- Change ctx.Get to ctx.GetHeader; Change ctx.Set to ctx.SetHeader.

## [1.7.13] - 2017-08-06

**Fixed:**

- Fix empty status detecting.

## [1.7.12] - 2017-07-30

**Fixed:**

- Support Referrer header.


## [1.7.11] - 2017-07-27

**Fixed:**

- Fix const errors.

## [1.7.10] - 2017-07-19

**Fixed:**

- Fix gear.ValuesToStruct.

## [1.7.9] - 2017-07-19

**Changed:**

- Add Error.WithErr(name string).

## [1.7.8] - 2017-07-18

**Changed:**

- Update logging middleware: print time with brackets.

## [1.7.7] - 2017-07-09

**Changed:**

- Update defaultHeaderFilterReg for error response.
- Make end-hooks running in defer func.

## [1.7.6] - 2017-06-17

**Changed:**

- Add a QUIC example.

**Fixed:**

- Fix logging's development consume function.

## [1.7.5] - 2017-06-16

**Changed:**

- Add gear.LoggerFilterWriter to filter unnecessary server errors.

## [1.7.4] - 2017-06-14

**Fixed:**

- Fix root path for router with namepace.

## [1.7.3] - 2017-06-10

**Fixed:**

- Fix logging middleware, [issue 24](https://github.com/teambition/gear/issues/24).

## [1.7.2] - 2017-06-06

**Changed:**

- Add gear.SetServerName setting.

## [1.7.1] - 2017-06-01

**Changed:**

- Change development logger format to "Common Log Format".

## [1.7.0] - 2017-05-27

**Changed:**

- Improve gear.Error and logging.

## [1.6.2] - 2017-05-23

**Changed:**

- Improve ValuesToStruct function.

## [1.6.1] - 2017-05-22

**Fixed:**

- Add `X-Ratelimit-` to `defaultHeaderFilterReg`.

## [1.6.0] - 2017-05-20

**Changed:**

- Run "end hooks" in a goroutine, in order to not block current process.

## [1.5.3] - 2017-05-19

**Changed:**

- Add `Error.WithStack(skip ...int) *Error` and `Error.WithMsgf(format string, args ...interface{}) *Error`.

## [1.5.2] - 2017-05-18

**Changed:**

- Add `logging.Debugf(format string, args ...interface{})`.

## [1.5.1] - 2017-05-04

**Changed:**

- Add `ctx.Res.Body()` method.

## [1.5.0] - 2017-05-03

**Changed:**

- Add `ctx.ParseURL(body BodyTemplate)` method.
- Change `gear.FormToStruct(map[string][]string, interface{}) error` to `gear.ValuesToStruct(map[string][]string, interface{}, string) error`
- `gear.ValuesToStruct` support pointer fields, so `ctx.ParseURL` and `ctx.ParseBody` support pointer fields too.

## [1.4.3] - 2017-04-26

**Fixed:**

- Fix for multi-routers

## [1.4.2] - 2017-04-25

**Changed:**

- Improve `ctx.IP()`
- Add `ctx.Protocol()`

## [1.4.1] - 2017-04-24

**Changed:**

- Remove unnecessary error constants.

## [1.4.0] - 2017-04-23

**Changed:**

- Refactor `gear.Error` type. It is more powerful!

## [1.3.1] - 2017-04-12

**Changed:**

- Ignore `"context canceled"` error.
- Add more examples.

## [1.3.0] - 2017-03-28

**Changed:**

- Support Form body parse.
- Improve project structure.

## [1.2.0] - 2017-03-19

**Changed:**

- Add `Response.Status()`, `Response.Type()`.
- `Context.Status` and `Context.Type` will not return value now.

## [1.1.3] - 2017-03-15

**Changed:**

- Improve code.

## [1.1.2] - 2017-03-14

**Fixed:**

- Fix Logging middleware https://github.com/teambition/gear/pull/19 .

**Changed:**

- Add more document.

## [1.1.1] - 2017-03-11

**Changed:**

- Change default BodyParser's `MaxBytes` to `2MB`.
- Add more document.

## [1.1.0] - 2017-03-08

**Changed:**

- Simplify ctx.Timing method: `func (*Context) Timing(time.Duration, fn func(context.Context)) error`

## [1.0.5] - 2017-03-08

**Fixed:**

- Fix CORS middleware https://github.com/teambition/gear/pull/17 .

## [1.0.4] - 2017-03-05

**Fixed:**

- Fix logging middleware that should init structured log when start.

## [1.0.3] - 2017-03-05

**Fixed:**

- Fix Context.WithContext method that maybe recursive.

## [1.0.2] - 2017-03-02

**Fixed:**

- Fix Error.String() method that Error.Meta may be invalid utf-8 bytes.

## [1.0.1] - 2017-03-01

**Fixed:**

- Fix and improve Content-Disposition which is used by ctx.Attachment.
- Update Readme document.

-----

## [1.0.0] - 2017-03-01
