package main

import (
	"github.com/n04ln/diesirae.nvim/command"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	a, _ := command.NewAOJ()

	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "AojSubmit"}, a.SubmitAndCheckStatus)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AojRunSample"}, a.Trial)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AojDescription"}, a.Description)
		p.HandleCommand(&plugin.CommandOptions{Name: "AojStatus"}, a.Status)
		p.HandleCommand(&plugin.CommandOptions{Name: "AojSelf"}, a.Self)
		p.HandleCommand(&plugin.CommandOptions{Name: "AojSession"}, a.Session)
		p.HandleCommand(&plugin.CommandOptions{Name: "AojStatusList"}, a.StatusList)
		return nil
	})
}
