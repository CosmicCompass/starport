package add

import (
	"fmt"
	"strings"
	
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
)

// New ...
func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	g.RunFn(appModify(opts))
	// g.RunFn(cmdMainModify(opts))
	if err := g.Box(packr.New("wasm", "./wasm")); err != nil {
		return g, err
	}
	ctx := plush.NewContext()
	ctx.Set("AppName", opts.AppName)
	ctx.Set("title", strings.Title)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Replace("{{appName}}", opts.AppName))
	return g, nil
}

const placeholder = "// this line is used by starport scaffolding"
const placeholder2 = "// this line is used by starport scaffolding # 2"
const placeholder3 = "// this line is used by starport scaffolding # 3"
const placeholder4 = "// this line is used by starport scaffolding # 4"
const placeholder5 = "// this line is used by starport scaffolding # 5"
const placeholder6 = "// this line is used by starport scaffolding # 6"
const placeholder7 = "// this line is used by starport scaffolding # 7"
const placeholder8 = "// this line is used by starport scaffolding # 8"
const placeholder9 = "// this line is used by starport scaffolding # 9"
const placeholder10 = "// this line is used by starport scaffolding # 10"
const placeholder11 = "// this line is used by starport scaffolding # 11"

func appModify(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		
		path := "app/app.go"
		f, err := r.Disk.Find(path)
		if err != nil {
			return err
		}
		
		switch opts.Feature {
		case "wasm":
			template := `%[1]v
	"path/filepath"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/spf13/viper"`
			replacement := fmt.Sprintf(template, placeholder)
			content := strings.Replace(f.String(), placeholder, replacement, 1)
			
			template2 := `%[1]v
		wasm.AppModuleBasic{},`
			replacement2 := fmt.Sprintf(template2, placeholder2)
			content = strings.Replace(content, placeholder2, replacement2, 1)
			
			template3 := `%[1]v
		wasmKeeper    wasm.Keeper`
			replacement3 := fmt.Sprintf(template3, placeholder3)
			content = strings.Replace(content, placeholder3, replacement3, 1)
			
			template4 := placeholder4 + "\n" + "type WasmWrapper struct { Wasm wasm.WasmConfig `mapstructure:\"wasm\"`}" + `
	var wasmRouter = bApp.Router()
	homeDir := viper.GetString(cli.HomeFlag)
	wasmDir := filepath.Join(homeDir, "wasm")

	wasmWrap := WasmWrapper{Wasm: wasm.DefaultWasmConfig()}
	err := viper.Unmarshal(&wasmWrap)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}
	wasmConfig := wasmWrap.Wasm
	supportedFeatures := "staking"
	app.wasmKeeper = wasm.NewKeeper(app.cdc, keys[wasm.StoreKey], app.accountKeeper, app.bankKeeper, app.stakingKeeper, wasmRouter, wasmDir, wasmConfig, supportedFeatures, nil, nil)	`
			content = strings.Replace(content, placeholder4, template4, 1)
			
			template5 := `%[1]v
	wasm.StoreKey,`
			replacement5 := fmt.Sprintf(template5, placeholder5)
			content = strings.Replace(content, placeholder5, replacement5, 1)
			
			template6 := `%[1]v
	wasm.NewAppModule(app.wasmKeeper),`
			replacement6 := fmt.Sprintf(template6, placeholder6)
			content = strings.Replace(content, placeholder6, replacement6, 1)
			
			template7 := `%[1]v
	wasm.ModuleName,`
			replacement7 := fmt.Sprintf(template7, placeholder7)
			content = strings.Replace(content, placeholder7, replacement7, 1)
			
			newFile := genny.NewFileS(path, content)
			return r.File(newFile)
		
		case "nfts":
			
			template := `%[1]v
	"github.com/FreeFlixMedia/modules/nfts"
	"github.com/FreeFlixMedia/modules/xnfts"
`
			replacement := fmt.Sprintf(template, placeholder)
			content := strings.Replace(f.String(), placeholder, replacement, 1)
			
			template2 := `%[1]v
	nfts.AppModuleBasic{},
	xnfts.AppModuleBasic{},
`
			replacement2 := fmt.Sprintf(template2, placeholder2)
			content = strings.Replace(content, placeholder2, replacement2, 1)
			
			template3 := `%[1]v
	nftsKeeper    nfts.Keeper
	xnftsKeeper xnfts.Keeper

`
			replacement3 := fmt.Sprintf(template3, placeholder3)
			content = strings.Replace(content, placeholder3, replacement3, 1)
			
			template4 := `%[1]v
	scopedXNFTKeeper capability.ScopedKeeper
`
			replacement4 := fmt.Sprintf(template4, placeholder4)
			content = strings.Replace(content, placeholder4, replacement4, 1)
			
			template5 := `%[1]v
	nfts.StoreKey,
	xnfts.StoreKey,
`
			replacement5 := fmt.Sprintf(template5, placeholder5)
			content = strings.Replace(content, placeholder5, replacement5, 1)
			
			template6 := `%[1]v
	scopedXNFTKeeper := app.capabilityKeeper.ScopeToModule(xnfts.ModuleName)
`
			replacement6 := fmt.Sprintf(template6, placeholder6)
			content = strings.Replace(content, placeholder6, replacement6, 1)
			
			template7 := `%[1]v
	app.nftsKeeper = nfts.NewKeeper(
		app.cdc,
		keys[nfts.StoreKey],
	)
	
	app.xnftsKeeper = xnfts.NewKeeper(
		app.cdc,
		keys[xnfts.StoreKey],
		app.nftsKeeper,
		app.bankKeeper,
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		scopedXNFTKeeper,
	)

	xnftsModule := xnfts.NewAppModule(app.xnftsKeeper)
`
			replacement7 := fmt.Sprintf(template7, placeholder7)
			content = strings.Replace(content, placeholder7, replacement7, 1)
			
			template8 := `%[1]v
	ibcRouter.AddRoute(xnfts.ModuleName, xnftsModule)
`
			replacement8 := fmt.Sprintf(template8, placeholder8)
			content = strings.Replace(content, placeholder8, replacement8, 1)
			
			template9 := `%[1]v
	nfts.NewAppModule(app.nftsKeeper),
	xnftsModule,
`
			replacement9 := fmt.Sprintf(template9, placeholder9)
			content = strings.Replace(content, placeholder9, replacement9, 1)
			
			template10 := `%[1]v
	nfts.ModuleName,
	xnfts.ModuleName,
`
			replacement10 := fmt.Sprintf(template10, placeholder10)
			content = strings.Replace(content, placeholder10, replacement10, 1)
			
			template11 := `%[1]v
	app.scopedXNFTKeeper = scopedXNFTKeeper
`
			replacement11 := fmt.Sprintf(template11, placeholder11)
			content = strings.Replace(content, placeholder11, replacement11, 1)
			
			newFile := genny.NewFileS(path, content)
			return r.File(newFile)
			
		}
		return nil
	}
}

func cmdMainModify(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		path := fmt.Sprintf("cmd/%[1]vcli/main.go", opts.AppName)
		f, err := r.Disk.Find(path)
		if err != nil {
			return err
		}
		template := `%[1]v
	wasmrest "github.com/CosmWasm/wasmd/x/wasm/client/rest"`
		replacement := fmt.Sprintf(template, placeholder)
		content := strings.Replace(f.String(), placeholder, replacement, 1)
		
		template2 := `%[1]v
	wasmrest.RegisterRoutes(rs.CliCtx, rs.Mux)`
		replacement2 := fmt.Sprintf(template2, placeholder2)
		content = strings.Replace(content, placeholder2, replacement2, 1)
		newFile := genny.NewFileS(path, content)
		return r.File(newFile)
	}
}
