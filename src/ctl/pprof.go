// package ctl

// import (
// 	"net/http/pprof"

// 	"github.com/teambition/gear"
// )

// type Pprof struct {
// }

// // Init ...
// func (p *Pprof) Get(ctx *gear.Context) error {
// 	switch ctx.Param("pp") {
// 	default:
// 		pprof.Index(ctx.Res.w, ctx.Req)
// 	case "":
// 		pprof.Index(ctx.Res.w, ctx.Req)
// 	case "cmdline":
// 		pprof.Cmdline(ctx.Res.w, ctx.Req)
// 	case "profile":
// 		pprof.Profile(ctx.Res.w, ctx.Req)
// 	case "symbol":
// 		pprof.Symbol(ctx.Res.w, ctx.Req)
// 	}
// 	ctx.Res.WriteHeader(200)
// }
