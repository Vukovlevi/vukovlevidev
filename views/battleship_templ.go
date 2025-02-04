// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Battleship() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"hu\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Vukovlevi - torpedó</title><link rel=\"stylesheet\" href=\"/css/battleship.css\"><link rel=\"icon\" type=\"image/x-icon\" href=\"/images/vukov_logo.png\"></head><body><header><nav class=\"container\"><a href=\"#\" class=\"logo\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><circle cx=\"12\" cy=\"5\" r=\"3\"></circle> <line x1=\"12\" y1=\"22\" x2=\"12\" y2=\"8\"></line> <path d=\"M5 12H2a10 10 0 0 0 20 0h-3\"></path></svg> <span>Torpedó</span></a><ul><li><a href=\"#features\">Lehetőségek</a></li><li><a href=\"#download\">Letöltés</a></li><li><a href=\"#about\">Rólam</a></li></ul></nav></header><main><section class=\"hero\"><div class=\"container\"><h1>Üdvözöllek a Torpedó hivatalos weboldalán</h1><p>Játssz barátaiddal, éld újra a gyerekkori emlékeket! Próbáld ki a klasszikus társasjáték online verzióját!</p><a href=\"#download\" class=\"button\">Letöltés</a> <a href=\"#features\" class=\"button secondary\">Tudj meg többet</a></div></section><section id=\"features\" class=\"features\"><div class=\"container\"><h2>A játék részei</h2><div class=\"feature-grid\"><div class=\"feature\"><div class=\"feature-icon\">🛡️</div><h3>Stratégia és szerencse</h3><p>Ismerd ki ellenfeled és alkalmazz ellenstratégiákat a tippjeire, vagy legyél a szerencsés, akinek minden lövése találat!</p></div><div class=\"feature\"><div class=\"feature-icon\">🌎</div><h3>Meccskeresés az egész világon</h3><p>Indíts egy meccskeresést és engedd, hogy a játék bárkivel összedobjon a világon!</p></div><div class=\"feature\"><div class=\"feature-icon\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 640 512\"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.--><path d=\"M96 128a128 128 0 1 1 256 0A128 128 0 1 1 96 128zM0 482.3C0 383.8 79.8 304 178.3 304l91.4 0C368.2 304 448 383.8 448 482.3c0 16.4-13.3 29.7-29.7 29.7L29.7 512C13.3 512 0 498.7 0 482.3zM609.3 512l-137.8 0c5.4-9.4 8.6-20.3 8.6-32l0-8c0-60.7-27.1-115.2-69.8-151.8c2.4-.1 4.7-.2 7.1-.2l61.4 0C567.8 320 640 392.2 640 481.3c0 17-13.8 30.7-30.7 30.7zM432 256c-31 0-59-12.6-79.3-32.9C372.4 196.5 384 163.6 384 128c0-26.8-6.6-52.1-18.3-74.3C384.3 40.1 407.2 32 432 32c61.9 0 112 50.1 112 112s-50.1 112-112 112z\"></path></svg></div><h3>Játssz barátokkal</h3><p>Nem szeretnél véletlenszerűen ellenfelet kapni? Egy barátoddal akarsz játszani? Erre is van lehetőség! Találjatok ki egy kódot, amit csak ti ismertek, majd csatlakozzatok egy játékhoz ezzel a kóddal és garantáltan egymás ellen kerültök!</p></div></div></div></section><section id=\"download\" class=\"download\"><div class=\"container\"><h2>Készen állsz a játékra?</h2><p>Töltsd le a Torpedót és játssz pillanatokon belül.</p><a href=\"/content/battleship/Torpedó telepítő.exe\" download class=\"button\">Letöltés Windowsra</a></div></section><section id=\"about\" class=\"about\"><div class=\"container\"><h2>A fejlesztőről</h2><p>Egy átlagos középiskolás srác vagyok, akinek épp a weboldalán jársz. Nagyon szerettem kiskoromban torpedózni, ez jutott eszembe, amikor azon gondolkoztam mit is csinálhatnék, ami technikai kihívást jelentene. Ennek eredményeként született meg a Torpedó. Elsősorban technikai szempontból volt érdekes számomra, szóval ha ez téged is érdekel, akkor látogass <a title=\"Hamarosan\">ide!</a></p><p>Egyébként nem az egyetlen projektem, ha érdekel mit alkottam még, akkor <a title=\"Hamarosan\">nézd meg a többi művemet is!</a></p></div></section></main><footer><div class=\"container\"><p>&copy; 2024 Torpedó. All rights reserved.</p></div></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
